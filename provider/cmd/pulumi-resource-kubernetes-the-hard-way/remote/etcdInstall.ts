import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { archiveInstall } from './archiveInstall';

export type Architecture = 'amd64' | 'arm64';

export class EtcdInstall extends schema.EtcdInstall {
  public static readonly defaultArch: Architecture = 'amd64';
  public static readonly defaultdirectory: string = '/usr/local/bin';
  public static readonly defaultVersion: string = '3.4.15'; // TODO: Versioning

  constructor(name: string, args: schema.EtcdInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const architecture = output(args.architecture ?? EtcdInstall.defaultArch);
    const connection = output(args.connection);
    const directory = output(args.directory ?? EtcdInstall.defaultdirectory);
    const version = output(args.version ?? EtcdInstall.defaultVersion); // TODO: Stateful versioning?
    const archiveName = interpolate`etcd-v${version}-linux-${architecture}.tar.gz`;
    const url = interpolate`https://github.com/etcd-io/etcd/releases/download/v${version}/${archiveName}`;

    // TODO: Permission checks?
    // TODO: Caching? Put archive/bins into ~/.kthw/cache so i.e. directory changes, tarball doesn't need to be re-downloaded.
    // TODO: General update logic

    const install = archiveInstall(name, {
      archiveName,
      binaries: ['etcd', 'etcdctl'] as const,
      connection,
      directory,
      stripComponents: 1,
      url,
    }, this);

    this.architecture = architecture;
    this.archiveName = archiveName;
    this.download = install.download;
    this.etcdPath = install.paths.etcd;
    this.etcdctlPath = install.paths.etcdctl;
    this.directory = directory;
    this.mkdir = install.mkdir;
    this.etcdMv = install.mvs.etcd;
    this.etcdctlMv = install.mvs.etcdctl;
    // this.name = output(name);
    this.tar = install.tar;
    this.url = url;
    this.version = version;

    this.registerOutputs({
      architecture,
      archiveName,
      download: install.download,
      etcdPath: install.paths.etcd,
      etcdctlPath: install.paths.etcdctl,
      directory,
      mkdir: install.mkdir,
      mvEtcd: install.mvs.etcd,
      mvEtcdctl: install.mvs.etcdctl,
      name,
      tar: install.tar,
      url,
      version,
    });
  }

  // public etcdctl(inputs: schema.EtcdInstall_etcdctlInputs): Output<schema.EtcdInstall_etcdctlOutputs> {
  //   throw new Error('not implemented');
  // }
}
