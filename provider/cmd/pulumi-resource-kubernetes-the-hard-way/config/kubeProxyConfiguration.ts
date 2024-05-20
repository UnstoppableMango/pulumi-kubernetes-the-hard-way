import * as YAML from 'yaml';
import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { getKubeProxyConfiguration } from './getKubeProxyConfiguration';


export class KubeProxyConfiguration extends schema.KubeProxyConfiguration {
  constructor(name: string, args: schema.KubeProxyConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const { result } = output(getKubeProxyConfiguration(args));
    const yaml = result.apply(YAML.stringify);

    this.result = result as unknown as Output<schema.KubeProxyConfigurationOutputs>;
    this.yaml = yaml

    this.registerOutputs({ result, yaml });
  }
}
