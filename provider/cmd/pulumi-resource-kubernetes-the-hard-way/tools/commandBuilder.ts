import { Input, Output, all, output } from '@pulumi/pulumi';

type MaybeNumber = number | undefined;
type MaybeString = string | undefined;
type OptValue = boolean | MaybeString | MaybeNumber;

interface Opt {
  option: string;
  value: OptValue;
}

export class CommandBuilder {
  private readonly cmd: Output<string[]>;
  private opts: Output<Opt>[] = [];
  private args: Output<string[]> = output([]);

  public get command(): Output<string> {
    const opts = all(this.opts)
      .apply((opts) => opts.filter(x => x.value))
      .apply((opts) => opts.map(toString));

    return all([this.cmd, opts, this.args])
      .apply(([cmd, opts, args]) => [...cmd, ...opts, ...args].join(' '));
  }

  constructor(...cmd: Input<string>[]) {
    this.cmd = all(cmd);
  }

  public arg(arg?: Input<string> | Input<Input<string>[]>): CommandBuilder {
    if (arg) {
      this.args = output(arg)
        .apply(toArray)
        .apply(prepend(this.args));
    }

    return this;
  }

  public option(option: string, value: Input<OptValue>): CommandBuilder {
    const pair: Output<Opt> = output(value).apply(toOpt(option));
    this.opts.push(pair);
    return this;
  }

  public options(option: string, value: Input<OptValue>[]): CommandBuilder {
    throw new Error('TODO');
  }
}

// Ideally would like for this to not be exported
export function toArray(x: string | Input<string>[]): Output<string[]> {
  if (typeof x === 'string') {
    return output([x]);
  }

  return x?.length > 0 ? all(x) : output([]);
}

function prepend(a: Input<string[]>): (b: string[]) => Output<string[]> {
  return (b) => output(a).apply(a => [...a, ...b]);
}

function toOpt(option: string): (value: OptValue) => Opt {
  return (value) => ({ option, value });
}

function toString(opt: Opt): string {
  if (typeof opt.value === 'boolean') {
    return opt.value ? opt.option : '';
  }

  return opt.value ? `${opt.option} ${opt.value}` : '';
}
