import { ComponentResourceOptions } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { startSystemdService } from './startSystemdService';

export class StartKubelet extends schema.StartKubelet {
  constructor(name: string, args: schema.StartKubeletArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const { connection, daemonReload, enable, start } = startSystemdService(name, {
      unit: 'kubelet',
      connection: args.connection,
    }, this);

    this.connection = connection;
    this.daemonReload = daemonReload;
    this.enable = enable;
    this.start = start;

    this.registerOutputs({
      connection,
      daemonReload,
      enable,
      start,
    });
  }
}
