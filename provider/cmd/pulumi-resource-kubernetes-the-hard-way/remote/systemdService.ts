import { ComponentResourceOptions, Inputs, Output, interpolate, output } from '@pulumi/pulumi';
import * as ini from 'ini';
import * as schema from '../schema-types';
import { File } from './file';

export class SystemdService extends schema.SystemdService {
  constructor(name: string, args: schema.SystemdServiceArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const connection = output(args.connection);
    const directory = output(args.directory ?? '/etc/systemd/system');
    const install = output(args.install);
    const service = output(args.service);
    const unit = output(args.unit);

    const file = new File(name, {
      connection: args.connection,
      content: stringify({
        Unit: unit,
        Service: service,
        Install: install,
      }),
      path: interpolate`${directory}/${name}.service`,
    }, { parent: this });

    this.connection = connection;
    this.directory = directory;
    this.file = file;
    this.install = install as Output<schema.SystemdInstallSectionOutputs>;
    this.service = service as Output<schema.SystemdServiceSectionOutputs>;
    this.unit = unit as Output<schema.SystemdUnitSectionOutputs>;

    this.registerOutputs({
      connection,
      directory,
      file,
      install,
      service,
      unit,
    });
  }
}

function stringify(input: Inputs): Output<string> {
  return output(input).apply(x => ini.stringify(x, { whitespace: false, newline: true }));
}
