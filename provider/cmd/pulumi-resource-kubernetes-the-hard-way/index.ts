import { readFileSync } from 'fs';
import * as pulumi from '@pulumi/pulumi';
import { Provider } from './provider';

function main(args: string[]): Promise<void> {
    const schema: string = readFileSync(require.resolve('./schema.json'), { encoding: 'utf-8' });
    let version: string = require('./package.json').version;

    if (version === "${VERSION}" && require("inspector")?.url() !== undefined) {
      // We're running in debug mode so just use a valid placeholder version.
      version = "0.0.0";
    }

    // Node allows for the version to be prefixed by a "v",
    // while semver doesn't. If there is a v, strip it off.
    if (version.startsWith('v')) {
        version = version.slice(1);
    }

    return pulumi.provider.main(new Provider(version, schema), args);
}

main(process.argv.slice(2));
