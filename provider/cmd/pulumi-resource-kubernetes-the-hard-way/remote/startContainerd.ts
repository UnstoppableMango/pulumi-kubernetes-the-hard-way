import { ComponentResourceOptions } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { startSystemdService } from './startSystemdService';

export class StartContainerd extends schema.StartContainerd {
  constructor(name: string, args: schema.StartContainerdArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const { connection, daemonReload, enable, start } = startSystemdService(name, {
      unit: 'containerd',
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
