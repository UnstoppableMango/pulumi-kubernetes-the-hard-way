import { ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as schema from '../schema-types';
import { CommandBuilder, toArray } from './commandBuilder';

type SedArgs = schema.SedArgs & {
  files?: Input<string | Input<string>[]>;
  inputFiles?: Input<string | Input<string>[]>;
};

export class Sed extends schema.Sed {
  constructor(name: string, args: SedArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const binaryPath = output(args.binaryPath ?? 'sed');
    const connection = output(args.connection);
    const debug = output(args.debug ?? false);
    const environment = output(args.environment ?? {});
    const expressions = output(args.expressions) as Output<string | Output<string>[]>; // TODO
    const files = output(args.files ?? []).apply(toArray); // TODO
    const followSymlinks = output(args.followSymlinks ?? false);
    const help = output(args.help ?? false);
    const inPlace = output(args.inPlace);
    const inputFiles = output(args.inputFiles ?? []).apply(toArray);
    const lifecycle = args.lifecycle ?? 'create';
    const lineLength = output(args.lineLength ?? false);
    const nullData = output(args.nullData ?? false);
    const posix = output(args.posix ?? false);
    const quiet = output(args.quiet ?? false);
    const regexpExtended = output(args.regexpExtended ?? false);
    const sandbox = output(args.sandbox ?? false);
    const script = output(args.script);
    const separate = output(args.separate ?? false);
    const silent = output(args.silent ?? false);
    const stdin = output(args.stdin);
    const triggers = output(args.triggers ?? []);
    const unbuffered = output(args.unbuffered ?? false);
    const version = output(args.version ?? false);

    const builder = new CommandBuilder(binaryPath)
      .option('--debug', debug)
      .option('--follow-symlinks', followSymlinks)
      .option('--help', help)
      .option('--in-place', inPlace)
      .option('--line-length', lineLength)
      .option('--null-data', nullData)
      .option('--posix', posix)
      .option('--quiet', quiet)
      .option('--regexp-extended', regexpExtended)
      .option('--sandbox', sandbox)
      .option('--script', script)
      .option('--separate', separate)
      .option('--silent', silent)
      .option('--unbuffered', unbuffered)
      .option('--version', version)
      .arg(inputFiles);

    const command = new Command(name, {
      connection,
      environment,
      triggers,
      stdin: args.stdin,
      [lifecycle]: builder.command,
    }, { parent: this });

    this.binaryPath = binaryPath;
    this.command = command;
    this.connection = connection;
    this.debug = debug;
    this.environment = environment;
    this.expressions = expressions;
    this.files = files;
    this.followSymlinks = followSymlinks;
    this.help = help;
    this.inPlace = inPlace as Output<string>;
    this.inputFiles = inputFiles;
    this.lifecycle = lifecycle;
    this.lineLength = lineLength;
    this.nullData = nullData;
    this.posix = posix;
    this.quiet = quiet;
    this.regexpExtended = regexpExtended;
    this.sandbox = sandbox;
    this.script = script as Output<string>;
    this.separate = separate;
    this.silent = silent;
    this.stdin = stdin as Output<string>;
    this.triggers = triggers;
    this.unbuffered = unbuffered;
    this.version = version;

    this.registerOutputs({
      binaryPath,
      command,
      connection,
      debug,
      environment,
      expressions,
      files,
      followSymlinks,
      help,
      inPlace,
      inputFiles,
      lifecycle,
      lineLength,
      nullData,
      posix,
      quiet,
      regexpExtended,
      sandbox,
      script,
      separate,
      silent,
      stdin,
      triggers,
      unbuffered,
      version,
    });
  }
}
