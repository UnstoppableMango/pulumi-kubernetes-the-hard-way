import { Input } from '@pulumi/pulumi';
import { KubeconfigType } from '../types';

export interface Cluster {
  certificateAuthorityData: string;
  server: string;
}

export interface KubeconfigCluster {
  cluster: Cluster;
  name: string;
}

export interface Context {
  cluster: string;
  user: string;
}

export interface KubeconfigContext {
  context: Context;
  name: string;
}

export interface User {
  clientCertificateData: string;
  clientKeyData: string;
}

export interface KubeconfigUser {
  name: string;
  user: User;
}

export interface Kubeconfig {
  clusters: KubeconfigCluster[];
  contexts: KubeconfigContext[];
  users: KubeconfigUser[];
}

export interface KubeconfigAdminOptions {
  type: 'admin';
  publicIp?: Input<string>;
}

export interface KubeconfigKubeControllerManagerOptions {
  type: 'kube-controller-manager';
  publicIp?: Input<string>;
}

export interface KubeconfigKubeProxyOptions {
  type: 'kube-proxy';
  publicIp?: Input<string>;
}

export interface KubeconfigKubeSchedulerOptions {
  type: 'kube-scheduler';
  publicIp?: Input<string>;
}

export interface KubeconfigWorkerOptions {
  type: 'worker';
  name: Input<string>;
  publicIp: Input<string>;
}
