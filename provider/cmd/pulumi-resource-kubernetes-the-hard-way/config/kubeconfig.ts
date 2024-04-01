import { Input } from '@pulumi/pulumi';

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

export const KubeconfigType = {
  admin: 'admin',
  kubeControllerManager: 'kube-controller-manager',
  kubeProxy: 'kube-proxy',
  kubeScheduler: 'kube-scheduler',
  worker: 'worker',
} as const;

export type KubeconfigType = (typeof KubeconfigType)[keyof typeof KubeconfigType];

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

export type KubeconfigOptions =
  | KubeconfigAdminOptions
  | KubeconfigKubeControllerManagerOptions
  | KubeconfigKubeProxyOptions
  | KubeconfigKubeSchedulerOptions
  | KubeconfigWorkerOptions;
