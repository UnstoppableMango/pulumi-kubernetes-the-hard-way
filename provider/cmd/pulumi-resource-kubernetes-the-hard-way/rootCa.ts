import { ComponentResourceOptions, Input, Inputs, Output, output } from '@pulumi/pulumi';
import { ConstructResult, InvokeResult } from '@pulumi/pulumi/provider';
import { SelfSignedCert } from '@pulumi/tls';
import { SelfSignedCertSubject } from '@pulumi/tls/types/input';
import { KeyPair, KeyPairArgs } from './keypair';
import { Certificate, CertificateArgs } from './certificate';
import { AllowedUsage } from './types';
import { toAllowedUsage } from './util';

export type NewCertificateArgs = Omit<CertificateArgs, 'caCertPem' | 'caPrivateKeyPem'> & {
  name: string;
  opts?: ComponentResourceOptions;
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
      caCertPem: this.certPem,
      caPrivateKeyPem: this.privateKeyPem,
    }, args.opts);
  }
}

export async function newCertificate(inputs: Inputs): Promise<InvokeResult> {
  const result = (inputs.__self__ as RootCa).newCertificate({
    algorithm: inputs.algorithm,
    allowedUses: inputs.allowedUses,
    name: inputs.name,
    validityPeriodHours: inputs.validityPeriodHours,
    dnsNames: inputs.dnsNames,
    ecdsaCurve: inputs.ecdsaCurve,
    ipAddresses: inputs.ipAddresses,
    isCaCertificate: inputs.isCaCertificate,
    opts: inputs.opts,
    rsaBits: inputs.rsaBits,
    subject: inputs.subject,
    uris: inputs.uris,
  });

  return { outputs: { result } };
}

export async function installCert(inputs: Inputs): Promise<InvokeResult> {
  const result = (inputs.__self__ as RootCa).installCert({
    connection: inputs.connection,
    name: inputs.name,
    path: inputs.path,
    opts: inputs.opts,
  });

  return { outputs: { result } };
}

export async function installKey(inputs: Inputs): Promise<InvokeResult> {
  const result = (inputs.__self__ as RootCa).installCert({
    connection: inputs.connection,
    name: inputs.name,
    path: inputs.path,
    opts: inputs.opts,
  });

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
