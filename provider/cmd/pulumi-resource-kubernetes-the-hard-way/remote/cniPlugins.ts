import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { archiveInstall } from './archiveInstall';

export class CniPluginsInstall extends schema.CniPluginsInstall {
  constructor(name: string, args: schema.CniPluginsInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const architecture = output(args.architecture ?? 'amd64');
    const connection = output(args.connection);
    const directory = output(args.directory ?? '/usr/local/bin');
    const version = output(args.version ?? '0.9.1'); // TODO: Stateful versioning?
    const archiveName = interpolate`cni-plugins-linux-${architecture}-v${version}.tgz`;
    const url = interpolate`https://github.com/containernetworking/plugins/releases/download/v${version}/${archiveName}`;

    // TODO: Permission checks?
    // TODO: Caching? Put archive/bins into ~/.kthw/cache so i.e. directory changes, tarball doesn't need to be re-downloaded.
    // TODO: General update logic

    const { download, mkdir, mktemp, mvs, paths, rm, tar } = archiveInstall(name, {
      archiveName,
      binaries: ['cni-plugins'] as const,
      connection,
      directory,
      url,
    }, this);

    this.architecture = architecture;
    this.archiveName = archiveName;
    this.download = download;
    this.path = paths['cni-plugins'];
    this.directory = directory;
    this.mktemp = mktemp;
    this.mv = mvs['cni-plugins'];
    this.tar = tar;
    this.url = url;
    this.version = version;

    this.registerOutputs({
      architecture,
      archiveName,
      download,
      path: paths['cni-plugins'],
      directory,
      mkdir,
      mv: mvs['cni-plugins'],
      name,
      tar,
      url,
      version,
    });
  }
}
