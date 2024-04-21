import { ComponentResource, Input, Output, Unwrap, output } from '@pulumi/pulumi';
import { types } from '@pulumi/command';
import { Command } from '@pulumi/command/remote';
import { CommandBuilder } from './commandBuilder';

export type ApplyInputs<T> = {
  (builder: CommandBuilder, inputs: T): CommandBuilder;
}

export type ToolFactory<T, U> = {
  (name: string, args: ToolInputs<T>, resource: ToolResource<U>): ToolOutputs<U>;
};

interface ToolInputs<T> {
  binaryPath?: string | Input<string>;
  connection: Input<types.input.remote.ConnectionArgs>;
  create?: T;
  delete?: T;
  environment?: Input<Record<string, Input<string>>>;
  stdin?: string | Input<string>;
  triggers?: any[] | Input<any[]>;
  update?: T;
}

interface ToolOutputs<T> {
  binaryPath: string | Output<string>;
  command: Command | Output<Command>;
  connection: types.output.remote.Connection | Output<types.output.remote.Connection>;
  create?: T | Output<T>;
  delete?: T | Output<T>;
  environment: Record<string, string> | Output<Record<string, string>>;
  stdin?: string | Output<string>;
  triggers: any[] | Output<any[]>;
  update?: T | Output<T>;
}

type ToolResource<T> = ToolOutputs<T> & ComponentResource & {
  command: Command | Output<Command>;
  stderr: string | Output<string>;
  stdout: string | Output<string>;
}

export function factory<T, U>(
  defaultPath: string,
  apply: ApplyInputs<T>,
  toOutput: (i: T) => U
): ToolFactory<T, U> {
  return (name, args, resource) => makeTool(
    name,
    args,
    defaultPath,
    apply,
    toOutput,
    resource
  );
}

function makeTool<T, U>(
  name: string,
  args: ToolInputs<T>,
  defaultPath: string,
  apply: ApplyInputs<T>,
  toOutput: (i: T) => U,
  resource: ToolResource<U>,
): ToolOutputs<U> {
  const binaryPath = output(args.binaryPath ?? defaultPath);
  const connection = output(args.connection);
  const environment = output(args.environment ?? {});
  const stdin = args.stdin ? output(args.stdin) : undefined;
  const triggers = output(args.triggers ?? []);

  const builder = new CommandBuilder(binaryPath)
  const format = (o?: T) => o ? apply(builder, o).command : undefined;

  const command = new Command(name, {
    connection,
    environment,
    stdin,
    triggers,
    create: format(args.create),
    update: format(args.update),
    delete: format(args.delete),
  }, { parent: resource });

  const map = (i?: T) => i ? toOutput(i) : undefined;

  resource.binaryPath = binaryPath;
  resource.command = command;
  resource.connection = connection;
  resource.environment = environment;
  resource.stderr = command.stderr;
  resource.stdin = stdin;
  resource.stdout = command.stdout;
  resource.triggers = triggers;
  resource.create = map(args.create);
  resource.delete = map(args.delete);
  resource.update = map(args.update);

  return {
    binaryPath,
    command,
    connection,
    environment,
    triggers,
    stdin,
    create: map(args.create),
    delete: map(args.delete),
    update: map(args.delete),
  };
}

export function mapO<T>(i: Input<T> | undefined): Output<Unwrap<T>> | undefined {
  return i ? output(i) : undefined;
}
