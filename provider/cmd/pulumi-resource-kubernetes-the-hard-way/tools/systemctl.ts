import { ComponentResourceOptions } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.SystemctlOptsInputs,
  schema.SystemctlOptsOutputs
>(
  'systemctl',
  (builder, opts) => {
    const b = builder.arg(opts.command);

    // TODO: Little bit smarter check here
    if (opts.command !== 'daemon-reload') {
      b.arg(opts.unit);
    }

    return b.option('--now', opts.now);
  },
  (i) => ({
    command: i.command,
    unit: tool.mapO(i.unit),
    pattern: tool.mapO(i.pattern),
    now: tool.mapO(i.now),
  }),
);

export class Systemctl extends schema.Systemctl {
  constructor(name: string, args: schema.SystemctlArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
