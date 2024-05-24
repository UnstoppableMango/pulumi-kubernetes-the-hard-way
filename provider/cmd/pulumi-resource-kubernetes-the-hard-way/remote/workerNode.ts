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
import { KubeProxyInstall } from './kubeProxy';

export class WorkerNode extends schema.WorkerNode {
  constructor(name: string, args: schema.WorkerNodeArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const architecture = output(args.architecture);
    const clusterCIDR = output(args.clusterCIDR ?? '10.200.0.0/16');
    const connection = output(args.connection);
    const cniConfigurationDirectory = output(args.cniConfigurationDirectory ?? '/etc/cni/net.d');
    const containerdConfigurationDirectory = output(args.containerdConfigurationDirectory ?? '/etc/containerd');
    const kubeletConfigurationDirectory = output(args.kubeletConfigurationDirectory ?? '/var/lib/kubelet');
    const kubeletKubeconfigPath = output(args.kubeletKubeconfigPath ?? '/var/lib/kubelet/kubeconfig');
    const kubeProxyConfigurationDirectory = output(args.kubeProxyConfigurationDirectory ?? '/var/lib/kube-proxy');
    const kubeProxyKubeconfigPath = output(args.kubeProxyKubeconfigPath ?? '/var/lib/kube-proxy/kubeconfig');
    const kubernetesVersion = output(args.kubernetesVersion ?? '1.29.2');
    const podCIDR = output(args.podCIDR);
    const subnet = output(args.subnet);

    const cniMkdir = new Mkdir(`${name}-cni`, {
      connection,
      create: {
        directory: cniConfigurationDirectory,
        parents: true,
      },
      delete: interpolate`rm -rf ${cniConfigurationDirectory}`,
    }, { parent: this });

    const containerdMkdir = new Mkdir(`${name}-containerd`, {
      connection,
      create: {
        directory: containerdConfigurationDirectory,
        parents: true,
      },
      delete: interpolate`rm -rf ${containerdConfigurationDirectory}`,
    }, { parent: this });

    const kubeletMkdir = new Mkdir(`${name}-kubelet`, {
      connection,
      create: {
        directory: kubeletConfigurationDirectory,
        parents: true,
      },
      delete: interpolate`rm -rf ${kubeletConfigurationDirectory}`,
    }, { parent: this });

    const kubeProxyMkdir = new Mkdir(`${name}-kube-proxy`, {
      connection,
      create: {
        directory: kubeProxyConfigurationDirectory,
        parents: true,
      },
      delete: interpolate`rm -rf ${kubeProxyConfigurationDirectory}`,
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
        directory: `/var/run/kubernetes`,
        parents: true,
      },
      delete: `rm -rf /var/run/kubernetes`,
    }, { parent: this });

    const cniPluginsInstall = new CniPluginsInstall(name, {
      connection,
      architecture,
      directory: args.cniInstallDirectory,
      version: args.cniVersion,
    }, { parent: this });

    const cniBridgeConfiguration = new CniBridgePluginConfiguration(name, {
      subnet,
    }, { parent: this });

    const cniBridgeConfigurationFile = new File(`${name}-cni-bridge`, {
      connection,
      content: cniBridgeConfiguration.yaml,
      path: interpolate`${cniConfigurationDirectory}/10-bridge.conf`,
    }, { parent: this, dependsOn: cniMkdir });

    const cniLoopbackConfiguration = new CniLoopbackPluginConfiguration(name, {
    }, { parent: this });

    const cniLoopbackConfigurationFile = new File(`${name}-cni-loopback`, {
      connection,
      content: cniLoopbackConfiguration.yaml,
      path: interpolate`${cniConfigurationDirectory}/99-loopback.conf`,
    }, { parent: this, dependsOn: cniMkdir });

    const containerdConfiguration = new ContainerdConfiguration(name, {
      cri: {
        cni: {
          binDir: cniPluginsInstall.directory,
          confDir: cniConfigurationDirectory,
        },
        containerd: {}, // TODO
      },
    }, { parent: this });

    const containerdConfigurationFile = new File(`${name}-containerd`, {
      connection,
      content: containerdConfiguration.toml,
      path: interpolate`${containerdConfigurationDirectory}/config.toml`,
    }, { parent: this, dependsOn: containerdMkdir });

    const containerdInstall = new ContainerdInstall(name, {
      connection,
      architecture,
      directory: args.containerdInstallDirectory,
      version: args.containerdVersion,
    }, { parent: this });

    const containerdService = new ContainerdService(name, {
      connection,
      configuration: output(containerdConfiguration.result),
      containerdPath: containerdInstall.path,
    }, { parent: this });

    const crictlInstall = new CrictlInstall(name, {
      connection,
      architecture,
      directory: args.crictlInstallDirectory,
      version: kubernetesVersion,
    }, { parent: this });

    const kubectlInstall = new KubectlInstall(name, {
      connection,
      architecture,
      directory: args.kubectlInstallDirectory,
      version: kubernetesVersion,
    }, { parent: this });

    const kubeletInstall = new KubeletInstall(name, {
      connection,
      architecture,
      directory: args.kubeletInstallDirectory,
      version: kubernetesVersion,
    }, { parent: this });

    const kubeletConfiguration = new KubeletConfiguration(name, {
      podCIDR,
      // TODO: Rest of the config
    }, { parent: this });

