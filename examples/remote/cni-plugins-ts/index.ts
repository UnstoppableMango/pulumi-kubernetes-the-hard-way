import { Config } from '@pulumi/pulumi';
import { CniBridgePluginConfiguration, CniLoopbackPluginConfiguration } from '@unmango/pulumi-kubernetes-the-hard-way/remote';
import { Mkdir } from '@unmango/pulumi-kubernetes-the-hard-way/tools';

const config = new Config();
const host = config.require('host');
const port = config.requireNumber('port');
const user = config.require('user');
const password = config.require('password');

const mkdir = new Mkdir('var-lib', {
  connection: { host, port, user, password },
  create: {
    parents: true,
    directory: '/var/lib/kubernetes',
  },
});

const bridgeConfig = new CniBridgePluginConfiguration('simple', {
  connection: { host, port, user, password },
  subnet: '10.0.69.0/24',
});

const loopbackConfig = new CniLoopbackPluginConfiguration('simple', {
  connection: { host, port, user, password },
});
