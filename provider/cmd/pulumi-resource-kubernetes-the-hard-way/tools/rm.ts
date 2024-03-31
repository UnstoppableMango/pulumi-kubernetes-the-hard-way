import { ComponentResource, ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import { remote } from '@pulumi/command/types/input';
import { CommandBuilder, toArray } from './commandBuilder';

export interface RmArgs {
  connection: Input<remote.ConnectionArgs>;
  dir?: Input<boolean>;
  files: Input<string | Input<string>[]>;
  force?: Input<boolean>;
  onDelete?: boolean;
  recursive?: Input<boolean>;
  verbose?: Input<boolean>;
}

export class Rm extends ComponentResource {
  public readonly command!: Command;
  public readonly dir!: Output<boolean>;
  public readonly files!: Output<string[]>;
  public readonly force!: Output<boolean>;
  public readonly onDelete!: Output<boolean>;
  public readonly recursive!: Output<boolean>;
  public readonly verbose!: Output<boolean>;

  constructor(name: string, args: RmArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:tools:Rm', name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const dir = output(args.dir ?? false);
    const files = output(args.files).apply(toArray);
    const force = output(args.force ?? false);
    const onDelete = args.onDelete ?? false;
    const recursive = output(args.recursive ?? false);
    const verbose = output(args.verbose ?? false);

    const builder = new CommandBuilder('rm')
      .option('--dir', dir)
      .option('--force', force)
      .option('--recursive', recursive)
      .option('--verbose', verbose)
      .arg(files);

    const command = new Command(name, {
      connection: args.connection,
      create: !onDelete ? builder.command : undefined,
      delete: onDelete ? builder.command : undefined,
    }, { parent: this });

    this.command = command;
    this.dir = dir;
    this.files = files;
    this.force = force;
    this.onDelete = output(onDelete);
    this.recursive = recursive;
    this.verbose = verbose;

    this.registerOutputs({
      command,
      dir,
      files,
      force,
      onDelete,
      recursive,
      verbose,
    });
  }
}
