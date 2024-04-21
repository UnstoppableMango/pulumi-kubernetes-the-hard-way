import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.WgetOptsInputs,
  schema.WgetOptsOutputs
>(
  'wget',
  (builder, opts) => builder
    .option('--directory-prefix', opts.directoryPrefix)
    .option('--https-only', opts.httpsOnly)
    .option('--no-verbose', opts.noVerbose)
    .option('--outputDocument', opts.outputDocument)
    .option('--quiet', opts.quiet)
    .option('--timestamping', opts.timestamping)
    .arg(opts.url),
  (i) => ({
    url: output(i.url),
    directoryPrefix: tool.mapO(i.directoryPrefix),
    httpsOnly: tool.mapO(i.httpsOnly),
    noVerbose: tool.mapO(i.noVerbose),
    outputDocument: tool.mapO(i.outputDocument),
    quiet: tool.mapO(i.quiet),
    timestamping: tool.mapO(i.timestamping),
  }),
);

export class Wget extends schema.Wget {
  constructor(name: string, args: schema.WgetArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
