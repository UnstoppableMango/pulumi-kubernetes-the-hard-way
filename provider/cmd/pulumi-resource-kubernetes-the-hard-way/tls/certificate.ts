import { ComponentResourceOptions, Inputs, Output, all, output } from '@pulumi/pulumi';
import { ConstructResult } from '@pulumi/pulumi/provider';
import { CertRequest, LocallySignedCert } from '@pulumi/tls';
import { AllowedUsage, CertificateArgs } from './sdk';
import { KeyPair } from './keypair';
import { AllowedUsage } from '../types';
import { toAllowedUsage } from '../util';

export class Certificate extends KeyPair<LocallySignedCert> {
  public readonly allowedUses!: Output<AllowedUsage[]>;
  public readonly certPem!: Output<string>;
  public readonly cert!: LocallySignedCert;
  public readonly csr!: CertRequest;

  constructor(name: string, args: CertificateArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:tls:Certificate', name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const key = this.key;
    const allowedUses = output(args.allowedUses);

    const csr = new CertRequest(name, {
      privateKeyPem: key.privateKeyPem,
      ipAddresses: args.ipAddresses,
      dnsNames: args.dnsNames,
      uris: args.uris,
      subject: args.subject,
    }, { parent: this });

    const cert = new LocallySignedCert(name, {
      isCaCertificate: args.isCaCertificate,
      allowedUses,
      validityPeriodHours: args.validityPeriodHours,
      caCertPem: args.caCertPem,
      caPrivateKeyPem: args.caPrivateKeyPem,
      certRequestPem: csr.certRequestPem,
    }, { parent: this });

    this.allowedUses = allowedUses;
    this.cert = cert;
    this.certPem = cert.certPem;
    this.csr = csr;

    this.registerOutputs({
      allowedUses,
      cert, certPem: cert.certPem, csr, key,
      privateKeyPem: key.privateKeyPem,
      publicKeyPem: key.publicKeyPem,
    });
  }
}
