import * as YAML from 'yaml';
import { Config } from '@pulumi/pulumi';
import { ControlPlaneNode, File } from '@unmango/pulumi-kubernetes-the-hard-way/remote';
import { ClusterPki } from '@unmango/pulumi-kubernetes-the-hard-way/tls';

const config = new Config();
const connection = {
  host: config.require('host'),
  port: config.requireNumber('port'),
  user: config.require('user'),
  password: config.require('password')
};
const ip = config.require('ip');

// Create the ClusterPki
// - getKubeconfig(options: { type: 'kube-scheduler' })
// Create the EncryptionConfig
// output(object).apply(YAML.stringify)

const pki = new ClusterPki('simple', {
  clusterName: 'simple',
  publicIp: ip,
  nodes: {
    'simple': {
      ip,
      role: 'controlplane'
    },
  },
});

const file = new File('simple', {
  connection,
  content: pki.ca.certPem,
  path: '/tmp/caPem',
});

const controlPlane = new ControlPlaneNode('simple', {
  connection,
  architecture: 'amd64',
  apiServerCount: 1,
  caCertificatePath: file.path,
  caPrivateKeyPath: 'TODO',
  encryptionConfigYaml: 'TODO',
  kubeApiServerCertificatePath: 'TODO',
  kubeApiServerPrivateKeyPath: 'TODO',
  kubeControllerManagerKubeconfigPath: 'TODO',
  kubeSchedulerConfigYaml: 'TODO',
  kubeSchedulerKubeconfigPath: 'TODO',
  serviceAccountsCertificatePath: 'TODO',
  serviceAccountsPrivateKeyPath: 'TODO',
});

controlPlane.kubeApiServerInstallDirectory;
