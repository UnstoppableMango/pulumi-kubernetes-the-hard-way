import { Config } from '@pulumi/pulumi';
import { RemoteFile } from '@unmango/pulumi-kubernetes-the-hard-way';
import { Wget } from '@unmango/pulumi-kubernetes-the-hard-way/tools';
import path = require('path');

const config = new Config();
const host = config.require('host');
const port = config.requireNumber('port');
const user = config.require('user');
const password = config.require('password');
const basePath = config.require('basePath');

const file = new RemoteFile('remote', {
  connection: { host, port, user, password },
  content: config.require('content'),
  path: path.join(basePath, 'remote-file'),
});

const wget = new Wget('remote', {
  connection: { host, port, user, password },
  url: 'https://www.example.com',
  directoryPrefix: basePath,
  // The container image seems to have an old version
  httpsOnly: false,
  timestamping: false,
});

export const fileStderr = file.stderr;
export const fileStdout = file.stdout;

export const wgetStderr = wget.stderr;
export const wgetStdout = wget.stdout;
export const wgetCommand = wget.command;
