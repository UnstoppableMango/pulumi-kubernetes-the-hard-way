import { ComponentResourceOptions, Input, Inputs, Output, output } from '@pulumi/pulumi';
import { SelfSignedCert } from '@pulumi/tls';
import { SelfSignedCertSubject } from '@pulumi/tls/types/input';
import * as schema from '../schema-types';
import { KeyPair, KeyPairArgs } from './keypair';
import { Certificate } from './certificate';
import { AllowedUsage } from '../types';
import { toAllowedUsage } from '../util';

export type NewCertificateInputs = Omit<schema.CertificateArgs, 'caCertPem' | 'caPrivateKeyPem'> & {
  name: string;
  options?: ComponentResourceOptions;
};

export interface NewCertificateOutputs {
  result: Certificate;
}

export interface RootCaArgs extends KeyPairArgs {
  subject?: Input<SelfSignedCertSubject>;
}

export class RootCa extends KeyPair<SelfSignedCert> {
  public static readonly __pulumiType: string = 'kubernetes-the-hard-way:tls:RootCa';

  public readonly allowedUses!: Output<AllowedUsage[]>;
  public readonly cert!: SelfSignedCert;
  public readonly certPem!: Output<string>;

  constructor(name: string, args: RootCaArgs, opts?: ComponentResourceOptions) {
    const props = {
      allowedUses: undefined,
      cert: undefined,
      certPem: undefined,
      key: undefined,
      privateKeyPem: undefined,
      publicKeyPem: undefined,
    };

    super(RootCa.__pulumiType, name, opts?.urn ? props : args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const key = this.key;

    const cert = new SelfSignedCert(name, {
      isCaCertificate: true,
      allowedUses: [
        AllowedUsage.CertSigning,
        AllowedUsage.KeyEncipherment,
        AllowedUsage.ServerAuth,
        AllowedUsage.ClientAuth,
      ],
      privateKeyPem: key.privateKeyPem,
      validityPeriodHours: args.validityPeriodHours,
      subject: output(args.subject).apply(subject => ({
        commonName: subject?.commonName ?? 'Kubernetes',
        country: subject?.country,
        locality: subject?.locality,
        organization: subject?.organization,
        organizationalUnit: subject?.organizationalUnit,
        postalCode: subject?.postalCode,
        province: subject?.province,
        serialNumber: subject?.serialNumber,
        streetAddresses: subject?.streetAddresses,
      })),
    }, { parent: this });

    this.allowedUses = cert.allowedUses.apply(toAllowedUsage);
    this.cert = cert;
    this.certPem = cert.certPem;

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
