import { Config } from '@pulumi/pulumi';
import { ContainerdConfiguration, KubeletService } from '@unmango/pulumi-kubernetes-the-hard-way/remote';
import { Mkdir } from '@unmango/pulumi-kubernetes-the-hard-way/tools';

const config = new Config();
const connection = {
  host: config.require('host'),
  port: config.requireNumber('port'),
  user: config.require('user'),
  password: config.require('password')
};

const dir = new Mkdir('config', {
  connection,
  create: {
    parents: true,
    directory: '/etc/containerd',
  },
});

const containerdConfiguration = new ContainerdConfiguration('simple', {
  connection,
}, { dependsOn: dir });

const kubeletService = new KubeletService('simple', {
  connection,
  after: ['containerd.service'],
  requires: ['containerd.service'],
  configuration: {
    kubeletPath: '/usr/local/bin/kubelet',
    configurationFilePath: '/var/lib/kubelet/kubelet-config.yaml',
    kubeconfigPath: '/var/lib/kubelet/kubeconfig',
    registerNode: true,
    v: 69,
  },
});
