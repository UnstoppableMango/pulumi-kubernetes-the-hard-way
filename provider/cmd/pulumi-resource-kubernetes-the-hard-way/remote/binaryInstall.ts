import { Input, Output, Resource, interpolate, output } from '@pulumi/pulumi';
import { remote } from '@pulumi/command/types/input';
import { Mkdir, Mktemp, Mv, Rm } from '../tools';
import { Download } from './download';

interface BinaryInstallArgs {
  binName: Input<string>;
  connection: Input<remote.ConnectionArgs>;
  directory: Input<string>;
  url: Input<string>;
  finalBin?: Input<string>;
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
    create: { directory: true },
  }, { parent });

  const tmpDir = mktemp.stdout;
  const binName = output(args.binName);

  const download = new Download(name, {
    connection: args.connection,
    destination: tmpDir,
    url: args.url,
  }, { parent, dependsOn: mktemp });

  const mkdir = new Mkdir(name, {
    connection: args.connection,
    create: {
      directory: args.directory,
      parents: true,
    },
  }, { parent });

  const finalBin = output(args.finalBin ?? binName);
  const binPath = interpolate`${args.directory}/${finalBin}`;

  const mv = new Mv(name, {
    connection: args.connection,
    create: {
      source: [interpolate`${download.destination}/${binName}`],
      dest: binPath,
    },
    delete: interpolate`rm -rf ${binPath}`,
  }, { parent, dependsOn: [download, mkdir] });

  const rm = new Rm(name, {
    connection: args.connection,
    create: {
      files: [tmpDir],
      force: true,
      recursive: true,
    },
  }, { parent, dependsOn: mv });

  return { mktemp, download, mkdir, mv, path: binPath, rm };
}
