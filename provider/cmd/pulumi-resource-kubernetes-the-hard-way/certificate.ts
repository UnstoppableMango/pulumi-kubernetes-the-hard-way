import { ComponentResourceOptions, Input, Inputs, Output, log } from '@pulumi/pulumi';
import { ConstructResult } from '@pulumi/pulumi/provider';
import { CertRequest, LocallySignedCert } from '@pulumi/tls';
import { CertRequestSubject } from '@pulumi/tls/types/input';
import { KeyPair, KeyPairArgs } from './keypair';
import { AllowedUsage } from './types';
import { toAllowedUsage } from './util';

export interface CertificateArgs extends KeyPairArgs {
  allowedUses: Input<Input<AllowedUsage>[]>;
  dnsNames?: Input<Input<string>[]>;
  caCertPem: Input<string>;
  caPrivateKeyPem: Input<string>;
  ipAddresses?: Input<Input<string>[]>;
  isCaCertificate?: Input<boolean>;
  uris?: Input<Input<string>[]>;
  subject?: Input<CertRequestSubject>;
}

export class Certificate extends KeyPair<LocallySignedCert> {
  public readonly allowedUses: Output<AllowedUsage[]>;
  public readonly certPem: Output<string>;
  public readonly cert: LocallySignedCert;
  public readonly csr: CertRequest;

  constructor(name: string, args: CertificateArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:index:Certificate', name, args, opts);

    const key = this.key;

    const csr = new CertRequest(name, {
      privateKeyPem: key.privateKeyPem,
      ipAddresses: args.ipAddresses,
      dnsNames: args.dnsNames,
      uris: args.uris,
      subject: args.subject,
    }, { parent: this });

    const cert = new LocallySignedCert(name, {
      isCaCertificate: args.isCaCertificate,
      allowedUses: args.allowedUses,
      validityPeriodHours: args.validityPeriodHours,
      caCertPem: args.caCertPem,
      caPrivateKeyPem: args.caPrivateKeyPem,
      certRequestPem: csr.certRequestPem,
    }, { parent: this });

    this.allowedUses = cert.allowedUses.apply(toAllowedUsage);
    this.cert = cert;
    this.certPem = cert.certPem;
    this.csr = csr;

    this.registerOutputs({
      allowedUses: this.allowedUses,
      cert, certPem: cert.certPem, csr, key,
      privateKeyPem: key.privateKeyPem,
      publicKeyPem: key.publicKeyPem,
    });
  }
}

export async function construct(
  name: string,
  inputs: Inputs,
  options: ComponentResourceOptions,
): Promise<ConstructResult> {
  const cert = new Certificate(name, inputs as CertificateArgs, options);
  return {
    urn: cert.urn,
    state: {
      allowedUses: cert.allowedUses,
      cert: cert.cert,
      certPem: cert.certPem,
      csr: cert.csr,
      key: cert.key,
      privateKeyPem: cert.privateKeyPem,
      publicKeyPem: cert.publicKeyPem,
    },
  };
}
