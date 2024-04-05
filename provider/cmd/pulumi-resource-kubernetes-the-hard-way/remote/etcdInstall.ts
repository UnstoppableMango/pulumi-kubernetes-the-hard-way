import { ComponentResourceOptions, Input, Output, interpolate, output } from '@pulumi/pulumi';
import { RandomString } from '@pulumi/random';
import * as schema from '../schema-types';
import { Etcdctl, Mkdir, Mv, Tar } from '../tools';
import { Download } from './download';
import { File } from './file';
import { CommandBuilder } from '../tools/commandBuilder';
import { SystemdService } from './systemdService';

export type Architecture = 'amd64' | 'arm64';

export class EtcdInstall extends schema.EtcdInstall {
  public static readonly defaultArch: Architecture = 'amd64';
  public static readonly defaultInstallDirectory: string = '/usr/local/bin';
  public static readonly defaultVersion: string = '3.4.15'; // TODO: Versioning
  private readonly _nameInput: string;

  constructor(name: string, args: schema.EtcdInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    this._nameInput = name;

    // Rehydrating
    if (opts?.urn) return;

    const architecture = output(args.architecture ?? EtcdInstall.defaultArch);
    const configurationDirectory = output(args.configurationDirectory ?? '/etc/etcd'); // Default value from schema?
    const dataDirectory = output(args.dataDirectory ?? '/var/lib/etcd'); // Default value from schema?
    const downloadDirectory = this.getDownloadDirectory(args.downloadDirectory);
    const installDirectory = output(args.installDirectory ?? EtcdInstall.defaultInstallDirectory);
    const internalIp = output(args.internalIp);
    const systemdDirectory = output(args.systemdDirectory ?? '/etc/systemd/system');
    const version = output(args.version ?? EtcdInstall.defaultVersion); // TODO: Stateful versioning?
    const archiveName = interpolate`etcd-v${version}-linux-${architecture}.tar.gz`;
    const url = interpolate`https://github.com/etcd-io/etcd/releases/download/v${version}/${archiveName}`;

    // TODO: Permission checks?
    // TODO: Caching? Put archive/bins into ~/.kthw/cache so i.e. installDirectory changes, tarball doesn't need to be re-downloaded.
    // TODO: General update logic

    const downloadMkdir = new Mkdir(`${name}-download`, {
      connection: args.connection,
      directory: downloadDirectory,
      parents: true,
      removeOnDelete: true,
    }, { parent: this });

    const download = new Download(name, {
      connection: args.connection,
      destination: downloadDirectory,
      url,
    }, { parent: this, dependsOn: downloadMkdir });

    const tar = new Tar(name, {
      connection: args.connection,
      archive: interpolate`${download.destination}/${archiveName}`,
      directory: download.destination,
      gzip: true,
      stripComponents: 1,
    }, { parent: this, dependsOn: download });

    const installMkdir = new Mkdir(`${name}-install`, {
      connection: args.connection,
      directory: installDirectory,
      parents: true,
    }, { parent: this });

    const etcdPath = interpolate`${installDirectory}/etcd`;
    const etcdctlPath = interpolate`${installDirectory}/etcdctl`;

    const mvEtcd = new Mv(`${name}-etcd`, {
      connection: args.connection,
      source: interpolate`${download.destination}/etcd`,
      dest: etcdPath,
    }, { parent: this, dependsOn: [tar, installMkdir] });

    const mvEtcdctl = new Mv(`${name}-etcdctl`, {
      connection: args.connection,
      source: interpolate`${download.destination}/etcdctl`,
      dest: etcdctlPath,
    }, { parent: this, dependsOn: [tar, installMkdir] });

    // TODO: Rm archive or tmp dir?

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
      etcdPath,
      name, // TODO: Review
      caFilePath,
      certFilePath,
      keyFilePath,
      dataDirectory,
      internalIp,
      {}, // TODO: Peers
    );

    const systemdFile = new SystemdService(name, {
      connection: args.connection,
      directory: systemdDirectory,
      unit: {
        description: 'etcd',
        documentation: ['https://github.com/etcd-io/etcd'],
      },
      service: {
        type: 'notify',
        execStart,
        restart: 'on-failure',
        restartSec: '5',
      },
      install: {
        wantedBy: ['multi-user.target'],
      },
    }, { parent: this });

    this.architecture = architecture;
    this.archiveName = archiveName;
    this.caFile = caFile;
    this.certFile = certFile;
    this.configurationDirectory = configurationDirectory;
    this.configurationMkdir = configurationMkdir;
    this.dataDirectory = dataDirectory;
    this.dataMkdir = dataMkdir;
    this.download = download;
    this.downloadDirectory = downloadDirectory;
    this.downloadMkdir = downloadMkdir;
    this.etcdPath = etcdPath;
    this.etcdctlPath = etcdctlPath;
    this.installDirectory = installDirectory;
    this.installMkdir = installMkdir;
    this.internalIp = internalIp;
    this.keyFile = keyFile;
    this.mvEtcd = mvEtcd;
    this.mvEtcdctl = mvEtcdctl;
    this.name = output(name);
    this.systemdService = systemdFile,
    this.tar = tar;
    this.url = url;
    this.version = version;

    this.registerOutputs({
      architecture,
      archiveName,
      download,
      downloadDirectory,
      downloadMkdir,
      etcdPath,
      etcdctlPath,
      installDirectory,
      installMkdir,
      mvEtcd,
      mvEtcdctl,
      name,
      tar,
      url,
      version,
    });
  }

  public etcdctl(inputs: schema.EtcdInstall_etcdctlInputs): Output<schema.EtcdInstall_etcdctlOutputs> {
    throw new Error('not implemented');
  }

  private getDownloadDirectory(input?: Input<string>): Output<string> {
    if (input) {
      return output(input);
    }

    // TODO: Mktemp?
    const name = this._nameInput;
    const random = new RandomString(name, {
      length: 8,
      special: false,
      upper: false,
    }, { parent: this });

    return interpolate`/tmp/${random.result}`;
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
