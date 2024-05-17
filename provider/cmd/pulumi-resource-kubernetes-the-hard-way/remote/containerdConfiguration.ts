import { ComponentResourceOptions, Output, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { File } from './file';

export class ContainerdConfiguration extends schema.ContainerdConfiguration {
  constructor(name: string, args: schema.ContainerdConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const connection = output(args.connection);
    const cri = criDefaults(args.cri);
    const path = output(args.path ?? '/etc/containerd');

    const file = new File(name, {
      connection,
      content: formatAsToml(cri),
      path,
    })

    this.connection = connection;;
    this.cri = cri;
    this.file = file;
  }
}

function criDefaults(cri?: schema.ContainerdCriPluginConfigurationInputs): schema.ContainerdCriPluginConfigurationOutputs {
  return {
    cni: cniDefaults(cri?.cni),
    containerd: containerdDefaults(cri?.containerd),
  };
}

function cniDefaults(cni?: schema.ContainerdCriPluginConfigurationCniInputs): schema.ContainerdCriPluginConfigurationCniOutputs {
  return {
    binDir: output(cni?.binDir ?? '/opt/cni/bin'),
    confDir: output(cni?.confDir ?? '/etc/cni/net.d'),
  };
}

function containerdDefaults(containerd?: schema.ContainerdCriPluginConfigurationContainerdInputs): schema.ContainerdCriPluginConfigurationContainerdOutputs {
  return {
    defaultRuntimeName: output(containerd?.defaultRuntimeName ?? 'runc'),
    snapshotter: output(containerd?.snapshotter ?? 'overlayfs'),
    runtimes: runcDefaults(containerd?.runtimes),
  };
}

function runcDefaults(runtimes?: schema.ContainerdCriPluginConfigurationContainerdRuncInputs): schema.ContainerdCriPluginConfigurationContainerdRuncOutputs {
  return {
    options: {
      systemdCgroup: output(runtimes?.options.systemdCgroup ?? true),
    },
    runtimeType: output(runtimes?.runtimeType ?? 'io.containerd.runc.v2'),
  };
}

// https://github.com/kelseyhightower/kubernetes-the-hard-way/blob/master/configs/containerd-config.toml
function formatAsToml(cri: schema.ContainerdCriPluginConfigurationOutputs): Output<string> {
  return interpolate`# DO NOT MODIFY - Managed by Pulumi
version = 2

[plugins."io.containerd.grpc.v1.cri"]
  [plugins."io.containerd.grpc.v1.cri".containerd]
    snapshotter = "${cri.containerd.snapshotter}"
    default_runtime_name = "${cri.containerd.defaultRuntimeName}"
  [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
    runtime_type = "${cri.containerd.runtimes?.runtimeType}"
  [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
    SystemdCgroup = ${cri.containerd.runtimes?.options.systemdCgroup}
[plugins."io.containerd.grpc.v1.cri".cni]
  bin_dir = "${cri.cni.binDir}"
  conf_dir = "${cri.cni.confDir}"
`;
}
