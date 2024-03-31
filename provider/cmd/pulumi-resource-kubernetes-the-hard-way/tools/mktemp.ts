import { ComponentResource, ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import { remote } from '@pulumi/command/types/input';
import { CommandBuilder } from './commandBuilder';

export interface MktempArgs {
  connection: Input<remote.ConnectionArgs>;
  directory?: Input<boolean>;
  dryRun?: Input<boolean>;
  quiet?: Input<boolean>;
  suffix?: Input<string>;
  template: Input<string>;
  tmpdir?: Input<string>;
}

export class Mktemp extends ComponentResource {
  public readonly command!: Command;
  public readonly directory!: Output<boolean>;
  public readonly dryRun!: Output<boolean>;
  public readonly quiet!: Output<boolean>;
  public readonly suffix!: Output<string | undefined>;
  public readonly template!: Output<string>;
  public readonly tmpdir!: Output<string | undefined>;

  constructor(name: string, args: MktempArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:tools:Mktemp', name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

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
      .arg(template);

    const command = new Command(name, {
      connection: args.connection,
      create: builder.command,
    }, { parent: this });

    this.command = command;
    this.directory = directory;
    this.dryRun = dryRun;
    this.quiet = quiet;
    this.suffix = suffix;
    this.template = template;
    this.tmpdir = tmpdir;

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
