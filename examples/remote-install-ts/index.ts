import * as path from 'node:path';
import { Config } from '@pulumi/pulumi';
import {
  CniPluginsInstall,
  ContainerdInstall,
  CrictlInstall,
  EtcdInstall,
  KubeApiServerInstall,
  KubeControllerManagerInstall,
  KubeProxyInstall,
  KubeSchedulerInstall,
  KubectlInstall,
  KubeletInstall,
  RuncInstall,
} from '@unmango/pulumi-kubernetes-the-hard-way/remote';

const config = new Config();
const host = config.require('host');
const port = config.requireNumber('port');
const user = config.require('user');
const password = config.require('password');
const basePath = config.require('basePath');

const etcd = new EtcdInstall('remote', {
  connection: { host, port, user, password },
  directory: path.join(basePath, 'etcd'),
});

// Still a little bit more work needed here
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

const kubelet = new KubeletInstall('remote', {
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
