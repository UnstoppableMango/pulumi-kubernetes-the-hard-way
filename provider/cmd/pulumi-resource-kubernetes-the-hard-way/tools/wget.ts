import { ComponentResource, ComponentResourceOptions, Input, Output, all, interpolate, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import { remote } from '@pulumi/command/types/input';

export interface WgetArgs {
  connection: Input<remote.ConnectionArgs>;
  directoryPrefix?: Input<string>;
  httpsOnly?: Input<boolean>;
  noVerbose?: Input<boolean>;
  outputDocument?: Input<string>;
  quiet?: Input<boolean>;
  timestamping?: Input<boolean>;
  url: Input<string>;
}

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
    super('kubernetes-the-hard-way:tools:wget', name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const directoryprefix = output(args.directoryPrefix);
    const httpsOnly = output(args.httpsOnly ?? true);
    const noVerbose = output(args.noVerbose ?? true);
    const outputDocument = output(args.outputDocument);
    const quiet = output(args.quiet ?? false);
    const timestamping = output(args.timestamping ?? true);
    const url = output(args.url);

    const options: Output<string> = all([
      directoryprefix,
      httpsOnly,
      noVerbose,
      outputDocument,
      quiet,
      timestamping,
    ]).apply(([directoryPrefix, httpsOnly, noVerbose, outputDocument, quiet, timestamping]) => {
      const options: string[] = [];

      if (directoryPrefix) options.push(`--directory-prefix '${directoryPrefix}'`);
      if (httpsOnly) options.push('--https-only');
      if (noVerbose) options.push('--no-verbose');
      if (outputDocument) options.push(`--output-document '${outputDocument}'`);
      if (quiet) options.push('--quiet');
      if (timestamping ?? true) options.push('--timestamping');

      return options.join(' ');
    });

    const command = new Command(name, {
      connection: args.connection,
      create: interpolate`wget ${options} '${url}'`,
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
