import { ComponentResourceOptions, Input, Output, interpolate, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as schema from '../schema-types';

export type InstallInputs = Omit<schema.FileArgs, 'content'> & {
  name: string;
  options?: ComponentResourceOptions;
};

export interface InstallOutputs {
  result: File;
}

export class File extends schema.File {
  constructor(name: string, args: schema.FileArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

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
    this.stdin = command.stdin as Output<string>;
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
