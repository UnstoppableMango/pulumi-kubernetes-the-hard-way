import { ComponentResource, ComponentResourceOptions, Input, Output } from '@pulumi/pulumi';
import { LocallySignedCert, PrivateKey, SelfSignedCert } from '@pulumi/tls';
import { InstallArgs, RemoteFile, install } from './remoteFile';
import { Algorithm, EcdsaCurve } from './types';

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
  public readonly key: PrivateKey;
  public readonly privateKeyPem: Output<string>;
  public readonly publicKeyPem: Output<string>;

  protected constructor(type: string, name: string, args: KeyPairArgs, opts?: ComponentResourceOptions) {
    super(type, name, args, opts);

    const key = new PrivateKey(name, {
      algorithm: args.algorithm,
      rsaBits: args.rsaBits,
      ecdsaCurve: args.ecdsaCurve,
    }, { parent: this });

    this.key = key;
    this.privateKeyPem = key.privateKeyPem;
    this.publicKeyPem = key.publicKeyPem;
  }

  public installCert(args: InstallArgs): RemoteFile {
    return installCert(this, args);
  }

  public installKey(args: InstallArgs): RemoteFile {
    return installKey(this, args);
  }
}

export function installCert<T extends CertType>(pair: KeyPair<T>, args: InstallArgs): RemoteFile {
  return install(args, pair.certPem);
}

export function installKey<T extends CertType>(pair: KeyPair<T>, args: InstallArgs): RemoteFile {
  return install(args, pair.publicKeyPem);
}
