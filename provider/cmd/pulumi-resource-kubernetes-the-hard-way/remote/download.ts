import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import { Mkdir, Rm, Wget } from '../tools';
import * as types from '../schema-types';

export class Download extends types.Download {
  constructor(name: string, args: types.DownloadArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const destination = output(args.destination);
    const url = output(args.url);

    // TODO: Remove on delete

    const mkdir = new Mkdir(name, {
      connection: args.connection,
      create: {
        directory: destination,
        parents: true,
      },
    }, { parent: this });

    const wget = new Wget(name, {
      connection: args.connection,
      create: {
        url: [args.url],
        directoryPrefix: destination,
      },
    }, { parent: this, dependsOn: mkdir });

    this.mkdir = mkdir;
    this.destination = destination;
    this.url = url;
    this.wget = wget;

    this.registerOutputs({ mkdir, destination, url, wget });
  }
}
