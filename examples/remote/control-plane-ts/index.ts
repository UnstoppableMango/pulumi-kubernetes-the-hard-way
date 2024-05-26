import { Config } from '@pulumi/pulumi';
import { ControlPlaneNode } from '@unmango/pulumi-kubernetes-the-hard-way/remote';

const config = new Config();
const connection = {
  host: config.require('host'),
  port: config.requireNumber('port'),
  user: config.require('user'),
  password: config.require('password')
};

const worker = new ControlPlaneNode('simple', {
  connection,
  architecture: 'amd64',
  apiServerCount: 1,
  caCertificatePath: 'TODO',
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
