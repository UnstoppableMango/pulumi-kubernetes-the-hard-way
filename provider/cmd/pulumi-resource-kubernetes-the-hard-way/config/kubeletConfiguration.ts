import * as YAML from 'yaml';
import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { getKubeletConfiguration } from './getKubeletConfiguration';

export class KubeletConfiguration extends schema.KubeletConfiguration {
  constructor(name: string, args: schema.KubeletConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const { result } = output(getKubeletConfiguration(args));
    const yaml = result.apply(YAML.stringify);

    this.result = result as unknown as Output<schema.KubeletConfigurationOutputs>;
    this.yaml = yaml

    this.registerOutputs({ result, yaml });
  }
}
