import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as schema from '../schema-types';
import { CommandBuilder } from './commandBuilder';

export class Mktemp extends schema.Mktemp {
  constructor(name: string, args: schema.MktempArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const connection = output(args.connection);
    const directory = output(args.directory ?? false);
    const dryRun = output(args.dryRun ?? false);
    const quiet = output(args.quiet ?? false);
    const suffix = output(args.suffix);
    const template = output(args.template);
    const tmpdir = output(args.tmpdir);

    const builder = new CommandBuilder('mktemp')
      .option('--directory', directory)
      .option('--dry-run', dryRun)
      .option('--quiet', quiet)
      .option('--suffix', suffix)
      .option('--tmpdir', tmpdir)
      .arg(args.template);

    const command = new Command(name, {
      connection,
      [args.lifecycle ?? 'create']: builder.command,
      triggers: args.triggers,
    }, { parent: this });

    this.command = command;
    this.directory = directory;
    this.dryRun = dryRun;
    this.quiet = quiet;
    this.stderr = command.stderr;
    this.stdout = command.stdout;

    if (args.suffix) this.suffix = output(args.suffix);
    if (args.template) this.template = output(args.template);
    if (args.tmpdir) this.tmpdir = output(args.tmpdir);

    this.registerOutputs({
      command,
      directory,
      dryRun,
      quiet,
      suffix,
      template,
      tmpdir,
    });
  }
}
