import { ComponentResourceOptions, Input, Inputs, Output, output } from '@pulumi/pulumi';
import { ConstructResult, InvokeResult } from '@pulumi/pulumi/provider';
import { SelfSignedCert } from '@pulumi/tls';
import { SelfSignedCertSubject } from '@pulumi/tls/types/input';
import { KeyPair, KeyPairArgs, installCert, installKey } from './keypair';
import { Certificate, CertificateArgs } from './certificate';
import { InstallArgs } from './remoteFile';
import { AllowedUsage } from './types';
import { toAllowedUsage } from './util';

export type NewCertificateArgs = Omit<CertificateArgs, 'caCertPem' | 'caPrivateKeyPem'> & {
  name: string;
  options?: ComponentResourceOptions;
};

export interface RootCaArgs extends KeyPairArgs {
  subject?: Input<SelfSignedCertSubject>;
}

export class RootCa extends KeyPair<SelfSignedCert> {
  public readonly allowedUses: Output<AllowedUsage[]>;
  public readonly cert: SelfSignedCert;
  public readonly certPem: Output<string>;

  constructor(name: string, args: RootCaArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:index:RootCa', name, args, opts);

    const key = this.key;

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

  public newCertificate(args: NewCertificateArgs): Certificate {
    return newCertificate(this, args);
  }
}

export function newCertificate(ca: RootCa, args: NewCertificateArgs): Certificate {
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
    caCertPem: ca.certPem,
    caPrivateKeyPem: ca.privateKeyPem,
  }, args.options);
}

export async function callNewCertificateInstance(inputs: Inputs): Promise<InvokeResult> {
  const result = newCertificate(inputs.__self__, inputs as NewCertificateArgs);
  return { outputs: { result } };
}

export async function callNewCertificateStatic(inputs: Inputs): Promise<InvokeResult> {
  const result = newCertificate(inputs.ca, inputs as NewCertificateArgs);
  return { outputs: { result } };
}

export async function callInstallCertStatic(inputs: Inputs): Promise<InvokeResult> {
  const result = installCert(inputs.ca, inputs as InstallArgs);
  return { outputs: { result } };
}

export async function callInstallKeyStatic(inputs: Inputs): Promise<InvokeResult> {
  const result = installKey(inputs.ca, inputs as InstallArgs);
  return { outputs: { result } };
}

export async function construct(
  name: string,
  inputs: Inputs,
  options: ComponentResourceOptions,
): Promise<ConstructResult> {
  const rootCa = new RootCa(name, inputs as RootCaArgs, options);
  return {
    urn: rootCa.urn,
    state: {
      allowedUses: rootCa.allowedUses,
      cert: rootCa.cert,
      certPem: rootCa.certPem,
      key: rootCa.key,
      privateKeyPem: rootCa.privateKeyPem,
      publicKeyPem: rootCa.publicKeyPem,
    },
  };
}
