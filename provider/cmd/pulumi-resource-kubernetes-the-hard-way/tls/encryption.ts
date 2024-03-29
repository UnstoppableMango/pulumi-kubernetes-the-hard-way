import * as path from 'node:path';
import * as YAML from 'yaml';
import { ComponentResource, ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import { RandomBytes } from '@pulumi/random';
import { remote } from '@pulumi/command/types/input';
import { EncryptionKeyArgs } from './sdk';
import { File } from '../remote/file';

export class EncryptionKey extends ComponentResource {
  public static readonly defaultBytes = 24;

  public readonly config: Output<string>;
  public readonly key: RandomBytes;

  constructor(name: string, args: EncryptionKeyArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:tls:encryptionKey', name, args, opts);

    const bytes = output(args.bytes ?? EncryptionKey.defaultBytes);
    const key = new RandomBytes(name, { length: bytes }, { parent: this });
    const config = key.base64.apply(keyData => YAML.stringify({
      kind: 'EncryptionConfig',
      apiVersion: 'v1',
      resources: [{ secrets: [] }],
      providers: [
        {
          aescbc: {
            keys: [{
              name: 'key1',
              secret: keyData,
            }],
          },
        },
        { identity: {} },
      ],
    }));

    this.config = config;
    this.key = key;

    this.registerOutputs({ config, key });
  }

  public install(name: string, connection: remote.ConnectionArgs, opts?: ComponentResourceOptions): File {
    return install(this, name, connection, opts);
  }
}

export function install(
  key: EncryptionKey,
  name: string,
  connection: remote.ConnectionArgs,
  opts?: ComponentResourceOptions,
): File {
  const target = path.join('home', 'kthw'); // TODO
  return new File(name, {
    connection,
    content: key.config,
    path: path.join(target, 'encryption-config.yaml'),
  }, opts);
}
