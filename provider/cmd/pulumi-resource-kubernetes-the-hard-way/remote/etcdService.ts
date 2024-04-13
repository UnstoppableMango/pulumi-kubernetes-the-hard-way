import { ComponentResourceOptions, Input, Output, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { SystemdService } from './systemdService';
import { CommandBuilder } from '../tools/commandBuilder';

export class EtcdService extends schema.EtcdService {
  constructor(name: string, args: schema.EtcdServiceArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const configuration = output(args.configuration);
    const connection = output(args.connection);
    const description = output(args.description ?? 'etcd');
    const documentation = output(args.documentation ?? 'https://github.com/etcd-io/etcd');
    const restart = output(args.restart ?? 'on-failure');
    const restartSec = output(args.restartSec ?? '5');
    const wantedBy = output(args.wantedBy ?? 'multi-user.target');

    // TODO: Pretty print like the guide?
    const execStart = formatExecStart(
      output(configuration.etcdPath), // TODO
      name, // TODO: Review
      output(configuration.caFile.path),
      output(configuration.certFile.path),
      output(configuration.keyFile.path),
      output(configuration.dataDirectory),
      output(configuration.internalIp),
      {}, // TODO: Peers
    );

    const service = new SystemdService(name, {
      connection,
      directory: args.directory,
      unit: {
        description,
        documentation: [documentation],
      },
      service: {
        type: 'notify',
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
  etcdPath: Input<string>,
  nodeName: Input<string>,
  caPath: Input<string>,
  certPath: Input<string>,
  keyPath: Input<string>,
  dataDirectory: Input<string>,
  internalIp: Input<string>,
  peers: Record<string, Input<string>>,
  peerPort: number = 2380,
  clientPort: number = 2379
): Output<string> => {
  const peerUrl = interpolate`https://${internalIp}:${peerPort}`;
  const clientUrl = interpolate`https://${internalIp}:${clientPort}`;
  const localhostUrl = interpolate`https://127.0.0.1:${clientPort}`;

  const peerMapping = Object.entries(peers).map(([name, ip]) => {
    return interpolate`${name}=https://${ip}:${peerPort}`;
  }).concat(interpolate`${nodeName}=${peerUrl}`);

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
    .option('--listen-client-urls', interpolate`${clientUrl},${localhostUrl}`)
    .option('--advertise-client-urls', clientUrl)
    .option('--initial-cluster-token', 'etcd-cluster-0') // TODO
    .option('--initial-cluster', initialCluster)
    .option('--initial-cluster-state', 'new')
    .option('--data-dir', dataDirectory)
    .command;
};
