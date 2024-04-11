import { Input, Output, interpolate, output } from '@pulumi/pulumi';
import { core } from '@pulumi/kubernetes/types/input';
import * as schema from '../schema-types';

export async function getKubeVipManifest(inputs: schema.getKubeVipManifestInputs): Promise<schema.getKubeVipManifestOutputs> {
  const { image, kubeconfigPath, version } = inputs;

  // TODO: Validate required props depending on mode
  // TODO: Validate one of bgp or arp is enabled

  const env: core.v1.EnvVar[] = [];

  if (inputs.address) push(env, 'address', inputs.address);
  if (inputs.bgpAs) push(env, 'bgp_as', inputs.bgpAs);
  if (inputs.bgpEnable) push(env, 'bgp_enable', inputs.bgpEnable);
  if (inputs.bgpPeerAddress) push(env, 'bgp_peeraddress', inputs.bgpPeerAddress);
  if (inputs.bgpPeerAs) push(env, 'bgp_peeras', inputs.bgpPeerAs);
  if (inputs.bgpPeerPass) push(env, 'bgp_peerpass', inputs.bgpPeerPass);
  if (inputs.bgpPeers) push(env, 'bgp_peers', inputs.bgpPeers);
  if (inputs.bgpRouterId) push(env, 'bgp_routerid', inputs.bgpRouterId);
  if (inputs.cpEnable) push(env, 'cp_enable', inputs.cpEnable);
  if (inputs.cpNamespace) push(env, 'cp_namespace', inputs.cpNamespace);
  if (inputs.port) push(env, 'port', inputs.port);
  if (inputs.svcEnable) push(env, 'svc_enable', inputs.svcEnable);
  if (inputs.vipArp) push(env, 'vip_arp', inputs.vipArp);
  if (inputs.vipCidr) push(env, 'vip_cidr', inputs.vipCidr);
  if (inputs.vipDdns) push(env, 'vip_ddns', inputs.vipDdns);
  if (inputs.vipInterface) push(env, 'vip_interface', inputs.vipInterface);
  if (inputs.vipLeaseDuration) push(env, 'vip_leaseduration', inputs.vipLeaseDuration);
  if (inputs.vipLeaderElection) push(env, 'vip_leaderelection', inputs.vipLeaderElection);
  if (inputs.vipRenewDeadline) push(env, 'vip_renewdeadline', inputs.vipRenewDeadline);
  if (inputs.vipRetryPeriod) push(env, 'vip_retryperiod', inputs.vipRetryPeriod);

  const result: schema.PodManifestInputs = {
    spec: {
      containers: [{
        name: 'kube-vip',
        image: image ?? interpolate`ghcr.io/kube-vip/kube-vip:v${version}`,
        imagePullPolicy: 'Always',
        args: ['manager'],
        env,
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

function push(env: core.v1.EnvVar[], name: string, value: Input<string | number | boolean>): void {
  env.push({ name, value: interpolate`${value}` });
}
