import * as path from 'node:path';
import { Config } from '@pulumi/pulumi';
import { Mkdir, Mktemp, Tar, Wget } from '@unmango/pulumi-kubernetes-the-hard-way/tools';
import { Download, EtcdConfiguration, File, SystemdService } from '@unmango/pulumi-kubernetes-the-hard-way/remote';

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
});

const mkdir = new Mkdir('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'test-dir', 'subdir'),
  parents: true,
});

const tmp = new Mktemp('remote', {
  connection: { host, port, user, password },
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

const etcdConfig = new EtcdConfiguration('remote', {
  connection: { host, port, user, password },
  caPem: 'pretend theres pem data here',
  certPem: 'pretend theres pem data here',
  keyPem: 'pretend theres pem data here',
  internalIp: '10.240.0.10',
  configurationDirectory: path.join(basePath, 'etc', 'etcd'),
  dataDirectory: path.join(basePath, 'var', 'lib', 'etcd'),
  // systemdDirectory: basePath,
  etcdPath: '/some/path/that/probably/should/exist/etcd',
});

// const etcdctl = new Etcdctl('remote', {
//   connection: { host, port, user, password },
//   binaryPath: etcd.etcdctlPath,
//   commands: ['member', 'list'],
// }, { dependsOn: etcd });

const systemdService = new SystemdService('remote-test', {
  connection: { host, port, user, password },
  directory: basePath,
  service: {
    execStart: 'test',
  },
});

export const fileStderr = file.stderr;
export const fileStdout = file.stdout;

export const wgetStderr = wget.stderr;
export const wgetStdout = wget.stdout;
export const wgetCommand = wget.command;

// export const tarStderr = tar.stderr;
// export const tarStdout = tar.stdout;

export const mktemp = tmp.stdout;
