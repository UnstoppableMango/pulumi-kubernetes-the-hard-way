// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

import * as pulumiTls from "@pulumi/tls";

/**
 * Creates a Certificate configured for the cluster admin.
 */
export function getAdminCertificate(args: GetAdminCertificateArgs, opts?: pulumi.InvokeOptions): Promise<void> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("kubernetes-the-hard-way:index:getAdminCertificate", {
        "algorithm": args.algorithm,
        "allowedUses": args.allowedUses,
        "dnsNames": args.dnsNames,
        "earlyRenewalHours": args.earlyRenewalHours,
        "ecdsaCurve": args.ecdsaCurve,
        "ipAddresses": args.ipAddresses,
        "isCaCertificate": args.isCaCertificate,
        "rsaBits": args.rsaBits,
        "setAuthorityKeyId": args.setAuthorityKeyId,
        "setSubjectKeyId": args.setSubjectKeyId,
        "subject": args.subject,
        "uris": args.uris,
        "validityPeriodHours": args.validityPeriodHours,
    }, opts);
}

export interface GetAdminCertificateArgs {
    /**
     * Name of the algorithm to use when generating the private key.
     */
    algorithm?: string;
    allowedUses: string[];
    /**
     * List of DNS names for which a certificate is being requested.
     */
    dnsNames?: string[];
    /**
     * TODO
     */
    earlyRenewalHours?: number;
    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
     */
    ecdsaCurve?: string;
    /**
     * List of IP addresses for which a certificate is being requested.
     */
    ipAddresses?: string[];
    isCaCertificate?: boolean;
    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
     */
    rsaBits?: number;
    /**
     * Should the generated certificate include an authority key identifier.
     */
    setAuthorityKeyId?: boolean;
    /**
     * Should the generated certificate include a subject key identifier.
     */
    setSubjectKeyId?: boolean;
    subject?: pulumiTls.types.input.CertRequestSubject;
    /**
     * List of URIs for which a certificate is being requested.
     */
    uris?: string[];
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid.
     */
    validityPeriodHours: number;
}
