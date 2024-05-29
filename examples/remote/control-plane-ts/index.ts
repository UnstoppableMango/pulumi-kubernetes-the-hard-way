import * as YAML from 'yaml';
import { Config } from '@pulumi/pulumi';
import { ControlPlaneNode, File } from '@unmango/pulumi-kubernetes-the-hard-way/remote';
import { ClusterPki, EncryptionKey } from '@unmango/pulumi-kubernetes-the-hard-way/tls';

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

const caCertFile = new File('caCertFile', {
  connection,
  content: pki.ca.certPem,
  path: '/tmp/caPem',
});

const caPrivateKeyFile = new File('caPrivateKeyFile', {
  connection,
  content: pki.ca.privateKeyPem,
  path: '/tmp/caKey',
});

const encryptionConfig = new EncryptionKey('simple', {
  bytes: 32,
});

const kubeApiServerCertFile = new File('kubeApiServerCertFile', {
  connection,
  content: pki.kubernetes.certPem,
  path: '/tmp/kubeApiPem',
});

const kubeApiServerPrivateKeyFile = new File('kubeApiServerPrivateKeyFile', {
  connection,
  content: pki.kubernetes.privateKeyPem,
  path: '/tmp/kubeApiKey',
});

const kubeControllerManagerKubeconfigFile = new File('kubeControllerManagerKubeconfigFile', {
  connection,
  content: pki.getKubeconfig({
    options: {
      type: 'kube-controller-manager',
    },
  }).apply(YAML.stringify),
  path: '/tmp/kubeControllerManagerKubeconfig',
});

const kubeSchedulerKubeconfigFile = new File('kubeSchedulerKubeconfigFile', {
  connection,
  content: pki.getKubeconfig({
    options: {
      type: 'kube-scheduler',
    },
  }).apply(YAML.stringify),
  path: '/tmp/kubeSchedulerKubeconfig',
});

const serviceAccountsCertFile = new File('serviceAccountsCertFile', {
  connection,
  content: pki.serviceAccounts.certPem,
  path: '/tmp/saPem',
});

const serviceAccountsPrivateKeyFile = new File('serviceAccountsPrivateKeyFile', {
  connection,
  content: pki.serviceAccounts.privateKeyPem,
  path: '/tmp/saKey',
});

const controlPlane = new ControlPlaneNode('simple', {
  connection,
  architecture: 'amd64',
  apiServerCount: 1,
  caCertificatePath: caCertFile.path,
  caPrivateKeyPath: caPrivateKeyFile.path,
  encryptionConfigYaml: encryptionConfig.config,
  kubeApiServerCertificatePath: kubeApiServerCertFile.path,
  kubeApiServerPrivateKeyPath: kubeApiServerPrivateKeyFile.path,
  kubeControllerManagerKubeconfigPath: kubeControllerManagerKubeconfigFile.path,
  kubeSchedulerConfigYaml: 'TODO',
  kubeSchedulerKubeconfigPath: kubeSchedulerKubeconfigFile.path,
  serviceAccountsCertificatePath: serviceAccountsCertFile.path,
  serviceAccountsPrivateKeyPath: serviceAccountsPrivateKeyFile.path,
});
