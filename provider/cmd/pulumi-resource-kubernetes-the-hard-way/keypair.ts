import { ComponentResource, ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { LocallySignedCert, PrivateKey, SelfSignedCert } from '@pulumi/tls';
import { remote } from '@pulumi/command/types/input';
import { InstallArgs, RemoteFile } from './remoteFile';
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

  constructor(type: string, name: string, args: KeyPairArgs, opts?: ComponentResourceOptions) {
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

  public installCert(name: string, args: InstallArgs, opts?: ComponentResourceOptions): RemoteFile {
    return installCert(this, name, args, opts);
  }

  public installKey(name: string, args: InstallArgs, opts?: ComponentResourceOptions): RemoteFile {
    return installKey(this, name, args, opts);
  }
}

export function installCert<T extends CertType>(
  pair: KeyPair<T>,
  name: string,
  args: InstallArgs,
  opts?: ComponentResourceOptions
): RemoteFile {
  return install(name, args.connection, pair.certPem, args.path, opts);
}

export function installKey<T extends CertType>(
  pair: KeyPair<T>,
  name: string,
  args: InstallArgs,
  opts?: ComponentResourceOptions
): RemoteFile {
  return install(name, args.connection, pair.publicKeyPem, args.path, opts);
}

function install(
  name: string,
  connection: Input<remote.ConnectionArgs>,
  content: Input<string>,
  path: Input<string>,
  opts?: ComponentResourceOptions
): RemoteFile {
  return new RemoteFile(name, { connection, path, content }, opts);
}