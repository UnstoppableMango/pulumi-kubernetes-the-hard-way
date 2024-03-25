// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.inputs.InstallControlPlaneArgs;
import com.unmango.kubernetesthehardway.inputs.InstallControlPlanePlainArgs;
import java.lang.Void;
import java.util.concurrent.CompletableFuture;

public final class KubernetesthehardwayFunctions {
    /**
     * Install PKI onto a controlplane node.
     * 
     */
    public static Output<Void> installControlPlane(InstallControlPlaneArgs args) {
        return installControlPlane(args, InvokeOptions.Empty);
    }
    /**
     * Install PKI onto a controlplane node.
     * 
     */
    public static CompletableFuture<Void> installControlPlanePlain(InstallControlPlanePlainArgs args) {
        return installControlPlanePlain(args, InvokeOptions.Empty);
    }
    /**
     * Install PKI onto a controlplane node.
     * 
     */
    public static Output<Void> installControlPlane(InstallControlPlaneArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("kubernetes-the-hard-way:index:installControlPlane", TypeShape.of(Void.class), args, Utilities.withVersion(options));
    }
    /**
     * Install PKI onto a controlplane node.
     * 
     */
    public static CompletableFuture<Void> installControlPlanePlain(InstallControlPlanePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("kubernetes-the-hard-way:index:installControlPlane", TypeShape.of(Void.class), args, Utilities.withVersion(options));
    }
}
