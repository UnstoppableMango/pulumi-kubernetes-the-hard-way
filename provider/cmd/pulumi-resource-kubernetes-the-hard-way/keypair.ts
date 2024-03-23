import { ComponentResource, ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { LocallySignedCert, PrivateKey, SelfSignedCert } from '@pulumi/tls';
import { remote } from '@pulumi/command/types/input';
import { InstallArgs, RemoteFile } from './remoteFile';
import { Algorithm, AllowedUsage, EcdsaCurve } from './types';

export interface KeyPairArgs {
  algorithm: Input<Algorithm>;
  allowedUses: Input<Input<AllowedUsage>[]>;
  ecdsaCurve?: Input<EcdsaCurve>;
  commonName: Input<string>;
  country?: Input<string>;
  validityPeriodHours: Input<number>;
  location?: Input<string>;
  organization?: Input<string>;
  organizationalUnit?: Input<string>;
  size?: Input<number>;
  state?: Input<string>;
}

type CertType = SelfSignedCert | LocallySignedCert;

export abstract class KeyPair<TCert extends CertType> extends ComponentResource {
  public readonly allowedUses: Output<Output<AllowedUsage>[]>;
  public abstract readonly cert: TCert;
  public abstract readonly certPem: Output<string>;
  public readonly key: PrivateKey;
  public readonly keyPem: Output<string>;

  constructor(type: string, name: string, args: KeyPairArgs, opts?: ComponentResourceOptions) {
    super(type, name, args, opts);
  
    // TODO: Can this be cleaned up?
    const allowedUses = output(args.allowedUses).apply(x => x.map(y => output(y)));

    const key = new PrivateKey(name, {
      algorithm: args.algorithm,
      rsaBits: args.size,
      ecdsaCurve: args.ecdsaCurve,
    }, { parent: this });

    this.allowedUses = allowedUses;
    this.key = key;
    this.keyPem = key.publicKeyPem;

    this.registerOutputs({ allowedUses, key, keyPem: key.publicKeyPem });
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
  return install(name, args.connection, pair.keyPem, args.path, opts);
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
