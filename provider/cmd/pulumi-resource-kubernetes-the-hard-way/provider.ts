import * as pulumi from '@pulumi/pulumi';
import * as provider from '@pulumi/pulumi/provider';
import { RootCa, RootCaArgs } from './rootCa';

export class Provider implements provider.Provider {
    constructor(readonly version: string, readonly schema: string) { }

    async construct(name: string, type: string, inputs: pulumi.Inputs,
        options: pulumi.ComponentResourceOptions): Promise<provider.ConstructResult> {
        switch (type) {
            case 'kubernetes-the-hard-way:index:RootCa':
                return await constructRootCa(name, inputs, options);
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    }
}

async function constructRootCa(name: string, inputs: pulumi.Inputs,
    options: pulumi.ComponentResourceOptions): Promise<provider.ConstructResult> {
    const rootCa = new RootCa(name, inputs as RootCaArgs, options);
    return {
        urn: rootCa.urn,
        state: {
            allowedUses: rootCa.allowedUses,
            cert: rootCa.cert,
            certPem: rootCa.certPem,
            key: rootCa.key,
            keyPem: rootCa.keyPem,
        },
    };
}
