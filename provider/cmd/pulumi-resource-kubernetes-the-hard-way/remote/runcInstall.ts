import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { binaryInstall } from './binaryInstall';

export class RuncInstall extends schema.RuncInstall {
  constructor(name: string, args: schema.RuncInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const architecture = output(args.architecture ?? 'amd64');
    const binName = interpolate`runc.${architecture}`;
    const connection = output(args.connection);
    const directory = output(args.directory ?? '/usr/local/bin');
    const version = output(args.version ?? '1.1.13');
    const url = interpolate`https://github.com/opencontainers/runc/releases/download/v${version}/${binName}`;

    const { download, mkdir, mktemp, mv, path, rm } = binaryInstall(name, {
      binName,
      connection,
      directory,
      url,
      finalBin: 'runc',
    }, this);

    this.architecture = architecture;
    this.connection = connection;
    this.directory = directory;
    this.download = download;
    this.mkdir = mkdir;
    this.mktemp = mktemp;
    this.mv = mv;
    this.path = path;
    this.rm = rm;
    this.version = version;

    this.registerOutputs({
      architecture,
      connection,
      directory,
      download,
      mkdir,
      mktemp,
      mv,
      path,
      rm,
      version,
    });
  }
}
