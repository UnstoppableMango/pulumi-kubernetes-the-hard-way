import { ComponentResourceOptions, Input, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { KubeApiServerInstall } from './kubeApiServerInstall';
import { SystemdService } from './systemdService';
import { remote } from '@pulumi/command/types/input';
import { CommandBuilder } from '../tools/commandBuilder';
import { File } from './file';
import { Mkdir } from '../tools';
import { KubeControllerManagerInstall } from './kubeControllerManagerInstall';
import { KubectlInstall } from './kubectlInstall';
import { KubeSchedulerInstall } from './kubeSchedulerInstall';

export class ControlPlaneNode extends schema.ControlPlaneNode {
  constructor(name: string, args: schema.ControlPlaneNodeArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const apiServerCount = output(args.apiServerCount);
    const architecture = output(args.architecture);
    const auditLogPath = output(args.audiLogPath ?? '/var/log/audit.log');
    const caCertificatePath = output(args.caCertificatePath);
    const caPrivateKeyPath = output(args.caPrivateKeyPath);
    const clusterCIDR = output(args.clusterCIDR ?? '10.200.0.0/16');
    const clusterName = output(args.clusterName ?? 'kubernetes');
    const connection = output(args.connection);
    const encryptionConfigYaml = output(args.encryptionConfigYaml);
    const kubeApiServerCertificatePath = output(args.kubeApiServerCertificatePath);
    const kubeApiServerInstallDirectory = output(args.kubeApiServerInstallDirectory);
    const kubeApiServerPrivateKeyPath = output(args.kubeApiServerPrivateKeyPath);
    const kubeControllerManagerKubeconfigPath = output(args.kubeControllerManagerKubeconfigPath);
    const kubeControllerManagerInstallDirectory = output(args.kubeControllerManagerInstallDirectory);
    const kubeSchedulerConfigYaml = output(args.kubeSchedulerConfigYaml);
    const kubeSchedulerInstallDirectory = output(args.kubeSchedulerInstallDirectory);
    const kubeSchedulerKubeconfigPath = output(args.kubeSchedulerKubeconfigPath);
    const kubernetesVersion = output(args.kubernetesVersion ?? '1.30.0');
    const nodeName = output(args.nodeName ?? 'server');
    const serviceAccountsCertificatePath = output(args.serviceAccountsCertificatePath);
    const serviceAccountsPrivateKeyPath = output(args.serviceAccountsPrivateKeyPath);
    const serviceClusterIpRange = output(args.serviceClusterIpRange ?? '10.32.0.0/24');

    const kubeApiServerInstall = new KubeApiServerInstall(name, {
      connection,
      architecture,
      directory: args.kubeApiServerInstallDirectory,
      version: kubernetesVersion,
    }, { parent: this });

    const varLibKubernetesMkdir = new Mkdir(`${name}-var-lib-kubernetes`, {
      connection,
      create: {
        directory: '/var/lib/kubernetes',
        parents: true,
      },
      delete: 'rm -rf /var/lib/kubernetes',
    }, { parent: this });

    const encryptionConfigFile = new File(`${name}-encryption-config`, {
      connection,
      content: encryptionConfigYaml,
      path: '/var/lib/kubernetes/encryption-config.yaml',
    }, { parent: this, dependsOn: varLibKubernetesMkdir });

    const kubeApiServerService = this.getKubeApiServerService(
      name,
      connection,
      kubeApiServerInstall.path,
      apiServerCount,
      auditLogPath,
      kubeApiServerCertificatePath,
      encryptionConfigFile.path,
      caCertificatePath,
      kubeApiServerCertificatePath,
      kubeApiServerPrivateKeyPath,
      serviceAccountsCertificatePath,
      serviceAccountsPrivateKeyPath,
      interpolate`https://${nodeName}.kubernetes.local:6443`,
      serviceClusterIpRange,
      kubeApiServerCertificatePath,
      kubeApiServerPrivateKeyPath,
    );

    const kubeControllerManagerInstall = new KubeControllerManagerInstall(name, {
      connection,
      architecture,
      directory: args.kubeControllerManagerInstallDirectory,
      version: kubernetesVersion,
    }, { parent: this });

    const kubeControllerManagerService = this.getKubeControllerManagerService(
      name,
      connection,
      kubeApiServerInstall.path,
      clusterName,
      clusterCIDR,
      caCertificatePath,
      caCertificatePath,
      caPrivateKeyPath,
      kubeControllerManagerKubeconfigPath,
      serviceAccountsPrivateKeyPath,
      serviceClusterIpRange,
    );

    const kubectlInstall = new KubectlInstall(name, {
      connection,
      architecture,
      directory: args.kubectlInstallDirectory,
      version: kubernetesVersion,
    }, { parent: this });

    const kubernetesConfigurationMkdir = new Mkdir(`${name}-kubernetes-config`, {
      connection,
      create: {
        directory: '/etc/kubernetes/config',
        parents: true,
      },
      delete: `rm -rf /etc/kubernetes/config`,
    }, { parent: this });

    const kubeSchedulerInstall = new KubeSchedulerInstall(name, {
      connection,
      architecture,
      directory: args.kubeSchedulerInstallDirectory,
      version: kubernetesVersion,
    }, { parent: this });

    const kubeSchedulerConfigFile = new File(`${name}-kube-scheduler-config`, {
      connection,
      content: kubeSchedulerConfigYaml,
      path: interpolate`/etc/kubernetes/config/kube-scheduler.yaml`,
    }, { parent: this });

    const kubeSchedulerService = this.getKubeSchedulerService(
      name,
      connection,
      kubeSchedulerInstall.path,
      kubeSchedulerConfigFile.path,
    );

    this.apiServerCount = apiServerCount;
    this.architecture = architecture;
    this.audiLogPath = auditLogPath;
    this.caCertificatePath = caCertificatePath;
    this.caPrivateKeyPath = caPrivateKeyPath;
    this.clusterCIDR = clusterCIDR;
    this.clusterName = clusterName;
    this.connection = connection;
    this.encryptionConfigYaml = encryptionConfigYaml;
    this.kubeApiServerCertificatePath = kubeApiServerCertificatePath;
    this.kubeApiServerInstall = kubeApiServerInstall;
    this.kubeApiServerInstallDirectory = kubeApiServerInstall.directory;
    this.kubeApiServerPrivateKeyPath = kubeApiServerPrivateKeyPath;
    this.kubeApiServerService = kubeApiServerService;
    this.kubeControllerManagerInstall = kubeControllerManagerInstall;
    this.kubeControllerManagerInstallDirectory = kubeApiServerInstall.directory;
    this.kubeControllerManagerKubeconfigPath = kubeControllerManagerKubeconfigPath;
    this.kubeControllerManagerService = kubeControllerManagerService;
    this.kubectlInstall = kubectlInstall;
    this.kubectlInstallDirectory = kubectlInstall.directory;
    this.kubernetesConfigurationMkdir = kubernetesConfigurationMkdir;
    this.kubernetesVersion = kubernetesVersion;
    this.kubeSchedulerConfigYaml = kubeSchedulerConfigYaml;
    this.kubeSchedulerInstall = kubeSchedulerInstall;
    this.kubeSchedulerInstallDirectory = kubeSchedulerInstall.directory;
    this.kubeSchedulerKubeconfigPath = kubeSchedulerKubeconfigPath;
    this.nodeName = nodeName;
    this.serviceAccountsCertificatePath = serviceAccountsCertificatePath;
    this.serviceAccountsPrivateKeyPath = serviceAccountsPrivateKeyPath;
    this.serviceClusterIpRange = serviceClusterIpRange;
    this.varLibKubernetesMkdir = varLibKubernetesMkdir;

    this.registerOutputs({
      apiServerCount,
      architecture,
      auditLogPath,
      caCertificatePath,
      caPrivateKeyPath,
      clusterCIDR,
      clusterName,
      connection,
      encryptionConfigYaml,
      kubeApiServerCertificatePath,
      kubeApiServerInstall,
      kubeApiServerInstallDirectory,
      kubeApiServerPrivateKeyPath,
      kubeApiServerService,
      kubeControllerManagerInstall,
      kubeControllerManagerInstallDirectory,
      kubeControllerManagerKubeconfigPath,
      kubeControllerManagerService,
      kubectlInstall,
      kubectlInstallDirectory: this.kubectlInstallDirectory,
      kubernetesConfigurationMkdir,
      kubernetesVersion,
      kubeSchedulerConfigYaml,
      kubeSchedulerInstall,
      kubeSchedulerInstallDirectory,
      kubeSchedulerKubeconfigPath,
      kubeSchedulerService,
      nodeName,
      serviceAccountsCertificatePath,
      serviceAccountsPrivateKeyPath,
      serviceClusterIpRange,
      varLibKubernetesMkdir,
    });
  }

  private getKubeApiServerService(
    name: string,
    connection: Input<remote.ConnectionArgs>,
    binaryPath: Input<string>,
    apiServerCount: Input<number>,
    auditLogPath: Input<string>,
    clientCaFile: Input<string>,
    encryptionProviderConfig: Input<string>,
    kubeletCertificateAuthority: Input<string>,
    kubeletClientCertificate: Input<string>,
    kubeletClientKey: Input<string>,
    serviceAccountKeyFile: Input<string>,
    serviceAccountSigningKeyFile: Input<string>,
    serviceAccountIssuer: Input<string>,
    serviceClusterIpRange: Input<string>,
    tlsCertFile: Input<string>,
    tlsPrivateKeyFile: Input<string>,
  ): SystemdService {
    const execStart = new CommandBuilder(binaryPath)
      .option('allow-privileged', 'true')
      .option('--apiserver-count', apiServerCount)
      .option('--audit-log-maxbackup', 3)
      .option('--audit-log-maxsize', 100)
      .option('--audit-log-path', auditLogPath)
      .option('--authorization-mode', 'Node,RBAC')
      .option('--bind-address', '0.0.0.0')
      .option('--client-ca-file', clientCaFile)
      .option('--enable-admission-plugins', 'NamespaceLifecycle,NodeRestriction,LimitRanger,ServiceAccount,DefaultStorageClass,ResourceQuota')
      .option('--etcd-servers', 'http://127.0.0.1:2379')
      .option('--event-ttl', '1h')
      .option('--encryption-provider-config', encryptionProviderConfig)
      .option('--kubelet-certificate-authority', kubeletCertificateAuthority)
      .option('--kubelet-client-certificate', kubeletClientCertificate)
      .option('--kubelet-client-key', kubeletClientKey)
      .option('--runtime-config', 'api/all=true')
      .option('--service-account-key-file', serviceAccountKeyFile)
      .option('--service-account-signing-key-file', serviceAccountSigningKeyFile)
      .option('--service-account-issuer', serviceAccountIssuer)
      .option('--service-cluster-ip-range', serviceClusterIpRange)
      .option('--service-node-port-range', '30000-32767')
      .option('--tls-cert-file', tlsCertFile)
      .option('--tls-private-key-file', tlsPrivateKeyFile)
      .option('--v', 2);

    return new SystemdService(`${name}`, {
      connection,
      unit: {
        description: 'Kubernetes API Server',
        documentation: ['https://github.com/kubernetes/kubernetes'],
      },
      service: {
        execStart: execStart.command,
        restart: 'on-failure',
        restartSec: '5',
      },
      install: {
        wantedBy: ['mult-user.target']
      },
    }, { parent: this });
  }

  private getKubeControllerManagerService(
    name: string,
    connection: Input<remote.ConnectionArgs>,
    binaryPath: Input<string>,
    clusterName: Input<string>,
    clusterCIDR: Input<string>,
    rootCaFile: Input<string>,
    clusterSigningCertFile: Input<string>,
    clusterSigningKeyFile: Input<string>,
    kubeconfig: Input<string>,
    serviceAccountPrivateKeyFile: Input<string>,
    serviceClusterIpRange: Input<string>,
  ): SystemdService {
    const execStart = new CommandBuilder(binaryPath)
      .option('--bind-address', '0.0.0.0')
      .option('--cluster-cidr', clusterCIDR)
      .option('--cluster-name', clusterName)
      .option('--cluster-signing-cert-file', clusterSigningCertFile)
      .option('--cluster-signing-key-file', clusterSigningKeyFile)
      .option('--kubeconfig', kubeconfig)
      .option('--root-ca-file', rootCaFile)
      .option('--service-account-private-key-file', serviceAccountPrivateKeyFile)
      .option('--service-cluster-ip-range', serviceClusterIpRange)
      .option('--use-service-account-credentials', 'true')
      .option('--v', '2');

    return new SystemdService(`${name}-kube-apiserver`, {
      connection,
      unit: {
        description: 'Kubernetes Controller Manager',
        documentation: ['https://github.com/kubernetes/kubernetes'],
      },
      service: {
        execStart: execStart.command,
        restart: 'on-failure',
        restartSec: '5',
      },
      install: {
        wantedBy: ['multi-user.target'],
      },
    }, { parent: this });
  }

  private getKubeSchedulerService(
    name: string,
    connection: Input<remote.ConnectionArgs>,
    binaryPath: Input<string>,
    config: Input<string>,
  ): SystemdService {
    const execStart = new CommandBuilder(binaryPath)
      .option('--config', config)
      .option('--v', 2);

    return new SystemdService(`${name}-kube-scheduler`, {
      connection,
      unit: {
        description: 'Kubernetes Scheduler',
        documentation: ['https://github.com/kubernetes/kubernetes'],
      },
      service: {
        execStart: execStart.command,
        restart: 'on-failure',
        restartSec: '5',
      },
      install: {
        wantedBy: ['mult-user.target'],
      },
    }, { parent: this });
  }
}
