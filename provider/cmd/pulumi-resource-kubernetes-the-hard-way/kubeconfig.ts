import * as path from 'node:path';
import * as YAML from 'yaml';
import { ComponentResource, ComponentResourceOptions, Input, Output, all, output } from '@pulumi/pulumi';
import { remote } from '@pulumi/command/types/input';
import * as config from './config';
import { ClusterPki } from './tls';
import { File } from './remote';
import { KubeconfigType } from './types';

type Options =
  | config.KubeconfigAdminOptions
  | config.KubeconfigKubeControllerManagerOptions
  | config.KubeconfigKubeProxyOptions
  | config.KubeconfigKubeSchedulerOptions
  | config.KubeconfigWorkerOptions;

export interface KubeconfigArgs {
  pki: Input<ClusterPki>;
  options: Input<Options>;
}

const localhost = '127.0.0.1';

export class Kubeconfig extends ComponentResource {
  public static readonly defaultContextName = 'default';

  public readonly type!: Output<KubeconfigType>;
  public readonly value!: Output<config.Kubeconfig>;
  public readonly yaml!: Output<string>;

  constructor(name: string, args: KubeconfigArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:index:Kubeconfig', name, args, opts);

    // Rehydrating
    if (opts?.urn) return;

    const pki = output(args.pki);
    const options = output(args.options);

    const cert = all([pki, options]).apply(([pki, options]) => {
      switch (options.type) {
        case 'admin':
          return pki.admin;
        case 'kube-controller-manager':
          return pki.controllerManager;
        case 'kube-proxy':
          return pki.kubeProxy;
        case 'kube-scheduler':
          return pki.kubeScheduler;
        case 'worker':
          return pki.kubelet[options.name];
      }
    });

    const server = options.apply(o => {
      switch (o.type) {
        case 'worker':
          return `https://${o.publicIp}:6443`;
        default:
          return `https://${o.publicIp ?? localhost}:6443`;
      }
    });

    const username = options.apply(o => {
      switch (o.type) {
        case 'admin':
          return 'admin';
        case 'kube-controller-manager':
          return 'system:kube-controller-manager';
        case 'kube-proxy':
          return 'system:kube-proxy';
        case 'kube-scheduler':
          return 'system:kube-scheduler';
        case 'worker':
          return `system:node:${o.name}`
        default:
          throw new Error('unsupported kubeconfig type');
      }
    });

    const clusterName = pki.clusterName;
    const caData = pki.rootCa.certPem;
    const clientCertPem = cert.certPem;
    const clientKeyPem = cert.privateKeyPem;

    const kubeconfig = all([clusterName, caData, server, username, clientCertPem, clientKeyPem])
      .apply<config.Kubeconfig>(([clusterName, caData, server, username, clientCertPem, clientKeyPem]) => ({
        clusters: [{
          name: clusterName,
          cluster: {
            certificateAuthorityData: caData,
            server,
          },
        }],
        contexts: [{
          name: Kubeconfig.defaultContextName,
          context: {
            cluster: clusterName,
            user: username,
          },
        }],
        users: [{
          name: username,
          user: {
            clientCertificateData: clientCertPem,
            clientKeyData: clientKeyPem,
          },
        }],
      }));

    const type = options.type;
    const value = kubeconfig;
    const yaml = kubeconfig.apply(YAML.stringify);

    this.type = type;
    this.value = value;
    this.yaml = yaml;

    this.registerOutputs({ type, value, yaml });
  }
}

export function installOn(
  config: Kubeconfig,
  name: string,
  connection: remote.ConnectionArgs,
  opts?: ComponentResourceOptions
): File {
  const target = path.join('home', 'kthw');
  return new File(name, {
    connection,
    content: output(config).apply(x => YAML.stringify(x)),
    path: path.join(target, `${name}.kubeconfig`),
  }, opts);
}
