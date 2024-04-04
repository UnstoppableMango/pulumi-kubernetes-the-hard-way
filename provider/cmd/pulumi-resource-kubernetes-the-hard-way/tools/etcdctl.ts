import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { CommandBuilder } from './commandBuilder';
import { Command } from '@pulumi/command/remote';

export class Etcdctl extends schema.Etcdctl {
  constructor(name: string, args: schema.EtcdctlArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const caCert = output(args.caCert);
    const cert = output(args.cert);
    const connection = output(args.connection);
    const endpoints = output(args.endpoints);
    const env = output(args.env ?? {}); // TODO: Fix generated type
    const key = output(args.key);

    const builder = new CommandBuilder('etcdctl')
      .arg(args.commands)
      .option('--ca-cert', caCert)
      .option('--cert', cert)
      .option('--endpoints', endpoints)
      .option('--key', key);

    const command = new Command(name, {
      connection,
      environment: env,
      create: builder.command,
    }, { parent: this });
  }
}