    const kubeletConfigurationFile = new File(`${name}-kubelet`, {
      connection,
      content: kubeletConfiguration.yaml,
      path: interpolate`${kubeletConfigurationDirectory}/kubelet-config.yaml`,
    }, { parent: this, dependsOn: kubeletMkdir });

    const kubeletService = new KubeletService(name, {
      connection,
      configuration: { // TODO
        configurationFilePath: kubeletConfigurationFile.path,
        kubeconfigPath: kubeletKubeconfigPath,
        kubeletPath: kubeletInstall.path,
        registerNode: true,
        v: 2,
      },
      after: ['containerd.service'],
      requires: ['containerd.service'],
    }, { parent: this });

    const kubeProxyConfiguration = new KubeProxyConfiguration(name, {
      clusterCIDR,
      kubeconfig: kubeProxyKubeconfigPath,
      mode: undefined, // TODO
    }, { parent: this });

    const kubeProxyConfigurationFile = new File(`${name}-kube-proxy`, {
      connection,
      content: kubeProxyConfiguration.yaml,
      path: interpolate`${kubeProxyConfigurationDirectory}/kube-proxy-config.yaml`,
    }, { parent: this, dependsOn: kubeProxyMkdir });

    const kubeProxyInstall = new KubeProxyInstall(name, {
      connection,
      architecture,
      directory: args.kubeProxyInstallDirectory,
      version: kubernetesVersion,
    }, { parent: this });

    const kubeProxyService = new KubeProxyService(name, {
      connection,
      configuration: {
        configurationFilePath: kubeProxyConfigurationFile.path,
        kubeProxyPath: kubeProxyInstall.path,
      },
    }, { parent: this });

    this.clusterCIDR = clusterCIDR;
    this.cniMkdir = cniMkdir;
    this.cniConfigurationDirectory = cniConfigurationDirectory;
    this.cniBridgeConfiguration = cniBridgeConfiguration;
    this.cniBridgeConfigurationFile = cniBridgeConfigurationFile;
    this.cniInstallDirectory = cniPluginsInstall.directory;
    this.cniLoopbackConfiguration = cniLoopbackConfiguration;
    this.cniLoopbackConfigurationFile = cniLoopbackConfigurationFile;
    this.cniVersion = cniPluginsInstall.version;
    this.connection = connection;
    this.containerdConfiguration = containerdConfiguration;
    this.containerdConfigurationDirectory = containerdConfigurationDirectory;
    this.containerdConfigurationFile = containerdConfigurationFile;
    this.containerdMkdir = containerdMkdir;
    this.containerdInstall = containerdInstall;
    this.containerdService = containerdService;
    this.containerdVersion = containerdInstall.version;
    this.crictlInstall = crictlInstall;
    this.crictlInstallDirectory = crictlInstall.directory;
    this.kubectlInstall = kubectlInstall;
    this.kubectlInstallDirectory = kubectlInstall.directory;
    this.kubeletConfiguration = kubeletConfiguration;
    this.kubeletConfigurationDirectory = kubeletConfigurationDirectory;
    this.kubeletConfigurationFile = kubeletConfigurationFile;
    this.kubeletInstall = kubeletInstall;
    this.kubeletInstallDirectory = kubeletInstall.directory;
    this.kubeletKubeconfigPath = kubeletKubeconfigPath;
    this.kubeletMkdir = kubeletMkdir;
    this.kubeletService = kubeletService;
    this.kubeProxyConfiguration = kubeProxyConfiguration;
    this.kubeProxyConfigurationDirectory = kubeProxyConfigurationDirectory;
    this.kubeProxyConfigurationFile = kubeProxyConfigurationFile;
    this.kubeProxyInstall = kubeProxyInstall;
    this.kubeProxyInstallDirectory = kubeProxyInstall.directory;
    this.kubeProxyMkdir = kubeProxyMkdir;
    this.kubeProxyService = kubeProxyService;
    this.kubernetesVersion = kubernetesVersion;
    this.podCIDR = podCIDR;
    this.subnet = subnet;
    this.varLibKubernetesMkdir = varLibKubernetesMkdir;
    this.varRunKubernetesMkdir = varRunKubernetesMkdir;

    this.registerOutputs({
      clusterCIDR,
      cniMkdir,
      cniBridgeConfiguration,
      cniConfigurationDirectory,
      cniBridgeConfigurationFile,
      cniInstallDirectory: this.cniInstallDirectory,
      cniLoopbackConfiguration,
      cniLoopbackConfigurationFile,
      cniVersion: this.cniVersion,
      connection,
      containerdConfiguration,
      containerdConfigurationDirectory,
      containerdConfigurationFile,
      containerdInstall,
      containerdInstallDirectory: this.containerdInstallDirectory,
      containerdMkdir,
      containerdVersion: this.containerdVersion,
      crictlInstall,
      crictlInstallDirectory: this.crictlInstallDirectory,
      kubeletConfigurationDirectory,
      kubeletInstallDirectory: this.kubeletInstallDirectory,
      kubeletMkdir,
      kubeletService,
      kubeProxyConfiguration,
      kubeProxyConfigurationDirectory,
      kubeProxyConfigurationFile,
      kubeProxyInstall,
      kubeProxyInstallDirectory: this.kubeProxyInstallDirectory,
      kubeProxyMkdir,
      kubeProxyService,
      varLibKubernetesMkdir,
      varRunKubernetesMkdir,
    });
  }
}
