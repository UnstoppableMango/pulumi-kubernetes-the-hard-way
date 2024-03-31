import { ComponentResource, ComponentResourceOptions, Inputs } from '@pulumi/pulumi';
import { Certificate } from './tls/certificate';
import { ClusterPki } from './tls/clusterPki';
import { File } from './remote/file';
import { RootCa } from './tls/rootCa';
import { Download } from './remote';
import { Mkdir, Mktemp, Rm, Tar, Wget } from './tools';
import { EncryptionKey } from './tls/encryptionKey';
import { Etcd } from './etcd';
import { Kubeconfig } from './kubeconfig';

export type ConstructComponent<T extends ComponentResource = ComponentResource>
  = (name: string, inputs: any, options: ComponentResourceOptions) => T;

export type ResourceConstructor = {
  readonly 'kubernetes-the-hard-way:index:Etcd': ConstructComponent<Etcd>;
  readonly 'kubernetes-the-hard-way:index:Kubeconfig': ConstructComponent<Kubeconfig>;
  readonly 'kubernetes-the-hard-way:remote:Download': ConstructComponent<Download>;
  readonly 'kubernetes-the-hard-way:remote:File': ConstructComponent<File>;
  readonly 'kubernetes-the-hard-way:tls:Certificate': ConstructComponent<Certificate>;
  readonly 'kubernetes-the-hard-way:tls:ClusterPki': ConstructComponent<ClusterPki>;
  readonly 'kubernetes-the-hard-way:tls:EncryptionKey': ConstructComponent<EncryptionKey>;
  readonly 'kubernetes-the-hard-way:tls:RootCa': ConstructComponent<RootCa>;
  readonly 'kubernetes-the-hard-way:tools:Mkdir': ConstructComponent<Mkdir>;
  readonly 'kubernetes-the-hard-way:tools:Mktemp': ConstructComponent<Mktemp>;
  readonly 'kubernetes-the-hard-way:tools:Rm': ConstructComponent<Rm>;
  readonly 'kubernetes-the-hard-way:tools:Tar': ConstructComponent<Tar>;
  readonly 'kubernetes-the-hard-way:tools:Wget': ConstructComponent<Wget>;
};

const resources: ResourceConstructor = {
  'kubernetes-the-hard-way:index:Etcd': (...args) => new Etcd(...args),
  'kubernetes-the-hard-way:index:Kubeconfig': (...args) => new Kubeconfig(...args),
  'kubernetes-the-hard-way:remote:Download': (...args) => new Download(...args),
  'kubernetes-the-hard-way:remote:File': (...args) => new File(...args),
  'kubernetes-the-hard-way:tls:Certificate': (...args) => new Certificate(...args),
  'kubernetes-the-hard-way:tls:ClusterPki': (...args) => new ClusterPki(...args),
  'kubernetes-the-hard-way:tls:EncryptionKey': (...args) => new EncryptionKey(...args),
  'kubernetes-the-hard-way:tls:RootCa': (...args) => new RootCa(...args),
  'kubernetes-the-hard-way:tools:Mkdir': (...args) => new Mkdir(...args),
  'kubernetes-the-hard-way:tools:Mktemp': (...args) => new Mktemp(...args),
  'kubernetes-the-hard-way:tools:Rm': (...args) => new Rm(...args),
  'kubernetes-the-hard-way:tools:Tar': (...args) => new Tar(...args),
  'kubernetes-the-hard-way:tools:Wget': (...args) => new Wget(...args),
};

export function construct(
  name: string,
  type: string,
  inputs: Inputs,
  options: ComponentResourceOptions,
): ComponentResource | undefined {
  const genericResources: Record<string, ConstructComponent> = resources;
  const resource = genericResources[type];
  if (resource === undefined) {
    return undefined;
  }
  return resource(name, inputs, options);
}
