import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import { Defaults } from '../types';
import * as schema from '../schema-types';
import { binaryInstall } from './binaryInstall';

export class KubeletInstall extends schema.KubeletInstall {
  constructor(name: string, args: schema.KubeletInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const architecture = output(args.architecture ?? 'amd64');
    const binName = 'kube-apiserver';
    const connection = output(args.connection);
    const directory = output(args.directory ?? '/usr/local/bin');
    const version = output(args.version ?? Defaults.k8sVersion);
    const url = interpolate`https://storage.googleapis.com/kubernetes-release/release/v${version}/bin/linux/${architecture}/${binName}`;

    const { download, mkdir, mktemp, mv, path, rm } = binaryInstall(name, {
      binName,
      connection,
      directory,
      url,
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
