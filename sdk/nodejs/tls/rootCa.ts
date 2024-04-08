// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import * as pulumiTls from "@pulumi/tls";

/**
 * Root certificate authority for a cluster.
 */
export class RootCa extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:tls:RootCa';

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

    /**
     * Name of the algorithm to use when generating the private key.
     */
    public readonly algorithm!: pulumi.Output<enums.tls.Algorithm>;
    /**
     * List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
     */
    public /*out*/ readonly allowedUses!: pulumi.Output<string[]>;
    /**
     * The certificate authority certificate.
     */
    public /*out*/ readonly cert!: pulumi.Output<pulumiTls.SelfSignedCert>;
    /**
     * Certificate data in PEM (RFC 1421).
     */
    public /*out*/ readonly certPem!: pulumi.Output<string>;
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
    public /*out*/ readonly isCaCertificate!: pulumi.Output<boolean>;
    /**
     * The certificate authority key.
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
     * Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    public readonly setAuthorityKeyId!: pulumi.Output<boolean>;
    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    public readonly setSubjectKeyId!: pulumi.Output<boolean>;
    /**
     * TODO
     */
    public readonly subject!: pulumi.Output<pulumiTls.types.output.SelfSignedCertSubject | undefined>;
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
            resourceInputs["validityEndTime"] = undefined /*out*/;
            resourceInputs["validityStartTime"] = undefined /*out*/;
        } else {
            resourceInputs["algorithm"] = undefined /*out*/;
            resourceInputs["allowedUses"] = undefined /*out*/;
            resourceInputs["cert"] = undefined /*out*/;
            resourceInputs["certPem"] = undefined /*out*/;
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
            resourceInputs["setAuthorityKeyId"] = undefined /*out*/;
            resourceInputs["setSubjectKeyId"] = undefined /*out*/;
            resourceInputs["subject"] = undefined /*out*/;
            resourceInputs["uris"] = undefined /*out*/;
            resourceInputs["validityEndTime"] = undefined /*out*/;
            resourceInputs["validityPeriodHours"] = undefined /*out*/;
            resourceInputs["validityStartTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["privateKeyOpenssh", "privateKeyPem", "privateKeyPemPkcs8"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(RootCa.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a RootCa resource.
 */
export interface RootCaArgs {
    /**
     * Name of the algorithm to use when generating the private key.
     */
    algorithm?: pulumi.Input<enums.tls.Algorithm>;
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
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
     */
    rsaBits?: pulumi.Input<number>;
    /**
     * Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    setAuthorityKeyId?: pulumi.Input<boolean>;
    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    setSubjectKeyId?: pulumi.Input<boolean>;
    /**
     * TODO
     */
    subject?: pulumi.Input<pulumiTls.types.input.SelfSignedCertSubject>;
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     */
    uris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid for.
     */
    validityPeriodHours: pulumi.Input<number>;
}
