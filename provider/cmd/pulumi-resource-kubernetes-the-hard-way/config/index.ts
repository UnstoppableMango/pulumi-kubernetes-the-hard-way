import { Input } from '@pulumi/pulumi';
import { KubeconfigType } from '../types';

export interface ClusterCluster {
  certificateAuthorityData: string;
  server: string;
}

export interface Cluster {
  cluster: ClusterCluster;
  name: string;
}

export interface ContextContext {
  cluster: string;
  user: string;
}

export interface Context {
  context: ContextContext;
  name: string;
}

export interface UserUser {
  clientCertificateData: string;
  clientKeyData: string;
}

export interface User {
  name: string;
  user: UserUser;
}

export interface Kubeconfig {
  clusters: Cluster[];
  contexts: Context[];
  users: User[];
}

export interface KubeconfigAdminOptions {
  type: Input<'admin'>;
  publicIp?: Input<string>;
}

export interface KubeconfigKubeControllerManagerOptions {
  type: Input<'kube-controller-manager'>;
  publicIp?: Input<string>;
}

export interface KubeconfigKubeProxyOptions {
  type: Input<'kube-proxy'>;
  publicIp?: Input<string>;
}

export interface KubeconfigKubeSchedulerOptions {
  type: Input<'kube-scheduler'>;
  publicIp?: Input<string>;
}

export interface KubeconfigWorkerOptions {
  type: Input<'worker'>;
  name: Input<string>;
  publicIp: Input<string>;
}
