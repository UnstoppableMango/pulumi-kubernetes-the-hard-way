import { ComponentResourceOptions, Input, Output, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Chmod, Mkdir } from '../tools';
import { File } from './file';

export class EtcdConfiguration extends schema.EtcdConfiguration {
  constructor(name: string, args: schema.EtcdConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const caPem = output(args.caPem);
    const certPem = output(args.certPem);
    const configurationDirectory = output(args.configurationDirectory ?? '/etc/etcd'); // Default value from schema?
    const connection = output(args.connection);
    const dataDirectory = output(args.dataDirectory ?? '/var/lib/etcd'); // Default value from schema?
    const etcdPath = output(args.etcdPath);
    const internalIp = output(args.internalIp);

    const configurationMkdir = new Mkdir(`${name}-config`, {
      connection,
      directory: configurationDirectory,
      parents: true,
    }, { parent: this });

    const configurationChmod = new Chmod(`${name}-data-chmod`, {
      connection,
      files: configurationDirectory,
      mode: '700',
    }, { parent: this, dependsOn: configurationMkdir });

    const dataMkdir = new Mkdir(`${name}-data`, {
      connection,
      directory: configurationDirectory,
      parents: true,
    }, { parent: this });

    const caFilePath = interpolate`${configurationDirectory}/ca.pem`;
    const certFilePath = interpolate`${configurationDirectory}/kubernetes.pem`;
    const keyFilePath = interpolate`${configurationDirectory}/kubernetes-key.pem`;

    const caFile = new File(`${name}-ca`, {
      connection,
      content: args.caPem,
      path: caFilePath,
    }, { parent: this, dependsOn: configurationChmod });

    const certFile = new File(`${name}-cert`, {
      connection,
      content: args.certPem,
      path: certFilePath,
    }, { parent: this, dependsOn: configurationChmod });

    const keyFile = new File(`${name}-key`, {
      connection,
      content: args.keyPem,
      path: keyFilePath,
    }, { parent: this, dependsOn: configurationMkdir });

    this.caFile = caFile;
    this.caPem = caPem;
    this.certFile = certFile;
    this.certPem = certPem;
    this.configurationChmod = configurationChmod;
    this.configurationDirectory = configurationDirectory;
    this.configurationMkdir = configurationMkdir;
    this.dataDirectory = dataDirectory;
    this.dataMkdir = dataMkdir;
    this.etcdPath = etcdPath;
    this.internalIp = internalIp;
    this.keyFile = keyFile;

    const value: schema.EtcdConfigurationPropsInputs = {
      name,
      etcdPath,
      internalIp,
      dataDirectory,
      caFilePath: caFile.path,
      certFilePath: certFile.path,
      keyFilePath: keyFile.path,
    };

    this.value = output(value);

    this.registerOutputs({
      caFile,
      caPem,
      certFile,
      certPem,
      configurationChmod,
      configurationDirectory,
      configurationMkdir,
      dataDirectory,
      dataMkdir,
      etcdPath,
      internalIp,
      keyFile,
      name,
      value,
    });
  }
}

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
