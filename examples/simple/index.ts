import * as kthw from '@unmango/pulumi-kubernetes-the-hard-way';

const ca = new kthw.RootCa('simple', {
    algorithm: 'RSA',
    validityPeriodHours: 256,
});

const cert = new kthw.Certificate('simple2', {
    algorithm: 'RSA',
    allowedUses: [
        kthw.AllowedUsage.Cert_signing,
    ],
    validityPeriodHours: 256,
    caCertPem: ca.certPem,
    caPrivateKeyPem: ca.privateKeyPem,
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
