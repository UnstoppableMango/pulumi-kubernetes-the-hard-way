import { ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { SystemdService } from './systemdService';
import { CommandBuilder } from '../tools/commandBuilder';

export class KubeProxyService extends schema.KubeProxyService {
  constructor(name: string, args: schema.KubeProxyServiceArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const connection = output(args.connection);
    const configuration = output(args.configuration);
    const description = output(args.description ?? 'Kubernetes Kube Proxy');
    const documentation = output(args.documentation ?? 'https://github.com/kubernetes/kubernetes');
    const restart = output(args.restart ?? 'on-failure');
    const restartSec = output(args.restartSec ?? '5');
    const wantedBy = output(args.wantedBy ?? 'multi-user.target');

    const service = new SystemdService(name, {
      connection,
      directory: args.directory,
      unitName: 'kube-proxy',
      unit: {
        description,
        documentation: [documentation],
      },
      service: {
        execStart: formatExecStart(
          configuration.kubeProxyPath,
          configuration.configurationFilePath,
        ),
        restart,
        restartSec,
      },
      install: {
        wantedBy: [wantedBy],
      },
    }, { parent: this });

    this.connection = connection;
    this.description = description;
    this.directory = service.directory;
    this.documentation = documentation;
    this.restart = restart;
    this.restartSec = restartSec;
    this.service = service;
    this.wantedBy = wantedBy;

    this.registerOutputs({
      connection,
      description,
      directory: service.directory,
      documentation,
      restart,
      restartSec,
      service,
      wantedBy,
    });
  }
}

const formatExecStart = (
  kubeProxyPath: Input<string>,
  configFilePath: Input<string>,
): Output<string> => {
  return new CommandBuilder(kubeProxyPath)
    .option('--config', configFilePath)
    .command;
};
