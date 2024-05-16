import { Config } from '@pulumi/pulumi';
import { CniBridgePluginConfiguration, CniLoopbackPluginConfiguration } from '@unmango/pulumi-kubernetes-the-hard-way/remote';

const config = new Config();
const host = config.require('host');
const port = config.requireNumber('port');
const user = config.require('user');
const password = config.require('password');

const bridgeConfig = new CniBridgePluginConfiguration('simple', {
  connection: { host, port, user, password },
});

const loopbackConfig = new CniLoopbackPluginConfiguration('simple', {
  connection: { host, port, user, password },
});