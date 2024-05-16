import { ComponentResourceOptions, Output, all, jsonStringify, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { File } from './file';

export class CniBridgePluginConfiguration extends schema.CniBridgePluginConfiguration {
  constructor(name: string, args: schema.CniBridgePluginConfigurationArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const bridge = output(args.bridge ?? 'cni0');
    const bridgeName = output(args.name ?? 'bridge');
    const cniVersion = output(args.cniVersion ?? '1.0.0');
    const connection = output(args.connection);
    const ipMasq = output(args.ipMasq ?? true);
    const ipam = all([args.ipam, args.subnet]).apply(x => ipamDefaults(...x));
    const isGateway = output(args.isGateway ?? true);
    const path = output(args.path ?? '/var/lib/kubernetes');
    const subnet = output(args.subnet);
    const type = output(args.type ?? 'bridge');

    const file = new File(name, {
      connection,
      content: jsonStringify({
        bridge,
        name: bridgeName,
        cniVersion,
        ipMasq,
        ipam,
        isGateway,
        subnet,
        type,
      }),
      path,
    }, { parent: this });

    this.bridge = bridge;
    this.cniVersion = cniVersion;
    this.connection = connection;
    this.file = file;
    this.ipMasq = ipMasq;
    this.ipam = ipam;
    this.isGateway = isGateway;
    this.name = bridgeName;
    this.path = path;
    this.subnet = subnet as Output<string> | undefined;
    this.type = type;

    this.registerOutputs({
      bridge,
      cniVersion,
      connection,
      file,
      ipMasq,
      ipam,
      isGateway,
      name,
      path,
      subnet,
      type,
    });
  }
}

function ipamDefaults(ipam?: schema.CniBridgeIpamInputs, subnet?: string): schema.CniBridgeIpamOutputs {
  const require = (x?: string): string => {
    if (!x) throw new Error('');
    return x;
  }

  return {
    type: output(ipam?.type ?? 'host-local'),
    ranges: output(ipam?.ranges).apply(x => x ?? [{
      subnet: require(subnet),
    }]),
    routes: output(ipam?.routes).apply(x => x ?? [{
      dst: '0.0.0.0/0',
    }]),
  };
}
