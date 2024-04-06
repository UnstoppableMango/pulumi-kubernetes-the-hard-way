import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { binaryInstall } from './binaryInstall';

export class KubeControllerManagerInstall extends schema.KubeControllerManagerInstall {
  constructor(name: string, args: schema.KubeControllerManagerInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const architecture = output(args.architecture ?? 'amd64');
    const binName = 'kube-controller-manager';
    const connection = output(args.connection);
    const directory = output(args.directory ?? '/usr/local/bin');
    const version = output(args.version ?? '1.29.2');
    const url = interpolate`https://storage.googleapis.com/kubernetes-release/release/v${version}/bin/linux/${architecture}/${binName}`;

    const install = binaryInstall(name, {
      binName,
      connection,
      directory,
      url,
    }, this);

    this.architecture = architecture;
    this.binName = binName;
    this.connection = connection;
    this.directory = directory;
    this.mkdir = install.mkdir;
    this.mv = install.mv;
    this.path = install.path;
    this.version = version;

    this.registerOutputs({
      architecture,
      binName,
      connection,
      directory,
      mkdir: install.mkdir,
      mv: install.mv,
      path: install.path,
      version,
      url,
    });
  }
}
