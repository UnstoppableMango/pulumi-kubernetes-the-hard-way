import { Inputs } from '@pulumi/pulumi';
import { InvokeResult } from '@pulumi/pulumi/provider';
import { Certificate, NewCertificateInputs, NewCertificateOutputs, RootCa } from './tls';
import { InstallInputs, InstallOutputs } from './remote';

type Functions = {
  'kubernetes-the-hard-way:tls:Certificate/installCert': (inputs: InstallInputs) => Promise<InstallOutputs>;
  'kubernetes-the-hard-way:tls:Certificate/installKey': (inputs: InstallInputs) => Promise<InstallOutputs>;
  // 'kubernetes-the-hard-way:tls:installCert': (inputs: InstallInputs) => Promise<InstallOutputs>;
  // 'kubernetes-the-hard-way:tls:newCertificate': (inputs: NewCertificateInputs) => Promise<NewCertificateOutputs>;
  // 'kubernetes-the-hard-way:tls:installKey': (inputs: InstallInputs) => Promise<InstallOutputs>;
  'kubernetes-the-hard-way:tls:RootCa/newCertificate': (inputs: NewCertificateInputs) => Promise<NewCertificateOutputs>;
  'kubernetes-the-hard-way:tls:RootCa/installCert': (inputs: InstallInputs) => Promise<InstallOutputs>;
  'kubernetes-the-hard-way:tls:RootCa/installKey': (inputs: InstallInputs) => Promise<InstallOutputs>;
};

export const functions: Functions = {
  'kubernetes-the-hard-way:tls:Certificate/installCert': (i) => self<Certificate>(i).installCert(i),
  'kubernetes-the-hard-way:tls:Certificate/installKey': (i) => self<Certificate>(i).installKey(i),
  // 'kubernetes-the-hard-way:tls:installCert': (i) => keypair.installCert()
  // 'kubernetes-the-hard-way:tls:newCertificate':
  // 'kubernetes-the-hard-way:tls:installKey':
  'kubernetes-the-hard-way:tls:RootCa/newCertificate': (i) => self<RootCa>(i).newCertificate(i),
  'kubernetes-the-hard-way:tls:RootCa/installCert': (i) => self<RootCa>(i).installCert(i),
  'kubernetes-the-hard-way:tls:RootCa/installKey': (i) => self<RootCa>(i).installKey(i),
};

export async function call(token: string, inputs: Inputs): Promise<InvokeResult> {
  const untypedFunctions: Record<string, (inputs: any) => Promise<any>> = functions;
  const handler = untypedFunctions[token];
  if (!handler) {
    throw new Error(`unknown method ${token}`);
  }
  const outputs = await handler(inputs);
  return { outputs };
}

function self<T>(inputs: Inputs): T {
  return inputs.__self__;
}
