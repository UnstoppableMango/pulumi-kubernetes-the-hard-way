import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as schema from '../schema-types';
import { CommandBuilder } from './commandBuilder';

export class Etcdctl extends schema.Etcdctl {
  constructor(name: string, args: schema.EtcdctlArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const binaryPath = output(args.binaryPath ?? 'etcdctl');
    const caCert = output(args.caCert);
    const cert = output(args.cert);
    const connection = output(args.connection);
    const endpoints = output(args.endpoints);
    const env = output(args.environment ?? {}); // TODO: Fix generated type
    const key = output(args.key);

    const builder = new CommandBuilder(binaryPath)
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

    this.binaryPath = binaryPath;
    this.command = command;
    this.connection = connection;

    this.registerOutputs({ command, connection });
  }
}
