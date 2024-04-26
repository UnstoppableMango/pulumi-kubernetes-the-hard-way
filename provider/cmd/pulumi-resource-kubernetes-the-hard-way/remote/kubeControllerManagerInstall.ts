import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { binaryInstall } from './binaryInstall';

export class KubeControllerManagerInstall extends schema.KubeControllerManagerInstall {
  constructor(name: string, args: schema.KubeControllerManagerInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const architecture = output(args.architecture ?? 'amd64');
    const binName = 'kube-controller-manager';
    const connection = output(args.connection);
    const directory = output(args.directory ?? '/usr/local/bin');
    const version = output(args.version ?? '1.29.2');
    const url = interpolate`https://storage.googleapis.com/kubernetes-release/release/v${version}/bin/linux/${architecture}/${binName}`;

    const { download, mkdir, mktemp, mv, path, rm } = binaryInstall(name, {
      binName,
      connection,
      directory,
      url,
    }, this);

    this.architecture = architecture;
    this.binName = binName;
    this.connection = connection;
    this.directory = directory;
    this.mkdir = mkdir;
    this.mv = mv;
    this.path = path;
    this.version = version;

    this.registerOutputs({
      architecture,
      binName,
      connection,
      directory,
      mkdir,
      mv,
      path,
      version,
      url,
    });
  }
}
