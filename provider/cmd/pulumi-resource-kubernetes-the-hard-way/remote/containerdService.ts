import { ComponentResourceOptions, Output, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { SystemdService } from './systemdService';

export class ContainerdService extends schema.ContainerdService {
  constructor(name: string, args: schema.ContainerdServiceArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const configuration = output(args.configuration);
    const connection = output(args.connection);
    const description = output(args.description ?? 'containerd container runtime');
    const documentation = output(args.documentation ?? 'https://containerd.io');
    const path = output(args.containerdPath ?? '/bin/containerd');
    const restart = output(args.restart ?? 'always');
    const restartSec = output(args.restartSec ?? '5');
    const wantedBy = output(args.wantedBy ?? 'multi-user.target');

    const service = new SystemdService(name, {
      connection,
      unit: {
        description,
        documentation: [documentation],
        after: ['network.target'],
      },
      service: {
        delegate: 'yes',
        execStartPre: '/sbin/modprobe overlay',
        execStart: path,
        restart,
        restartSec,
        killMode: 'process',
        oomScoreAdjust: -999,
        limitNoFile: 1048576,
        limitCore: 'infinity',
        limitNProc: 'infinity',
      },
      install: {
        wantedBy: [wantedBy],
      },
    }, { parent: this });

    this.configuration = configuration as Output<schema.ContainerdConfigurationOutputs>;
    this.connection = connection;
    this.description = description;
    this.documentation = documentation;
    this.containerdPath = path;
    this.restart = restart;
    this.restartSec = restartSec;
    this.service = service;
    this.wantedBy = wantedBy;

    this.registerOutputs({
      configuration,
      connection,
      description,
      documentation,
      containerdPath: path,
      restart,
      restartSec,
      service,
      wantedBy,
    });
  }
}
