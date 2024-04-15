import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as schema from '../schema-types';
import { CommandBuilder } from './commandBuilder';

export class Hostnamectl extends schema.Hostnamectl {
  constructor(name: string, args: schema.HostnamectlArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const arg = output(args.arg);
    const binaryPath = output(args.binaryPath ?? 'hostnamectl');
    const connection = output(args.connection);
    const hostnamectlCommand = output(args.command);
    const environment = output(args.environment ?? {});
    const help = output(args.help ?? false);
    const json = output(args.json);
    const machine = output(args.machine);
    const noAskPassword = output(args.noAskPassword ?? false);
    const pretty = output(args.pretty ?? false);
    const stc = output(args.static ?? false);
    const stdin = output(args.stdin);
    const transient = output(args.transient ?? false);
    const triggers = output(args.triggers ?? []);
    const version = output(args.version ?? false);
    const lifecycle = args.lifecycle ?? 'create';

    const builder = new CommandBuilder(binaryPath)
      .option('--help', help)
      .option('--json', json)
      .option('--machine', machine)
      .option('--no-ask-password', noAskPassword)
      .option('--pretty', pretty)
      .option('--static', stc)
      .option('--transient', transient)
      .option('--version', version)
      .arg(hostnamectlCommand)
      .arg(args.arg);

    const command = new Command(name, {
      connection,
      environment,
      triggers,
      stdin: args.stdin,
      [lifecycle]: builder.command,
    }, { parent: this });

    this.arg = arg as Output<string> | undefined;
    this.binaryPath = binaryPath;
    this.command = command;
    this.connection = connection;
    this.hostnamectlCommand = hostnamectlCommand;
    this.environment = environment;
    this.help = help;
    this.json = json as Output<schema.HostnamectlJsonModeOutputs> | undefined;
    this.machine = machine as Output<string> | undefined;
    this.noAskPassword = noAskPassword;
    this.pretty = pretty;
    this.static = stc;
    this.stderr = command.stderr;
    this.stdin = stdin as Output<string> | undefined;
    this.stdout = command.stdout;
    this.transient = transient;
    this.triggers = triggers;
    this.version = version;

    this.registerOutputs({
      arg,
      binaryPath,
      command,
      connection,
      hostnamectlCommand,
      environment,
      help,
      json,
      machine,
      noAskPassword,
      pretty,
      static: stc,
      stderr: command.stderr,
      stdin,
      stdout: command.stdout,
      transient,
      triggers,
      version,
    });
  }
}
