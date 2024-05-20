import { ComponentResourceOptions, Input, Output, Resource, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Systemctl } from '../tools';
import { input as inputs, output as outputs } from '@pulumi/command/types';

export interface StartSystemdServiceArgs {
  unit: string;
  connection: Input<inputs.remote.ConnectionArgs>;
}

export interface StartSystemdServiceResult {
  connection: Output<outputs.remote.Connection>;
  daemonReload: Systemctl;
  enable: Systemctl;
  start: Systemctl;
}

export function startSystemdService(name: string, args: StartSystemdServiceArgs, parent: Resource): StartSystemdServiceResult {
  const connection = output(args.connection);

  const daemonReload = new Systemctl(`${name}-daemon-reload`, {
    connection,
    create: { command: 'daemon-reload' },
  }, { parent });

  const enable = new Systemctl(`${name}-enable`, {
    connection,
    create: {
      command: 'enable',
      unit: args.unit,
    },
    delete: {
      command: 'disable',
      unit: args.unit,
    },
  }, { parent, dependsOn: daemonReload });

  const start = new Systemctl(`${name}-start`, {
    connection,
    create: {
      command: 'start',
      unit: args.unit,
    },
    delete: {
      command: 'stop',
      unit: args.unit,
    },
  }, { parent, dependsOn: enable });

  return {
    connection,
    daemonReload,
    enable,
    start,
  };
}
