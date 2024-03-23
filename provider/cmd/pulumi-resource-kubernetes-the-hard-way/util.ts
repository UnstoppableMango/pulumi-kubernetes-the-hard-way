import { AllowedUsage } from './types';

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
