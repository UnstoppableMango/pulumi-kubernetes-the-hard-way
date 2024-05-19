// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetKubeVipManifestArgs, GetKubeVipManifestResult, GetKubeVipManifestOutputArgs } from "./getKubeVipManifest";
export const getKubeVipManifest: typeof import("./getKubeVipManifest").getKubeVipManifest = null as any;
export const getKubeVipManifestOutput: typeof import("./getKubeVipManifest").getKubeVipManifestOutput = null as any;
utilities.lazyLoad(exports, ["getKubeVipManifest","getKubeVipManifestOutput"], () => require("./getKubeVipManifest"));

export { GetKubeconfigArgs, GetKubeconfigResult, GetKubeconfigOutputArgs } from "./getKubeconfig";
export const getKubeconfig: typeof import("./getKubeconfig").getKubeconfig = null as any;
export const getKubeconfigOutput: typeof import("./getKubeconfig").getKubeconfigOutput = null as any;
utilities.lazyLoad(exports, ["getKubeconfig","getKubeconfigOutput"], () => require("./getKubeconfig"));

export { GetKubeletConfigurationArgs, GetKubeletConfigurationResult, GetKubeletConfigurationOutputArgs } from "./getKubeletConfiguration";
export const getKubeletConfiguration: typeof import("./getKubeletConfiguration").getKubeletConfiguration = null as any;
export const getKubeletConfigurationOutput: typeof import("./getKubeletConfiguration").getKubeletConfigurationOutput = null as any;
utilities.lazyLoad(exports, ["getKubeletConfiguration","getKubeletConfigurationOutput"], () => require("./getKubeletConfiguration"));

export { KubeVipManifestArgs } from "./kubeVipManifest";
export type KubeVipManifest = import("./kubeVipManifest").KubeVipManifest;
export const KubeVipManifest: typeof import("./kubeVipManifest").KubeVipManifest = null as any;
utilities.lazyLoad(exports, ["KubeVipManifest"], () => require("./kubeVipManifest"));

export { KubeconfigArgs } from "./kubeconfig";
export type Kubeconfig = import("./kubeconfig").Kubeconfig;
export const Kubeconfig: typeof import("./kubeconfig").Kubeconfig = null as any;
utilities.lazyLoad(exports, ["Kubeconfig"], () => require("./kubeconfig"));

export { KubeletConfigurationArgs } from "./kubeletConfiguration";
export type KubeletConfiguration = import("./kubeletConfiguration").KubeletConfiguration;
export const KubeletConfiguration: typeof import("./kubeletConfiguration").KubeletConfiguration = null as any;
utilities.lazyLoad(exports, ["KubeletConfiguration"], () => require("./kubeletConfiguration"));


// Export enums:
export * from "../types/enums/config";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes-the-hard-way:config:KubeVipManifest":
                return new KubeVipManifest(name, <any>undefined, { urn })
            case "kubernetes-the-hard-way:config:Kubeconfig":
                return new Kubeconfig(name, <any>undefined, { urn })
            case "kubernetes-the-hard-way:config:KubeletConfiguration":
                return new KubeletConfiguration(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes-the-hard-way", "config", _module)
