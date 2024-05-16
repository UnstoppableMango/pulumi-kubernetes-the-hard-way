import { ComponentResourceOptions, jsonStringify, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { File } from './file';

export class CniLoopbackPluginConfiguration extends schema.CniLoopbackPluginConfiguration {
  constructor(name: string, args: schema.CniLoopbackPluginConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const cniVersion = output(args.cniVersion ?? '1.1.0');
    const connection = output(args.connection);
    const loopbackName = output(args.name ?? 'bridge');
    const path = output(args.path ?? '/var/lib/kubernetes');
    const type = output(args.type ?? 'loopback');

    const file = new File(name, {
      connection,
      content: jsonStringify({
        cniVersion,
        name: loopbackName,
        type,
      }),
      path,
    }, { parent: this });

    this.cniVersion = cniVersion;
    this.connection = connection;
    this.file = file;
    this.name = loopbackName;
    this.type = type;

    this.registerOutputs({
      cniVersion,
      connection,
      file,
      loopbackName,
      type,
    });
  }
}
