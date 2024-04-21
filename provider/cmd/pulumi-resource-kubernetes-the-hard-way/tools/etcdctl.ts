import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.EtcdctlOptsInputs,
  schema.EtcdctlOptsOutputs
>(
  'etcdctl',
  (builder, opts) => builder
    .arg(opts.commands)
    .option('--ca-cert', opts.caCert)
    .option('--cert', opts.cert)
    .option('--endpoints', opts.endpoints)
    .option('--key', opts.key),
  (i) => ({
    commands: output(i.commands),
    caCert: tool.mapO(i.caCert),
    cert: tool.mapO(i.cert),
    endpoints: tool.mapO(i.endpoints),
    key: tool.mapO(i.key),
  }),
);

export class Etcdctl extends schema.Etcdctl {
  constructor(name: string, args: schema.EtcdctlArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
