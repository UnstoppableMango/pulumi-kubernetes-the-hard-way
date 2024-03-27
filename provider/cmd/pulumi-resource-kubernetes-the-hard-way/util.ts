import { ComponentResource } from '@pulumi/pulumi';
import { AllowedUsage } from './types';
import { ConstructResult } from '@pulumi/pulumi/provider';

export function isAllowedUsage(value: string): value is AllowedUsage {
  return Object.values(AllowedUsage).some(x => x === value);
}

export function toAllowedUsage(value: string): AllowedUsage
export function toAllowedUsage(value: string[]): AllowedUsage[]
export function toAllowedUsage(value: string | string[]): AllowedUsage | AllowedUsage[] {
  if (Array.isArray(value)) {
    return value.filter(isAllowedUsage);
  } else if (isAllowedUsage(value)) {
    return value;
  } else {
    throw new Error(`unsupported allowedUsage: ${value}`);
  }
}

/**
 * https://github.com/pulumi/pulumi-awsx/blob/0390f0dd5b166861ef240e21fe676959fed041d5/awsx/utils.ts#L55
 * @internal
 * */
export function resourceToConstructResult<T extends ComponentResource>(
  resource: T,
): ConstructResult {
  const state = Object.fromEntries(
    Object.entries(resource).filter(
      ([key, value]) =>
        key !== "urn" &&
        !key.startsWith("get") &&
        !key.startsWith("_") &&
        typeof value !== "function" &&
        value !== undefined,
    ),
  );
  return {
    urn: resource.urn,
    state,
  };
}
