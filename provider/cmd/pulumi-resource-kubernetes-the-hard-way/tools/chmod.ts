import { ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as schema from '../schema-types';
import * as tool from './tool';
import { CommandBuilder } from './commandBuilder';

export class Chmod extends schema.Chmod {
  constructor(name: string, args: schema.ChmodArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const {
      binaryPath,
      connection,
      environment,
      triggers,
    } = tool.outputs(args, 'chmod');

    const command = new Command(name, {
      connection,
      environment,
      stdin: args.stdin,
      triggers,
      create: build(binaryPath, args.create),
      update: build(binaryPath, args.update),
      delete: build(binaryPath, args.delete),
    }, { parent: this });

    this.binaryPath = binaryPath;
    this.command = command;
    this.connection = connection;
    this.environment = environment;
    this.stderr = command.stderr;
    this.stdin = output(args.stdin) as Output<string> | undefined; // Ugh
    this.stdout = command.stdout;
    this.triggers = triggers;

    this.registerOutputs({
      binaryPath,
      command,
      connection,
      environment,
      stderr: this.stderr,
      stdin: this.stdin,
      stdout: this.stdout,
      triggers,
    });
  }
}

function build(path: Input<string>, opts?: schema.ChmodOptsInputs): Output<string> | undefined {
  if (!opts) return undefined;

  const builder = new CommandBuilder(path)
    .option('--changes', opts.changes)
    .option('--no-preserve-root', opts.noPreserveRoot)
    .option('--preserve-root', opts.preserveRoot)
    .option('--quiet', opts.quiet)
    .option('--silent', opts.silent)
    .option('--recursive', opts.recursive)
    .option('--refernce', opts.reference)
    .option('--help', opts.help)
    .option('--version', opts.version)
    .arg(opts.mode)
    .arg(opts.files);

  return builder.command;
}
