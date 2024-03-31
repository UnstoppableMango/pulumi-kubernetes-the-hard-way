import { ComponentResource, ComponentResourceOptions, Input, Output, interpolate, output } from '@pulumi/pulumi';
import { RandomString } from '@pulumi/random';
import { remote } from '@pulumi/command/types/input';
import { Mkdir, Tar } from './tools';
import { Download } from './remote';

export type Architecture = 'amd64' | 'arm64';

export interface EtcdArgs {
  architecture?: Input<Architecture>;
  connection: Input<remote.ConnectionArgs>;
  downloadDirectory?: Input<string>;
  version?: Input<string>;
}

export class Etcd extends ComponentResource {
  public static readonly defaultArch: Architecture = 'amd64';
  public static readonly defaultVersion: string = '3.4.15';

  public readonly architecture!: Output<Architecture>;
  public readonly download!: Download;
  public readonly downloadDirectory!: Output<string>;
  public readonly filename!: Output<string>;
  public readonly mkdir!: Mkdir;
  public readonly name!: string;
  public readonly tar!: Tar;
  public readonly url!: Output<string>;
  public readonly version!: Output<string>;

  constructor(name: string, args: EtcdArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:index:Etcd', name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    this.name = name;
    const architecture = output(args.architecture ?? Etcd.defaultArch);
    const downloadDirectory = this.getDownloadDirectory(args.downloadDirectory);
    const version = output(args.version ?? Etcd.defaultVersion);
    const filename = interpolate`etcd-v${version}-linux-${architecture}.tar.gz`;
    const url = interpolate`https://github.com/etcd-io/etcd/releases/download/v${version}/${filename}`;

    const mkdir = new Mkdir(name, {
      connection: args.connection,
      directory: downloadDirectory,
      parents: true,
      removeOnDelete: true,
    }, { parent: this });

    const download = new Download(name, {
      connection: args.connection,
      destination: downloadDirectory,
      url,
    }, { parent: this, dependsOn: mkdir });

    const tar = new Tar(name, {
      connection: args.connection,
      archive: interpolate`${download.destination}/${filename}`,
      directory: interpolate`/tmp`, // TODO: Allow customizing base dir
      gzip: true,
    }, { parent: this, dependsOn: download });

    // TODO: Move to /usr/local/bin?
    // Reminder: The archive contains etcdctl and etcd

    this.architecture = architecture;
    this.download = download;
    this.downloadDirectory = downloadDirectory;
    this.filename = filename;
    this.mkdir = mkdir;
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

  private getDownloadDirectory(input?: Input<string>): Output<string> {
    if (input) {
      return output(input);
    }

    const random = new RandomString(this.name, {
      length: 8,
      special: false,
      upper: false,
    }, { parent: this });

    return interpolate`/tmp/${random.result}`;
  }
}
