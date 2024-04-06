import { ComponentResourceOptions, interpolate, output } from '@pulumi/pulumi';
import * as schema from '../schema-types';
import { archiveInstall } from './archiveInstall';

export class CniPluginsInstall extends schema.CniPluginsInstall {
  constructor(name: string, args: schema.CniPluginsInstallArgs, opts?: ComponentResourceOptions) {
    super(name, args, opts);

    const architecture = output(args.architecture ?? 'amd64');
    const connection = output(args.connection);
    const directory = output(args.directory ?? '/opt/cni/bin/');
    const version = output(args.version ?? '0.9.1'); // TODO: Stateful versioning?
    const archiveName = interpolate`cni-plugins-linux-${architecture}-v${version}.tgz`;
    const url = interpolate`https://github.com/containernetworking/plugins/releases/download/v${version}/${archiveName}`;

    // TODO: Permission checks?
    // TODO: Caching? Put archive/bins into ~/.kthw/cache so i.e. directory changes, tarball doesn't need to be re-downloaded.
    // TODO: General update logic

    const { download, mkdir, mktemp, mvs, paths, rm, tar } = archiveInstall(name, {
      archiveName,
      binaries: [
        'bandwidth',
        'bridge',
        'dhcp',
        'dummy',
        'firewall',
        'host-device',
        'host-local',
        'ipvlan',
        'loopback',
        'macvlan',
        'portmap',
        'ptp',
        'sbr',
        'static',
        'tap',
        'tuning',
        'vlan',
        'vrf',
      ] as const,
      connection,
      directory,
      stripComponents: 1,
      url,
    }, this);

    this.architecture = architecture;
    this.archiveName = archiveName;
    this.bandwidthPath = paths.bandwidth;
    this.bridgePath = paths.bridge;
    this.connection = connection;
    this.dhcpPath = paths.dhcp;
    this.download = download;
    this.directory = directory;
    this.dummyPath = paths.dummy;
    this.firewallPath = paths.firewall;
    this.hostDevicePath = paths['host-device'];
    this.hostLocalPath = paths['host-local'];
    this.ipvlanPath = paths.ipvlan;
    this.loopbackPath = paths.loopback;
    this.macvlanPath = paths.macvlan;
    this.mktemp = mktemp;
    this.portmapPath = paths.portmap;
    this.ptpPath = paths.ptp;
    this.rm = rm;
    this.sbrPath = paths.sbr;
    this.staticPath = paths.static;
    this.tapPath = paths.tap
    this.tar = tar;
    this.tuningPath = paths.tuning;
    this.url = url;
    this.version = version;
    this.vlanPath = paths.vlan;
    this.vrfPath = paths.vrf;

    this.registerOutputs({
      architecture,
      archiveName,
      bandwidthPath: paths.bandwidth,
      bridgePath: paths.bridge,
      connection,
      dhcpPath: paths.dhcp,
      download,
      directory,
      dummyPath: paths.dummy,
      firewallPath: paths.firewall,
      hostDevicePath: paths['host-device'],
      hostLocalPath: paths['host-local'],
      ipvlanPath: paths.ipvlan,
      loopbackPath: paths.loopback,
      macvlanPath: paths.macvlan,
      mkdir,
      name,
      portmapPath: paths.portmap,
      ptpPath: paths.ptp,
      rm,
      sbrPath: paths.sbr,
      staticPath: paths.static,
      tapPath: paths.tap,
      tar,
      tuningPath: paths.tuning,
      url,
      version,
      vlanPath: paths.vlan,
      vrfPath: paths.vrf,
    });
  }
}
