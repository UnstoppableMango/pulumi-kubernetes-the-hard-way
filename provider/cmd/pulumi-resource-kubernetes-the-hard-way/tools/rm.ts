import { ComponentResourceOptions, Input, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as schema from '../schema-types';
import { CommandBuilder, toArray } from './commandBuilder';

export type RmArgs = schema.RmArgs & {
  files: Input<string | Input<string>[]>;
};

export class Rm extends schema.Rm {
  constructor(name: string, args: RmArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

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
