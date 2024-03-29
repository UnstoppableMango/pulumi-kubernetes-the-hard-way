import { ComponentResource, ComponentResourceOptions, Input, Inputs, Output } from '@pulumi/pulumi';
import { InvokeResult } from '@pulumi/pulumi/provider';
import { LocallySignedCert, PrivateKey, SelfSignedCert } from '@pulumi/tls';
import { InstallArgs, File, install } from '../remote/file';
import { Algorithm, EcdsaCurve } from 'sdk';

export interface KeyPairArgs {
  algorithm: Input<Algorithm>;
  ecdsaCurve?: Input<EcdsaCurve>;
  rsaBits?: Input<number>;
  validityPeriodHours: Input<number>;
}

type CertType = SelfSignedCert | LocallySignedCert;

export abstract class KeyPair<TCert extends CertType> extends ComponentResource {
  public abstract readonly cert: TCert;
  public abstract readonly certPem: Output<string>;
  public readonly key!: PrivateKey;
  public readonly privateKeyPem!: Output<string>;
  public readonly publicKeyPem!: Output<string>;

  protected constructor(type: string, name: string, args: Inputs, opts?: ComponentResourceOptions) {
    super(type, name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const key = new PrivateKey(name, {
      algorithm: args.algorithm,
      rsaBits: args.rsaBits,
      ecdsaCurve: args.ecdsaCurve,
    }, { parent: this });

    this.key = key;
    this.privateKeyPem = key.privateKeyPem;
    this.publicKeyPem = key.publicKeyPem;
  }

  public installCert(args: InstallArgs): File {
    return installCert(this, args);
  }

  public installKey(args: InstallArgs): File {
    return installKey(this, args);
  }
}

export function installCert(pair: KeyPair<CertType>, args: InstallArgs): File {
  return install(args, pair.certPem);
}

export function installKey(pair: KeyPair<CertType>, args: InstallArgs): File {
  return install(args, pair.publicKeyPem);
}

export async function callInstallCertInstance(inputs: Inputs): Promise<InvokeResult> {
  const result = installCert(inputs.__self__, inputs as InstallArgs);
  return { outputs: { result } };
}

export async function callInstallCertStatic(inputs: Inputs): Promise<InvokeResult> {
  const result = installCert(inputs.keypair, inputs as InstallArgs);
  return { outputs: { result } };
}

export async function callInstallKeyInstance(inputs: Inputs): Promise<InvokeResult> {
  const result = installKey(inputs.__self__, inputs as InstallArgs);
  return { outputs: { result } };
}

export async function callInstallKeyStatic(inputs: Inputs): Promise<InvokeResult> {
  const result = installCert(inputs.keypair, inputs as InstallArgs);
  return { outputs: { result } };
}
