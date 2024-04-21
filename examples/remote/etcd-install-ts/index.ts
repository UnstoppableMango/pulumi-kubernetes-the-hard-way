import { Config } from '@pulumi/pulumi';
import { EtcdInstall } from '@unmango/pulumi-kubernetes-the-hard-way/remote';
import { Etcdctl } from '@unmango/pulumi-kubernetes-the-hard-way/tools';

const config = new Config();
const host = config.require('host');
const port = config.requireNumber('port');
const user = config.require('user');
const password = config.require('password');

const simple = new EtcdInstall('simple', {
  connection: { host, port, user, password },
});

const etcdctl = new Etcdctl('simple', {
  connection: simple.connection,
  create: { commands: 'version' },
});

export const etcdctlVersion = etcdctl.command.stdout;
