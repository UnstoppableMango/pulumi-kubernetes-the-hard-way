import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as schema from '../schema-types';
import { CommandBuilder } from './commandBuilder';

export class Systemctl extends schema.Systemctl {
  constructor(name: string, args: schema.SystemctlArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const connection = output(args.connection);
    const commands = output(args.commands);
    const environment = output(args.environment ?? {});
    const lifecycle = args.lifecycle ?? 'create';
    const serviceName = output(args.serviceName);

    const builder = new CommandBuilder('systemctl')
      .arg(commands)
      .arg(args.serviceName);

    const command = new Command(name, {
      connection,
      environment,
      [lifecycle]: builder.command,
    }, { parent: this });

    this.command = command;
    this.commands = commands;
    this.connection = connection;
    this.serviceName = serviceName as Output<string>;
    this.stderr = command.stderr;
    this.stdin = command.stdin as Output<string>;
    this.stdout = command.stdout;

    this.registerOutputs({
      command,
      commands,
      connection,
      serviceName,
      stderr: command.stderr,
      stdin: command.stdin,
      stdout: command.stdout,
    });
  }
}
