import { ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { SystemdService } from './systemdService';
import { CommandBuilder } from '../tools/commandBuilder';

export class KubeletService extends schema.KubeletService {
  constructor(name: string, args: schema.KubeletServiceArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const after = output(args.after ?? ['containerd.service']);
    const connection = output(args.connection);
    const configuration = output(args.configuration);
    const description = output(args.description ?? 'Kubernetes Kubelet');
    const documentation = output(args.documentation ?? 'https://github.com/kubernetes/kubernetes');
    const requires = output(args.requires ?? ['containerd.service']);
    const restart = output(args.restart ?? 'on-failure');
    const restartSec = output(args.restartSec ?? '5');
    const wantedBy = output(args.wantedBy ?? 'multi-user.target');

    const service = new SystemdService(name, {
      connection,
      directory: args.directory,
      unitName: 'kubelet',
      unit: {
        description,
        documentation: [documentation],
        after,
        requires,
      },
      service: {
        execStart: formatExecStart(
          configuration.kubeletPath,
          configuration.configurationFilePath,
          configuration.kubeconfigPath,
          configuration.registerNode,
          configuration.v,
        ),
        restart,
        restartSec,
      },
      install: {
        wantedBy: [wantedBy],
      },
    }, { parent: this });

    this.after = after;
    this.connection = connection;
    this.description = description;
    this.directory = service.directory;
    this.documentation = documentation;
    this.requires = requires;
    this.restart = restart;
    this.restartSec = restartSec;
    this.service = service;
    this.wantedBy = wantedBy;

    this.registerOutputs({
      after,
      connection,
      description,
      directory: service.directory,
      documentation,
      requires,
      restart,
      restartSec,
      service,
      wantedBy,
    });
  }
}

const formatExecStart = (
  kubeletPath: Input<string>,
  configFilePath: Input<string>,
  kubeconfigPath: Input<string>,
  registerNode: Input<boolean>,
  v: Input<number>,
): Output<string> => {
  return new CommandBuilder(kubeletPath)
    .option('--config', configFilePath)
    .option('--kubeconfig', kubeconfigPath)
    .option('--register-node', registerNode)
    .option('--v', v)
    .command;
};
