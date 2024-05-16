import { ComponentResourceOptions, Input, Output, all, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { CommandBuilder } from '../tools/commandBuilder';
import { SystemdService } from './systemdService';

export class KubeApiServerService extends schema.KubeApiServerService {
  constructor(name: string, args: schema.KubeApiServerServiceArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const configuration = output(args.configuration);
    const connection = output(args.connection);
    const description = output(args.description ?? 'Kubernetes API Server');
    const documentation = output(args.documentation ?? 'https://github.com/kubernetes/kubernetes');
    const restart = output(args.restart ?? 'on-failure');
    const restartSec = output(args.restartSec ?? '5');
    const wantedBy = output(args.wantedBy ?? 'multi-user.target');

    // TODO: Pretty print like the guide?
    const execStart = formatExecStart(
      configuration.etcdPath,
      configuration.name, // TODO: Review
      configuration.caFilePath,
      configuration.certFilePath,
      configuration.keyFilePath,
      configuration.dataDirectory,
      configuration.internalIp,
    );

    const service = new SystemdService(name, {
      connection,
      directory: args.directory,
      unitName: 'kube-apiserver',
      unit: {
        description,
        documentation: [documentation],
      },
      service: {
        execStart,
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
  caPath: Input<string>,
  certPath: Input<string>,
  keyPath: Input<string>,
  path: Input<string>,
  dataDirectory: Input<string>,
): Output<string> => {
  const peerUrl = interpolate`https://${internalIp}:${peerPort}`;
  const clientUrl = interpolate`https://${internalIp}:${clientPort}`;
  const localhostUrl = interpolate`https://127.0.0.1:${clientPort}`;

  // TODO: Can we clean this up
  const peerMapping = output(peers).apply(x => x.map(c => {
    return interpolate`${c.name}=https://${c.internalIp}:${peerPort}`;
  }).concat(interpolate`${nodeName}=${peerUrl}`));

  const initialCluster = output(peerMapping).apply(m => m.join(','));

  return new CommandBuilder(etcdPath)
    .option('--name', nodeName)
    .option('--cert-file', certPath)
    .option('--key-file', keyPath)
    .option('--peer-cert-file', certPath)
    .option('--peer-key-file', keyPath)
    .option('--trusted-ca-file', caPath)
    .option('--peer-trusted-ca-file', caPath)
    .option('--peer-client-cert-auth', true)
    .option('--client-cert-auth', true)
    .option('--initial-advertise-peer-urls', peerUrl)
    .option('--listen-peer-urls', peerUrl)
    .option('--listen-client-urls', interpolate`${clientUrl}`) // TODO: We can be a little smarter about this
    .option('--advertise-client-urls', clientUrl)
    .option('--initial-cluster-token', 'etcd-cluster-0') // TODO
    .option('--initial-cluster', initialCluster)
    .option('--initial-cluster-state', 'new')
    .option('--data-dir', dataDirectory)
    .command;
};
