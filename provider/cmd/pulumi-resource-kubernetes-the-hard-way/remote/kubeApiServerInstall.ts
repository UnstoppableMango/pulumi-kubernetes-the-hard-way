import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { binaryInstall } from './binaryInstall';

export class KubeApiServerInstall extends schema.KubeApiServerInstall {
  constructor(name: string, args: schema.KubeApiServerInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const architecture = output(args.architecture ?? 'amd64');
    const binName = 'kube-apiserver';
    const connection = output(args.connection);
    const installDirectory = output(args.installDirectory ?? '/usr/local/bin');
    const version = output(args.version ?? '1.29.2');
    const url = interpolate`https://storage.googleapis.com/kubernetes-release/release/v${version}/bin/linux/${architecture}/${binName}`;

    binaryInstall(name, {
      binName,
      connection,
      installDirectory,
      url,
    }, this);

    this.architecture = architecture;
    this.connection = connection;
    this.installDirectory = installDirectory;
    this.version = version;

    this.registerOutputs({
      architecture,
      connection,
      installDirectory,
      version,
    });
  }
}
