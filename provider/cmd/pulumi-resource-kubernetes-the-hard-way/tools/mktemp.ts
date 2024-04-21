import { ComponentResourceOptions } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.MktempOptsInputs,
  schema.MktempOptsOutputs
>(
  'mktemp',
  (builder, opts) => builder
    .option('--directory', opts.directory)
    .option('--dry-run', opts.dryRun)
    .option('--quiet', opts.quiet)
    .option('--suffix', opts.suffix)
    .option('--tmpdir', opts.tmpdir)
    .arg(opts.template),
  (i) => ({
    directory: tool.mapO(i.directory),
    dryRun: tool.mapO(i.dryRun),
    quiet: tool.mapO(i.quiet),
    suffix: tool.mapO(i.suffix),
    template: tool.mapO(i.template),
    tmpdir: tool.mapO(i.tmpdir),
  }),
);

export class Mktemp extends schema.Mktemp {
  constructor(name: string, args: schema.MktempArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
