import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { CniPluginsInstall } from './cniPlugins';
import { Mkdir } from '../tools';
import {
  CniBridgePluginConfiguration,
  CniLoopbackPluginConfiguration,
  ContainerdConfiguration,
  KubeProxyConfiguration,
  KubeletConfiguration,
} from '../config';
import { File } from './file';
import { ContainerdInstall } from './containerd';
import { ContainerdService } from './containerdService';
import { CrictlInstall } from './crictl';
import { KubectlInstall } from './kubectl';
import { KubeletInstall } from './kubelet';
import { KubeletService } from './kubeletService';
import { KubeProxyService } from './kubeProxyService';

export class WorkerNode extends schema.WorkerNode {
  constructor(name: string, args: schema.WorkerNodeArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const connection = output(args.connection);

    const cniMkdir = new Mkdir(`${name}-cni`, {
      connection,
      create: {
        directory: `/etc/cni/net.d`,
        parents: true,
      },
      delete: `rm -rf /etc/cni/net.d`,
    }, { parent: this });

    const containerdMkdir = new Mkdir(`${name}-containerd`, {
      connection,
      create: {
        directory: interpolate`/etc/containerd`,
        parents: true,
      },
      delete: interpolate`rm -rf /etc/containerd`,
    }, { parent: this });

    const kubeletMkdir = new Mkdir(`${name}-kubelet`, {
      connection,
      create: {
        directory: `/var/lib/kubelet`,
        parents: true,
      },
      delete: `rm -rf /var/lib/kubelet`,
    }, { parent: this });

    const kubeProxyMkdir = new Mkdir(`${name}-kube-proxy`, {
      connection,
      create: {
        directory: `/var/lib/kube-proxy`,
        parents: true,
      },
      delete: `rm -rf /var/lib/kube-proxy`,
    }, { parent: this });

    const varLibKubernetesMkdir = new Mkdir(`${name}-var-lib-k8s`, {
      connection,
      create: {
        directory: `/var/lib/kubernetes`,
        parents: true,
      },
      delete: `rm -rf /var/lib/kubernetes`,
    }, { parent: this });

    const varRunKubernetesMkdir = new Mkdir(`${name}-var-run-k8s`, {
      connection,
      create: {
        directory: `/var/lib/kubernetes`,
        parents: true,
      },
      delete: `rm -rf /var/lib/kubernetes`,
    }, { parent: this });

    const cniPluginsInstall = new CniPluginsInstall(name, {
      connection,
      architecture: 'amd64', // TODO
    }, { parent: this });

    const cniBridgeConfiguration = new CniBridgePluginConfiguration(name, {
      subnet: '', // TODO
    }, { parent: this });

    const cniBridgeConfigurationFile = new File(`${name}-cni-bridge`, {
      connection,
      content: cniBridgeConfiguration.yaml,
      path: interpolate`/etc/cni/net.d/10-bridge.conf`,
    }, { parent: this, dependsOn: cniMkdir });

    const cniLoopbackConfiguration = new CniLoopbackPluginConfiguration(name, {
    }, { parent: this });

    const cniLoopbackConfigurationFile = new File(`${name}-cni-loopback`, {
      connection,
      content: cniLoopbackConfiguration.yaml,
      path: interpolate`/etc/cni/net.d/99-loopback.conf`,
    }, { parent: this, dependsOn: cniMkdir });

    const containerdConfiguration = new ContainerdConfiguration(name, {
    }, { parent: this });

    const containerdConfigurationFile = new File(`${name}-containerd`, {
      connection,
      content: containerdConfiguration.toml,
      path: interpolate`/etc/containerd/config.toml`,
    }, { parent: this, dependsOn: containerdMkdir });

    const containerdInstall = new ContainerdInstall(name, {
      connection,
    }, { parent: this });

    const containerdService = new ContainerdService(name, {
      connection,
      configuration: {},
      containerdPath: containerdInstall.path,
    }, { parent: this });

    const crictlInstall = new CrictlInstall(name, {
      connection,
    }, { parent: this });

    const kubectlInstall = new KubectlInstall(name, {
      connection,
    }, { parent: this });

    const kubeletInstall = new KubeletInstall(name, {
      connection,
    }, { parent: this });

    const kubeletConfiguration = new KubeletConfiguration(name, {
      podCIDR: '',
    }, { parent: this });

    const kubeletConfigurationFile = new File(`${name}-kubelet`, {
      connection,
      content: kubeletConfiguration.yaml,
      path: interpolate`/var/lib/kubelet/kubelet-config.yaml`,
    }, { parent: this, dependsOn: kubeletMkdir });

    const kubeletService = new KubeletService(name, {
      connection,
      configuration: { // TODO
        configurationFilePath: '',
        kubeconfigPath: '',
        kubeletPath: '',
        registerNode: true,
        v: 2,
      },
      after: [], // TODO
      requires: [], // TODO
    }, { parent: this });

    const kubeProxyConfiguration = new KubeProxyConfiguration(name, {
      clusterCIDR: '',
      kubeconfig: '',
    }, { parent: this });

    const kubeProxyConfigurationFile = new File(`${name}-kube-proxy`, {
      connection,
      content: kubeProxyConfiguration.yaml,
      path: interpolate`/var/lib/kube-proxy/kube-proxy-config.yaml`,
    }, { parent: this, dependsOn: kubeProxyMkdir });

    const kubeProxyService = new KubeProxyService(name, {
      connection,
      configuration: {
        configurationFilePath: '',
        kubeProxyPath: '',
      },
    }, { parent: this });

    this.cniMkdir = cniMkdir;
    this.cniBridgeConfiguration = cniBridgeConfiguration;
    this.cniBridgeConfigurationFile = cniBridgeConfigurationFile;
    this.cniLoopbackConfiguration = cniLoopbackConfiguration;
    this.cniLoopbackConfigurationFile = cniLoopbackConfigurationFile;
    this.connection = connection;
    this.containerdConfiguration = containerdConfiguration;
    this.containerdConfigurationFile = containerdConfigurationFile;
    this.containerdMkdir = containerdMkdir;
    this.containerdInstall = containerdInstall;
    this.containerdService = containerdService;
    this.crictlInstall = crictlInstall;
    this.kubectlInstall = kubectlInstall;
    this.kubeletConfiguration = kubeletConfiguration;
    this.kubeletConfigurationFile = kubeletConfigurationFile;
    this.kubeletInstall = kubeletInstall;
    this.kubeletMkdir = kubeletMkdir;
    this.kubeletService = kubeletService;
    this.kubeProxyConfiguration = kubeProxyConfiguration;
    this.kubeProxyConfigurationFile = kubeProxyConfigurationFile;
    this.kubeProxyMkdir = kubeProxyMkdir;
    this.kubeProxyService = kubeProxyService;
    this.varLibKubernetesMkdir = varLibKubernetesMkdir;
    this.varRunKubernetesMkdir = varRunKubernetesMkdir;

    this.registerOutputs({
      cniMkdir,
      cniBridgeConfiguration,
      cniBridgeConfigurationFile,
      cniLoopbackConfiguration,
      cniLoopbackConfigurationFile,
      connection,
      containerdConfiguration,
      containerdConfigurationFile,
      containerdInstall,
      containerdMkdir,
      crictlInstall,
      kubeletMkdir,
      kubeletService,
      kubeProxyConfiguration,
      kubeProxyConfigurationFile,
      kubeProxyMkdir,
      kubeProxyService,
      varLibKubernetesMkdir,
      varRunKubernetesMkdir,
    });
  }
}
