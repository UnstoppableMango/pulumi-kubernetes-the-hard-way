import { ComponentResource, ComponentResourceOptions, Inputs } from '@pulumi/pulumi';
import { Download, File } from './remote';
import { Certificate, ClusterPki, EncryptionKey, RootCa } from './tls';
import { Mkdir, Mktemp, Rm, Tar, Wget } from './tools';
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

type Functions = {
  'kubernetes-the-hard-way:index:Certificate/installCert': (inputs: InstallArgs) => Promise<RemoteFile>;
  'kubernetes-the-hard-way:index:Certificate/installKey': (inputs: InstallArgs) => Promise<RemoteFile>;
  'kubernetes-the-hard-way:index:installCert': (inputs: InstallArgs) => Promise<RemoteFile>;
  'kubernetes-the-hard-way:index:newCertificate': (inputs: NewCertificateArgs) => Promise<Certificate>;
  'kubernetes-the-hard-way:index:installKey': (inputs: InstallArgs) => Promise<RemoteFile>;
  'kubernetes-the-hard-way:index:RootCa/newCertificate': (inputs: NewCertificateArgs) => Promise<Certificate>;
  'kubernetes-the-hard-way:index:RootCa/installCert': (inputs: InstallArgs) => Promise<RemoteFile>;
  'kubernetes-the-hard-way:index:RootCa/installKey': (inputs: InstallArgs) => Promise<RemoteFile>;
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

const functions: Functions = {};

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
