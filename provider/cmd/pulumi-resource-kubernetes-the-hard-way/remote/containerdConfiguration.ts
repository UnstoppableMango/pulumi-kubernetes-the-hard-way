import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { File } from './file';

export class ContainerdConfiguration extends schema.ContainerdConfiguration {
  constructor(name: string, args: schema.ContainerdConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const connection = output(args.connection);
    const cri = criDefaults(args.cri);

    const file = new File(name, {
      connection,
      content: '', // TODO
      path: '', // TODO
    })

    this.connection = connection;;
    this.cri = cri;
    this.file = file;
  }
}

function criDefaults(cri?: schema.ContainerdCriPluginConfigurationInputs): schema.ContainerdCriPluginConfigurationOutputs {
  return {
    cni: {
      binDir: output(cri?.cni.binDir ?? ''),
      confDir: output(cri?.cni.confDir ?? ''),
    },
    containerd: {
      defaultRuntimeName: output(cri?.containerd.defaultRuntimeName ?? ''),
      snapshotter: output(cri?.containerd.snapshotter ?? ''),
    },
  };
}
