import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Mkdir, Mktemp, Mv, Rm } from '../tools';
import { Download } from './download';

export class KubeApiServerInstall extends schema.KubeApiServerInstall {
  constructor(name: string, args: schema.KubeApiServerInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const architecture = output(args.architecture ?? 'amd64');
    const connection = output(args.connection);
    const installDirectory = output(args.installDirectory ?? '/usr/local/bin');
    const version = output(args.version ?? '1.29.2');
    const url = interpolate`https://storage.googleapis.com/kubernetes-release/release/v${version}/bin/linux/${architecture}/kube-apiserver`

    const tmp = new Mktemp(name, {
      connection,
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
      removeOnDelete: true,
    }, { parent: this });

    const binPath = interpolate`${installDirectory}/kube-apiserver`;

    const mv = new Mv(name, {
      connection,
      source: interpolate`${download.destination}/kube-apiserver`,
      dest: binPath,
    }, { parent: this, dependsOn: [tmp, mkdir] });

    const rm = new Rm(name, {
      connection,
      files: tmpDir,
      force: true,
      recursive: true,
    }, { parent: this });

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
