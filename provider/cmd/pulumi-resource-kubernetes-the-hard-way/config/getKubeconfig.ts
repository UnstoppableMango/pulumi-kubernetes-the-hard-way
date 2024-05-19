import { all, interpolate } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { Kubeconfig } from '../schema-types';

export async function getKubeconfig(inputs: schema.getKubeconfigInputs): Promise<schema.getKubeconfigOutputs> {
  const cert = this.getCert(inputs.options);
  const ip = getIp(options);
  const username = getUsername(options);

  const caData = this.ca.certPem;
  const clientCertPem = cert.certPem;
  const clientKeyPem = cert.privateKeyPem;
  const clusterName = this.clusterName;
  const server = interpolate`https://${ip}:6443`;

  const kubeconfig = all([clusterName, caData, server, username, clientCertPem, clientKeyPem])
    .apply<Kubeconfig>(([clusterName, caData, server, username, clientCertPem, clientKeyPem]) => ({
      clusters: [{
        name: clusterName,
        cluster: {
          certificateAuthorityData: caData,
          server,
        },
      }],
      contexts: [{
        name: 'default',
        context: {
          cluster: clusterName,
          user: username,
        },
      }],
      users: [{
        name: username,
        user: {
          clientCertificateData: clientCertPem,
          clientKeyData: clientKeyPem,
        },
      }],
    }));

  return { result: kubeconfig };
}
