import { ClusterPki } from '@unmango/pulumi-kubernetes-the-hard-way/tls';

const simple = new ClusterPki('simple', {
  clusterName: 'my-cluster',
  publicIp: '10.0.69.2',
  nodes: {
    node0: {
      ip: '10.0.69.3',
      role: 'worker',
    },
    node1: {
      ip: '10.0.69.3',
      role: 'worker',
    },
  },
});

export const node0Pem = simple.kubelet.apply(x => x["node0"].certPem);
