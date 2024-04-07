// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import * as pulumiTls from "@pulumi/tls";

import {CertRequestSubject} from "..";

/**
 * A certificate key pair.
 */
export class Certificate extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:tls:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * Name of the algorithm to use when generating the private key.
     */
    public readonly algorithm!: pulumi.Output<enums.tls.Algorithm>;
    /**
     * List of key usages allowed for the issued certificate.
     */
    public readonly allowedUses!: pulumi.Output<enums.tls.AllowedUsage[]>;
    /**
     * Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    public readonly caCertPem!: pulumi.Output<string>;
    /**
     * Name of the algorithm used when generating the private key provided in `ca_private_key_pem`.
     */
    public /*out*/ readonly caKeyAlgorithm!: pulumi.Output<string>;
    /**
     * Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    public readonly caPrivateKeyPem!: pulumi.Output<string>;
    /**
     * The certificate.
     */
    public /*out*/ readonly cert!: pulumi.Output<pulumiTls.LocallySignedCert>;
    /**
     * Certificate data in PEM (RFC 1421).
     */
    public /*out*/ readonly certPem!: pulumi.Output<string>;
    /**
     * Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    public /*out*/ readonly certRequestPem!: pulumi.Output<string>;
    /**
     * The certificate signing request.
     */
    public /*out*/ readonly csr!: pulumi.Output<pulumiTls.CertRequest>;
    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     */
    public readonly dnsNames!: pulumi.Output<string[] | undefined>;
    /**
     * The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     */
    public readonly earlyRenewalHours!: pulumi.Output<number>;
    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
     */
    public readonly ecdsaCurve!: pulumi.Output<string>;
    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     */
    public readonly ipAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     */
    public readonly isCaCertificate!: pulumi.Output<boolean>;
    /**
     * The private key
     */
    public /*out*/ readonly key!: pulumi.Output<pulumiTls.PrivateKey>;
    /**
     * Name of the algorithm used when generating the private key provided in `private_key_pem`.
     */
    public /*out*/ readonly keyAlgorithm!: pulumi.Output<string>;
    /**
     * Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
     */
    public /*out*/ readonly privateKeyOpenssh!: pulumi.Output<string>;
    /**
     * Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    public /*out*/ readonly privateKeyPem!: pulumi.Output<string>;
    /**
     * Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
     */
    public /*out*/ readonly privateKeyPemPkcs8!: pulumi.Output<string>;
    /**
     * The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     */
    public /*out*/ readonly publicKeyFingerprintMd5!: pulumi.Output<string>;
    /**
     * The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     */
    public /*out*/ readonly publicKeyFingerprintSha256!: pulumi.Output<string>;
    /**
     * The public key data in "Authorized Keys".
     */
    public /*out*/ readonly publicKeyOpenssh!: pulumi.Output<string>;
    /**
     * Public key data in PEM (RFC 1421).
     */
    public /*out*/ readonly publicKeyPem!: pulumi.Output<string>;
    /**
     * Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
     */
    public /*out*/ readonly readyForRenewal!: pulumi.Output<boolean>;
    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
     */
    public readonly rsaBits!: pulumi.Output<number>;
    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    public readonly setSubjectKeyId!: pulumi.Output<boolean>;
    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     */
    public readonly subject!: pulumi.Output<CertRequestSubject | undefined>;
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     */
    public readonly uris!: pulumi.Output<string[] | undefined>;
    /**
     * The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     */
    public /*out*/ readonly validityEndTime!: pulumi.Output<string>;
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid for.
     */
    public readonly validityPeriodHours!: pulumi.Output<number>;
    /**
     * The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     */
    public /*out*/ readonly validityStartTime!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.algorithm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'algorithm'");
            }
            if ((!args || args.allowedUses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedUses'");
            }
            if ((!args || args.caCertPem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'caCertPem'");
            }
            if ((!args || args.caPrivateKeyPem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'caPrivateKeyPem'");
            }
            if ((!args || args.validityPeriodHours === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validityPeriodHours'");
            }
            resourceInputs["algorithm"] = args ? args.algorithm : undefined;
            resourceInputs["allowedUses"] = args ? args.allowedUses : undefined;
            resourceInputs["caCertPem"] = args ? args.caCertPem : undefined;
            resourceInputs["caPrivateKeyPem"] = args?.caPrivateKeyPem ? pulumi.secret(args.caPrivateKeyPem) : undefined;
            resourceInputs["dnsNames"] = args ? args.dnsNames : undefined;
            resourceInputs["earlyRenewalHours"] = args ? args.earlyRenewalHours : undefined;
            resourceInputs["ecdsaCurve"] = args ? args.ecdsaCurve : undefined;
            resourceInputs["ipAddresses"] = args ? args.ipAddresses : undefined;
            resourceInputs["isCaCertificate"] = args ? args.isCaCertificate : undefined;
            resourceInputs["rsaBits"] = args ? args.rsaBits : undefined;
            resourceInputs["setSubjectKeyId"] = args ? args.setSubjectKeyId : undefined;
            resourceInputs["subject"] = args ? args.subject : undefined;
            resourceInputs["uris"] = args ? args.uris : undefined;
            resourceInputs["validityPeriodHours"] = args ? args.validityPeriodHours : undefined;
            resourceInputs["caKeyAlgorithm"] = undefined /*out*/;
            resourceInputs["cert"] = undefined /*out*/;
            resourceInputs["certPem"] = undefined /*out*/;
            resourceInputs["certRequestPem"] = undefined /*out*/;
            resourceInputs["csr"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
            resourceInputs["keyAlgorithm"] = undefined /*out*/;
            resourceInputs["privateKeyOpenssh"] = undefined /*out*/;
            resourceInputs["privateKeyPem"] = undefined /*out*/;
            resourceInputs["privateKeyPemPkcs8"] = undefined /*out*/;
            resourceInputs["publicKeyFingerprintMd5"] = undefined /*out*/;
            resourceInputs["publicKeyFingerprintSha256"] = undefined /*out*/;
            resourceInputs["publicKeyOpenssh"] = undefined /*out*/;
            resourceInputs["publicKeyPem"] = undefined /*out*/;
            resourceInputs["readyForRenewal"] = undefined /*out*/;
            resourceInputs["validityEndTime"] = undefined /*out*/;
            resourceInputs["validityStartTime"] = undefined /*out*/;
        } else {
            resourceInputs["algorithm"] = undefined /*out*/;
            resourceInputs["allowedUses"] = undefined /*out*/;
            resourceInputs["caCertPem"] = undefined /*out*/;
            resourceInputs["caKeyAlgorithm"] = undefined /*out*/;
            resourceInputs["caPrivateKeyPem"] = undefined /*out*/;
            resourceInputs["cert"] = undefined /*out*/;
            resourceInputs["certPem"] = undefined /*out*/;
            resourceInputs["certRequestPem"] = undefined /*out*/;
            resourceInputs["csr"] = undefined /*out*/;
            resourceInputs["dnsNames"] = undefined /*out*/;
            resourceInputs["earlyRenewalHours"] = undefined /*out*/;
            resourceInputs["ecdsaCurve"] = undefined /*out*/;
            resourceInputs["ipAddresses"] = undefined /*out*/;
            resourceInputs["isCaCertificate"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
            resourceInputs["keyAlgorithm"] = undefined /*out*/;
            resourceInputs["privateKeyOpenssh"] = undefined /*out*/;
            resourceInputs["privateKeyPem"] = undefined /*out*/;
            resourceInputs["privateKeyPemPkcs8"] = undefined /*out*/;
            resourceInputs["publicKeyFingerprintMd5"] = undefined /*out*/;
            resourceInputs["publicKeyFingerprintSha256"] = undefined /*out*/;
            resourceInputs["publicKeyOpenssh"] = undefined /*out*/;
            resourceInputs["publicKeyPem"] = undefined /*out*/;
            resourceInputs["readyForRenewal"] = undefined /*out*/;
            resourceInputs["rsaBits"] = undefined /*out*/;
            resourceInputs["setSubjectKeyId"] = undefined /*out*/;
            resourceInputs["subject"] = undefined /*out*/;
            resourceInputs["uris"] = undefined /*out*/;
            resourceInputs["validityEndTime"] = undefined /*out*/;
            resourceInputs["validityPeriodHours"] = undefined /*out*/;
            resourceInputs["validityStartTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["caPrivateKeyPem", "privateKeyOpenssh", "privateKeyPem", "privateKeyPemPkcs8"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Certificate.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * Name of the algorithm to use when generating the private key.
     */
    algorithm: pulumi.Input<enums.tls.Algorithm>;
    /**
     * List of key usages allowed for the issued certificate.
     */
    allowedUses: pulumi.Input<pulumi.Input<enums.tls.AllowedUsage>[]>;
    /**
     * Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    caCertPem: pulumi.Input<string>;
    /**
     * Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    caPrivateKeyPem: pulumi.Input<string>;
    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     */
    dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     */
    earlyRenewalHours?: pulumi.Input<number>;
    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
     */
    ecdsaCurve?: pulumi.Input<string>;
    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     */
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     */
    isCaCertificate?: pulumi.Input<boolean>;
    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
     */
    rsaBits?: pulumi.Input<number>;
    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    setSubjectKeyId?: pulumi.Input<boolean>;
    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     */
    subject?: pulumi.Input<CertRequestSubject>;
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     */
    uris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid for.
     */
    validityPeriodHours: pulumi.Input<number>;
}
