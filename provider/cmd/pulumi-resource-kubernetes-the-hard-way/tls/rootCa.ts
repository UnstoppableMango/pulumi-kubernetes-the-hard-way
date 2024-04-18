import { ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { SelfSignedCert } from '@pulumi/tls';
import { SelfSignedCertSubject } from '@pulumi/tls/types/output';
import * as schema from '../schema-types';
import { AllowedUsage } from '../types';
import { toAllowedUsage } from '../util';
import { File, InstallInputs, InstallOutputs, install } from '../remote';
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
    if (opts?.urn) return;

    const algorithm = output(args.algorithm ?? 'RSA');
    const dnsNames = output(args.dnsNames ?? []);
    const earlyRenewalHours = output(args.earlyRenewalHours);
    const ecdsaCurve = output(args.ecdsaCurve ?? 'P256');
    const ipAddresses = output(args.ipAddresses ?? []);
    const rsaBits = output(args.rsaBits);
    const setAuthorityKeyId = output(args.setAuthorityKeyId);
    const setSubjectKeyId = output(args.setSubjectKeyId);
    const subject = output(args.subject);
    const uris = output(args.uris ?? []);
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
      dnsNames,
      ipAddresses,
      uris,
      setAuthorityKeyId: args.setAuthorityKeyId,
      setSubjectKeyId: args.setSubjectKeyId,
      earlyRenewalHours: args.earlyRenewalHours,
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
    this.dnsNames = dnsNames;
    this.earlyRenewalHours = cert.earlyRenewalHours;
    this.ecdsaCurve = ecdsaCurve;
    this.ipAddresses = ipAddresses;
    this.isCaCertificate = cert.isCaCertificate;
    this.key = key;
    this.keyAlgorithm = cert.keyAlgorithm;
    this.privateKeyOpenssh = key.privateKeyOpenssh;
    this.privateKeyPem = key.privateKeyPem;
    this.privateKeyPemPkcs8 = key.privateKeyPemPkcs8;
    this.publicKeyFingerprintMd5 = key.publicKeyFingerprintMd5;
    this.publicKeyFingerprintSha256 = key.publicKeyFingerprintSha256;
    this.publicKeyOpenssh = key.publicKeyOpenssh;
    this.publicKeyPem = key.publicKeyPem;
    this.readyForRenewal = cert.readyForRenewal;
    this.rsaBits = key.rsaBits;
    this.setAuthorityKeyId = cert.setAuthorityKeyId;
    this.setSubjectKeyId = cert.setSubjectKeyId;
    this.subject = subject as Output<SelfSignedCertSubject> | undefined;
    this.uris = uris;
    this.validityEndTime = cert.validityEndTime;
    this.validityPeriodHours = validityPeriodHours;
    this.validityStartTime = cert.validityStartTime;

    this.registerOutputs({
      algorithm,
      allowedUses: this.allowedUses,
      cert,
      certPem: this.certPem,
      dnsNames,
      earlyRenewalHours,
      ecdsaCurve,
      ipAddresses,
      isCaCertificate: this.isCaCertificate,
      key,
      keyAlgorithm: this.keyAlgorithm,
      privateKeyOpenSsh: this.privateKeyOpenssh,
      privateKeyPem: this.privateKeyPem,
      privateKeyPemPkcs8: this.privateKeyPemPkcs8,
      publicKeyFingerprintMd5: this.publicKeyFingerprintMd5,
      publicKeyFingerprintSha256: this.publicKeyFingerprintSha256,
      publicKeyOpenssh: this.publicKeyOpenssh,
      publicKeyPem: this.publicKeyPem,
      readyForRenewal: this.readyForRenewal,
      rsaBits,
      setAuthorityKeyId,
      setSubjectKeyId,
      subject,
      uris,
      validityEndTime: this.validityEndTime,
      validityPeriodHours,
      validityStartTime: this.validityStartTime,
    });
  }

  public async installCert(args: InstallInputs): Promise<InstallOutputs> {
    return { result: installCert(this, args) };
  }

  public async installKey(args: InstallInputs): Promise<InstallOutputs> {
    return { result: installKey(this, args) };
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

function installCert(ca: RootCa, args: InstallInputs): File {
  return install(args, ca.certPem);
}

function installKey(ca: RootCa, args: InstallInputs): File {
  return install(args, ca.publicKeyPem);
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
    caPrivateKeyPem: self.privateKeyPem as Output<string>, // TODO: Wtf
  }, args.options);
}
