import { ComponentResourceOptions, Input, Output, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Mkdir } from '../tools';
import { File } from './file';

export class EtcdConfiguration extends schema.EtcdConfiguration {
  constructor(name: string, args: schema.EtcdConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const configurationDirectory = output(args.configurationDirectory ?? '/etc/etcd'); // Default value from schema?
    const dataDirectory = output(args.dataDirectory ?? '/var/lib/etcd'); // Default value from schema?
    const internalIp = output(args.internalIp);

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
