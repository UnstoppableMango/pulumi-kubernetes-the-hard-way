import { ComponentResource, ComponentResourceOptions, Input, Output, interpolate, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import { remote } from '@pulumi/command/types/input';

export type InstallInputs = Omit<FileArgs, 'content'> & {
  name: string;
  options?: ComponentResourceOptions;
};

export interface InstallOutputs {
  result: File;
}

export interface FileArgs {
  connection: Input<remote.ConnectionArgs>;
  path: Input<string>;
  content: Input<string>;
}

export class File extends ComponentResource {
  public readonly command!: Command;
  public readonly content!: Output<string>;
  public readonly path!: Output<string>;
  public readonly stderr!: Output<string>;
  public readonly stdin!: Output<string | undefined>;
  public readonly stdout!: Output<string>;

  constructor(name: string, args: FileArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:remote:File', name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const content = output(args.content);
    const path = output(args.path);

    const command = new Command(name, {
      connection: args.connection,
      stdin: content,
      create: interpolate`tee ${path}`,
      delete: interpolate`rm ${path}`,
    }, { parent: this });

    this.command = command;
    this.content = content;
    this.path = path;
    this.stderr = command.stderr;
    this.stdin = command.stdin;
    this.stdout = command.stdout;

    this.registerOutputs({
      command, content, path,
      stderr: command.stderr,
      stdin: command.stdin,
      stdout: command.stdout,
    });
  }
}

export function install({ name, options, ...rest }: InstallInputs, content: Input<string>): File {
  return new File(name, {
    connection: rest.connection,
    path: rest.path,
    content,
  }, options);
}
