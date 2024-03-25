import * as pulumi from '@pulumi/pulumi';
import * as provider from '@pulumi/pulumi/provider';
import * as cert from './certificate';
import * as keypair from './keypair';
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

  async call(token: string, inputs: pulumi.Inputs): Promise<provider.InvokeResult> {
    switch (token) {
      case 'kubernetes-the-hard-way:index:Certificate/installCert':
        return call<cert.Certificate>(
          inputs,
          (self, i) => self.installCert({
            connection: i.connection,
            name: i.name,
            path: i.path,
            opts: i.opts,
          }),
        );
      case 'kubernetes-the-hard-way:index:Certificate/installKey':
        return call<cert.Certificate>(
          inputs,
          (self, i) => self.installKey({
            connection: i.connection,
            name: i.name,
            path: i.path,
            opts: i.opts,
          }),
        );
      case 'kubernetes-the-hard-way:index:RootCa/newCertificate':
        return call<rootCa.RootCa>(
          inputs,
          (self, i) => self.newCertificate({
            algorithm: i.algorithm,
            allowedUses: i.allowedUses,
            name: i.name,
            validityPeriodHours: i.validityPeriodHours,
            dnsNames: i.dnsNames,
            ecdsaCurve: i.ecdsaCurve,
            ipAddresses: i.ipAddresses,
            isCaCertificate: i.isCaCertificate,
            opts: i.opts,
            rsaBits: i.rsaBits,
            subject: i.subject,
            uris: i.uris,
          }),
        );
      case 'kubernetes-the-hard-way:index:RootCa/installCert':
        return call<rootCa.RootCa>(
          inputs,
          (self, i) => self.installCert({
            connection: i.connection,
            name: i.name,
            path: i.path,
            opts: i.opts,
          }),
        );
      default:
        throw new Error(`unknown call token ${token}`);
    }
  }
}

async function call<T>(i: pulumi.Inputs, c: (r: T, i: pulumi.Inputs) => any): Promise<provider.InvokeResult> {
  return { outputs: { result: c(i.__self__, i) } };
}
