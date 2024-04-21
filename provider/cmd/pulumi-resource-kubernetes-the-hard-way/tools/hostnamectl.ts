import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.HostnamectlOptsInputs,
  schema.HostnamectlOptsOutputs
>(
  'hostnamectl',
  (builder, opts) => builder
    .option('--help', opts.help)
    .option('--json', opts.json)
    .option('--machine', opts.machine)
    .option('--no-ask-password', opts.noAskPassword)
    .option('--pretty', opts.pretty)
    .option('--static', opts.static)
    .option('--transient', opts.transient)
    .option('--version', opts.version)
    .arg(opts.command)
    .arg(opts.arg),
  (i) => ({
    command: output(i.command),
    arg: tool.mapO(i.arg),
    help: tool.mapO(i.help),
    host: tool.mapO(i.host),
    json: tool.mapO(i.json),
    machine: tool.mapO(i.machine),
    noAskPassword: tool.mapO(i.noAskPassword),
    pretty: tool.mapO(i.pretty),
    static: tool.mapO(i.static),
    transient: tool.mapO(i.transient),
    version: tool.mapO(i.version),
  }),
);

export class Hostnamectl extends schema.Hostnamectl {
  constructor(name: string, args: schema.HostnamectlArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
