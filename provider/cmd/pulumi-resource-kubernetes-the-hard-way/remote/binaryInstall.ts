import { Input, Resource, interpolate } from '@pulumi/pulumi';
import { remote } from '@pulumi/command/types/input';
import { Mkdir, Mktemp, Mv, Rm } from '../tools';
import { Download } from './download';

interface BinaryInstallArgs {
  binName: Input<string>;
  connection: Input<remote.ConnectionArgs>;
  installDirectory: Input<string>;
  url: Input<string>;
}

interface InstallOperations {
  mktemp: Mktemp;
  download: Download;
  mkdir: Mkdir;
  mv: Mv;
  rm: Rm;
}

export function binaryInstall(name: string, args: BinaryInstallArgs, parent: Resource): InstallOperations {
  const mktemp = new Mktemp(name, {
    connection: args.connection,
    directory: true,
  }, { parent });

  const tmpDir = mktemp.stdout;

  const download = new Download(name, {
    connection: args.connection,
    destination: tmpDir,
    url: args.url,
  }, { parent, dependsOn: mktemp });

  const mkdir = new Mkdir(name, {
    connection: args.connection,
    directory: args.installDirectory,
    parents: true,
  }, { parent });

  const binPath = interpolate`${args.installDirectory}/${args.binName}`;

  const mv = new Mv(name, {
    connection: args.connection,
    source: interpolate`${download.destination}/${args.binName}`,
    dest: binPath,
  }, { parent, dependsOn: [download, mkdir] });

  const rm = new Rm(name, {
    connection: args.connection,
    files: tmpDir,
    force: true,
    recursive: true,
  }, { parent, dependsOn: mv });

  return { mktemp, download, mkdir, mv, rm };
}
