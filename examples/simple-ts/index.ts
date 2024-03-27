import * as kthw from '@unmango/pulumi-kubernetes-the-hard-way';

const ca = new kthw.RootCa('simple', {
    algorithm: 'RSA',
    validityPeriodHours: 256,
});

const cert = new kthw.Certificate('simple', {
    algorithm: 'RSA',
    allowedUses: [
        kthw.AllowedUsage.CertSigning,
    ],
    validityPeriodHours: 256,
    caCertPem: ca.certPem,
    caPrivateKeyPem: ca.privateKeyPem,
});

// TODO: Fix
// const cert2 = ca.newCertificate({
//   name: 'nested',
//   allowedUses: ['cert_signing'],
//   algorithm: 'RSA',
//   validityPeriodHours: 256,
// });

const pki = new kthw.ClusterPki('simple', {
  clusterName: 'simple-cluster',
  nodes: {
    myNode: {
      ip: '10.0.69.3',
      role: 'controlplane',
    },
  },
  publicIp: '10.0.69.2',
});

export const caAllowedUses = ca.allowedUses;
export const caCertPem = ca.certPem;
export const caKeyPem = ca.publicKeyPem;

export const certPem = cert.certPem;
export const keyPem = cert.privateKeyPem;

export const caCert = ca.cert;
export const caKey = ca.key;
export const certCert = cert.cert;
export const certKey = cert.key;
