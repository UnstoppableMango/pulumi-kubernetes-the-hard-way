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
