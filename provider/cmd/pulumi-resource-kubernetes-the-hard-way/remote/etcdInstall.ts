import { ComponentResourceOptions, Input, Output, interpolate, output } from '@pulumi/pulumi';
import { RandomString } from '@pulumi/random';
import * as schema from '../schema-types';
import { Mkdir, Mv, Tar } from '../tools';
import { Download } from './download';

export type Architecture = 'amd64' | 'arm64';

export class EtcdInstall extends schema.EtcdInstall {
  public static readonly defaultArch: Architecture = 'amd64';
  public static readonly defaultInstallDirectory: string = '/usr/local/bin';
  public static readonly defaultVersion: string = '3.4.15'; // TODO: Versioning
  private readonly _nameInput: string;

  constructor(name: string, args: schema.EtcdInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    this._nameInput = name;

    // Rehydrating
    if (opts?.urn) return;

    const architecture = output(args.architecture ?? EtcdInstall.defaultArch);
    const downloadDirectory = this.getDownloadDirectory(args.downloadDirectory);
    const installDirectory = output(args.installDirectory ?? EtcdInstall.defaultInstallDirectory);
    const version = output(args.version ?? EtcdInstall.defaultVersion); // TODO: Stateful versioning?
    const archiveName = interpolate`etcd-v${version}-linux-${architecture}.tar.gz`;
    const url = interpolate`https://github.com/etcd-io/etcd/releases/download/v${version}/${archiveName}`;

    // TODO: Permission checks?
    // TODO: Caching? Put archive/bins into ~/.kthw/cache so i.e. installDirectory changes, tarball doesn't need to be re-downloaded.
    // TODO: General update logic

    const downloadMkdir = new Mkdir(`${name}-download`, {
      connection: args.connection,
      directory: downloadDirectory,
      parents: true,
      removeOnDelete: true,
    }, { parent: this });

    const download = new Download(name, {
      connection: args.connection,
      destination: downloadDirectory,
      url,
    }, { parent: this, dependsOn: downloadMkdir });

    const tar = new Tar(name, {
      connection: args.connection,
      archive: interpolate`${download.destination}/${archiveName}`,
      directory: download.destination,
      gzip: true,
      stripComponents: 1,
    }, { parent: this, dependsOn: download });

    const installMkdir = new Mkdir(`${name}-install`, {
      connection: args.connection,
      directory: installDirectory,
      parents: true,
    }, { parent: this });

    const etcdPath = interpolate`${installDirectory}/etcd`;
    const etcdctlPath = interpolate`${installDirectory}/etcdctl`;

    const mvEtcd = new Mv(`${name}-etcd`, {
      connection: args.connection,
      source: interpolate`${download.destination}/etcd`,
      dest: etcdPath,
    }, { parent: this, dependsOn: [tar, installMkdir] });

    const mvEtcdctl = new Mv(`${name}-etcdctl`, {
      connection: args.connection,
      source: interpolate`${download.destination}/etcdctl`,
      dest: etcdctlPath,
    }, { parent: this, dependsOn: [tar, installMkdir] });

    // TODO: Rm archive or tmp dir?

    this.architecture = architecture;
    this.archiveName = archiveName;
    this.download = download;
    this.downloadDirectory = downloadDirectory;
    this.downloadMkdir = downloadMkdir;
    this.installDirectory = installDirectory;
    this.installMkdir = installMkdir;
    this.mvEtcd = mvEtcd;
    this.mvEtcdctl = mvEtcdctl;
    this.name = output(name);
    this.tar = tar;
    this.url = url;
    this.version = version;

    this.registerOutputs({
      architecture,
      archiveName,
      download,
      downloadDirectory,
      downloadMkdir,
      installDirectory,
      installMkdir,
      mvEtcd,
      mvEtcdctl,
      name,
      tar,
      url,
      version,
    });
  }

  private getDownloadDirectory(input?: Input<string>): Output<string> {
    if (input) {
      return output(input);
    }

    // TODO: Mktemp?
    const name = this._nameInput;
    const random = new RandomString(name, {
      length: 8,
      special: false,
      upper: false,
    }, { parent: this });

    return interpolate`/tmp/${random.result}`;
  }
}
