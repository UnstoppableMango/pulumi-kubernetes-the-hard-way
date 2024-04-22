import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import { Mkdir, Rm, Wget } from '../tools';
import * as types from '../schema-types';
import path = require('node:path');

export class Download extends types.Download {
  constructor(name: string, args: types.DownloadArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const destination = output(args.destination);
    const url = output(args.url);
    const file = url.apply(path.basename);

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
      delete: interpolate`rm -f ${file}`,
    }, { parent: this, dependsOn: mkdir });

    this.mkdir = mkdir;
    this.destination = destination;
    this.url = url;
    this.wget = wget;

    this.registerOutputs({ mkdir, destination, url, wget });
  }
}
