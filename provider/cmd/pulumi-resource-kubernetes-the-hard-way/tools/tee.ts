import { ComponentResource, ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import { remote } from '@pulumi/command/types/input';
import * as schema from '../schema-types';
import { CommandBuilder, toArray } from './commandBuilder';

export type TeeArgs = schema.TeeArgs & {
  files?: Input<string> | Input<Input<string>[]>;
}

export class Tee extends schema.Tee {
  constructor(name: string, args: TeeArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const append = output(args.append);
    const connection = output(args.connection);
    const files = output(args.files ?? []).apply(toArray); // TODO: Can we get types happy without the `toArray`?
    const ignoreInterrupts = output(args.ignoreInterrupts);
    const lifecycle = args.lifecycle ?? 'create';
    const outputError = output(args.outputError);
    const pipe = output(args.pipe);
    const stdin = output(args.stdin);
    const version = output(args.version);

    const builder = new CommandBuilder('tee')
      .option('--append', append)
      .option('--ignore-interrupts', ignoreInterrupts)
      .option('--output-error', outputError)
      .option('--pipe', pipe)
      .option('--version', version)
      .arg(files);

    const command = new Command(name, {
      connection,
      stdin,
      [lifecycle]: builder.command,
    }, { parent: this });

    this.command = command;
    this.files = files;
    this.stderr = command.stderr;
    this.stdin = command.stdin as Output<string>;
    this.stdout = command.stdout;

    this.registerOutputs({
      command,
      files,
      stderr: this.stderr,
      stdin: this.stdin,
      stdout: this.stdout,
    });
  }
}
