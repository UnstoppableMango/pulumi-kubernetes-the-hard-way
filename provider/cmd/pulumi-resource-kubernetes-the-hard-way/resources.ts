import { ComponentResource, ComponentResourceOptions, Inputs } from '@pulumi/pulumi';
import * as schema from './schema-types';
import { Download, EtcdInstall, File, KubeApiServerInstall, KubeControllerManagerInstall, KubeSchedulerInstall, SystemdService } from './remote';
import { Certificate, ClusterPki, EncryptionKey, RootCa } from './tls';
import { Etcdctl, Mkdir, Mktemp, Mv, Rm, Systemctl, Tar, Tee, Wget } from './tools';

const resources: schema.ResourceConstructor = {
  'kubernetes-the-hard-way:remote:Download': (...args) => new Download(...args),
  'kubernetes-the-hard-way:remote:EtcdInstall': (...args) => new EtcdInstall(...args),
  'kubernetes-the-hard-way:remote:File': (...args) => new File(...args),
  'kubernetes-the-hard-way:remote:KubeApiServerInstall': (...args) => new KubeApiServerInstall(...args),
  'kubernetes-the-hard-way:remote:KubeControllerManagerInstall': (...args) => new KubeControllerManagerInstall(...args),
  'kubernetes-the-hard-way:remote:KubeSchedulerInstall': (...args) => new KubeSchedulerInstall(...args),
  'kubernetes-the-hard-way:remote:SystemdService': (...args) => new SystemdService(...args),
  'kubernetes-the-hard-way:tls:Certificate': (...args) => new Certificate(...args),
  'kubernetes-the-hard-way:tls:ClusterPki': (...args) => new ClusterPki(...args),
  'kubernetes-the-hard-way:tls:EncryptionKey': (...args) => new EncryptionKey(...args),
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
