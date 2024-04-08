import { ComponentResourceOptions, Input, Output, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Mkdir } from '../tools';
import { File } from './file';
import { SystemdService } from './systemdService';
import { CommandBuilder } from '../tools/commandBuilder';

export class EtcdConfiguration extends schema.EtcdConfiguration {
  constructor(name: string, args: schema.EtcdConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const configurationDirectory = output(args.configurationDirectory ?? '/etc/etcd'); // Default value from schema?
    const dataDirectory = output(args.dataDirectory ?? '/var/lib/etcd'); // Default value from schema?
    const internalIp = output(args.internalIp);
    // const systemdDirectory = output(args.systemdDirectory ?? '/etc/systemd/system');

    const configurationMkdir = new Mkdir(`${name}-config`, {
      connection: args.connection,
      directory: configurationDirectory,
      parents: true,
    }, { parent: this });

    const dataMkdir = new Mkdir(`${name}-data`, {
      connection: args.connection,
      directory: configurationDirectory,
      parents: true,
    }, { parent: this });

    const caFilePath = interpolate`${configurationDirectory}/ca.pem`;
    const certFilePath = interpolate`${configurationDirectory}/kubernetes.pem`;
    const keyFilePath = interpolate`${configurationDirectory}/kubernetes-key.pem`;

    const caFile = new File(`${name}-ca`, {
      connection: args.connection,
      content: args.caPem,
      path: caFilePath,
    }, { parent: this, dependsOn: configurationMkdir });

    const certFile = new File(`${name}-cert`, {
      connection: args.connection,
      content: args.certPem,
      path: certFilePath,
    }, { parent: this, dependsOn: configurationMkdir });

    const keyFile = new File(`${name}-key`, {
      connection: args.connection,
      content: args.keyPem,
      path: keyFilePath,
    }, { parent: this, dependsOn: configurationMkdir });

    // TODO: Pretty print like the guide?
    const execStart = formatExecStart(
      args.etcdPath, // TODO
      name, // TODO: Review
      caFilePath,
      certFilePath,
      keyFilePath,
      dataDirectory,
      internalIp,
      {}, // TODO: Peers
    );

    // const systemdService = new SystemdService(name, {
    //   connection: args.connection,
    //   directory: systemdDirectory,
    //   unit: {
    //     description: 'etcd',
    //     documentation: ['https://github.com/etcd-io/etcd'],
    //   },
    //   service: {
    //     type: 'notify',
    //     execStart,
    //     restart: 'on-failure',
    //     restartSec: '5',
    //   },
    //   install: {
    //     wantedBy: ['multi-user.target'],
    //   },
    // }, { parent: this });

    this.caFile = caFile;
    this.certFile = certFile;
    this.configurationDirectory = configurationDirectory;
    this.configurationMkdir = configurationMkdir;
    this.dataDirectory = dataDirectory;
    this.dataMkdir = dataMkdir;
    this.internalIp = internalIp;
    this.keyFile = keyFile;
    // this.systemdService = systemdService,

    this.registerOutputs({
      caFile,
      certFile,
      configurationDirectory,
      configurationMkdir,
      dataDirectory,
      dataMkdir,
      internalIp,
      keyFile,
      name,
      // systemdService,
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

function formatSystemdFile(execStart: Input<string>): Output<string> {
  // Would be nice to not have [Unit] hangout up here all alone
  return interpolate`[Unit]
Description=etcd
Documentation=https://github.com/etcd-io/etcd

[Service]
Type=notify
ExecStart=${execStart}
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
`;
}
