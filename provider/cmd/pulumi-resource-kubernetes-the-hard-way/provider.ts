import * as pulumi from '@pulumi/pulumi';
import * as provider from '@pulumi/pulumi/provider';
import * as cert from './certificate';
import * as pki from './clusterPki';
import * as keypair from './keypair';
import * as remoteFile from './remoteFile';
import * as rootCa from './rootCa';

export class Provider implements provider.Provider {
  constructor(readonly version: string, readonly schema: string) {
    pulumi.runtime.registerResourceModule('kubernetes-the-hard-way', 'index', {
      version: version,
      construct(name, type, urn) {
        pulumi.log.error(`REHYDRATING ${name} MAYBE`);
        pulumi.log.error(`REHYDRATE URN: ${urn}`)
        switch (type) {
          case 'kubernetes-the-hard-way:index:Certificate':
            return new cert.Certificate(name, {} as cert.CertificateArgs, { urn });
          case 'kubernetes-the-hard-way:index:ClusterPki':
            return new pki.ClusterPki(name, {} as pki.ClusterPkiArgs, { urn });
          case 'kubernetes-the-hard-way:index:RemoteFile':
            return new remoteFile.RemoteFile(name, {} as remoteFile.RemoteFileArgs, { urn });
          case 'kubernetes-the-hard-way:index:RootCa':
            return new rootCa.RootCa(name, {} as rootCa.RootCaArgs, { urn });
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
    switch (type) {
      case 'kubernetes-the-hard-way:index:Certificate':
        return await cert.construct(name, inputs, options);
      case 'kubernetes-the-hard-way:index:ClusterPki':
        return await pki.construct(name, inputs, options);
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
        return await keypair.callInstallCertInstance(inputs);
      case 'kubernetes-the-hard-way:index:Certificate/installKey':
        return await keypair.callInstallKeyInstance(inputs);
      case 'kubernetes-the-hard-way:index:installCert':
        return await keypair.callInstallCertStatic(inputs);
      case 'kubernetes-the-hard-way:index:newCertificate':
        return await rootCa.callNewCertificateStatic(inputs);
      case 'kubernetes-the-hard-way:index:installKey':
        return await keypair.callInstallKeyStatic(inputs);
      case 'kubernetes-the-hard-way:index:RootCa/newCertificate':
        return await rootCa.callNewCertificateInstance(inputs);
      case 'kubernetes-the-hard-way:index:RootCa/installCert':
        return await keypair.callInstallCertInstance(inputs);
      case 'kubernetes-the-hard-way:index:RootCa/installKey':
        return await keypair.callInstallKeyInstance(inputs);
      default:
        throw new Error(`unknown call token ${token}`);
    }
  }

  async invoke(token: string, inputs: pulumi.Inputs): Promise<provider.InvokeResult> {
    return this.call(token, inputs);
  }
}
