// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.config.inputs.GetKubeVipManifestArgs;
import com.unmango.kubernetesthehardway.config.inputs.GetKubeVipManifestPlainArgs;
import com.unmango.kubernetesthehardway.config.inputs.GetKubeconfigArgs;
import com.unmango.kubernetesthehardway.config.inputs.GetKubeconfigPlainArgs;
import com.unmango.kubernetesthehardway.config.outputs.GetKubeVipManifestResult;
import com.unmango.kubernetesthehardway.config.outputs.GetKubeconfigResult;
import java.util.concurrent.CompletableFuture;

public final class ConfigFunctions {
    /**
     * Gets the static pod manifests for KubeVip.
     * 
     */
    public static Output<GetKubeVipManifestResult> getKubeVipManifest(GetKubeVipManifestArgs args) {
        return getKubeVipManifest(args, InvokeOptions.Empty);
    }
    /**
     * Gets the static pod manifests for KubeVip.
     * 
     */
    public static CompletableFuture<GetKubeVipManifestResult> getKubeVipManifestPlain(GetKubeVipManifestPlainArgs args) {
        return getKubeVipManifestPlain(args, InvokeOptions.Empty);
    }
    /**
     * Gets the static pod manifests for KubeVip.
     * 
     */
    public static Output<GetKubeVipManifestResult> getKubeVipManifest(GetKubeVipManifestArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("kubernetes-the-hard-way:config:getKubeVipManifest", TypeShape.of(GetKubeVipManifestResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Gets the static pod manifests for KubeVip.
     * 
     */
    public static CompletableFuture<GetKubeVipManifestResult> getKubeVipManifestPlain(GetKubeVipManifestPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("kubernetes-the-hard-way:config:getKubeVipManifest", TypeShape.of(GetKubeVipManifestResult.class), args, Utilities.withVersion(options));
    }
    public static Output<GetKubeconfigResult> getKubeconfig(GetKubeconfigArgs args) {
        return getKubeconfig(args, InvokeOptions.Empty);
    }
    public static CompletableFuture<GetKubeconfigResult> getKubeconfigPlain(GetKubeconfigPlainArgs args) {
        return getKubeconfigPlain(args, InvokeOptions.Empty);
    }
    public static Output<GetKubeconfigResult> getKubeconfig(GetKubeconfigArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("kubernetes-the-hard-way:config:getKubeconfig", TypeShape.of(GetKubeconfigResult.class), args, Utilities.withVersion(options));
    }
    public static CompletableFuture<GetKubeconfigResult> getKubeconfigPlain(GetKubeconfigPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("kubernetes-the-hard-way:config:getKubeconfig", TypeShape.of(GetKubeconfigResult.class), args, Utilities.withVersion(options));
    }
}
