import * as path from 'node:path';
import { Config } from '@pulumi/pulumi';
import { Etcdctl, Mkdir, Mktemp, Tar, Wget } from '@unmango/pulumi-kubernetes-the-hard-way/tools';
import { CniPluginsInstall, ContainerdInstall, CrictlInstall, Download, EtcdConfiguration, EtcdInstall, File, KubeApiServerInstall, KubeControllerManagerInstall, KubeProxyInstall, KubeSchedulerInstall, KubectlInstall, RuncInstall, SystemdService } from '@unmango/pulumi-kubernetes-the-hard-way/remote';

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

const etcd = new EtcdInstall('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'etcd'),
});

const etcdConfig = new EtcdConfiguration('remote', {
  connection: { host, port, user, password },
  caPem: 'pretend theres pem data here',
  certPem: 'pretend theres pem data here',
  keyPem: 'pretend theres pem data here',
  internalIp: '10.240.0.10',
  configurationDirectory: path.join(basePath, 'etc', 'etcd'),
  dataDirectory: path.join(basePath, 'var', 'lib', 'etcd'),
  systemdDirectory: basePath,
  etcdPath: etcd.etcdPath,
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

// const cniPlugins = new CniPluginsInstall('remote', {
//   connection: { host, port, user, password },
//   directory: path.join(basePath, 'cni-plugins'),
// });

const containerd = new ContainerdInstall('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'containerd'),
});

const crictl = new CrictlInstall('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'crictl'),
});

const apiServer = new KubeApiServerInstall('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'kube-apiserver'),
});

const controllerManager = new KubeControllerManagerInstall('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'kube-controller-manager'),
});

const kubectl = new KubectlInstall('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'kubectl'),
});

const kubelet = new KubeApiServerInstall('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'kubelet'),
});

const kubeProxy = new KubeProxyInstall('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'kube-proxy'),
});

const scheduler = new KubeSchedulerInstall('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'kube-scheduler'),
});

const runc = new RuncInstall('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'runc'),
});

export const fileStderr = file.stderr;
export const fileStdout = file.stdout;

export const wgetStderr = wget.stderr;
export const wgetStdout = wget.stdout;
export const wgetCommand = wget.command;

// export const tarStderr = tar.stderr;
// export const tarStdout = tar.stdout;

export const mktemp = tmp.stdout;
