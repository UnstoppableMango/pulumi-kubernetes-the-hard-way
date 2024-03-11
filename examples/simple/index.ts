import * as kthw from "@unmango/kubernetes-the-hard-way";

const file = new kthw.RemoteFile("page", {
    connection: {
        host: "localhost",
    },
    content: "Some file content",
    path: "/home/pulumi/test.txt",
});

export const command = file.command;
