import { ComponentResourceOptions, Input, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as types from '../schema-types';
import { CommandBuilder, toArray } from './commandBuilder';

export type MvArgs = types.MvArgs & {
  source: Input<string | Input<string>[]>;
};

export class Mv extends types.Mv {
  constructor(name: string, args: MvArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const backup = output(args.backup ?? false);
    const context = output(args.context ?? false);
    const control = output(args.control);
    const dest = output(args.dest);
    const directory = output(args.directory);
    const force = output(args.force ?? false);
    const noClobber = output(args.noClobber ?? false);
    const noTargetDirectory = output(args.noTargetDirectory ?? false);
    const source = output(args.source).apply(toArray);
    const stripTrailingSlashes = output(args.stripTrailingSlashes ?? false);
    const suffix = output(args.suffix);
    const targetDirectory = output(args.targetDirectory);
    const update = output(args.update ?? false);
    const verbose = output(args.verbose ?? false);

    const builder = new CommandBuilder('mv')
      .option('--context', context)
      .option('--force', force)
      .option('--no-clobber', noClobber)
      .option('--no-target-directory', noTargetDirectory)
      .option('--strip-trailing-slashed', stripTrailingSlashes)
      .option('--suffix', suffix)
      .option('--target-directory', targetDirectory)
      .option('--update', update)
      .option('--verbose', verbose)
      .arg(source)
      .arg(args.dest)
      .arg(args.directory);

    if (backup) {
      if (args.control) {
        builder.option('--backup', control);
      } else {
        builder.option('-b', backup);
      }
    }

    const command = new Command(name, {
      connection: args.connection,
      create: builder.command,
    }, { parent: this });

    this.backup = backup;
    this.context = context;
    this.force = force;
    this.noClobber = noClobber;
    this.noTargetDirectory = noTargetDirectory;
    this.source = source;
    this.stripTrailingSlashes = stripTrailingSlashes;
    this.update = update;
    this.verbose = verbose;

    if (args.control) this.control = output(args.control);
    if (args.dest) this.dest = output(args.dest);
    if (args.directory) this.directory = output(args.directory);
    if (args.suffix) this.suffix = output(args.suffix);
    if (args.targetDirectory) this.targetDirectory = output(args.targetDirectory);

    this.registerOutputs({
      backup,
      command,
      context,
      control,
      dest,
      directory,
      force,
      noClobber,
      noTargetDirectory,
      source,
      stripTrailingSlashes,
      suffix,
      targetDirectory,
      update,
      verbose,
    });
  }
}
