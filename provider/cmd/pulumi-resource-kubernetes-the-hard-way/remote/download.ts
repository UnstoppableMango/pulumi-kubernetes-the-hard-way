import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import { Mkdir, Wget } from '../tools';
import * as types from '../schema-types';

export class Download extends types.Download {
  constructor(name: string, args: types.DownloadArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const destination = output(args.destination);
    const removeOnDelete = output(args.removeOnDelete ?? false);
    const url = output(args.url);

    const mkdir = new Mkdir(name, {
      connection: args.connection,
      directory: destination,
      parents: true,
      removeOnDelete,
    }, { parent: this });

    const wget = new Wget(name, {
      connection: args.connection,
      url: args.url,
      directoryPrefix: mkdir.directory,
    }, { parent: this });

    this.mkdir = mkdir;
    this.destination = destination;
    this.url = url;
    this.wget = wget;

    this.registerOutputs({ mkdir, destination, url, wget });
  }
}
