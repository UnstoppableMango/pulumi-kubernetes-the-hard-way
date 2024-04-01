import { Inputs, output } from '@pulumi/pulumi';
import { InvokeResult } from '@pulumi/pulumi/provider';
import { ClusterPki } from './tls';
import * as keypair from './tls/keypair';
import * as rootCa from './tls/rootCa';

export async function call(token: string, inputs: Inputs): Promise<InvokeResult> {
  switch (token) {
    case 'kubernetes-the-hard-way:tls:Certificate/installCert':
      return await keypair.callInstallCertInstance(inputs);
    case 'kubernetes-the-hard-way:tls:Certificate/installKey':
      return await keypair.callInstallKeyInstance(inputs);
    case 'kubernetes-the-hard-way:tls:ClusterPki/getKubeconfig':
      return await callGetKubeconfig(inputs);
    case 'kubernetes-the-hard-way:remote:installCert':
      return await keypair.callInstallCertStatic(inputs);
    case 'kubernetes-the-hard-way:tls:newCertificate':
      return await rootCa.callNewCertificateStatic(inputs);
    case 'kubernetes-the-hard-way:remote:installKey':
      return await keypair.callInstallKeyStatic(inputs);
    case 'kubernetes-the-hard-way:tls:RootCa/newCertificate':
      return await rootCa.callNewCertificateInstance(inputs);
    case 'kubernetes-the-hard-way:tls:RootCa/installCert':
      return await keypair.callInstallCertInstance(inputs);
    case 'kubernetes-the-hard-way:tls:RootCa/installKey':
      return await keypair.callInstallKeyInstance(inputs);
    default:
      throw new Error(`unknown call token ${token}`);
  }
}

async function callGetKubeconfig(inputs: Inputs): Promise<InvokeResult> {
  const self = output<ClusterPki>(inputs.__self__);
  const result = self.apply(x => x.getKubeconfig(inputs.options));
  return { outputs: { result } };
}
