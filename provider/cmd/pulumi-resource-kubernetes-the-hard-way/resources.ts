import { ComponentResource, ComponentResourceOptions, Inputs } from '@pulumi/pulumi';
import { Certificate } from './certificate';
import { ClusterPki } from './clusterPki';
import { RemoteFile } from './remote/file';
import { RootCa } from './rootCa';
import { Download } from './remote';
import { Mkdir, Tar, Wget } from './tools';

export type ConstructComponent<T extends ComponentResource = ComponentResource>
  = (name: string, inputs: any, options: ComponentResourceOptions) => T;

export type ResourceConstructor = {
  readonly 'kubernetes-the-hard-way:index:Certificate': ConstructComponent<Certificate>;
  readonly 'kubernetes-the-hard-way:index:ClusterPki': ConstructComponent<ClusterPki>;
  readonly 'kubernetes-the-hard-way:remote:File': ConstructComponent<RemoteFile>;
  readonly 'kubernetes-the-hard-way:index:RootCa': ConstructComponent<RootCa>;
  readonly 'kubernetes-the-hard-way:remote:Download': ConstructComponent<Download>;
  readonly 'kubernetes-the-hard-way:tools:Mkdir': ConstructComponent<Mkdir>;
  readonly 'kubernetes-the-hard-way:tools:Tar': ConstructComponent<Tar>;
  readonly 'kubernetes-the-hard-way:tools:Wget': ConstructComponent<Wget>;
};

const resources: ResourceConstructor = {
  'kubernetes-the-hard-way:index:Certificate': (...args) => new Certificate(...args),
  'kubernetes-the-hard-way:index:ClusterPki': (...args) => new ClusterPki(...args),
  'kubernetes-the-hard-way:remote:File': (...args) => new RemoteFile(...args),
  'kubernetes-the-hard-way:index:RootCa': (...args) => new RootCa(...args),
  'kubernetes-the-hard-way:remote:Download': (...args) => new Download(...args),
  'kubernetes-the-hard-way:tools:Mkdir': (...args) => new Mkdir(...args),
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
