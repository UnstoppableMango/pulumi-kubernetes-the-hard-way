import { RootCa } from '@unmango/pulumi-kubernetes-the-hard-way/tls';

const simple = new RootCa('simple', {
  validityPeriodHours: 256,
});
