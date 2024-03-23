import { ComponentResourceOptions, Input, Output } from '@pulumi/pulumi';
import { SelfSignedCert } from '@pulumi/tls';
import { KeyPair, KeyPairArgs } from './keypair';
import { Certificate, CertificateArgs } from './certificate';

export interface RootCaArgs extends KeyPairArgs { }

export class RootCa extends KeyPair<SelfSignedCert> {
  public readonly cert: SelfSignedCert;
  public readonly certPem: Output<string>;

  constructor(name: string, args: RootCaArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:index:RootCa', name, args, opts);

    const cert = new SelfSignedCert(name, {
      isCaCertificate: true,
      allowedUses: args.allowedUses,
      privateKeyPem: this.key.privateKeyPem,
      validityPeriodHours: args.validityPeriodHours,
      subject: {
        commonName: 'Kubernetes',
        country: args.country,
        organization: args.organization,
        organizationalUnit: args.organizationalUnit,
        province: args.state, // Eh
      },
    }, { parent: this });

    this.cert = cert;
    this.certPem = cert.certPem;

    this.registerOutputs({ cert, certPem: cert.certPem });
  }

  public newCertificate(
    name: string,
    args: Omit<CertificateArgs, 'caCertPem' | 'caPrivateKeyPem'>,
    opts?: ComponentResourceOptions,
  ): Certificate {
    return new Certificate(name, {
      ...args,
      caCertPem: this.cert.certPem,
      caPrivateKeyPem: this.key.privateKeyPem,
    }, opts);
  }

  public newAdminCertificate(
    name: string,
    args: Omit<
      CertificateArgs,
      'caCertPem' | 'caPrivateKeyPem' | 'commonName' | 'organization'
    >,
    opts?: ComponentResourceOptions,
  ): Certificate {
    return this.newCertificate(`${name}-admin`, {
      ...args,
      commonName: 'admin',
      organization: 'system:masters',
    }, opts);
  }

  public newKubeletCertificate(
    name: string,
    index: number,
    ip: string,
    args: Omit<
      CertificateArgs,
      'caCertPem' | 'caPrivateKeyPem' | 'commonName' | 'organization'
    >,
    opts?: ComponentResourceOptions,
  ): Certificate {
    return this.newCertificate(`${name}-node${index}`, {
      ...args,
      commonName: `system:node:${ip}`,
      organization: 'system:nodes',
    }, opts);
  }
}
