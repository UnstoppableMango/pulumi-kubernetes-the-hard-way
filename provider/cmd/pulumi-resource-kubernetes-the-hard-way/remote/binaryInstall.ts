import { Input, Output, Resource, interpolate } from '@pulumi/pulumi';
import { remote } from '@pulumi/command/types/input';
import { Mkdir, Mktemp, Mv, Rm } from '../tools';
import { Download } from './download';

interface BinaryInstallArgs {
  binName: Input<string>;
  connection: Input<remote.ConnectionArgs>;
  directory: Input<string>;
  url: Input<string>;
}

interface BinaryInstallResult {
  mktemp: Mktemp;
  download: Download;
  mkdir: Mkdir;
  mv: Mv;
  path: Output<string>;
  rm: Rm;
}

export function binaryInstall(name: string, args: BinaryInstallArgs, parent: Resource): BinaryInstallResult {
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
    directory: args.directory,
    parents: true,
  }, { parent });

  const binPath = interpolate`${args.directory}/${args.binName}`;

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

  return { mktemp, download, mkdir, mv, path: binPath, rm };
}
