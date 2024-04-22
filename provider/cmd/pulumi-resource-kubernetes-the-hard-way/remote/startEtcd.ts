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
      create: { command: 'daemon-reload' },
    }, { parent: this });

    const enable = new Systemctl(`${name}-enable`, {
      connection,
      create: {
        command: 'enable',
        unit,
      },
      delete: {
        command: 'disable',
        unit,
      },
    }, { parent: this, dependsOn: daemonReload });

    const start = new Systemctl(`${name}-start`, {
      connection,
      create: {
        command: 'start',
        unit,
      },
      delete: {
        command: 'stop',
        unit,
      },
    }, { parent: this, dependsOn: enable });

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
