import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as types from '../schema-types';
import { CommandBuilder } from './commandBuilder';

export class Wget extends types.Wget {
  constructor(name: string, args: types.WgetArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    // TODO: Account for wget version and supported options
    const directoryprefix = output(args.directoryPrefix);
    const httpsOnly = output(args.httpsOnly ?? false);
    const noVerbose = output(args.noVerbose ?? true);
    const outputDocument = output(args.outputDocument);
    const quiet = output(args.quiet ?? false);
    const timestamping = output(args.timestamping ?? false);
    const url = output(args.url);

    const builder = new CommandBuilder('wget')
      .option('--directory-prefix', directoryprefix)
      .option('--https-only', httpsOnly)
      .option('--no-verbose', noVerbose)
      .option('--outputDocument', outputDocument)
      .option('--quiet', quiet)
      .option('--timestamping', timestamping)
      .arg(url);

    const command = new Command(name, {
      connection: args.connection,
      create: builder.command,
    }, { parent: this });

    this.command = command;
    this.directoryPrefix = directoryprefix as Output<string>;
    this.httpsOnly = httpsOnly;
    this.outputDocument = outputDocument as Output<string>;
    this.quiet = quiet;
    this.stderr = command.stderr;
    this.stdin = command.stdin as Output<string>;
    this.stdout = command.stdout;
    this.timestamping = timestamping;
    this.url = url;

    this.registerOutputs({
      command, httpsOnly, outputDocument,
      quiet, timestamping, url,
      stderr: this.stderr,
      stdin: this.stdin,
      stdout: this.stdout,
    });
  }
}
