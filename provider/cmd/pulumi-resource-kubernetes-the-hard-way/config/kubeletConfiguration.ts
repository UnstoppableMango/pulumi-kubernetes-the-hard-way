import { ComponentResourceOptions } from '@pulumi/pulumi';
import * as schema from '../schema-types';

export class KubeletConfiguration extends schema.KubeletConfiguration {
  constructor(name: string, args: schema.KubeletConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    
  }
}
