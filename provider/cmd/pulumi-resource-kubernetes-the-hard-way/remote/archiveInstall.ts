import { Input, Resource, interpolate } from '@pulumi/pulumi';
import { remote } from '@pulumi/command/types/input';
import { Mkdir, Mktemp, Mv, Rm, Tar } from '../tools';
import { Download } from './download';

type Maps<T extends ReadonlyArray<string>, V> = {
  [K in (T extends ReadonlyArray<infer U> ? U : never)]: V;
}

export interface ArchiveInstallArgs<T extends ReadonlyArray<string>> {
  archiveName: Input<string>;
  binaries: T;
  connection: Input<remote.ConnectionArgs>;
  installDirectory: Input<string>;
  url: Input<string>;
}

export interface ArchiveInstallResult<T extends ReadonlyArray<string>> {
  mvs: Maps<T, Mv>;
  download: Download;
  mkdir: Mkdir;
  mktemp: Mktemp;
  paths: Maps<T, string>;
  rm: Rm;
  tar: Tar;
}

export function archiveInstall<T extends ReadonlyArray<string>>(
  name: string,
  args: ArchiveInstallArgs<T>,
  parent: Resource,
): ArchiveInstallResult<T> {
  const { archiveName, connection, installDirectory, url } = args;

  const mktemp = new Mktemp(name, {
    connection,
    directory: true,
  }, { parent });

  const tmpDir = mktemp.stdout;

  const download = new Download(name, {
    connection,
    destination: tmpDir,
    url,
  }, { parent, dependsOn: mktemp });

  const tar = new Tar(name, {
    connection: args.connection,
    archive: interpolate`${download.destination}/${archiveName}`,
    directory: download.destination,
    gzip: true,
    stripComponents: 1,
  }, { parent, dependsOn: download });

  const mkdir = new Mkdir(name, {
    connection: args.connection,
    directory: installDirectory,
    parents: true,
  }, { parent });

  const mvs = args.binaries.reduce((b, k): Maps<T, Mv> => ({
    ...b,
    [k]: new Mv(`${name}-${k}`, {
      connection,
      source: interpolate`${download.destination}/${k}`,
      dest: interpolate`${installDirectory}/${k}`,
    }, { parent, dependsOn: [tar, mkdir] })
  }), {} as Maps<T, Mv>);

  const paths = args.binaries.reduce((b, k): Maps<T, string> => ({
    ...b, [k]: interpolate`${installDirectory}/${k}`,
  }), {} as Maps<T, string>);

  const rm = new Rm(name, {
    connection,
    files: tmpDir,
    force: true,
    recursive: true,
  }, { parent, dependsOn: Object.values(mvs) });

  return { download, mkdir, mktemp, mvs, paths, rm, tar };
}
