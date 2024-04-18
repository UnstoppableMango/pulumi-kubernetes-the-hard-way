import { Certificate, RootCa } from '@unmango/pulumi-kubernetes-the-hard-way/tls';

const ca = new RootCa('certificate', {
  validityPeriodHours: 420,
});

const simple = new Certificate('simple', {
  algorithm: 'RSA',
  allowedUses: ['digital_signature'],
  caCertPem: ca.certPem,
  caPrivateKeyPem: ca.privateKeyPem,
  validityPeriodHours: 69,
});
