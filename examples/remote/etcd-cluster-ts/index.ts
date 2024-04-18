import { Config } from '@pulumi/pulumi';
import { EtcdCluster } from '@unmango/pulumi-kubernetes-the-hard-way/remote';
import { ClusterPki } from '@unmango/pulumi-kubernetes-the-hard-way/tls';

const config = new Config();
const host = config.require('host');
const port = config.requireNumber('port');
const user = config.require('user');
const password = config.require('password');

const pki = new ClusterPki('etcd-cluster', {
  clusterName: 'my-cluster',
  nodes: {
    node0: {
      // TODO: Make these two required
      ip: '10.69.0.3',
      role: 'controlplane',
    }
  },
  publicIp: '10.69.0.1',
});

const simple = new EtcdCluster('simple', {
  bundle: {
    caPem: pki.kubernetes.caCertPem,
    certPem: pki.kubernetes.certPem,
    keyPem: pki.kubernetes.privateKeyPem,
  },
  nodes: {
    node0: {
      connection: { host, port, user, password },
      internalIp: '10.69.0.2',
      architecture: 'amd64',
    },
  },
});
