import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Mkdir } from '../tools';
import { CniBridgePluginConfiguration } from './cniBridgePluginConfiguration';
import { CniLoopbackPluginConfiguration } from './cniLoopbackPluginConfiguration';

export class CniPluginConfiguration extends schema.CniPluginConfiguration {
  constructor(name: string, args: schema.CniPluginConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const connection = output(args.connection);
    const directory = output(args.directory ?? '/etc/cni/net.d');
    const subnet = output(args.subnet);

    const mkdir = new Mkdir(name, {
      connection,
      create: {
        parents: true,
        directory: directory,
      },
      delete: interpolate`rm -rf ${directory}`,
    }, { parent: this });

    const bridge = new CniBridgePluginConfiguration(name, {
      connection,
      subnet,
      path: interpolate`${directory}/10-bridge.conf`,
    }, { parent: this, dependsOn: mkdir });

    const loopback = new CniLoopbackPluginConfiguration(name, {
      connection,
      path: interpolate`${directory}/99-loopback.conf`,
    }, { parent: this, dependsOn: mkdir });

    this.bridge = bridge;
    this.connection = connection;
    this.directory = directory;
    this.mkdir = mkdir;
    this.loopback = loopback;
    this.subnet = subnet;

    this.registerOutputs({
      bridge,
      connection,
      directory,
      mkdir,
      loopback,
      subnet,
    });
  }
}
