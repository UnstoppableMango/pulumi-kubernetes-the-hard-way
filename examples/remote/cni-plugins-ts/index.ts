import { Config } from '@pulumi/pulumi';
import {
  CniBridgePluginConfiguration,
  CniLoopbackPluginConfiguration,
} from '@unmango/pulumi-kubernetes-the-hard-way/config';
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
  subnet: '10.0.69.0/24',
}, { dependsOn: mkdir });

const loopbackConfig = new CniLoopbackPluginConfiguration('simple', {
}, { dependsOn: mkdir });
