import * as pulumi from '@pulumi/pulumi';
import { ComponentResourceOptions, Inputs } from '@pulumi/pulumi';
import * as provider from '@pulumi/pulumi/provider';
import { ConstructResult, InvokeResult } from '@pulumi/pulumi/provider';
import { EtcdInstall, File } from './remote';
import { Certificate, ClusterPki, RootCa } from './tls';
import { construct } from './resources';
import { functions } from './functions';
import { resourceToConstructResult } from './util';

export class Provider implements provider.Provider {
  constructor(readonly version: string, readonly schema: string) {
    pulumi.runtime.registerResourceModule('kubernetes-the-hard-way', 'tls', {
      version: version,
      construct(name, type, urn) {
        switch (type) {
          case 'kubernetes-the-hard-way:tls:Certificate':
            return new Certificate(name, <any>undefined, { urn });
          case 'kubernetes-the-hard-way:tls:ClusterPki':
            return new ClusterPki(name, <any>undefined, { urn });
          case 'kubernetes-the-hard-way:tls:RootCa':
            return new RootCa(name, <any>undefined, { urn });
          default:
            throw new Error(`registerResourceModule: unknown resource type ${type}`);
        }
      },
    });

    pulumi.runtime.registerResourceModule('kubernetes-the-hard-way', 'remote', {
      version: version,
      construct(name, type, urn) {
        switch (type) {
          case 'kubernetes-the-hard-way:remote:EtcdInstall':
            return new EtcdInstall(name, <any>undefined, { urn });
          case 'kubernetes-the-hard-way:remote:File':
            return new File(name, <any>undefined, { urn });
          default:
            throw new Error(`registerResourceModule: unknown resource type ${type}`);
        }
      }
    });
  }

  async construct(name: string, type: string, inputs: Inputs, options: ComponentResourceOptions): Promise<ConstructResult> {
    const resource = construct(name, type, inputs, options);
    if (resource === undefined) {
      throw new Error(`construct: unknown resource type ${type}`);
    }

    return resourceToConstructResult(resource);
  }

  async call(token: string, inputs: pulumi.Inputs): Promise<InvokeResult> {
    const untypedFunctions: Record<string, (inputs: any) => Promise<any>> = functions;
    const handler = untypedFunctions[token];
    if (!handler) {
      throw new Error(`unknown method ${token}`);
    }
    const outputs = await handler(inputs);
    return { outputs };
  }

  async invoke(token: string, inputs: pulumi.Inputs): Promise<InvokeResult> {
    return this.call(token, inputs); // Why do both of these functions exist
  }
}
