import { ComponentResource, ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import { remote } from '@pulumi/command/types/input';
import { CommandBuilder, toArray } from './commandBuilder';

export interface MvArgs {
  backup?: boolean;
  connection: Input<remote.ConnectionArgs>;
  context?: Input<boolean>;
  control?: Input<string>;
  dest?: Input<string>;
  directory?: Input<string>;
  force?: Input<boolean>;
  noClobber?: Input<boolean>;
  noTargetDirectory?: Input<boolean>;
  source: Input<string | Input<string>[]>;
  stripTrailingSlashes?: Input<boolean>;
  suffix?: Input<string>;
  targetDirectory?: Input<string>;
  update?: Input<boolean>;
  verbose?: Input<boolean>;
}

export class Mv extends ComponentResource {
  public readonly backup!: Output<boolean>;
  public readonly command!: Output<Command>;
  public readonly context!: Output<boolean>;
  public readonly control!: Output<string | undefined>;
  public readonly dest!: Output<string | undefined>;
  public readonly directory!: Output<string | undefined>;
  public readonly force!: Output<boolean>;
  public readonly noClobber!: Output<boolean>;
  public readonly noTargetDirectory!: Output<boolean>;
  public readonly source!: Output<string[]>;
  public readonly stripTrailingSlashes!: Output<boolean>;
  public readonly suffix!: Output<string | undefined>;
  public readonly targetDirectory!: Output<string | undefined>;
  public readonly update!: Output<boolean>;
  public readonly verbose!: Output<boolean>;

  constructor(name: string, args: MvArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:tools:Mv', name, args, opts);

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
    this.control = control;
    this.dest = dest;
    this.directory = directory;
    this.force = force;
    this.noClobber = noClobber;
    this.noTargetDirectory = noTargetDirectory;
    this.source = source;
    this.stripTrailingSlashes = stripTrailingSlashes;
    this.suffix = suffix;
    this.targetDirectory = targetDirectory;
    this.update = update;
    this.verbose = verbose;

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
