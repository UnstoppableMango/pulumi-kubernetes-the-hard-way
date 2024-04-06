import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Download } from './download';
import { Mkdir, Mktemp, Mv, Rm } from '../tools';

export class KubeControllerManagerInstall extends schema.KubeControllerManagerInstall {
  constructor(name: string, args: schema.KubeControllerManagerInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const architecture = output(args.architecture ?? 'amd64');
    const binName = 'kube-controller-manager';
    const connection = output(args.connection);
    const installDirectory = output(args.installDirectory ?? '/usr/local/bin');
    const version = output(args.version ?? '1.29.2');
    const url = interpolate`https://storage.googleapis.com/kubernetes-release/release/v${version}/bin/linux/${architecture}/${binName}`;

    const tmp = new Mktemp(name, {
      connection,
      directory: true,
    }, { parent: this });

    const tmpDir = tmp.stdout;

    const download = new Download(name, {
      connection,
      destination: tmpDir,
      url,
    }, { parent: this, dependsOn: tmp });

    const mkdir = new Mkdir(name, {
      connection,
      directory: installDirectory,
      parents: true,
    }, { parent: this });

    const binPath = interpolate`${installDirectory}/${binName}`;

    const mv = new Mv(name, {
      connection,
      source: interpolate`${download.destination}/${binName}`,
      dest: binPath,
    }, { parent: this, dependsOn: [tmp, mkdir, download] });

    const rm = new Rm(name, {
      connection,
      files: tmpDir,
      force: true,
      recursive: true,
    }, { parent: this, dependsOn: mv });

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
