// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

import * as pulumiTls from "@pulumi/tls";

import {Certificate} from "./index";
import {File} from "remote";

export class RootCa extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:index:RootCa';

    /**
     * Returns true if the given object is an instance of RootCa.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RootCa {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RootCa.__pulumiType;
    }

    public /*out*/ readonly allowedUses!: pulumi.Output<enums.AllowedUsage[]>;
    public /*out*/ readonly cert!: pulumi.Output<pulumiTls.SelfSignedCert>;
    public /*out*/ readonly certPem!: pulumi.Output<string>;
    public /*out*/ readonly key!: pulumi.Output<pulumiTls.PrivateKey>;
    public /*out*/ readonly privateKeyPem!: pulumi.Output<string>;
    public /*out*/ readonly publicKeyPem!: pulumi.Output<string>;

    /**
     * Create a RootCa resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RootCaArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.validityPeriodHours === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validityPeriodHours'");
            }
            resourceInputs["algorithm"] = args ? args.algorithm : undefined;
            resourceInputs["dnsNames"] = args ? args.dnsNames : undefined;
            resourceInputs["earlyRenewalHours"] = args ? args.earlyRenewalHours : undefined;
            resourceInputs["ecdsaCurve"] = args ? args.ecdsaCurve : undefined;
            resourceInputs["ipAddresses"] = args ? args.ipAddresses : undefined;
            resourceInputs["rsaBits"] = args ? args.rsaBits : undefined;
            resourceInputs["setAuthorityKeyId"] = args ? args.setAuthorityKeyId : undefined;
            resourceInputs["setSubjectKeyId"] = args ? args.setSubjectKeyId : undefined;
            resourceInputs["subject"] = args ? args.subject : undefined;
            resourceInputs["uris"] = args ? args.uris : undefined;
            resourceInputs["validityPeriodHours"] = args ? args.validityPeriodHours : undefined;
            resourceInputs["allowedUses"] = undefined /*out*/;
            resourceInputs["cert"] = undefined /*out*/;
            resourceInputs["certPem"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
            resourceInputs["privateKeyPem"] = undefined /*out*/;
            resourceInputs["publicKeyPem"] = undefined /*out*/;
        } else {
            resourceInputs["allowedUses"] = undefined /*out*/;
            resourceInputs["cert"] = undefined /*out*/;
            resourceInputs["certPem"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
            resourceInputs["privateKeyPem"] = undefined /*out*/;
            resourceInputs["publicKeyPem"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RootCa.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }

    /**
     * Creates a RemoteFile resource representing the copy operation.
     */
    installCert(args: RootCa.InstallCertArgs): pulumi.Output<File> {
        const result: pulumi.Output<RootCa.InstallCertResult> = pulumi.runtime.call("kubernetes-the-hard-way:index:RootCa/installCert", {
            "__self__": this,
            "connection": args.connection,
            "name": args.name,
            "options": args.options,
            "path": args.path,
        }, this);
        return result.result;
    }

    /**
     * Creates a RemoteFile resource representing the copy operation.
     */
    installKey(args: RootCa.InstallKeyArgs): pulumi.Output<File> {
        const result: pulumi.Output<RootCa.InstallKeyResult> = pulumi.runtime.call("kubernetes-the-hard-way:index:RootCa/installKey", {
            "__self__": this,
            "connection": args.connection,
            "name": args.name,
            "options": args.options,
            "path": args.path,
        }, this);
        return result.result;
    }

    /**
     * Creates a Certificate configured for the current authority.
     */
    newCertificate(args: RootCa.NewCertificateArgs): pulumi.Output<Certificate> {
        const result: pulumi.Output<RootCa.NewCertificateResult> = pulumi.runtime.call("kubernetes-the-hard-way:index:RootCa/newCertificate", {
            "__self__": this,
            "algorithm": args.algorithm,
            "allowedUses": args.allowedUses,
            "dnsNames": args.dnsNames,
            "earlyRenewalHours": args.earlyRenewalHours,
            "ecdsaCurve": args.ecdsaCurve,
            "ipAddresses": args.ipAddresses,
            "isCaCertificate": args.isCaCertificate,
            "name": args.name,
            "options": args.options,
            "rsaBits": args.rsaBits,
            "setAuthorityKeyId": args.setAuthorityKeyId,
            "setSubjectKeyId": args.setSubjectKeyId,
            "subject": args.subject,
            "uris": args.uris,
            "validityPeriodHours": args.validityPeriodHours,
        }, this);
        return result.result;
    }
}

/**
 * The set of arguments for constructing a RootCa resource.
 */
export interface RootCaArgs {
    /**
     * Name of the algorithm to use when generating the private key.
     */
    algorithm?: pulumi.Input<enums.Algorithm>;
    /**
     * List of DNS names for which a certificate is being requested.
     */
    dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * TODO
     */
    earlyRenewalHours?: pulumi.Input<number>;
    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
     */
    ecdsaCurve?: pulumi.Input<enums.EcdsaCurve>;
    /**
     * List of IP addresses for which a certificate is being requested.
     */
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
     */
    rsaBits?: pulumi.Input<number>;
    /**
     * Should the generated certificate include an authority key identifier.
     */
    setAuthorityKeyId?: pulumi.Input<boolean>;
    /**
     * Should the generated certificate include a subject key identifier.
     */
    setSubjectKeyId?: pulumi.Input<boolean>;
    subject?: pulumi.Input<pulumiTls.types.input.SelfSignedCertSubject>;
    /**
     * List of URIs for which a certificate is being requested.
     */
    uris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid.
     */
    validityPeriodHours: pulumi.Input<number>;
}

export namespace RootCa {
    /**
     * The set of arguments for the RootCa.installCert method.
     */
    export interface InstallCertArgs {
        /**
         * The connection details.
         */
        connection: pulumi.Input<inputs.ConnectionArgs>;
        name: string;
        options?: inputs.ResourceOptionsArgs;
        /**
         * The path to install to.
         */
        path?: pulumi.Input<string>;
    }

    /**
     * The results of the RootCa.installCert method.
     */
    export interface InstallCertResult {
        readonly result: File;
    }

    /**
     * The set of arguments for the RootCa.installKey method.
     */
    export interface InstallKeyArgs {
        /**
         * The connection details.
         */
        connection: pulumi.Input<inputs.ConnectionArgs>;
        name: string;
        options?: inputs.ResourceOptionsArgs;
        /**
         * The path to install to.
         */
        path?: pulumi.Input<string>;
    }

    /**
     * The results of the RootCa.installKey method.
     */
    export interface InstallKeyResult {
        readonly result: File;
    }

    /**
     * The set of arguments for the RootCa.newCertificate method.
     */
    export interface NewCertificateArgs {
        /**
         * Name of the algorithm to use when generating the private key.
         */
        algorithm: pulumi.Input<enums.Algorithm>;
        allowedUses: pulumi.Input<pulumi.Input<enums.AllowedUsage>[]>;
        /**
         * List of DNS names for which a certificate is being requested.
         */
        dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * TODO
         */
        earlyRenewalHours?: pulumi.Input<number>;
        /**
         * When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
         */
        ecdsaCurve?: pulumi.Input<enums.EcdsaCurve>;
        /**
         * List of IP addresses for which a certificate is being requested.
         */
        ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
        isCaCertificate?: pulumi.Input<boolean>;
        name: string;
        options?: inputs.ResourceOptionsArgs;
        /**
         * When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
         */
        rsaBits?: pulumi.Input<number>;
        /**
         * Should the generated certificate include an authority key identifier.
         */
        setAuthorityKeyId?: pulumi.Input<boolean>;
        /**
         * Should the generated certificate include a subject key identifier.
         */
        setSubjectKeyId?: pulumi.Input<boolean>;
        subject?: pulumi.Input<inputs.CertRequestSubjectArgs>;
        /**
         * List of URIs for which a certificate is being requested.
         */
        uris?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * Number of hours, after initial issuing, that the certificate will remain valid.
         */
        validityPeriodHours: pulumi.Input<number>;
    }

    /**
     * The results of the RootCa.newCertificate method.
     */
    export interface NewCertificateResult {
        /**
         * The issued certificate.
         */
        readonly result: Certificate;
    }

}
