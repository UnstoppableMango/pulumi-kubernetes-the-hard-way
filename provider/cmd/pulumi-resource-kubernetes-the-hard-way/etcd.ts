import { ComponentResource, ComponentResourceOptions, Input, Output, interpolate, output } from '@pulumi/pulumi';
import { remote } from '@pulumi/command/types/input';
import { Tar } from './tools';
import { Download } from './remote';

export type Architecture = 'amd64' | 'arm64';

export interface EtcdArgs {
  architecture?: Input<Architecture>;
  connection: Input<remote.ConnectionArgs>;
  version?: Input<string>;
}

export class Etcd extends ComponentResource {
  public static readonly defaultArch: Architecture = 'amd64';
  public static readonly defaultVersion: string = '3.4.15';

  public readonly architecture!: Output<Architecture>;
  public readonly download!: Download;
  public readonly downloadDirectory!: Output<string>;
  public readonly filename!: Output<string>;
  public readonly tar!: Tar;
  public readonly url!: Output<string>;
  public readonly version!: Output<string>;

  constructor(name: string, args: EtcdArgs, opts?: ComponentResourceOptions) {
    super('thecluster:index:etcd', name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const architecture = output(args.architecture ?? Etcd.defaultArch);
    const downloadDirectory = interpolate`/tmp`; // TODO: Temp directory
    const version = output(args.version ?? Etcd.defaultVersion);
    const filename = interpolate`etcd-v${version}-linux-${architecture}.tar.gz`;
    const url = interpolate`https://github.com/etcd-io/etcd/releases/download/v${version}/${filename}`;

    const download = new Download(name, {
      connection: args.connection,
      destination: downloadDirectory,
      url,
    }, { parent: this });

    const tar = new Tar(name, {
      connection: args.connection,
      archive: interpolate`${download.destination}/${filename}`,
      directory: interpolate`/tmp`, // TODO: Allow customizing base dir
      gzip: true,
    }, { parent: this });

    // TODO: Move to /usr/local/bin?
    // Reminder: The archive contains etcdctl and etcd

    this.architecture = architecture;
    this.download = download;
    this.downloadDirectory = downloadDirectory;
    this.filename = filename;
    this.tar = tar;
    this.url = url;
    this.version = version;

    this.registerOutputs({
      architecture,
      download,
      downloadDirectory,
      filename,
      tar,
      url,
      version,
    });
  }
}
