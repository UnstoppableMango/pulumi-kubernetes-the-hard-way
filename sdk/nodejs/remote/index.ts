// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DownloadArgs } from "./download";
export type Download = import("./download").Download;
export const Download: typeof import("./download").Download = null as any;
utilities.lazyLoad(exports, ["Download"], () => require("./download"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes-the-hard-way:remote:Download":
                return new Download(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes-the-hard-way", "remote", _module)
