import { Config } from '@pulumi/pulumi';
import {
  CniBridgePluginConfiguration,
  CniLoopbackPluginConfiguration,
  CniPluginConfiguration,
} from '@unmango/pulumi-kubernetes-the-hard-way/remote';
import { Mkdir } from '@unmango/pulumi-kubernetes-the-hard-way/tools';

const config = new Config();
const host = config.require('host');
const port = config.requireNumber('port');
const user = config.require('user');
const password = config.require('password');

const mkdir = new Mkdir('simple', {
  connection: { host, port, user, password },
  create: {
    parents: true,
    directory: '/etc/cni/net.d',
  },
});

const bridgeConfig = new CniBridgePluginConfiguration('simple', {
  connection: { host, port, user, password },
  subnet: '10.0.69.0/24',
}, { dependsOn: mkdir });

const loopbackConfig = new CniLoopbackPluginConfiguration('simple', {
  connection: { host, port, user, password },
}, { dependsOn: mkdir });

const pluginConfig = new CniPluginConfiguration('all', {
  connection: { host, port, user, password },
  subnet: '10.0.69.0/24',
  directory: '/etc/cni2/net.d',
});
