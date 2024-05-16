import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Chmod, Mkdir } from '../tools';
import { File } from './file'

export class KubeApiServerConfiguration extends schema.KubeApiServerConfiguration {
  constructor(name: string, args: schema.KubeApiServerConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const configurationDirectory = output(args.configurationDirectory ?? '/etc/kubernetes/config');
    const caPem = output(args.caPem);
    const caKey = output(args.caKey);
    const connection = output(args.connection);
    const certPem = output(args.certPem);
    const keyPem = output(args.keyPem);
    const serviceAccountsPem = output(args.serviceAccountsPem);
    const serviceAccountsKey = output(args.serviceAccountsKey);
    const encryptionConfig = output(args.encryptionConfig);
    const path = output(args.path ?? '/usr/local/bin');
    const kubectlPath = output(args.kubectlPath);

    const dataDirectory = '/var/lib/kubernetes';

    const configurationMkdir = new Mkdir(`${name}-config`, {
      connection,
      create: {
        directory: configurationDirectory,
        parents: true,
      },
      delete: interpolate`rm -rf ${configurationDirectory}`,
    }, {parent: this });

    const dataMkdir = new Mkdir(`${name}-config`, {
      connection,
      create: {
        directory: dataDirectory,
        parents: true,
      },
      delete: interpolate`rm -rf ${dataDirectory}`,
    }, {parent: this });

    const certFilePath = interpolate`${dataDirectory}/kubernetes.pem`;
    const keyFilePath = interpolate`${dataDirectory}/kubernetes-key.pem`;

    const certFile = new File(`${name}-cert`, {
      connection,
      content: args.certPem,
      path: certFilePath,
    }, { parent: this, dependsOn: dataMkdir });

    const keyFile = new File(`${name}-key`, {
      connection,
      content: args.keyPem,
      path: keyFilePath,
    }, { parent: this, dependsOn: dataMkdir });

    const encryptionConfigFile = new File(`${name}-key`, {
      connection,
      content: args.encryptionConfig,
      path: dataDirectory,
    }, { parent: this, dependsOn: dataMkdir });

    const configurationKubeApiServerChmod = new Chmod(`${name}-kube-apiserver-chmod`, {
      connection,
      create: {
        files: [path],
        mode: '+x',
      },
    }, { parent: this });

    this.caPem = caPem;
    this.certPem = certPem;
    // this.certFilePath = certFilePath;
    // this.configurationKubeApiServerChmod = configurationKubeApiServerChmod;
    this.configurationDirectory = configurationDirectory;
    this.configurationMkdir = configurationMkdir;
    // this.dataDirectory = dataDirectory;
    // this.dataMkdir = dataMkdir;
    this.path = path;
    // this.keyFilePath = keyFilePath;

    this.registerOutputs({
      caPem,
      certPem,
      certFilePath,
      configurationKubeApiServerChmod,
      configurationDirectory,
      configurationMkdir,
      dataDirectory,
      dataMkdir,
      path,
      keyFilePath,
    });
  }
}
