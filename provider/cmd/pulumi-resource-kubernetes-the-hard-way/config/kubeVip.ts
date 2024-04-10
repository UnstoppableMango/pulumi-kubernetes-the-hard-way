import { Output, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';

export async function getKubeVipManifest(inputs: schema.getKubeVipManifestInputs): Promise<schema.getKubeVipManifestOutputs> {
  const { image, kubeconfigPath, version } = inputs;

  const result: schema.PodManifestInputs = {
    spec: {
      containers: [{
        name: 'kube-vip',
        image: image ?? interpolate`ghcr.io/kube-vip/kube-vip:v${version}`,
        imagePullPolicy: 'Always',
        args: ['manager'],
        env: [
          { name: 'vip_arp', value: interpolate`${inputs.vipArp}` },
          { name: 'port', value: interpolate`${inputs.port}` },
          { name: 'vip_interface', value: inputs.vipInterface },
          { name: 'vip_cidr', value: interpolate`${inputs.vipCidr}` },
          { name: 'cp_enable', value: interpolate`${inputs.cpEnable}` },
          { name: 'cp_namespace', value: inputs.cpNamespace },
          { name: 'vip_ddns', value: interpolate`${inputs.vipDdns}` },
          { name: 'svc_enable', value: interpolate`${inputs.svcEnable}` },
          { name: 'vip_leaderelection', value: interpolate`${inputs.vipLeaderElection}` },
          { name: 'vip_leaseduration', value: interpolate`${inputs.vipLeaseDuration}` },
          { name: 'vip_renewdeadline', value: interpolate`${inputs.vipRenewDeadline}` },
          { name: 'vip_retryperiod', value: interpolate`${inputs.vipRetryPeriod}` },
          { name: 'address', value: inputs.address },
        ],
        securityContext: {
          capabilities: {
            add: [
              'NET_ADMIN',
              'NET_RAW',
              'SYS_TIME',
            ],
          },
        },
        volumeMounts: [{
          name: 'kubernetes',
          mountPath: '/etc/kubernetes/admin.conf',
        }],
      }],
      hostAliases: [{
        hostnames: ['kubernetes'],
        ip: '127.0.0.1',
      }],
      hostNetwork: true,
      volumes: [{
        name: 'kubeconfig',
        hostPath: {
          path: kubeconfigPath,
        },
      }],
    },
  };

  // Why does this have to be terrible
  return { result: output(result) as Output<schema.PodManifestOutputs> };
}
