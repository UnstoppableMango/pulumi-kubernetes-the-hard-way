import * as YAML from 'yaml';
import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Mkdir } from '../tools';
import { File } from './file';

export class StaticPod extends schema.StaticPod {
  constructor(name: string, args: schema.StaticPodArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const connection = output(args.connection);
    const directory = '/etc/kubernetes/manifests';
    const fileName = output(args.fileName ?? 'kube-vip-pod.yaml');
    const path = interpolate`${directory}/${fileName}`;
    const pod = output(args.pod);

    const mkdir = new Mkdir(name, {
      connection,
      create: {
        directory,
        parents: true,
      },
    }, { parent: this });

    const file = new File(name, {
      connection,
      content: pod.apply(YAML.stringify),
      path,
    }, { parent: this, dependsOn: mkdir });

    this.connection = connection;
    this.file = file;
    this.fileName = fileName;
    this.mkdir = mkdir;
    this.path = path;

    this.registerOutputs({
      connection,
      file,
      fileName,
      mkdir,
      path,
      pod,
    });
  }
}
