import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.TarOptsInputs,
  schema.TarOptsOutputs
>(
  'tar',
  (builder, opts) => builder
    .option('--extract', opts.extract)
    .option('--gzip', opts.gzip)
    .option('--file', opts.archive)
    .option('--directory', opts.directory)
    .option('--strip-components', output(opts.stripComponents).apply(x => {
      return x ? x.toString() : undefined; // TODO: Move login into command builder
    }))
    .arg(opts.files),
  (i) => ({
    archive: output(i.archive),
    directory: tool.mapO(i.directory),
    extract: tool.mapO(i.extract),
    files: tool.mapO(i.files),
    gzip: tool.mapO(i.gzip),
    recursive: tool.mapO(i.recursive),
    stripComponents: tool.mapO(i.stripComponents),
  }),
);

export class Tar extends schema.Tar {
  constructor(name: string, args: schema.TarArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
