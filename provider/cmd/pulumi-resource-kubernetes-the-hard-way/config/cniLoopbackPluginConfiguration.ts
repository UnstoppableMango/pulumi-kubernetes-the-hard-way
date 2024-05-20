import * as YAML from 'yaml';
import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';

export async function getCniLoopabckPluginConfiguration(
  inputs: schema.getCniLoopbackPluginConfigurationInputs,
): Promise<schema.getCniLoopbackPluginConfigurationOutputs> {
  const cniVersion = output(inputs.cniVersion ?? '1.1.0');
  const loopbackName = output(inputs.name ?? 'lo');
  const path = output(inputs.path ?? '/etc/cni/net.d/99-loopback.conf');
  const type = output(inputs.type ?? 'loopback');

  // TODO: Why is this like this?
  const result: schema.CniLoopbackPluginConfigurationInputs = {};

  return { result: result as Output<schema.CniLoopbackPluginConfigurationOutputs> };
}

export class CniLoopbackPluginConfiguration extends schema.CniLoopbackPluginConfiguration {
  constructor(name: string, args: schema.CniLoopbackPluginConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const { result } = output(getCniLoopabckPluginConfiguration(args));
    const yaml = result.apply(YAML.stringify);

    this.result = result as unknown as Output<schema.CniBridgePluginConfigurationOutputs>;
    this.yaml = yaml;

    this.registerOutputs({ result, yaml });
  }
}
