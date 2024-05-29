import { Output, all, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Kubeconfig } from './kubeconfig';

export async function getKubeconfig(inputs: schema.getKubeconfigInputs): Promise<schema.getKubeconfigOutputs> {
  const caData = output(inputs.caPem);
  const clientCert = output(inputs.clientCert);
  const clientKey = output(inputs.clientKey);
  const clusterName = output(inputs.clusterName);
  const contextName = output(inputs.contextName ?? 'default');
  const server = output(inputs.server);
  const username = output(inputs.username);

  const kubeconfig = all([clusterName, contextName, caData, server, username, clientCert, clientKey])
    .apply<Kubeconfig>(([clusterName, contextName, caData, server, username, clientCert, clientKey]) => ({
      clusters: [{
        name: clusterName,
        cluster: {
          certificateAuthorityData: caData,
          server,
        },
      }],
      contexts: [{
        name: contextName,
        context: {
          cluster: clusterName,
          user: username,
        },
      }],
      users: [{
        name: username,
        user: {
          clientCertificateData: clientCert,
          clientKeyData: clientKey,
        },
      }],
    }));

  return { result: kubeconfig as unknown as Output<schema.KubeconfigOutputs> };
}
