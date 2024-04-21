import { ComponentResourceOptions } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.SedOptsInputs,
  schema.SedOptsOutputs
>(
  'sed',
  (builder, opts) => builder
    .option('--debug', opts.debug)
    .option('--follow-symlinks', opts.followSymlinks)
    .option('--help', opts.help)
    .option('--in-place', opts.inPlace)
    .option('--line-length', opts.lineLength)
    .option('--null-data', opts.nullData)
    .option('--posix', opts.posix)
    .option('--quiet', opts.quiet)
    .option('--regexp-extended', opts.regexpExtended)
    .option('--sandbox', opts.sandbox)
    .option('--script', opts.script)
    .option('--separate', opts.separate)
    .option('--silent', opts.silent)
    .option('--unbuffered', opts.unbuffered)
    .option('--version', opts.version)
    .arg(opts.inputFiles),
  (i) => ({
    debug: tool.mapO(i.debug),
    expressions: tool.mapO(i.expressions),
    files: tool.mapO(i.files),
    followSymlinks: tool.mapO(i.followSymlinks),
    help: tool.mapO(i.help),
    inPlace: tool.mapO(i.inPlace),
    inputFiles: tool.mapO(i.inputFiles),
    lineLength: tool.mapO(i.lineLength),
    nullData: tool.mapO(i.nullData),
    posix: tool.mapO(i.posix),
    quiet: tool.mapO(i.quiet),
    regexpExtended: tool.mapO(i.regexpExtended),
    sandbox: tool.mapO(i.sandbox),
    script: tool.mapO(i.script),
    separate: tool.mapO(i.separate),
    silent: tool.mapO(i.silent),
    unbuffered: tool.mapO(i.unbuffered),
    version: tool.mapO(i.version),
  }),
);

export class Sed extends schema.Sed {
  constructor(name: string, args: schema.SedArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
