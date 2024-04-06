import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { binaryInstall } from './binaryInstall';

export class KubeletInstall extends schema.KubeletInstall {
  constructor(name: string, args: schema.KubeletInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const architecture = output(args.architecture ?? 'amd64');
    const binName = 'kube-apiserver';
    const connection = output(args.connection);
    const directory = output(args.directory ?? '/usr/local/bin');
    const version = output(args.version ?? '1.29.2');
    const url = interpolate`https://storage.googleapis.com/kubernetes-release/release/v${version}/bin/linux/${architecture}/${binName}`;

    binaryInstall(name, {
      binName,
      connection,
      directory,
      url,
    }, this);

    this.architecture = architecture;
    this.connection = connection;
    this.directory = directory;
    this.version = version;

    this.registerOutputs({
      architecture,
      connection,
      directory,
      version,
    });
  }
}
