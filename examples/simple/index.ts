import * as kthw from "@unmango/pulumi-kubernetes-the-hard-way";

const ca = new kthw.RootCa("yea", {
    algorithm: "RSA",
    validityPeriodHours: 69,
});

const result = ca.createCertificate({
    algorithm: "RSA",
    validityPeriodHours: 69,
    allowedUses: [ // TODO
        "cert_signing",
        "client_auth",
    ],
});

const caFile = ca.installOn({
    connection: {
        host: '',
    },
    path: '',
});

const file = new kthw.RemoteFile("page", {
    connection: {
        host: "localhost",
    },
    content: "Some file content",
    path: "/home/pulumi/test.txt",
});

export const command = file.command;
