import { ComponentResource, ComponentResourceOptions, Output, all, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import { MkdirArgs } from 'sdk/tools';
import { CommandBuilder } from './commandBuilder';

export class Mkdir extends ComponentResource {
  public readonly command!: Command;
  public readonly directory!: Output<string>;
  public readonly parents!: Output<boolean>;
  public readonly removeOnDelete!: Output<boolean>;

  constructor(name: string, args: MkdirArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard:tools:Mkdir', name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const directory = output(args.directory);
    const parents = output(args.parents ?? false);
    const deleteCmd = all([args.removeOnDelete, directory])
      .apply(([remove, dir]) => remove ? `rm -rf ${dir}` : '');

    const builder = new CommandBuilder('mkdir')
      .option('--parents', parents)
      .arg(directory);

    const command = new Command(name, {
      connection: args.connection,
      create: builder.command,
      delete: deleteCmd,
    }, { parent: this });

    this.command = command;
    this.directory = directory;
    this.parents = parents;

    this.registerOutputs({ command, directory });
  }
}
