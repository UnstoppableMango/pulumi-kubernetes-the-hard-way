import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { archiveInstall } from './archiveInstall';

export class ContainerdInstall extends schema.ContainerdInstall {
  constructor(name: string, args: schema.ContainerdInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const architecture = output(args.architecture ?? 'amd64');
    const connection = output(args.connection);
    const directory = output(args.directory ?? '/usr/local/bin');
    const version = output(args.version ?? '1.4.4'); // TODO: Stateful versioning?
    const archiveName = interpolate`containerd-${version}-linux-${architecture}.tar.gz`;
    const url = interpolate`https://github.com/containerd/containerd/releases/download/v${version}/${archiveName}`;

    // TODO: Permission checks?
    // TODO: Caching? Put archive/bins into ~/.kthw/cache so i.e. directory changes, tarball doesn't need to be re-downloaded.
    // TODO: General update logic

    const install = archiveInstall(name, {
      archiveName,
      binaries: ['containerd'] as const,
      connection,
      directory,
      stripComponents: 1,
      url,
    }, this);

    this.architecture = architecture;
    this.archiveName = archiveName;
    this.download = install.download;
    this.path = install.paths.containerd;
    this.directory = directory;
    this.mkdir = install.mkdir;
    this.mv = install.mvs.containerd;
    this.tar = install.tar;
    this.url = url;
    this.version = version;

    this.registerOutputs({
      architecture,
      archiveName,
      download: install.download,
      path: install.paths.containerd,
      directory,
      mkdir: install.mkdir,
      mv: install.mvs.containerd,
      name,
      tar: install.tar,
      url,
      version,
    });
  }
}
