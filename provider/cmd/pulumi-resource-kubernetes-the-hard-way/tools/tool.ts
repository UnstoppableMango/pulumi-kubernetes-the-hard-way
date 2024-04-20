import { Input, Output, output } from '@pulumi/pulumi';
import { types } from '@pulumi/command';

interface ToolInputs {
  binaryPath?: string | Input<string>;
  connection: Input<types.input.remote.ConnectionArgs>;
  environment?: Input<Record<string, Input<string>>>;
  stdin?: string | Input<string>;
  triggers?: any[] | Input<any[]>;
}

interface ToolOutputs {
  binaryPath: string | Output<string>;
  connection: types.output.remote.Connection | Output<types.output.remote.Connection>;
  environment: Record<string, string> | Output<Record<string, string>>;
  stdin?: string | Output<string>;
  triggers: any[] | Output<any[]>;
}

export function outputs(args: ToolInputs, defaultPath: string): ToolOutputs {
  return {
    binaryPath: output(args.binaryPath ?? defaultPath),
    connection: output(args.connection),
    environment: output(args.environment ?? {}),
    stdin: stdin(args.stdin),
    triggers: output(args.triggers ?? []),
  };
}

function stdin(stdin?: Input<string>): Output<string> | undefined {
  return !stdin ? undefined : output(stdin);
}
