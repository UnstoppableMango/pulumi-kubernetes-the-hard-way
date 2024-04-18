import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import { CertRequest, LocallySignedCert } from '@pulumi/tls';
import { CertRequestSubject } from '@pulumi/tls/types/output';
import * as schema from '../schema-types';
import { toAllowedUsage } from '../util';
import { File, InstallInputs, InstallOutputs, install } from '../remote';
import { KeyPair } from './keypair';

export class Certificate extends schema.Certificate {
  constructor(name: string, args: schema.CertificateArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const algorithm = output(args.algorithm ?? 'RSA');
    const allowedUses = output(args.allowedUses);
    const caCertPem = output(args.caCertPem);
    const caPrivateKeyPem = output(args.caPrivateKeyPem);
    const dnsNames = output(args.dnsNames ?? []);
    const ecdsaCurve = output(args.ecdsaCurve ?? 'P256');
    const ipAddresses = output(args.ipAddresses ?? []);
    const isCaCertificate = output(args.isCaCertificate ?? false);
    const subject = output(args.subject);
    const uris = output(args.uris ?? []);
    const validityPeriodHours = output(args.validityPeriodHours);

    const key = KeyPair.key(name, {
      algorithm,
      validityPeriodHours,
      ecdsaCurve,
      rsaBits: args.rsaBits,
    }, this);

    const csr = new CertRequest(name, {
      privateKeyPem: key.privateKeyPem,
      ipAddresses,
      dnsNames,
      uris,
      subject: args.subject,
    }, { parent: this });

    const cert = new LocallySignedCert(name, {
      isCaCertificate: args.isCaCertificate,
      allowedUses,
      validityPeriodHours,
      caCertPem,
      caPrivateKeyPem,
      certRequestPem: csr.certRequestPem,
      earlyRenewalHours: args.earlyRenewalHours,
      setSubjectKeyId: args.setSubjectKeyId
    }, { parent: this });

    this.algorithm = algorithm;
    this.allowedUses = cert.allowedUses.apply(toAllowedUsage);
    this.caCertPem = caCertPem;
    this.caPrivateKeyPem = caPrivateKeyPem;
    this.cert = cert;
    this.certPem = cert.certPem;
    this.csr = csr;
    this.dnsNames = dnsNames;
    this.earlyRenewalHours = cert.earlyRenewalHours;
    this.isCaCertificate = isCaCertificate;
    this.key = key;
    this.privateKeyPem = key.privateKeyPem;
    this.publicKeyPem = key.publicKeyPem;
    this.rsaBits = key.rsaBits;
    this.subject = subject as Output<CertRequestSubject> | undefined;
    this.uris = uris;
    this.validityPeriodHours = validityPeriodHours;

    this.registerOutputs({
      algorithm,
      allowedUses: this.allowedUses,
      caCertPem,
      caPrivateKeyPem,
      cert,
      certPem: this.certPem,
      csr,
      dnsNames,
      earlyRenewalHours: cert.earlyRenewalHours,
      isCaCertificate,
      key,
      privateKeyPem: this.privateKeyPem,
      publicKeyPem: this.publicKeyPem,
      rsaBits: this.rsaBits,
      subject: this.subject,
      uris,
      validityPeriodHours,
    });
  }

  public async installCert(args: InstallInputs): Promise<InstallOutputs> {
    return { result: installCert(this, args) };
  }

  public async installKey(args: InstallInputs): Promise<InstallOutputs> {
    return { result: installKey(this, args) };
  }
}

function installCert(cert: Certificate, args: InstallInputs): File {
  return install(args, cert.certPem);
}

function installKey(cert: Certificate, args: InstallInputs): File {
  return install(args, cert.publicKeyPem);
}
