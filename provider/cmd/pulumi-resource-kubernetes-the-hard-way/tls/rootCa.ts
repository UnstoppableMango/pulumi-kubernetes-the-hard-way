import { ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { SelfSignedCert } from '@pulumi/tls';
import * as schema from '../schema-types';
import { AllowedUsage } from '../types';
import { toAllowedUsage } from '../util';
import { KeyPair } from './keypair';
import { Certificate } from './certificate';

export type NewCertificateInputs = Omit<schema.CertificateArgs, 'caCertPem' | 'caPrivateKeyPem'> & {
  name: string;
  options?: ComponentResourceOptions;
};

export interface NewCertificateOutputs {
  result: Certificate;
}

export class RootCa extends schema.RootCa {
  public readonly allowedUses!: Output<AllowedUsage[]>;
  public readonly cert!: SelfSignedCert;
  public readonly certPem!: Output<string>;

  constructor(name: string, args: schema.RootCaArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const algorithm = output(args.algorithm ?? 'RSA');
    const dnsNames = output(args.dnsNames ?? []);
    const earlyRenewalHours = output(args.earlyRenewalHours);
    const ecdsaCurve = output(args.ecdsaCurve ?? 'P256');
    const ipAddresses = output(args.ipAddresses ?? []);
    const rsaBits = output(args.rsaBits);
    const setAuthorityKeyId = output(args.setAuthorityKeyId);
    const subject = output(args.subject);
    const validityPeriodHours = output(args.validityPeriodHours);

    const key = KeyPair.key(name, {
      algorithm,
      validityPeriodHours,
      ecdsaCurve,
      rsaBits: args.rsaBits,
    }, this);

    const cert = new SelfSignedCert(name, {
      isCaCertificate: true,
      allowedUses: [
        AllowedUsage.CertSigning,
        AllowedUsage.KeyEncipherment,
        AllowedUsage.ServerAuth,
        AllowedUsage.ClientAuth,
      ],
      privateKeyPem: key.privateKeyPem,
      validityPeriodHours: validityPeriodHours,
      subject: subject.apply(s => ({
        commonName: s?.commonName ?? 'Kubernetes',
        country: s?.country,
        locality: s?.locality,
        organization: s?.organization,
        organizationalUnit: s?.organizationalUnit,
        postalCode: s?.postalCode,
        province: s?.province,
        serialNumber: s?.serialNumber,
        streetAddresses: s?.streetAddresses,
      })),
    }, { parent: this });

    this.algorithm = algorithm;
    this.allowedUses = cert.allowedUses.apply(toAllowedUsage);
    this.cert = cert;
    this.certPem = cert.certPem;
    this.ecdsaCurve = ecdsaCurve;
    this.key = key;
    this.keyAlgorithm = key.algorithm; // TODO
    this.validityPeriodHours = validityPeriodHours;
    this.dnsNames = dnsNames;

    this.registerOutputs({
      allowedUses: this.allowedUses,
      cert, certPem: cert.certPem, key,
      privateKeyPem: key.privateKeyPem,
      publicKeyPem: key.publicKeyPem,
    });
  }

  public async newCertificate(args: NewCertificateInputs): Promise<NewCertificateOutputs> {
    const cert = new Certificate(args.name, {
      algorithm: args.algorithm,
      allowedUses: args.allowedUses,
      validityPeriodHours: args.validityPeriodHours,
      dnsNames: args.dnsNames,
      ecdsaCurve: args.ecdsaCurve,
      ipAddresses: args.ipAddresses,
      isCaCertificate: args.isCaCertificate,
      rsaBits: args.rsaBits,
      subject: args.subject,
      uris: args.uris,
      caCertPem: this.certPem,
      caPrivateKeyPem: this.privateKeyPem,
    }, args.options);

    return { result: cert };
  }
}

export function newCertificate(ca: Input<RootCa>, args: NewCertificateInputs): Certificate {
  const self = output(ca);
  return new Certificate(args.name, {
    algorithm: args.algorithm,
    allowedUses: args.allowedUses,
    validityPeriodHours: args.validityPeriodHours,
    dnsNames: args.dnsNames,
    ecdsaCurve: args.ecdsaCurve,
    ipAddresses: args.ipAddresses,
    isCaCertificate: args.isCaCertificate,
    rsaBits: args.rsaBits,
    subject: args.subject,
    uris: args.uris,
    caCertPem: self.certPem,
    caPrivateKeyPem: self.privateKeyPem,
  }, args.options);
}
