import * as pulumi from '@pulumi/pulumi';
import * as provider from '@pulumi/pulumi/provider';
import * as cert from './certificate';
import * as remoteFile from './remoteFile';
import * as rootCa from './rootCa';

export class Provider implements provider.Provider {
    constructor(readonly version: string, readonly schema: string) { }

    async construct(
        name: string,
        type: string,
        inputs: pulumi.Inputs,
        options: pulumi.ComponentResourceOptions
    ): Promise<provider.ConstructResult> {
        switch (type) {
            case 'kubernetes-the-hard-way:index:Certificate':
                return await cert.construct(name, inputs, options);
            case 'kubernetes-the-hard-way:index:RemoteFile':
                return await remoteFile.construct(name, inputs, options);
            case 'kubernetes-the-hard-way:index:RootCa':
                return await rootCa.construct(name, inputs, options);
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    }
}
