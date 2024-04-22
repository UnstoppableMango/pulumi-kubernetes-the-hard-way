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
      internalIp: '0.0.0.0',
      architecture: 'amd64',
    },
  },
});

const multi = new EtcdCluster('multi', {
  bundle: {
    caPem: pki.kubernetes.caCertPem,
    certPem: pki.kubernetes.certPem,
    keyPem: pki.kubernetes.privateKeyPem,
  },
  nodes: {
    node1: {
      connection: {
        host: config.require('node1-host'),
        port: config.requireNumber('node1-port'),
        user: config.require('node1-user'),
        password: config.require('node1-password'),
      },
      internalIp: '0.0.0.0',
      architecture: 'amd64',
    },
    node2: {
      connection: {
        host: config.require('node2-host'),
        port: config.requireNumber('node2-port'),
        user: config.require('node2-user'),
        password: config.require('node2-password'),
      },
      internalIp: '0.0.0.0',
      architecture: 'amd64',
    },
    node3: {
      connection: {
        host: config.require('node3-host'),
        port: config.requireNumber('node3-port'),
        user: config.require('node3-user'),
        password: config.require('node3-password'),
      },
      internalIp: '0.0.0.0',
      architecture: 'amd64',
    },
  },
});
