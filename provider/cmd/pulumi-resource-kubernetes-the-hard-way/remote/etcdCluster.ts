import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { EtcdInstall } from './etcdInstall';
import { EtcdConfiguration } from './etcdConfiguration';
import { EtcdService } from './etcdService';
import { StartEtcd } from './startEtcd';

type MapFunc<T> = {
  (node: string, props: schema.EtcdNodeOutputs): T;
};

export class EtcdCluster extends schema.EtcdCluster {
  constructor(name: string, args: schema.EtcdClusterArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const architecture = output(args.architecture);
    const binaryDirectory = output(args.architecture);
    const bundle = output(args.bundle);
    const configurationDirectory = output(args.configurationDirectory);
    const dataDirectory = output(args.dataDirectory);
    const nodes = nodesOutput(args.nodes);
    const version = output(args.version);

    const install = mapNodes((node, props) => {
      return new EtcdInstall(`${name}-${node}`, {
        connection: props.connection,
        architecture: props.architecture,
        directory: args.binaryDirectory,
        version: args.version,
      }, { parent: this });
    }, nodes);

    const configuration = mapNodes((node, props) => {
      return new EtcdConfiguration(`${name}-${node}`, {
        connection: props.connection,
        caPem: bundle.caPem,
        certPem: bundle.certPem,
        etcdPath: install[node].etcdPath,
        internalIp: props.internalIp,
        keyPem: bundle.keyPem,
        configurationDirectory: args.configurationDirectory,
        dataDirectory: args.dataDirectory,
      }, { parent: this });
    }, nodes);

    const service = mapNodes((node, props) => {
      // This feels like a whole lotta jank in here
      const peers = Object.keys(configuration)
        .filter(k => k !== node)
        .map(k => configuration[k].value);

      return new EtcdService(`${name}-${node}`, {
        connection: props.connection,
        configuration: configuration[node].value,
        peers,
      }, { parent: this });
    }, nodes);

    const start = mapNodes((node, props) => {
      return new StartEtcd(`${name}-${node}`, {
        connection: props.connection,
      }, { parent: this });
    }, nodes);

    this.architecture = architecture as Output<schema.ArchitectureOutputs> | undefined;
    this.binaryDirectory = binaryDirectory as Output<string> | undefined;
    this.bundle = bundle;
    this.configuration = configuration;
    this.configurationDirectory = configurationDirectory as Output<string> | undefined;
    this.dataDirectory = dataDirectory as Output<string> | undefined;
    this.install = install;
    this.service = service;
    this.start = start;
    this.nodes = nodes;
    this.version = version as Output<string> | undefined;

    this.registerOutputs({
      architecture,
      binaryDirectory,
      bundle,
      configuration,
      configurationDirectory,
      dataDirectory,
      install,
      service,
      start,
      nodes,
      version,
    });
  }
}

function nodesOutput(nodes: Record<string, schema.EtcdNodeInputs>): Record<string, schema.EtcdNodeOutputs> {
  const result: Record<string, schema.EtcdNodeOutputs> = {};

  for (const key in nodes) {
    const node = nodes[key];
    result[key] = {
      connection: output(node.connection),
      internalIp: output(node.internalIp),
      architecture: output(node.architecture) as Output<schema.ArchitectureOutputs> | undefined, // Ugh
    };
  }

  return result;
}

function mapNodes<T>(f: MapFunc<T>, nodes: Record<string, schema.EtcdNodeOutputs>): Record<string, T> {
  const entries = Object.entries(nodes).map(([node, props]) => {
    const result = f(node, props);
    return [node, result] as const;
  });

  return Object.fromEntries(entries);
}
