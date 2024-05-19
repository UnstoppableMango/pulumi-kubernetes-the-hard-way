import * as YAML from 'yaml';
import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { getKubeconfig } from '../config';

export class Kubeconfig extends schema.Kubeconfig {
  constructor(name: string, args: schema.KubeconfigArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const result = output(getKubeconfig(args));

    this.result = result as unknown as Output<schema.KubeconfigOutputs>;
    const yaml = result.apply(YAML.stringify);

    this.registerOutputs({ result, yaml });
  }
}
