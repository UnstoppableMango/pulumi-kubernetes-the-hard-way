import {
  AllowedUsage,
  Certificate,
  ClusterPki,
  RootCa,
} from '@unmango/pulumi-kubernetes-the-hard-way/tls';

const ca = new RootCa('tls', {
  algorithm: 'RSA',
  validityPeriodHours: 256,
});

const cert = new Certificate('tls', {
  algorithm: 'RSA',
  allowedUses: [
      AllowedUsage.CertSigning,
  ],
  validityPeriodHours: 256,
  caCertPem: ca.certPem,
  caPrivateKeyPem: ca.privateKeyPem,
});

const pki = new ClusterPki('tls', {
  clusterName: 'simple-cluster',
  nodes: {
    myNode: {
      ip: '10.0.69.3',
      role: 'controlplane',
    },
  },
  publicIp: '10.0.69.2',
});
