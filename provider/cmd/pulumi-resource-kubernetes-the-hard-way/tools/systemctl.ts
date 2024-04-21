import { ComponentResourceOptions, output } from '@pulumi/pulumi';
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
      builder.arg(opts.unit);
    }

    return b;
  },
  (i) => ({
    command: i.command,
    unit: output(i.unit),
    pattern: tool.mapO(i.pattern),
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
