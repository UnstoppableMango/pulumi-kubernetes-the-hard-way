import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Chmod, Mkdir } from '../tools';

export class KubernetesControlPlaneConfiguration extends schema.KubernetesControlPlaneConfiguration {
  constructor(name: string, args: schema.KubernetesControlPlaneConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const configurationDirectory = output(args.configurationDirectory ?? '/etc/kubernetes/config');
    const caPem = output(args.caPem);
    const caKey = output(args.caKey);
    const connection = output(args.connection);
    const kubeApiServerPem = output(args.kubeApiServerPem);
    const kubeApiServerKey = output(args.kubeApiServerKey);
    const serviceAccountsPem = output(args.serviceAccountsPem);
    const serviceAccountsKey = output(args.serviceAccountsKey);
    const encryptionConfig = output(args.encryptionConfig);
    const kubeSchedulerConfig = output(args.kubeSchedulerConfig);
    const kubeControllerManagerKubeconfig = output(args.kubeControllerManagerKubeconfig);
    const kubeSchedulerKubeconfig = output(args.kubeSchedulerKubeconfig);
    const kubeApiServerPath = output(args.kubeApiServerPath);
    const kubeControllerManagerPath = output(args.kubeControllerManagerPath);
    const kubeSchedulerPath = output(args.kubeSchedulerPath);
    const kubectlPath = output(args.kubectlPath);

    const configurationMkdir = new Mkdir(`${name}-config`, {
      connection,
      create: {
        directory: configurationDirectory,
        parents: true,
      },
      delete: interpolate`rm -rf ${configurationDirectory}`,
    }, {parent: this });

    // const configurationKubeApiServerChmod = new Chmod(`${name}-kube-apiserver-chmod`, {
    //   connection,
    //   create: {
    //     files: [kubeApiServerPath],
    //     mode: '+x',
    //   },
    // }, { parent: this, dependsOn: configurationMkdir });
  }
}
