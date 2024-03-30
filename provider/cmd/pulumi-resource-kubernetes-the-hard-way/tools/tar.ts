import { ComponentResource, ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import { remote } from '@pulumi/command/types/input';
import { CommandBuilder, toArray } from './commandBuilder';

export interface TarArgs {
  archive?: Input<string>;
  connection: Input<remote.ConnectionArgs>;
  directory?: Input<string>;
  extract?: Input<boolean>;
  files?: Input<string> | Input<Input<string>[]>;
  gzip?: Input<boolean>;
}

export class Tar extends ComponentResource {
  public readonly archive!: Output<string | undefined>;
  public readonly command!: Command;
  public readonly directory!: Output<string | undefined>;
  public readonly extract!: Output<boolean>;
  public readonly files!: Output<string[]>;
  public readonly stderr!: Output<string>;
  public readonly stdin!: Output<string | undefined>;
  public readonly stdout!: Output<string>;

  constructor(name: string, args: TarArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:tools:Tar', name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const archive = output(args.archive);
    const directory = output(args.directory);
    const extract = output(args.extract ?? true); // Is this a sane default?
    const files = output(args.files ?? []).apply(toArray); // TODO: Can we get types happy without the `toArray`?
    const gzip = output(args.gzip ?? false);

    const builder = new CommandBuilder('tar')
      .option('--extract', extract)
      .option('--file', archive)
      .option('--directory', directory)
      .option('--gzip', gzip)
      .arg(files);

    const command = new Command(name, {
      connection: args.connection,
      create: builder.command,
      // TODO: Should we clean anything up on delete/update?
    }, { parent: this });

    this.archive = archive;
    this.command = command;
    this.directory = directory;
    this.extract = extract;
    this.files = files;
    this.stderr = command.stderr;
    this.stdin = command.stdin;
    this.stdout = command.stdout;

    this.registerOutputs({
      archive, command, extract,
      stderr: this.stderr,
      stdin: this.stdin,
      stdout: this.stdout,
    });
  }
}
