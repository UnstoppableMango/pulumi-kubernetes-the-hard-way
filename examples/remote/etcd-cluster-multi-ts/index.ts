import { Config } from '@pulumi/pulumi';
import { EtcdCluster } from '@unmango/pulumi-kubernetes-the-hard-way/remote';
import { ClusterPki } from '@unmango/pulumi-kubernetes-the-hard-way/tls';

const config = new Config();

const pki = new ClusterPki('etcd-cluster', {
  clusterName: 'my-cluster',
  nodes: {
    node1: {
      ip: config.require('node1-ip'),
      role: 'controlplane',
    },
    node2: {
      ip: config.require('node2-ip'),
      role: 'controlplane',
    },
    node3: {
      ip: config.require('node3-ip'),
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
    node1: {
      connection: {
        host: config.require('node1-host'),
        port: config.requireNumber('node1-port'),
        user: config.require('node1-user'),
        password: config.require('node1-password'),
      },
      internalIp: config.require('node1-ip'),
      architecture: 'amd64',
    },
    node2: {
      connection: {
        host: config.require('node2-host'),
        port: config.requireNumber('node2-port'),
        user: config.require('node2-user'),
        password: config.require('node2-password'),
      },
      internalIp: config.require('node2-ip'),
      architecture: 'amd64',
    },
    node3: {
      connection: {
        host: config.require('node3-host'),
        port: config.requireNumber('node3-port'),
        user: config.require('node3-user'),
        password: config.require('node3-password'),
      },
      internalIp: config.require('node3-ip'),
      architecture: 'amd64',
    },
  },
});
