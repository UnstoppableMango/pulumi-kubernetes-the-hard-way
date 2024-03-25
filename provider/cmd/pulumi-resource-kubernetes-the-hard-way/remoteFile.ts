import { ComponentResource, ComponentResourceOptions, Input, Inputs, Output, interpolate, output } from '@pulumi/pulumi';
import { ConstructResult } from '@pulumi/pulumi/provider';
import { Command } from '@pulumi/command/remote';
import { remote } from '@pulumi/command/types/input';

export type InstallArgs = Omit<RemoteFileArgs, 'content'> & {
  name: string;
  opts?: ComponentResourceOptions;
};

export interface RemoteFileArgs {
  connection: Input<remote.ConnectionArgs>;
  path: Input<string>;
  content: Input<string>;
}

export class RemoteFile extends ComponentResource {
  public readonly command: Command;
  public readonly content: Output<string>;
  public readonly path: Output<string>;
  public readonly stderr: Output<string>;
  public readonly stdin: Output<string | undefined>;
  public readonly stdout: Output<string>;

  constructor(name: string, args: RemoteFileArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:index:RemoteFile', name, args, opts);

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

export function install({ name, opts, ...rest }: InstallArgs, content: Input<string>): RemoteFile {
  return new RemoteFile(name, {
    connection: rest.connection,
    path: rest.path,
    content,
  }, opts);
}

export async function construct(
  name: string,
  inputs: Inputs,
  options: ComponentResourceOptions,
): Promise<ConstructResult> {
  const file = new RemoteFile(name, inputs as RemoteFileArgs, options);
  return {
    urn: file.urn,
    state: {
      command: file.command,
      content: file.content,
      path: file.path,
      stderr: file.stderr,
      stdin: file.stdin,
      stdout: file.stdout,
    },
  };
}
