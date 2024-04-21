import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.TeeOptsInputs,
  schema.TeeOptsOutputs
>(
  'tee',
  (builder, opts) => builder
    .option('--append', opts.append)
    .option('--ignore-interrupts', opts.ignoreInterrupts)
    .option('--output-error', opts.outputError)
    .option('--pipe', opts.pipe)
    .option('--version', opts.version)
    .arg(opts.files),
  (i) => ({
    append: tool.mapO(i.append),
    files: output(i.files),
    ignoreInterrupts: tool.mapO(i.ignoreInterrupts),
    outputError: tool.mapO(i.outputError),
    pipe: tool.mapO(i.pipe),
    version: tool.mapO(i.version),
  }),
);

export class Tee extends schema.Tee {
  constructor(name: string, args: schema.TeeArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
