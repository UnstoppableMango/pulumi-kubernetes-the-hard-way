import { Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';

export async function getKubeProxyConfiguration(inputs: schema.getKubeProxyConfigurationInputs): Promise<schema.getKubeProxyConfigurationOutputs> {
  const result: schema.KubeProxyConfigurationInputs = {
    apiVersion: 'kubeproxy.config.k8s.io/v1alpha1',
    kind: 'KubeProxyConfiguration',
    clusterCIDR: output(inputs.clusterCIDR),
    clientConnection: {
      kubeconfig: output(inputs.kubeconfig),
    },
    mode: output(inputs.mode ?? 'iptables'),
  };

  return { result: result as Output<schema.KubeProxyConfigurationOutputs> };
}
