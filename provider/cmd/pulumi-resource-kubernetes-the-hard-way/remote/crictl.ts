import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { archiveInstall } from './archiveInstall';

export class CrictlInstall extends schema.CrictlInstall {
  constructor(name: string, args: schema.CrictlInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const architecture = output(args.architecture ?? 'amd64');
    const connection = output(args.connection);
    const directory = output(args.directory ?? '/usr/local/bin');
    const version = output(args.version ?? '1.29.0'); // TODO: Stateful versioning?
    const archiveName = interpolate`crictl-v${version}-linux-${architecture}.tar.gz`;
    const url = interpolate`https://github.com/kubernetes-sigs/cri-tools/releases/download/v${version}/crictl-v${version}-linux-${architecture}.tar.gz`;

    // TODO: Permission checks?
    // TODO: Caching? Put archive/bins into ~/.kthw/cache so i.e. directory changes, tarball doesn't need to be re-downloaded.
    // TODO: General update logic

    const install = archiveInstall(name, {
      archiveName,
      binaries: ['crictl'] as const,
      connection,
      directory,
      stripComponents: 0,
      url,
    }, this);

    this.architecture = architecture;
    this.archiveName = archiveName;
    this.download = install.download;
    this.path = install.paths.crictl;
    this.directory = directory;
    this.mktemp = install.mktemp;
    this.mv = install.mvs.crictl;
    this.tar = install.tar;
    this.url = url;
    this.version = version;

    this.registerOutputs({
      architecture,
      archiveName,
      download: install.download,
      path: install.paths.crictl,
      directory,
      mkdir: install.mkdir,
      mv: install.mvs.crictl,
      name,
      tar: install.tar,
      url,
      version,
    });
  }
}
