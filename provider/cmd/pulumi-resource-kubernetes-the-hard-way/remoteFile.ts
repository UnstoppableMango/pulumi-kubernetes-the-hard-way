import { ComponentResource, ComponentResourceOptions, Input, Inputs, Output, interpolate, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import { remote } from '@pulumi/command/types/input';
import { ConstructResult } from '@pulumi/pulumi/provider';

export type InstallArgs = Omit<RemoteFileArgs, 'content'>;

export interface RemoteFileArgs {
  connection: Input<remote.ConnectionArgs>;
  path: Input<string>;
  content: Input<string>;
}

export class RemoteFile extends ComponentResource {
  public readonly command: Command;
  public readonly content: Output<string>;
  public readonly path: Output<string>;

  public get stderr(): Output<string> {
    return this.command.stderr;
  }

  public get stdin(): Output<string | undefined> {
    return this.command.stdin;
  }

  public get stdout(): Output<string> {
    return this.command.stdout;
  }

  constructor(name: string, args: RemoteFileArgs, opts?: ComponentResourceOptions) {
    super('thecluster:index:remoteFile', name, args, opts);

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

    this.registerOutputs({ command, content, path });
  }
}

export async function construct(name: string, inputs: Inputs, options: ComponentResourceOptions): Promise<ConstructResult> {
  const file = new RemoteFile(name, inputs as RemoteFileArgs, options);
  return {
    urn: file.urn,
    state: {
      allowedUses: file.command,
      content: file.content,
      path: file.path,
    },
  };
}
