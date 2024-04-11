import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as schema from '../schema-types';
import { CommandBuilder } from './commandBuilder';

export class Chmod extends schema.Chmod {
  constructor(name: string, args: schema.ChmodArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const binaryPath = output(args.binaryPath ?? 'chmod');
    const changes = output(args.changes ?? false);
    const connection = output(args.connection);
    const environment = output(args.environment ?? {});
    const files = output(args.files);
    const help = output(args.help ?? false);
    const lifecycle = args.lifecycle ?? 'create';
    const mode = output(args.mode);
    const noPreserveRoot = output(args.noPreserveRoot ?? false);
    const preserveRoot = output(args.preserveRoot ?? false);
    const quiet = output(args.quiet ?? false);
    const recursive = output(args.recursive ?? false);
    const reference = output(args.reference);
    const silent = output(args.silent ?? false);
    const triggers = output(args.triggers ?? []);
    const version = output(args.version ?? false);

    const builder = new CommandBuilder(binaryPath)
      .option('--changes', changes)
      .option('--no-preserve-root', noPreserveRoot)
      .option('--preserve-root', preserveRoot)
      .option('--quiet', quiet)
      .option('--silent', silent)
      .option('--recursive', recursive)
      .option('--refernce', reference)
      .option('--help', help)
      .option('--version', version)
      .arg(mode)
      .arg(files);

    const command = new Command(name, {
      connection,
      environment,
      stdin: args.stdin,
      [lifecycle]: builder.command,
      triggers,
    }, { parent: this });

    this.binaryPath = binaryPath;
    this.changes = changes;
    this.command = command;
    this.connection = connection;
    this.environment = environment;
    this.files = files;
    this.help = help;
    this.lifecycle = lifecycle;
    this.mode = mode;
    this.noPreserveRoot = noPreserveRoot;
    this.preserveRoot = preserveRoot;
    this.quiet = quiet;
    this.recursive = recursive;
    this.reference = reference as Output<string> | undefined; // Ugh
    this.silent = silent;
    this.stderr = command.stderr;
    this.stdin = output(args.stdin) as Output<string> | undefined; // Ugh
    this.stdout = command.stdout;
    this.triggers = triggers;
    this.version = version;

    this.registerOutputs({
      binaryPath,
      changes,
      command,
      connection,
      environment,
      files,
      help,
      lifecycle,
      mode,
      noPreserveRoot,
      preserveRoot,
      quiet,
      recursive,
      reference,
      silent,
      stderr: this.stderr,
      stdin: this.stdin,
      stdout: this.stdout,
      triggers,
      version,
    });
  }
}
