import { Config } from '@pulumi/pulumi';
import { RemoteFile } from '@unmango/pulumi-kubernetes-the-hard-way';

const config = new Config();
const host = config.require('host');

const file = new RemoteFile('remote-file', {
  connection: { host },
  content: config.require('content'),
  path: config.require('path'),
});

export const stderr = file.stderr;
export const stdout = file.stdout;
