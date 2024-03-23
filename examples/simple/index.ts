import * as kthw from "@unmango/pulumi-kubernetes-the-hard-way";

const ca = new kthw.RootCa("yea", {
    algorithm: "RSA",
    validityPeriodHours: 256,
});

export const pem = ca.certPem;
