import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.ChmodOptsInputs,
  schema.ChmodOptsOutputs
>(
  'chmod',
  (builder, opts) => builder
    .option('--changes', opts.changes)
    .option('--no-preserve-root', opts.noPreserveRoot)
    .option('--preserve-root', opts.preserveRoot)
    .option('--quiet', opts.quiet)
    .option('--silent', opts.silent)
    .option('--recursive', opts.recursive)
    .option('--refernce', opts.reference)
    .option('--help', opts.help)
    .option('--version', opts.version)
    .arg(opts.mode)
    .arg(opts.files),
  (i) => ({
    files: output(i.files),
    mode: output(i.mode),
    changes: tool.mapO(i.changes),
    help: tool.mapO(i.help),
    noPreserveRoot: tool.mapO(i.noPreserveRoot),
    preserveRoot: tool.mapO(i.preserveRoot),
    quiet: tool.mapO(i.quiet),
    recursive: tool.mapO(i.recursive),
    reference: tool.mapO(i.reference),
    silent: tool.mapO(i.silent),
    version: tool.mapO(i.version),
  }),
);

export class Chmod extends schema.Chmod {
  constructor(name: string, args: schema.ChmodArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
