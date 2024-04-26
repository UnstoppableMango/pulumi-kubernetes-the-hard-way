import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { EtcdInstall } from './etcdInstall';
import { EtcdConfiguration } from './etcdConfiguration';
import { StartEtcd } from './startEtcd';
import { EtcdService } from './etcdService';

export class ProvisionEtcd extends schema.ProvisionEtcd {
  constructor(name: string, args: schema.ProvisionEtcdArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const bundle = output(args.bundle);
    const connection = output(args.connection);
    const internalIp = output(args.internalIp);

    const install = new EtcdInstall(name, {
      connection,
      architecture: args.architecture,
      directory: args.binaryDirectory,
      version: args.version,
    }, { parent: this });

    const configuration = new EtcdConfiguration(name, {
      connection,
      caPem: bundle.caPem,
      certPem: bundle.certPem,
      etcdPath: install.etcdPath,
      internalIp,
      keyPem: bundle.keyPem,
      configurationDirectory: args.configurationDirectory,
      dataDirectory: args.dataDirectory,
    }, { parent: this, dependsOn: install });

    // TODO: Big chicken and egg problem here with peers
    const service = new EtcdService(name, {
      connection,
      configuration: configuration.value,
      directory: configuration.configurationDirectory,
      // TODO: Other options
    }, { parent: this, dependsOn: configuration });

    const start = new StartEtcd(name, {
      connection,
    }, { parent: this, dependsOn: service });

    this.architecture = install.architecture
    this.binaryDirectory = install.directory;
    this.bundle = bundle;
    this.configuration = configuration;
    this.configurationDirectory = configuration.configurationDirectory
    this.connection = connection;
    this.dataDirectory = configuration.dataDirectory;
    this.install = install;
    this.internalIp = internalIp;
    this.start = start;

    this.registerOutputs({
      architecture: this.architecture,
      binaryDirectory: this.binaryDirectory,
      bundle,
      configuration,
      configurationDirectory: this.configurationDirectory,
      connection,
      dataDirectory: this.dataDirectory,
      install,
      internalIp,
      start,
    });
  }
}
