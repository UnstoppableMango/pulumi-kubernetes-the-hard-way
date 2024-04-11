import * as YAML from 'yaml';
import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { File } from './file';

export class StaticPod extends schema.StaticPod {
  constructor(name: string, args: schema.StaticPodArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const connection = output(args.connection);
    const fileName = output(args.fileName ?? 'kube-vip-pod.yaml');
    const path = interpolate`/etc/kubernetes/manifests/${fileName}`;
    const pod = output(args.pod);

    const file = new File(name, {
      connection,
      content: pod.apply(YAML.stringify),
      path,
    }, { parent: this });

    this.connection = connection;
    this.file = file;
    this.fileName = fileName;
    this.path = path;

    this.registerOutputs({
      connection,
      file,
      fileName,
      path,
      pod,
    });
  }
}
