import * as path from 'node:path';
import { Config } from '@pulumi/pulumi';
import { Mkdir, Tar, Wget } from '@unmango/pulumi-kubernetes-the-hard-way/tools';
import { Download, EtcdInstall, File } from '@unmango/pulumi-kubernetes-the-hard-way/remote';

const config = new Config();
const host = config.require('host');
const port = config.requireNumber('port');
const user = config.require('user');
const password = config.require('password');
const basePath = config.require('basePath');

const file = new File('remote', {
  connection: { host, port, user, password },
  content: config.require('content'),
  path: path.join(basePath, 'remote-file'),
});

const wget = new Wget('remote', {
  connection: { host, port, user, password },
  url: 'https://www.example.com',
  directoryPrefix: basePath,
  // The container image seems to have an old version
  // These two options aren't supported
  httpsOnly: false,
  timestamping: false,
});

const mkdir = new Mkdir('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'test-dir', 'subdir'),
  parents: true,
});

const download = new Download('remote', {
  connection: { host, port, user, password },
  destination: path.join(basePath, 'download'),
  url: 'https://www.example.com',
  removeOnDelete: true,
});

// This has permissions issues at the moment for some reason
// const tar = new Tar('remote', {
//   connection: { host, port, user, password },
//   archive: path.join(basePath, 'text-file.tar.gz'),
//   extract: true,
//   gzip: true,
// });

const etcd = new EtcdInstall('remote', {
  connection: { host, port, user, password },
  caPem: 'pretend theres pem data here',
  certPem: 'pretend theres pem data here',
  keyPem: 'pretend theres pem data here',
  internalIp: '10.240.0.10',
  installDirectory: path.join(basePath, 'etcd'),
  configurationDirectory: path.join(basePath, 'etc', 'etcd'),
  dataDirectory: path.join(basePath, 'var', 'lib', 'etcd'),
  systemdDirectory: basePath,
});

export const fileStderr = file.stderr;
export const fileStdout = file.stdout;

export const wgetStderr = wget.stderr;
export const wgetStdout = wget.stdout;
export const wgetCommand = wget.command;

// export const tarStderr = tar.stderr;
// export const tarStdout = tar.stdout;
