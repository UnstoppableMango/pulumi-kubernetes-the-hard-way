import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.RmOptsInputs,
  schema.RmOptsOutputs
>(
  'rm',
  (builder, opts) => builder
    .option('--dir', opts.dir)
    .option('--force', opts.force)
    .option('--recursive', opts.recursive)
    .option('--verbose', opts.verbose)
    .arg(opts.files),
  (i) => ({
    files: output(i.files),
    dir: tool.mapO(i.dir),
    force: tool.mapO(i.force),
    onDelete: tool.mapO(i.onDelete),
    recursive: tool.mapO(i.recursive),
    verbose: tool.mapO(i.verbose),
  }),
);

export class Rm extends schema.Rm {
  constructor(name: string, args: schema.RmArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
