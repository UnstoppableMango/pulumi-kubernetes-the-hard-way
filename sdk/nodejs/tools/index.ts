// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { EtcdctlArgs } from "./etcdctl";
export type Etcdctl = import("./etcdctl").Etcdctl;
export const Etcdctl: typeof import("./etcdctl").Etcdctl = null as any;
utilities.lazyLoad(exports, ["Etcdctl"], () => require("./etcdctl"));

export { MkdirArgs } from "./mkdir";
export type Mkdir = import("./mkdir").Mkdir;
export const Mkdir: typeof import("./mkdir").Mkdir = null as any;
utilities.lazyLoad(exports, ["Mkdir"], () => require("./mkdir"));

export { MktempArgs } from "./mktemp";
export type Mktemp = import("./mktemp").Mktemp;
export const Mktemp: typeof import("./mktemp").Mktemp = null as any;
utilities.lazyLoad(exports, ["Mktemp"], () => require("./mktemp"));

export { MvArgs } from "./mv";
export type Mv = import("./mv").Mv;
export const Mv: typeof import("./mv").Mv = null as any;
utilities.lazyLoad(exports, ["Mv"], () => require("./mv"));

export { RmArgs } from "./rm";
export type Rm = import("./rm").Rm;
export const Rm: typeof import("./rm").Rm = null as any;
utilities.lazyLoad(exports, ["Rm"], () => require("./rm"));

export { SystemctlArgs } from "./systemctl";
export type Systemctl = import("./systemctl").Systemctl;
export const Systemctl: typeof import("./systemctl").Systemctl = null as any;
utilities.lazyLoad(exports, ["Systemctl"], () => require("./systemctl"));

export { TarArgs } from "./tar";
export type Tar = import("./tar").Tar;
export const Tar: typeof import("./tar").Tar = null as any;
utilities.lazyLoad(exports, ["Tar"], () => require("./tar"));

export { WgetArgs } from "./wget";
export type Wget = import("./wget").Wget;
export const Wget: typeof import("./wget").Wget = null as any;
utilities.lazyLoad(exports, ["Wget"], () => require("./wget"));


// Export enums:
export * from "../types/enums/tools";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes-the-hard-way:tools:Etcdctl":
                return new Etcdctl(name, <any>undefined, { urn })
            case "kubernetes-the-hard-way:tools:Mkdir":
                return new Mkdir(name, <any>undefined, { urn })
            case "kubernetes-the-hard-way:tools:Mktemp":
                return new Mktemp(name, <any>undefined, { urn })
            case "kubernetes-the-hard-way:tools:Mv":
                return new Mv(name, <any>undefined, { urn })
            case "kubernetes-the-hard-way:tools:Rm":
                return new Rm(name, <any>undefined, { urn })
            case "kubernetes-the-hard-way:tools:Systemctl":
                return new Systemctl(name, <any>undefined, { urn })
            case "kubernetes-the-hard-way:tools:Tar":
                return new Tar(name, <any>undefined, { urn })
            case "kubernetes-the-hard-way:tools:Wget":
                return new Wget(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes-the-hard-way", "tools", _module)
