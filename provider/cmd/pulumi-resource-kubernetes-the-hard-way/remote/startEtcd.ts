import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Systemctl } from '../tools';

export class StartEtcd extends schema.StartEtcd {
  constructor(name: string, args: schema.StartEtcdArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const connection = output(args.connection);
    const unit = 'etcd';

    const daemonReload = new Systemctl(`${name}-daemon-reload`, {
      connection,
      command: 'daemon-reload',
      unit,
    }, { parent: this });

    const enable = new Systemctl(`${name}-enable`, {
      connection,
      command: 'enable',
      unit,
    }, { parent: this });

    const start = new Systemctl(`${name}-start`, {
      connection,
      command: 'start',
      unit,
    }, { parent: this });

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