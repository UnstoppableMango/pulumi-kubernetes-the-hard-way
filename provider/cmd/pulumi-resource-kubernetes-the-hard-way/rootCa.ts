import { ComponentResourceOptions, Input, Inputs, Output } from '@pulumi/pulumi';
import { SelfSignedCert } from '@pulumi/tls';
import { KeyPair, KeyPairArgs } from './keypair';
import { Certificate, CertificateArgs } from './certificate';
import { AllowedUsage } from './types';
import { ConstructResult } from '@pulumi/pulumi/provider';

export interface RootCaArgs extends KeyPairArgs { }

export class RootCa extends KeyPair<SelfSignedCert> {
  public readonly cert: SelfSignedCert;
  public readonly certPem: Output<string>;

  constructor(name: string, args: RootCaArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:index:RootCa', name, args, opts);

    const cert = new SelfSignedCert(name, {
      isCaCertificate: true,
      allowedUses: [
        AllowedUsage.Cert_signing,
        AllowedUsage.Key_encipherment,
        AllowedUsage.Server_auth,
        AllowedUsage.Client_auth,
      ],
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
}

export async function construct(name: string, inputs: Inputs, options: ComponentResourceOptions): Promise<ConstructResult> {
  const rootCa = new RootCa(name, inputs as RootCaArgs, options);
  return {
      urn: rootCa.urn,
      state: {
          allowedUses: rootCa.allowedUses,
          cert: rootCa.cert,
          certPem: rootCa.certPem,
          key: rootCa.key,
          keyPem: rootCa.keyPem,
      },
  };
}
