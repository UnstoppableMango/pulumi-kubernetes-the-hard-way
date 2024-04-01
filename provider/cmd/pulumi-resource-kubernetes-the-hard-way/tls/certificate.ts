import { ComponentResourceOptions, Input, Inputs, Output } from '@pulumi/pulumi';
import { ConstructResult } from '@pulumi/pulumi/provider';
import { CertRequest, LocallySignedCert } from '@pulumi/tls';
import { CertRequestSubject } from '@pulumi/tls/types/input';
import { KeyPair, KeyPairArgs } from './keypair';
import { AllowedUsage } from '../types';
import { toAllowedUsage } from '../util';

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
  public static readonly __pulumiType: string = 'kubernetes-the-hard-way:tls:Certificate';

  public readonly allowedUses!: Output<AllowedUsage[]>;
  public readonly certPem!: Output<string>;
  public readonly cert!: LocallySignedCert;
  public readonly csr!: CertRequest;

  constructor(name: string, args: CertificateArgs, opts?: ComponentResourceOptions) {
    const props = {
      allowedUses: undefined,
      certPem: undefined,
      cert: undefined,
      csr: undefined,
      key: undefined,
      privateKeyPem: undefined,
      publicKeyPem: undefined,
    };

    super(Certificate.__pulumiType, name, opts?.urn ? props : args, opts);

    // Rehydrating
    if (opts?.urn) return;

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
