import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { SystemdService } from './systemdService';

export class EtcdService extends schema.EtcdService {
  constructor(name: string, args: schema.EtcdServiceArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const connection = output(args.connection);
    const description = output(args.description ?? 'etcd');
    const documentation = output(args.documentation ?? 'https://github.com/etcd-io/etcd');
    const restart = output(args.restart ?? 'on-failure');
    const restartSec = output(args.restartSec ?? '5');
    const wantedBy = output(args.wantedBy ?? 'multi-user.target');

    const service = new SystemdService(name, {
      connection,
      directory: args.directory,
      unit: {
        description,
        documentation: [documentation],
      },
      service: {
        type: 'notify',
        restart,
        restartSec,
      },
      install: {
        wantedBy: [wantedBy],
      },
    }, { parent: this });

    this.connection = connection;
    this.description = description;
    this.directory = service.directory;
    this.documentation = documentation;
    this.restart = restart;
    this.restartSec = restartSec;
    this.service = service;
    this.wantedBy = wantedBy;

    this.registerOutputs({
      connection,
      description,
      directory: service.directory,
      documentation,
      restart,
      restartSec,
      service,
      wantedBy,
    });
  }
}
