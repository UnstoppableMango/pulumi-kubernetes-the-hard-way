import { ComponentResource, ComponentResourceOptions, Inputs } from '@pulumi/pulumi';
import * as schema from './schema-types';
import {
  CniPluginsInstall,
  ContainerdInstall,
  CrictlInstall,
  Download,
  EtcdConfiguration,
  EtcdInstall,
  File,
  KubeApiServerInstall,
  KubeControllerManagerInstall,
  KubeProxyInstall,
  KubeSchedulerInstall,
  KubectlInstall,
  KubeletInstall,
  RuncInstall,
  StaticPod,
  SystemdService,
} from './remote';
import { Certificate, ClusterPki, EncryptionKey, RootCa } from './tls';
import { Etcdctl, Mkdir, Mktemp, Mv, Rm, Systemctl, Tar, Tee, Wget } from './tools';
import { KubeVipManifest } from './config';

const resources: schema.ResourceConstructor = {
  'kubernetes-the-hard-way:config:KubeVipManifest': (...args) => new KubeVipManifest(...args),
  'kubernetes-the-hard-way:remote:Download': (...args) => new Download(...args),
  'kubernetes-the-hard-way:remote:EtcdConfiguration': (...args) => new EtcdConfiguration(...args),
  'kubernetes-the-hard-way:remote:EtcdInstall': (...args) => new EtcdInstall(...args),
  'kubernetes-the-hard-way:remote:File': (...args) => new File(...args),
  'kubernetes-the-hard-way:remote:KubeApiServerInstall': (...args) => new KubeApiServerInstall(...args),
  'kubernetes-the-hard-way:remote:KubeControllerManagerInstall': (...args) => new KubeControllerManagerInstall(...args),
  'kubernetes-the-hard-way:remote:KubeSchedulerInstall': (...args) => new KubeSchedulerInstall(...args),
  'kubernetes-the-hard-way:remote:CniPluginsInstall': (...args) => new CniPluginsInstall(...args),
  'kubernetes-the-hard-way:remote:ContainerdInstall': (...args) => new ContainerdInstall(...args),
  'kubernetes-the-hard-way:remote:CrictlInstall': (...args) => new CrictlInstall(...args),
  'kubernetes-the-hard-way:remote:KubectlInstall': (...args) => new KubectlInstall(...args),
  'kubernetes-the-hard-way:remote:KubeletInstall': (...args) => new KubeletInstall(...args),
  'kubernetes-the-hard-way:remote:KubeProxyInstall': (...args) => new KubeProxyInstall(...args),
  'kubernetes-the-hard-way:remote:RuncInstall': (...args) => new RuncInstall(...args),
  'kubernetes-the-hard-way:remote:StaticPod': (...args) => new StaticPod(...args),
  'kubernetes-the-hard-way:remote:SystemdService': (...args) => new SystemdService(...args),
  // @ts-ignore Due to the keypair subclass this refactor is a little harder
  'kubernetes-the-hard-way:tls:Certificate': (...args) => new Certificate(...args),
  // @ts-ignore This refactor loosely depends on the KeyPair refactors
  'kubernetes-the-hard-way:tls:ClusterPki': (...args) => new ClusterPki(...args),
  'kubernetes-the-hard-way:tls:EncryptionKey': (...args) => new EncryptionKey(...args),
  // @ts-ignore Due to the keypair subclass this refactor is a little harder
  'kubernetes-the-hard-way:tls:RootCa': (...args) => new RootCa(...args),
  'kubernetes-the-hard-way:tools:Etcdctl': (...args) => new Etcdctl(...args),
  'kubernetes-the-hard-way:tools:Mkdir': (...args) => new Mkdir(...args),
  'kubernetes-the-hard-way:tools:Mktemp': (...args) => new Mktemp(...args),
  'kubernetes-the-hard-way:tools:Mv': (...args) => new Mv(...args),
  'kubernetes-the-hard-way:tools:Rm': (...args) => new Rm(...args),
  'kubernetes-the-hard-way:tools:Systemctl': (...args) => new Systemctl(...args),
  'kubernetes-the-hard-way:tools:Tar': (...args) => new Tar(...args),
  'kubernetes-the-hard-way:tools:Tee': (...args) => new Tee(...args),
  'kubernetes-the-hard-way:tools:Wget': (...args) => new Wget(...args),
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
