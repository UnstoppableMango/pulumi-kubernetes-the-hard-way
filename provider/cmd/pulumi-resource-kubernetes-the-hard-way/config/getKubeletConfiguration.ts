import { Output } from '@pulumi/pulumi';
import * as schema from '../schema-types';

export async function getKubeletConfiguration(inputs: schema.getKubeletConfigurationInputs): Promise<schema.getKubeletConfigurationOutputs> {
  const result: schema.KubeletConfigurationInputs = {
    apiVersion: 'kubelet.config.k8s.io/v1beta1',
    kind: 'KubeletConfiguration',
    authentication: {
      anonymous: {
        enabled: inputs.anonymous ?? false,
      },
      webhook: {
        enabled: inputs.webhook ?? true,
      },
      x509: {
        clientCAFile: inputs.clientCAFile ?? '/var/lib/kubelet/ca.crt',
      },
    },
    authorization: {
      mode: inputs.authorizationMode ?? 'Webhook',
    },
    clusterDomain: inputs.clusterDomain ?? 'cluster.local',
    clusterDNS: inputs.clusterDNS ?? ['10.32.0.10'],
    cgroupDriver: inputs.cgroupDriver ?? 'systemd',
    containerRuntimeEndpoint: inputs.containerRuntimeEndpoint ?? 'unix:///var/run/containerd/containerd.sock',
    podCIDR: inputs.podCIDR,
    resolvConf: inputs.resolvConf ?? '/etc/resolv.conf',
    runtimeRequestTimeout: inputs.runtimeRequestTimeout ?? '15m',
    tlsCertFile: inputs.tlsCertFile ?? '/var/lib/kubelet/kubelet.crt',
    tlsPrivateKeyFile: inputs.tlsPrivateKeyFile ?? '/var/lib/kubelet/kubelet.key',
  };

  return { result: result as Output<schema.KubeletConfigurationOutputs> };
}
