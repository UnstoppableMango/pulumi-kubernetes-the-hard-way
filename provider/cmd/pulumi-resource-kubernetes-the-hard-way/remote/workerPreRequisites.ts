import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import { Command } from '@pulumi/command/remote';
import * as schema from '../schema-types';

export class WorkerPreRequisites extends schema.WorkerPreRequisites {
  constructor(name: string, args: schema.WorkerPreRequisitesArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const connection = output(args.connection);
    const triggers = output(args.triggers ?? []);

    const conntrack = new Command(`${name}-conntrack`, {
      connection,
      create: 'which conntrack',
      triggers,
    }, { parent: this });

    const ipset = new Command(`${name}-ipset`, {
      connection,
      create: 'which ipset',
      triggers,
    }, { parent: this });

    const socat = new Command(`${name}-socat`, {
      connection,
      create: 'which socat',
      triggers,
    }, { parent: this });

    const swap = new Command(`${name}-swap`, {
      connection,
      create: 'swapon --show',
      triggers,
    }, { parent: this });

    this.connection = connection;
    this.conntrack = conntrack;
    this.ipset = ipset;
    this.socat = socat;
    this.swap = swap;
    this.triggers = triggers;

    this.registerOutputs({
      connection,
      conntrack,
      ipset,
      socat,
      swap,
      triggers,
    });
  }
}
