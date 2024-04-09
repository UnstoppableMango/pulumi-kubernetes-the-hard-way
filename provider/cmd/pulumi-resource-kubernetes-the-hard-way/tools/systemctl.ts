import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as schema from '../schema-types';
import { CommandBuilder } from './commandBuilder';

export class Systemctl extends schema.Systemctl {
  constructor(name: string, args: schema.SystemctlArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const connection = output(args.connection);
    const systemctlCommand = output(args.command);
    const environment = output(args.environment ?? {});
    const lifecycle = args.lifecycle ?? 'create';
    const unit = output(args.unit);

    const builder = new CommandBuilder('systemctl')
      .arg(systemctlCommand)
      .arg(args.unit);

    const command = new Command(name, {
      connection,
      environment,
      [lifecycle]: builder.command,
    }, { parent: this });

    this.command = command;
    this.systemctlCommand = args.command;
    this.connection = connection;
    this.unit = unit as Output<string>;
    this.stderr = command.stderr;
    this.stdin = command.stdin as Output<string>;
    this.stdout = command.stdout;

    this.registerOutputs({
      command,
      systemctlCommand,
      connection,
      unit,
      stderr: command.stderr,
      stdin: command.stdin,
      stdout: command.stdout,
    });
  }
}
