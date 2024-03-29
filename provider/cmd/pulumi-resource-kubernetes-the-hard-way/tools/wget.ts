import { ComponentResource, ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import { WgetArgs } from '../sdk';
import { CommandBuilder } from './commandBuilder';

export class Wget extends ComponentResource {
  public readonly command!: Command;
  public readonly directoryPrefix!: Output<string | undefined>;
  public readonly httpsOnly!: Output<boolean>;
  public readonly noVerbose!: Output<boolean>;
  public readonly outputDocument!: Output<string | undefined>;
  public readonly quiet!: Output<boolean>;
  public readonly stderr!: Output<string>;
  public readonly stdin!: Output<string | undefined>;
  public readonly stdout!: Output<string>;
  public readonly timestamping!: Output<boolean>;
  public readonly url!: Output<string>;

  constructor(name: string, args: WgetArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:tools:Wget', name, args, opts);

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
    this.directoryPrefix = directoryprefix;
    this.httpsOnly = httpsOnly;
    this.outputDocument = outputDocument;
    this.quiet = quiet;
    this.stderr = command.stderr;
    this.stdin = command.stdin;
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
