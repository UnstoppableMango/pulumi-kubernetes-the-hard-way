import { Config } from '@pulumi/pulumi';
import { RemoteFile } from '@unmango/pulumi-kubernetes-the-hard-way';

const config = new Config();
const host = config.require('host');
const port = config.requireNumber('port');
const user = config.require('user');
const password = config.require('password');

const file = new RemoteFile('remote-file', {
  connection: { host, port, user, password },
  content: config.require('content'),
  path: config.require('path'),
});

export const stderr = file.stderr;
export const stdout = file.stdout;
