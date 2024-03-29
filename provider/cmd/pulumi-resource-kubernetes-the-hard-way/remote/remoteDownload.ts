import { ComponentResource, ComponentResourceOptions, Input, Output, output } from '@pulumi/pulumi';
import { remote } from '@pulumi/command/types/input';
import { Mkdir, Wget } from '../tools';

export interface RemoteDownloadArgs {
  connection: Input<remote.ConnectionArgs>;
  destination: Input<string>;
  removeOnDelete?: Input<boolean>;
  url: Input<string>;
}

export class RemoteDownload extends ComponentResource {
  public readonly mkdir: Mkdir;
  public readonly destination: Output<string>
  public readonly url: Output<string>;
  public readonly wget: Wget;

  constructor(name: string, args: RemoteDownloadArgs, opts?: ComponentResourceOptions) {
    super('kubernetes-the-hard-way:remote:RemoteDownload', name, args, opts);

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
