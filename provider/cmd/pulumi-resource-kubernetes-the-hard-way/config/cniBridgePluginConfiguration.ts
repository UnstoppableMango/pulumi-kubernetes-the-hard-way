import * as YAML from 'yaml';
import { ComponentResourceOptions, Output, all, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';

export async function getCniBridgePluginConfiguration(
  inputs: schema.getCniBridgePluginConfigurationInputs,
): Promise<schema.getCniBridgePluginConfigurationOutputs> {
  const ipam = output(inputs.ipam);
  const subnet = output(inputs.subnet);

  const result: schema.CniBridgePluginConfigurationInputs = {
    cniVersion: '1.0.0',
    bridge: inputs.bridge ?? 'cni0',
    ipam: all([ipam, subnet]).apply(x => ipamDefaults(...x)),
    ipMasq: inputs.ipMasq ?? true,
    isGateway: inputs.isGateway ?? true,
    subnet,
  };

  return { result: result as Output<schema.CniBridgePluginConfigurationOutputs> };
}

export class CniBridgePluginConfiguration extends schema.CniBridgePluginConfiguration {
  constructor(name: string, args: schema.CniBridgePluginConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const { result } = output(getCniBridgePluginConfiguration(args));
    const yaml = result.apply(YAML.stringify);

    this.result = result as unknown as Output<schema.CniBridgePluginConfigurationOutputs>;
    this.yaml = yaml;

    this.registerOutputs({ result, yaml });
  }
}

function ipamDefaults(ipam?: schema.CniBridgeIpamInputs, subnet?: string): schema.CniBridgeIpamOutputs {
  const required = (x?: string): string => {
    if (!x) throw new Error('');
    return x;
  }

  return {
    type: output(ipam?.type ?? 'host-local'),
    ranges: output(ipam?.ranges).apply(x => x ?? [{
      subnet: required(subnet),
    }]),
    routes: output(ipam?.routes).apply(x => x ?? [{
      dst: '0.0.0.0/0',
    }]),
  };
}
