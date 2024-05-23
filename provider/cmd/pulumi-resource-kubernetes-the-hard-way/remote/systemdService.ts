import { ComponentResourceOptions, Inputs, Output, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { File } from './file';

export class SystemdService extends schema.SystemdService {
  constructor(name: string, args: schema.SystemdServiceArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;

    const connection = output(args.connection);
    const directory = output(args.directory ?? '/etc/systemd/system');
    const install = output(args.install);
    const service = output(args.service);
    const unit = output(args.unit);
    const unitName = output(args.unitName ?? name);

    const file = new File(name, {
      connection: args.connection,
      content: stringify({ unit, service, install }),
      path: interpolate`${directory}/${unitName}.service`,
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
  return output(input).apply(x => {
    const { unit, service, install } = x;
    return `# DO NOT MODIFY - Managed by Pulumi
[Unit]
${sprint(unit)}

[Service]
${sprint(service)}

[Install]
${sprint(install)}
`;
  });
}

function sprint(value: Record<string, unknown>): string {
  return Object.entries(value)
    .map(([k, v]) => `${capitalize(k)}=${v}`)
    .join('\n');
}

function capitalize(s: string): string {
    return s[0].toUpperCase() + s.slice(1);
}
