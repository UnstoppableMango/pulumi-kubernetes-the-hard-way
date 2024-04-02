import { ComponentResourceOptions, Input, Output, interpolate, output } from '@pulumi/pulumi';
import { RandomString } from '@pulumi/random';
import * as schema from '../schema-types';
import { Mkdir, Mv, Tar } from '../tools';
import { Download } from './download';
import { File } from './file';

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
    const configurationDirectory = output(args.configurationDirectory); // Default value?
    const downloadDirectory = this.getDownloadDirectory(args.downloadDirectory);
    const installDirectory = output(args.installDirectory ?? EtcdInstall.defaultInstallDirectory);
    const internalIp = output(args.internalIp);
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

    const caFilePath = interpolate`${configurationDirectory}/ca.pem`;
    const certFilePath = interpolate`${configurationDirectory}/kubernetes.pem`;
    const keyFilePath = interpolate`${configurationDirectory}/kubernetes-key.pem`;

    const caFile = new File(`${name}-ca`, {
      connection: args.connection,
      content: args.caPem,
      path: caFilePath,
    }, { parent: this });

    const certFile = new File(`${name}-cert`, {
      connection: args.connection,
      content: args.certPem,
      path: certFilePath,
    }, { parent: this });

    const keyFile = new File(`${name}-key`, {
      connection: args.connection,
      content: args.keyPem,
      path: keyFilePath,
    }, { parent: this });

    this.architecture = architecture;
    this.archiveName = archiveName;
    this.caFile = caFile;
    this.certFile = certFile;
    this.download = download;
    this.downloadDirectory = downloadDirectory;
    this.downloadMkdir = downloadMkdir;
    this.installDirectory = installDirectory;
    this.installMkdir = installMkdir;
    this.internalIp = internalIp;
    this.keyFile = keyFile;
    this.mvEtcd = mvEtcd;
    this.mvEtcdctl = mvEtcdctl;
    this.name = output(name);
    this.tar = tar;
    this.url = url;
    this.version = version;

    this.registerOutputs({
      architecture,
      archiveName,
      download,
      downloadDirectory,
      downloadMkdir,
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

function systemdFile(nodeName: string, internalIp: Input<string>): Output<string> {
  return interpolate`
[Unit]
Description=etcd
Documentation=https://github.com/etcd-io/etcd

[Service]
Type=notify
ExecStart=/usr/local/bin/etcd \\
  --name ${nodeName} \\
  --cert-file=/etc/etcd/kubernetes.pem \\
  --key-file=/etc/etcd/kubernetes-key.pem \\
  --peer-cert-file=/etc/etcd/kubernetes.pem \\
  --peer-key-file=/etc/etcd/kubernetes-key.pem \\
  --trusted-ca-file=/etc/etcd/ca.pem \\
  --peer-trusted-ca-file=/etc/etcd/ca.pem \\
  --peer-client-cert-auth \\
  --client-cert-auth \\
  --initial-advertise-peer-urls https://${internalIp}:2380 \\
  --listen-peer-urls https://${internalIp}:2380 \\
  --listen-client-urls https://${internalIp}:2379,https://127.0.0.1:2379 \\
  --advertise-client-urls https://${internalIp}:2379 \\
  --initial-cluster-token etcd-cluster-0 \\
  --initial-cluster controller-0=https://10.240.0.10:2380,controller-1=https://10.240.0.11:2380,controller-2=https://10.240.0.12:2380 \\
  --initial-cluster-state new \\
  --data-dir=/var/lib/etcd
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
`;
}
