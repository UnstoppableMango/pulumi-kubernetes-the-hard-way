import * as pulumi from '@pulumi/pulumi';
import * as provider from '@pulumi/pulumi/provider';
import * as cert from './tls/certificate';
import * as pki from './tls/clusterPki';
import * as remoteFile from './remote/file';
import * as rootCa from './tls/rootCa';
import { construct } from './resources';
import { resourceToConstructResult } from './util';
import { functions } from './functions';

export class Provider implements provider.Provider {
  constructor(readonly version: string, readonly schema: string) {
    pulumi.runtime.registerResourceModule('kubernetes-the-hard-way', 'index', {
      version: version,
      construct(name, type, urn) {
        switch (type) {
          case 'kubernetes-the-hard-way:tls:Certificate':
            return new cert.Certificate(name, <any>undefined, { urn });
          case 'kubernetes-the-hard-way:tls:ClusterPki':
            return new pki.ClusterPki(name, <any>undefined, { urn });
          case 'kubernetes-the-hard-way:remote:File':
            return new remoteFile.File(name, <any>undefined, { urn });
          case 'kubernetes-the-hard-way:tls:RootCa':
            return new rootCa.RootCa(name, <any>undefined, { urn });
          default:
            throw new Error(`unknown resource type ${type}`);
        }
      },
    })
  }

  async construct(
    name: string,
    type: string,
    inputs: pulumi.Inputs,
    options: pulumi.ComponentResourceOptions
  ): Promise<provider.ConstructResult> {
    const resource = construct(name, type, inputs, options);
    if (resource === undefined) {
      throw new Error(`unknown resource type ${type}`);
    }

    return resourceToConstructResult(resource);
  }

  async call(token: string, inputs: pulumi.Inputs): Promise<provider.InvokeResult> {
    const untypedFunctions: Record<string, (inputs: any) => Promise<any>> = functions;
    const handler = untypedFunctions[token];
    if (!handler) {
      throw new Error(`unknown method ${token}`);
    }
    const outputs = await handler(inputs);
    return { outputs };
  }

  async invoke(token: string, inputs: pulumi.Inputs): Promise<provider.InvokeResult> {
    return this.call(token, inputs); // Why do both of these functions exist
  }
}
