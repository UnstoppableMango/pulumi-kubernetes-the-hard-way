import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { getKubeVipManifest } from './getKubeVipManifest';

export class KubeVipManifest extends schema.KubeVipManifest {
  constructor(name: string, args: schema.KubeVipManifestArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const { result } = output(getKubeVipManifest(args));

    // Ugh... can this be better?
    this.result = result as Output<schema.PodManifestOutputs>;

    this.registerOutputs({ result });
  }
}
