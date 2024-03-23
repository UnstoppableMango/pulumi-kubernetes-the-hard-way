import * as kthw from '@unmango/pulumi-kubernetes-the-hard-way';

const ca = new kthw.RootCa('simple', {
    algorithm: 'RSA',
    validityPeriodHours: 256,
});

export const allowedUses = ca.allowedUses;
export const cert = ca.cert;
export const certPem = ca.certPem;
export const key = ca.key;
export const keyPem = ca.keyPem;
