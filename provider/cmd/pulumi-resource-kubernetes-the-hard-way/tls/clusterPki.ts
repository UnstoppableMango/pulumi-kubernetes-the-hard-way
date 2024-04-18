import { ComponentResourceOptions, Input, Output, all, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { RootCa, newCertificate } from './rootCa';
import { Certificate } from './certificate';
import { Kubeconfig, KubeconfigOptions } from '../config';
import { Algorithm } from '../types';

export interface GetKubeconfigInputs {
  options: KubeconfigOptions;
}

export interface GetKubeconfigOutputs {
  result: Output<Kubeconfig>;
}

export type NodeRole = 'worker' | 'controlplane';
export type NodeMapInput = Record<string, Input<NodeArgs>>;

export interface NodeArgs {
  ip: Input<string>;
  role: Input<NodeRole>;
}

export interface ClusterPkiArgs<T extends NodeMapInput = NodeMapInput> extends schema.ClusterPkiArgs {
  nodes: T;
}

type CertMap<T> = {
  [P in keyof T]: Certificate;
}

export class ClusterPki<T extends NodeMapInput = NodeMapInput> extends schema.ClusterPki {
  public static readonly defaultAlgorithm: Algorithm = 'RSA';
  public static readonly defaultExpiry: number = 8760;
  public static readonly defaultRsaBits: number = 2048;

  constructor(private name: string, args: ClusterPkiArgs<T>, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const algorithm = output(args.algorithm ?? ClusterPki.defaultAlgorithm);
    const clusterName = output(args.clusterName);
    const ecdsaCurve = output(args.ecdsaCurve ?? 'P256');
    const publicIp = output(args.publicIp);
    const rsaBits = output(args.rsaBits ?? ClusterPki.defaultRsaBits);
    const validityPeriodHours = output(args.validityPeriodHours ?? ClusterPki.defaultExpiry);

    const ca = new RootCa(name, {
      algorithm,
      rsaBits,
      validityPeriodHours,
      subject: {
        commonName: clusterName,
      },
    }, { parent: this });

    const admin = newCertificate(ca, {
      name: this.certName('admin'),
      algorithm,
      rsaBits,
      validityPeriodHours,
      allowedUses: ca.allowedUses, // TODO
      subject: {
        commonName: 'admin',
        organization: 'system:masters',
      },
      options: { parent: this },
    });

    const controllerManager = newCertificate(ca, {
      name: this.certName('controller-manager'),
      algorithm,
      rsaBits,
      validityPeriodHours,
      allowedUses: ca.allowedUses, // TODO
      subject: {
        commonName: 'system:kube-controller-manager',
        organization: 'system:kube-controller-manager',
      },
      options: { parent: this },
    });

    const kubelet: Partial<CertMap<T>> = {};
    for (const key in args.nodes) {
      const node = output(args.nodes[key]);
      kubelet[key] = newCertificate(ca, {
        name: this.certName(`${key}-worker`),
        algorithm,
        rsaBits,
        validityPeriodHours,
        allowedUses: ca.allowedUses, // TODO
        ipAddresses: [node.ip],
        subject: {
          commonName: interpolate`system:node:${node.ip}`,
          organization: 'system:nodes',
        },
        options: { parent: this },
      });
    }

    const kubeProxy = newCertificate(ca, {
      name: this.certName('kube-proxy'),
      algorithm,
      rsaBits,
      validityPeriodHours,
      allowedUses: ca.allowedUses, // TODO
      subject: {
        commonName: 'system:kube-proxy',
        organization: 'system:node-proxier',
      },
      options: { parent: this },
    });

    const kubeScheduler = newCertificate(ca, {
      name: this.certName('kube-scheduler'),
      algorithm,
      rsaBits,
      validityPeriodHours,
      allowedUses: ca.allowedUses, // TODO
      subject: {
        commonName: 'system:kube-scheduler',
        organization: 'system:kube-scheduler',
      },
      options: { parent: this },
    });

    const kubernetes = newCertificate(ca, {
      name: this.certName('kubernetes'),
      algorithm,
      rsaBits,
      validityPeriodHours,
      allowedUses: ca.allowedUses, // TODO
      subject: {
        commonName: 'kubernetes',
        organization: 'Kubernetes',
      },
      dnsNames: [
        'kubernetes',
        'kubernetes.default',
        'kubernetes.default.svc',
        'kubernetes.default.svc.cluster',
        'kubernetes.svc.cluster.local',
      ],
      ipAddresses: [
        // Internal cluster service IPs
        '10.32.0.1', '10.240.0.10', '10.240.0.11', '10.240.0.12', '127.0.0.1',
        // TODO: Input IPs
        ...Object.values(args.nodes).map(x => output(x).ip),
        publicIp,
      ],
      options: { parent: this },
    });

    const serviceAccounts = newCertificate(ca, {
      name: this.certName('service-accounts'),
      algorithm,
      rsaBits,
      validityPeriodHours,
      allowedUses: ca.allowedUses, // TODO
      subject: {
        commonName: 'service-accounts',
        organization: 'Kubernetes',
      },
      options: { parent: this },
    });

    this.admin = admin;
    this.algorithm = algorithm;
    this.ca = ca;
    this.clusterName = clusterName;
    this.controllerManager = controllerManager;
    this.ecdsaCurve = ecdsaCurve;
    this.kubelet = kubelet as CertMap<T>; // TODO: Can we refactor away from a cast?
    this.kubeProxy = kubeProxy;
    this.kubernetes = kubernetes;
    this.kubeScheduler = kubeScheduler;
    this.publicIp = publicIp;
    this.rsaBits = rsaBits;
    this.serviceAccounts = serviceAccounts;
    this.validityPeriodHours = validityPeriodHours;

    this.registerOutputs({
      admin,
      algorithm,
      ecdsaCurve,
      ca,
      clusterName,
      controllerManager,
      kubelet,
      kubeProxy,
      kubernetes,
      kubeScheduler,
      publicIp,
      rsaBits,
      serviceAccounts,
      validityPeriodHours,
    });
  }

  public async getKubeconfig({ options }: GetKubeconfigInputs): Promise<GetKubeconfigOutputs> {
    const cert = this.getCert(options);
    const ip = getIp(options);
    const username = getUsername(options);

    const caData = this.ca.certPem;
    const clientCertPem = cert.certPem;
    const clientKeyPem = cert.privateKeyPem;
    const clusterName = this.clusterName;
    const server = interpolate`https://${ip}:6443`;

    const kubeconfig = all([clusterName, caData, server, username, clientCertPem, clientKeyPem])
      .apply<Kubeconfig>(([clusterName, caData, server, username, clientCertPem, clientKeyPem]) => ({
        clusters: [{
          name: clusterName,
          cluster: {
            certificateAuthorityData: caData,
            server,
          },
        }],
        contexts: [{
          name: 'default',
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

    return { result: kubeconfig };
  }

  private certName(type: string): string {
    return `${this.name}-${type}`;
  }

  private getCert(options: KubeconfigOptions): Output<schema.Certificate> {
    switch (options.type) {
      case 'admin':
        return output(this.admin);
      case 'kube-controller-manager':
        return output(this.controllerManager);
      case 'kube-proxy':
        return output(this.kubeProxy);
      case 'kube-scheduler':
        return output(this.kubeScheduler);
      case 'worker':
        return all([options.name, this.kubelet]).apply(([n, k]) => k[n]);
      default:
        throw new Error('unsupported kubeconfig type');
    }
  }
}

function getIp(options: KubeconfigOptions): Output<string> {
  // TODO: Is the ternary really necessary?
  const result = options.type === 'worker'
    ? options.publicIp
    : options.publicIp ?? '127.0.0.1';

  return output(result);
}

function getUsername(options: KubeconfigOptions): Output<string> {
  switch (options.type) {
    case 'admin':
      return output('admin');
    case 'kube-controller-manager':
      return output('system:kube-controller-manager');
    case 'kube-proxy':
      return output('system:kube-proxy');
    case 'kube-scheduler':
      return output('system:kube-scheduler');
    case 'worker':
      return interpolate`system:node:${options.name}`
    default:
      throw new Error('unsupported kubeconfig type');
  }
}
