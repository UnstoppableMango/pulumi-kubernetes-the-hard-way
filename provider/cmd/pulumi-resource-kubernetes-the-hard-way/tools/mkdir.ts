import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.MkdirOptsInputs,
  schema.MkdirOptsOutputs
>(
  'mkdir',
  (builder, opts) => builder
    .option('--parents', opts.parents)
    .arg(opts.directory),
  (i) => ({
    directory: output(i.directory),
    parents: tool.mapO(i.parents),
  }),
);

export class Mkdir extends schema.Mkdir {
  constructor(name: string, args: schema.MkdirArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
