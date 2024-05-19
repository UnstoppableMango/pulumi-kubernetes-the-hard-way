import { ComponentResource, ComponentResourceOptions, Inputs } from '@pulumi/pulumi';
import * as schema from './schema-types';
import {
  CniPluginsInstall,
  ContainerdInstall,
  ContainerdService,
  ControlPlaneNode,
  CrictlInstall,
  Download,
  EtcdCluster,
  EtcdConfiguration,
  EtcdInstall,
  EtcdService,
  File,
  KubeApiServerInstall,
  KubeControllerManagerInstall,
  KubeProxyInstall,
  KubeProxyService,
  KubeSchedulerInstall,
  KubeadmInstall,
  KubectlInstall,
  KubeletInstall,
  KubeletService,
  RuncInstall,
  StartContainerd,
  StartEtcd,
  StartKubeProxy,
  StartKubelet,
  StaticPod,
  SystemdService,
  WorkerNode,
  WorkerPreRequisites,
} from './remote';
import { ProvisionEtcd } from './remote/provisionEtcd';
import { Certificate, ClusterPki, EncryptionKey, RootCa } from './tls';
import {
  CniBridgePluginConfiguration,
  CniLoopbackPluginConfiguration,
  ContainerdConfiguration,
  KubeProxyConfiguration,
  KubeVipManifest,
  KubeletConfiguration,
} from './config';

const resources: schema.ResourceConstructor = {
  'kubernetes-the-hard-way:config:CniBridgePluginConfiguration': (...args) => new CniBridgePluginConfiguration(...args),
  'kubernetes-the-hard-way:config:CniLoopbackPluginConfiguration': (...args) => new CniLoopbackPluginConfiguration(...args),
  'kubernetes-the-hard-way:config:ContainerdConfiguration': (...args) => new ContainerdConfiguration(...args),
  'kubernetes-the-hard-way:config:KubeletConfiguration': (...args) => new KubeletConfiguration(...args),
  'kubernetes-the-hard-way:config:KubeProxyConfiguration': (...args) => new KubeProxyConfiguration(...args),
  'kubernetes-the-hard-way:config:KubeVipManifest': (...args) => new KubeVipManifest(...args),
  'kubernetes-the-hard-way:remote:ContainerdService': (...args) => new ContainerdService(...args),
  'kubernetes-the-hard-way:remote:ControlPlaneNode': (...args) => new ControlPlaneNode(...args),
  'kubernetes-the-hard-way:remote:Download': (...args) => new Download(...args),
  'kubernetes-the-hard-way:remote:EtcdCluster': (...args) => new EtcdCluster(...args),
  'kubernetes-the-hard-way:remote:EtcdConfiguration': (...args) => new EtcdConfiguration(...args),
  'kubernetes-the-hard-way:remote:EtcdInstall': (...args) => new EtcdInstall(...args),
  'kubernetes-the-hard-way:remote:EtcdService': (...args) => new EtcdService(...args),
  'kubernetes-the-hard-way:remote:File': (...args) => new File(...args),
  'kubernetes-the-hard-way:remote:KubeadmInstall': (...args) => new KubeadmInstall(...args),
  'kubernetes-the-hard-way:remote:KubeApiServerInstall': (...args) => new KubeApiServerInstall(...args),
  'kubernetes-the-hard-way:remote:KubeControllerManagerInstall': (...args) => new KubeControllerManagerInstall(...args),
  'kubernetes-the-hard-way:remote:KubeSchedulerInstall': (...args) => new KubeSchedulerInstall(...args),
  'kubernetes-the-hard-way:remote:CniPluginsInstall': (...args) => new CniPluginsInstall(...args),
  'kubernetes-the-hard-way:remote:ContainerdInstall': (...args) => new ContainerdInstall(...args),
  'kubernetes-the-hard-way:remote:CrictlInstall': (...args) => new CrictlInstall(...args),
  'kubernetes-the-hard-way:remote:KubectlInstall': (...args) => new KubectlInstall(...args),
  'kubernetes-the-hard-way:remote:KubeletInstall': (...args) => new KubeletInstall(...args),
  'kubernetes-the-hard-way:remote:KubeletService': (...args) => new KubeletService(...args),
  'kubernetes-the-hard-way:remote:KubeProxyInstall': (...args) => new KubeProxyInstall(...args),
  'kubernetes-the-hard-way:remote:KubeProxyService': (...args) => new KubeProxyService(...args),
  'kubernetes-the-hard-way:remote:ProvisionEtcd': (...args) => new ProvisionEtcd(...args),
  'kubernetes-the-hard-way:remote:RuncInstall': (...args) => new RuncInstall(...args),
  'kubernetes-the-hard-way:remote:StartContainerd': (...args) => new StartContainerd(...args),
  'kubernetes-the-hard-way:remote:StartEtcd': (...args) => new StartEtcd(...args),
  'kubernetes-the-hard-way:remote:StartKubelet': (...args) => new StartKubelet(...args),
  'kubernetes-the-hard-way:remote:StartKubeProxy': (...args) => new StartKubeProxy(...args),
  'kubernetes-the-hard-way:remote:StaticPod': (...args) => new StaticPod(...args),
  'kubernetes-the-hard-way:remote:SystemdService': (...args) => new SystemdService(...args),
  'kubernetes-the-hard-way:remote:WorkerNode': (...args) => new WorkerNode(...args),
  'kubernetes-the-hard-way:remote:WorkerPreRequisites': (...args) => new WorkerPreRequisites(...args),
  'kubernetes-the-hard-way:tls:Certificate': (...args) => new Certificate(...args),
  'kubernetes-the-hard-way:tls:ClusterPki': (...args) => new ClusterPki(...args),
  'kubernetes-the-hard-way:tls:EncryptionKey': (...args) => new EncryptionKey(...args),
  'kubernetes-the-hard-way:tls:RootCa': (...args) => new RootCa(...args),
};

export function construct(
  name: string,
  type: string,
  inputs: Inputs,
  options: ComponentResourceOptions,
): ComponentResource | undefined {
  const genericResources: Record<string, schema.ConstructComponent> = resources;
  const resource = genericResources[type];
  if (resource === undefined) {
    return undefined;
  }
  return resource(name, inputs, options);
}
