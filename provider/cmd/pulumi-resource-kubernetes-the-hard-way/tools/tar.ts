import { ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as types from '../schema-types';
import { CommandBuilder, toArray } from './commandBuilder';

export type TarArgs = types.TarArgs & {
  files?: Input<string> | Input<Input<string>[]>;
}

export class Tar extends types.Tar {
  constructor(name: string, args: TarArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const archive = output(args.archive);
    const directory = output(args.directory);
    const extract = output(args.extract ?? true); // Is this a sane default?
    const files = output(args.files ?? []).apply(toArray); // TODO: Can we get types happy without the `toArray`?
    const gzip = output(args.gzip ?? false);
    const stripComponents = output(args.stripComponents);

    const builder = new CommandBuilder('tar')
      .option('--extract', extract)
      .option('--gzip', gzip)
      .option('--file', archive)
      .option('--directory', directory)
      .option('--strip-components', stripComponents.apply(x => x?.toString()))
      .arg(files);

    const command = new Command(name, {
      connection: args.connection,
      create: builder.command,
      // TODO: Should we clean anything up on delete/update?
    }, { parent: this });

    this.archive = archive;
    this.command = command;
    this.directory = directory as Output<string>;
    this.extract = extract;
    this.files = files;
    this.stderr = command.stderr;
    this.stdin = command.stdin as Output<string>;
    this.stdout = command.stdout;
    this.stripComponents = stripComponents as Output<number>;

    this.registerOutputs({
      archive, command, directory, extract,
      files, stripComponents,
      stderr: this.stderr,
      stdin: this.stdin,
      stdout: this.stdout,
    });
  }
}
