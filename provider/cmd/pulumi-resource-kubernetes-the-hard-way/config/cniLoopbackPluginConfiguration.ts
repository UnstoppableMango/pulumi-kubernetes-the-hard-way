import * as YAML from 'yaml';
import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';

export async function getCniLoopabckPluginConfiguration(
  inputs: schema.getCniLoopbackPluginConfigurationInputs,
): Promise<schema.getCniLoopbackPluginConfigurationOutputs> {
  const result: schema.CniLoopbackPluginConfigurationInputs = {
    cniVersion: inputs.cniVersion ?? '1.1.0',
    name: inputs.name ?? 'lo',
    type: inputs.type ?? 'loopback',
  };

  return { result: result as Output<schema.CniLoopbackPluginConfigurationOutputs> };
}

export class CniLoopbackPluginConfiguration extends schema.CniLoopbackPluginConfiguration {
  constructor(name: string, args: schema.CniLoopbackPluginConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const { result } = output(getCniLoopabckPluginConfiguration(args));
    const yaml = result.apply(YAML.stringify);

    this.result = result as unknown as Output<schema.CniLoopbackPluginConfigurationOutputs>;
    this.yaml = yaml;

    this.registerOutputs({ result, yaml });
  }
}
