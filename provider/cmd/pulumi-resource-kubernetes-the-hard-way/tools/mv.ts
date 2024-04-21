import { ComponentResourceOptions, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import * as tool from './tool';

const apply = tool.factory<
  schema.MvOptsInputs,
  schema.MvOptsOutputs
>(
  'mv',
  (builder, opts) => builder
    .option('--context', opts.context)
    .option('--force', opts.force)
    .option('--no-clobber', opts.noClobber)
    .option('--no-target-directory', opts.noTargetDirectory)
    .option('--strip-trailing-slashed', opts.stripTrailingSlashes)
    .option('--suffix', opts.suffix)
    .option('--target-directory', opts.targetDirectory)
    .option('--update', opts.update)
    .option('--verbose', opts.verbose)
    .arg(opts.source)
    .arg(opts.dest)
    .arg(opts.directory),
  (i) => ({
    source: output(i.source),
    backup: i.backup,
    context: tool.mapO(i.context),
    control: tool.mapO(i.control),
    dest: tool.mapO(i.dest),
    directory: tool.mapO(i.directory),
    force: tool.mapO(i.force),
    noClobber: tool.mapO(i.noClobber),
    noTargetDirectory: tool.mapO(i.noTargetDirectory),
    stripTrailingSlashes: tool.mapO(i.stripTrailingSlashes),
    suffix: tool.mapO(i.suffix),
    targetDirectory: tool.mapO(i.targetDirectory),
    update: tool.mapO(i.update),
    verbose: tool.mapO(i.verbose),
  }),
);

export class Mv extends schema.Mv {
  constructor(name: string, args: schema.MvArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);
    if (opts?.urn) return;
    const outputs = apply(name, args, this);
    this.registerOutputs(outputs);
  }
}
