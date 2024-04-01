/* tslint:disable */
/**
 * This file was automatically generated by pulumi-provider-scripts.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source Pulumi Schema file,
 * and run "pulumi-provider-scripts gen-provider-types" to regenerate this file. */
import * as pulumi from "@pulumi/pulumi";
export type ConstructComponent<T extends pulumi.ComponentResource = pulumi.ComponentResource> = (name: string, inputs: any, options: pulumi.ComponentResourceOptions) => T;
export type ResourceConstructor = {
    readonly "kubernetes-the-hard-way:tls:Certificate": ConstructComponent<Certificate>;
    readonly "kubernetes-the-hard-way:tls:ClusterPki": ConstructComponent<ClusterPki>;
    readonly "kubernetes-the-hard-way:tls:EncryptionKey": ConstructComponent<EncryptionKey>;
    readonly "kubernetes-the-hard-way:remote:EtcdInstall": ConstructComponent<EtcdInstall>;
    readonly "kubernetes-the-hard-way:remote:Download": ConstructComponent<Download>;
    readonly "kubernetes-the-hard-way:remote:File": ConstructComponent<File>;
    readonly "kubernetes-the-hard-way:tls:RootCa": ConstructComponent<RootCa>;
    readonly "kubernetes-the-hard-way:tools:Mkdir": ConstructComponent<Mkdir>;
    readonly "kubernetes-the-hard-way:tools:Mktemp": ConstructComponent<Mktemp>;
    readonly "kubernetes-the-hard-way:tools:Mv": ConstructComponent<Mv>;
    readonly "kubernetes-the-hard-way:tools:Rm": ConstructComponent<Rm>;
    readonly "kubernetes-the-hard-way:tools:Tar": ConstructComponent<Tar>;
    readonly "kubernetes-the-hard-way:tools:Wget": ConstructComponent<Wget>;
};
export type Functions = {
    "kubernetes-the-hard-way:tls:ClusterPki/getKubeconfig": (inputs: ClusterPki_getKubeconfigInputs) => Promise<ClusterPki_getKubeconfigOutputs>;
};
import * as command from "@pulumi/command";
import * as random from "@pulumi/random";
import * as tls from "@pulumi/tls";
export abstract class Certificate<TData = any> extends pulumi.ComponentResource<TData> {
    public cert!: tls.index.LocallySignedCert | pulumi.Output<tls.index.LocallySignedCert>;
    public certPem!: string | pulumi.Output<string>;
    public csr!: tls.index.CertRequest | pulumi.Output<tls.index.CertRequest>;
    public key!: tls.index.PrivateKey | pulumi.Output<tls.index.PrivateKey>;
    public privateKeyPem!: string | pulumi.Output<string>;
    public publicKeyPem!: string | pulumi.Output<string>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:tls:Certificate", name, opts.urn ? { cert: undefined, certPem: undefined, csr: undefined, key: undefined, privateKeyPem: undefined, publicKeyPem: undefined } : { name, args, opts }, opts);
    }
}
export interface CertificateArgs {
    readonly algorithm: pulumi.Input<AlgorithmInputs>;
    readonly allowedUses: pulumi.Input<pulumi.Input<AllowedUsageInputs>[]>;
    readonly caCertPem: pulumi.Input<string>;
    readonly caPrivateKeyPem: pulumi.Input<string>;
    readonly ecdsaCurve?: pulumi.Input<EcdsaCurveInputs>;
    readonly rsaBits?: pulumi.Input<number>;
    readonly dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    readonly earlyRenewalHours?: pulumi.Input<number>;
    readonly ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly isCaCertificate?: pulumi.Input<boolean>;
    readonly setAuthorityKeyId?: pulumi.Input<boolean>;
    readonly setSubjectKeyId?: pulumi.Input<boolean>;
    readonly subject?: pulumi.Input<tls.types.input.index.CertRequestSubject>;
    readonly uris?: pulumi.Input<pulumi.Input<string>[]>;
    readonly validityPeriodHours: pulumi.Input<number>;
}
export abstract class ClusterPki<TData = any> extends pulumi.ComponentResource<TData> {
    public admin!: Certificate | pulumi.Output<Certificate>;
    public algorithm!: AlgorithmOutputs | pulumi.Output<AlgorithmOutputs>;
    public ca!: RootCa | pulumi.Output<RootCa>;
    public clusterName!: string | pulumi.Output<string>;
    public controllerManager!: Certificate | pulumi.Output<Certificate>;
    public kubelet!: Record<string, Certificate> | pulumi.Output<Record<string, Certificate>>;
    public kubeProxy!: Certificate | pulumi.Output<Certificate>;
    public kubernetes!: Certificate | pulumi.Output<Certificate>;
    public kubeScheduler!: Certificate | pulumi.Output<Certificate>;
    public publicIp!: string | pulumi.Output<string>;
    public rsaBits!: number | pulumi.Output<number>;
    public serviceAccounts!: Certificate | pulumi.Output<Certificate>;
    public validityPeriodHours!: number | pulumi.Output<number>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:tls:ClusterPki", name, opts.urn ? { admin: undefined, algorithm: undefined, ca: undefined, clusterName: undefined, controllerManager: undefined, kubelet: undefined, kubeProxy: undefined, kubernetes: undefined, kubeScheduler: undefined, publicIp: undefined, rsaBits: undefined, serviceAccounts: undefined, validityPeriodHours: undefined } : { name, args, opts }, opts);
    }
}
export interface ClusterPkiArgs {
    readonly algorithm?: pulumi.Input<AlgorithmInputs>;
    readonly clusterName: pulumi.Input<string>;
    readonly ecdsaCurve?: pulumi.Input<EcdsaCurveInputs>;
    readonly nodes: pulumi.Input<Record<string, pulumi.Input<ClusterPkiNodeInputs>>>;
    readonly publicIp: pulumi.Input<string>;
    readonly rsaBits?: pulumi.Input<number>;
    readonly validityPeriodHours?: pulumi.Input<number>;
}
export abstract class EncryptionKey<TData = any> extends pulumi.ComponentResource<TData> {
    public config!: string | pulumi.Output<string>;
    public key!: random.index.RandomBytes | pulumi.Output<random.index.RandomBytes>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:tls:EncryptionKey", name, opts.urn ? { config: undefined, key: undefined } : { name, args, opts }, opts);
    }
}
export interface EncryptionKeyArgs {
    readonly bytes?: pulumi.Input<number>;
}
export abstract class EtcdInstall<TData = any> extends pulumi.ComponentResource<TData> {
    public architecture!: ArchitectureOutputs | pulumi.Output<ArchitectureOutputs>;
    public archiveName!: string | pulumi.Output<string>;
    public download!: Download | pulumi.Output<Download>;
    public downloadDirectory!: string | pulumi.Output<string>;
    public downloadMkdir!: Mkdir | pulumi.Output<Mkdir>;
    public etcdPath!: string | pulumi.Output<string>;
    public etcdctlPath!: string | pulumi.Output<string>;
    public installDirectory!: string | pulumi.Output<string>;
    public installMkdir!: Mkdir | pulumi.Output<Mkdir>;
    public mvEtcd!: Mv | pulumi.Output<Mv>;
    public mvEtcdctl!: Mv | pulumi.Output<Mv>;
    public name!: string | pulumi.Output<string>;
    public tar!: Tar | pulumi.Output<Tar>;
    public url!: string | pulumi.Output<string>;
    public version!: string | pulumi.Output<string>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:remote:EtcdInstall", name, opts.urn ? { architecture: undefined, archiveName: undefined, download: undefined, downloadDirectory: undefined, downloadMkdir: undefined, etcdPath: undefined, etcdctlPath: undefined, installDirectory: undefined, installMkdir: undefined, mvEtcd: undefined, mvEtcdctl: undefined, name: undefined, tar: undefined, url: undefined, version: undefined } : { name, args, opts }, opts);
    }
}
export interface EtcdInstallArgs {
    readonly architecture?: pulumi.Input<ArchitectureInputs>;
    readonly connection: pulumi.Input<command.types.input.Connection>;
    readonly downloadDirectory?: pulumi.Input<string>;
    readonly installDirectory?: pulumi.Input<string>;
    readonly version?: pulumi.Input<string>;
}
export abstract class Download<TData = any> extends pulumi.ComponentResource<TData> {
    public connection!: command.types.output.Connection | pulumi.Output<command.types.output.Connection>;
    public destination!: string | pulumi.Output<string>;
    public mkdir!: Mkdir | pulumi.Output<Mkdir>;
    public url!: string | pulumi.Output<string>;
    public wget!: Wget | pulumi.Output<Wget>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:remote:Download", name, opts.urn ? { connection: undefined, destination: undefined, mkdir: undefined, url: undefined, wget: undefined } : { name, args, opts }, opts);
    }
}
export interface DownloadArgs {
    readonly connection: pulumi.Input<command.types.input.Connection>;
    readonly destination: pulumi.Input<string>;
    readonly removeOnDelete?: pulumi.Input<boolean>;
    readonly url: pulumi.Input<string>;
}
export abstract class File<TData = any> extends pulumi.ComponentResource<TData> {
    public command!: command.Command | pulumi.Output<command.Command>;
    public content!: string | pulumi.Output<string>;
    public path!: string | pulumi.Output<string>;
    public stderr!: string | pulumi.Output<string>;
    public stdin?: string | pulumi.Output<string>;
    public stdout!: string | pulumi.Output<string>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:remote:File", name, opts.urn ? { command: undefined, content: undefined, path: undefined, stderr: undefined, stdin: undefined, stdout: undefined } : { name, args, opts }, opts);
    }
}
export interface FileArgs {
    readonly connection: pulumi.Input<command.types.input.Connection>;
    readonly content: pulumi.Input<string>;
    readonly path: pulumi.Input<string>;
}
export abstract class RootCa<TData = any> extends pulumi.ComponentResource<TData> {
    public allowedUses!: AllowedUsageOutputs[] | pulumi.Output<AllowedUsageOutputs[]>;
    public cert!: tls.index.SelfSignedCert | pulumi.Output<tls.index.SelfSignedCert>;
    public certPem!: string | pulumi.Output<string>;
    public key!: tls.index.PrivateKey | pulumi.Output<tls.index.PrivateKey>;
    public privateKeyPem!: string | pulumi.Output<string>;
    public publicKeyPem!: string | pulumi.Output<string>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:tls:RootCa", name, opts.urn ? { allowedUses: undefined, cert: undefined, certPem: undefined, key: undefined, privateKeyPem: undefined, publicKeyPem: undefined } : { name, args, opts }, opts);
    }
}
export interface RootCaArgs {
    readonly algorithm?: pulumi.Input<AlgorithmInputs>;
    readonly ecdsaCurve?: pulumi.Input<EcdsaCurveInputs>;
    readonly rsaBits?: pulumi.Input<number>;
    readonly dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    readonly earlyRenewalHours?: pulumi.Input<number>;
    readonly ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly setAuthorityKeyId?: pulumi.Input<boolean>;
    readonly setSubjectKeyId?: pulumi.Input<boolean>;
    readonly uris?: pulumi.Input<pulumi.Input<string>[]>;
    readonly validityPeriodHours: pulumi.Input<number>;
    readonly subject?: pulumi.Input<tls.types.input.index.SelfSignedCertSubject>;
}
export abstract class Mkdir<TData = any> extends pulumi.ComponentResource<TData> {
    public command!: command.Command | pulumi.Output<command.Command>;
    public directory!: string | pulumi.Output<string>;
    public parents!: boolean | pulumi.Output<boolean>;
    public removeOnDelete!: boolean | pulumi.Output<boolean>;
    public stderr!: string | pulumi.Output<string>;
    public stdout!: string | pulumi.Output<string>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:tools:Mkdir", name, opts.urn ? { command: undefined, directory: undefined, parents: undefined, removeOnDelete: undefined, stderr: undefined, stdout: undefined } : { name, args, opts }, opts);
    }
}
export interface MkdirArgs {
    readonly connection: pulumi.Input<command.types.input.Connection>;
    readonly directory: pulumi.Input<string>;
    readonly parents?: pulumi.Input<boolean>;
    readonly removeOnDelete?: pulumi.Input<boolean>;
}
export abstract class Mktemp<TData = any> extends pulumi.ComponentResource<TData> {
    public command!: command.Command | pulumi.Output<command.Command>;
    public directory!: boolean | pulumi.Output<boolean>;
    public dryRun!: boolean | pulumi.Output<boolean>;
    public quiet!: boolean | pulumi.Output<boolean>;
    public suffix?: string | pulumi.Output<string>;
    public template?: string | pulumi.Output<string>;
    public tmpdir?: string | pulumi.Output<string>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:tools:Mktemp", name, opts.urn ? { command: undefined, directory: undefined, dryRun: undefined, quiet: undefined, suffix: undefined, template: undefined, tmpdir: undefined } : { name, args, opts }, opts);
    }
}
export interface MktempArgs {
    readonly connection: pulumi.Input<command.types.input.Connection>;
    readonly directory?: pulumi.Input<boolean>;
    readonly dryRun?: pulumi.Input<boolean>;
    readonly quiet?: pulumi.Input<boolean>;
    readonly suffix?: pulumi.Input<string>;
    readonly template?: pulumi.Input<string>;
    readonly tmpdir?: pulumi.Input<string>;
}
export abstract class Mv<TData = any> extends pulumi.ComponentResource<TData> {
    public backup!: boolean | pulumi.Output<boolean>;
    public command!: command.Command | pulumi.Output<command.Command>;
    public context!: boolean | pulumi.Output<boolean>;
    public control?: string | pulumi.Output<string>;
    public dest?: string | pulumi.Output<string>;
    public directory?: string | pulumi.Output<string>;
    public force!: boolean | pulumi.Output<boolean>;
    public noClobber!: boolean | pulumi.Output<boolean>;
    public noTargetDirectory!: boolean | pulumi.Output<boolean>;
    public source!: string[] | pulumi.Output<string[]>;
    public stripTrailingSlashes!: boolean | pulumi.Output<boolean>;
    public suffix?: string | pulumi.Output<string>;
    public targetDirectory?: string | pulumi.Output<string>;
    public update!: boolean | pulumi.Output<boolean>;
    public verbose!: boolean | pulumi.Output<boolean>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:tools:Mv", name, opts.urn ? { backup: undefined, command: undefined, context: undefined, control: undefined, dest: undefined, directory: undefined, force: undefined, noClobber: undefined, noTargetDirectory: undefined, source: undefined, stripTrailingSlashes: undefined, suffix: undefined, targetDirectory: undefined, update: undefined, verbose: undefined } : { name, args, opts }, opts);
    }
}
export interface MvArgs {
    readonly backup?: boolean;
    readonly connection: pulumi.Input<command.types.input.Connection>;
    readonly context?: pulumi.Input<boolean>;
    readonly control?: pulumi.Input<string>;
    readonly dest?: pulumi.Input<string>;
    readonly directory?: pulumi.Input<string>;
    readonly force?: pulumi.Input<boolean>;
    readonly noClobber?: pulumi.Input<boolean>;
    readonly noTargetDirectory?: pulumi.Input<boolean>;
    readonly source: pulumi.Input<unknown>;
    readonly stripTrailingSlashes?: pulumi.Input<boolean>;
    readonly suffix?: pulumi.Input<string>;
    readonly targetDirectory?: pulumi.Input<string>;
    readonly update?: pulumi.Input<boolean>;
    readonly verbose?: pulumi.Input<boolean>;
}
export abstract class Rm<TData = any> extends pulumi.ComponentResource<TData> {
    public dir!: string | pulumi.Output<string>;
    public files!: string[] | pulumi.Output<string[]>;
    public force!: boolean | pulumi.Output<boolean>;
    public onDelete!: boolean | pulumi.Output<boolean>;
    public recursive!: boolean | pulumi.Output<boolean>;
    public verbose!: boolean | pulumi.Output<boolean>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:tools:Rm", name, opts.urn ? { dir: undefined, files: undefined, force: undefined, onDelete: undefined, recursive: undefined, verbose: undefined } : { name, args, opts }, opts);
    }
}
export interface RmArgs {
    readonly connection: pulumi.Input<command.types.input.Connection>;
    readonly dir?: pulumi.Input<boolean>;
    readonly files: pulumi.Input<unknown>;
    readonly force?: pulumi.Input<boolean>;
    readonly onDelete?: boolean;
    readonly recursive?: pulumi.Input<boolean>;
    readonly verbose?: pulumi.Input<boolean>;
}
export abstract class Tar<TData = any> extends pulumi.ComponentResource<TData> {
    public archive!: string | pulumi.Output<string>;
    public command!: command.Command | pulumi.Output<command.Command>;
    public directory?: string | pulumi.Output<string>;
    public extract!: boolean | pulumi.Output<boolean>;
    public files!: string[] | pulumi.Output<string[]>;
    public gzip?: boolean | pulumi.Output<boolean>;
    public stderr!: string | pulumi.Output<string>;
    public stdin?: string | pulumi.Output<string>;
    public stdout!: string | pulumi.Output<string>;
    public stripComponents?: number | pulumi.Output<number>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:tools:Tar", name, opts.urn ? { archive: undefined, command: undefined, directory: undefined, extract: undefined, files: undefined, gzip: undefined, stderr: undefined, stdin: undefined, stdout: undefined, stripComponents: undefined } : { name, args, opts }, opts);
    }
}
export interface TarArgs {
    readonly archive: pulumi.Input<string>;
    readonly connection: pulumi.Input<command.types.input.Connection>;
    readonly directory?: pulumi.Input<string>;
    readonly extract?: pulumi.Input<boolean>;
    readonly files?: pulumi.Input<unknown>;
    readonly gzip?: pulumi.Input<boolean>;
    readonly stripComponents?: pulumi.Input<number>;
}
export abstract class Wget<TData = any> extends pulumi.ComponentResource<TData> {
    public command!: command.Command | pulumi.Output<command.Command>;
    public directoryPrefix?: string | pulumi.Output<string>;
    public httpsOnly!: boolean | pulumi.Output<boolean>;
    public noVerbose?: boolean | pulumi.Output<boolean>;
    public outputDocument?: string | pulumi.Output<string>;
    public quiet!: boolean | pulumi.Output<boolean>;
    public stderr!: string | pulumi.Output<string>;
    public stdin?: string | pulumi.Output<string>;
    public stdout!: string | pulumi.Output<string>;
    public timestamping!: boolean | pulumi.Output<boolean>;
    public url!: string | pulumi.Output<string>;
    constructor(name: string, args: pulumi.Inputs, opts: pulumi.ComponentResourceOptions = {}) {
        super("kubernetes-the-hard-way:tools:Wget", name, opts.urn ? { command: undefined, directoryPrefix: undefined, httpsOnly: undefined, noVerbose: undefined, outputDocument: undefined, quiet: undefined, stderr: undefined, stdin: undefined, stdout: undefined, timestamping: undefined, url: undefined } : { name, args, opts }, opts);
    }
}
export interface WgetArgs {
    readonly connection: pulumi.Input<command.types.input.Connection>;
    readonly directoryPrefix?: pulumi.Input<string>;
    readonly httpsOnly?: pulumi.Input<boolean>;
    readonly noVerbose?: pulumi.Input<boolean>;
    readonly outputDocument?: pulumi.Input<string>;
    readonly quiet?: pulumi.Input<boolean>;
    readonly timestamping?: pulumi.Input<boolean>;
    readonly url: pulumi.Input<string>;
}
export type AlgorithmInputs = "RSA" | "ECDSA" | "ED25519";
export type AlgorithmOutputs = "RSA" | "ECDSA" | "ED25519";
export type AllowedUsageInputs = "cert_signing" | "client_auth" | "crl_signing" | "digital_signature" | "key_encipherment" | "server_auth";
export type AllowedUsageOutputs = "cert_signing" | "client_auth" | "crl_signing" | "digital_signature" | "key_encipherment" | "server_auth";
export type ArchitectureInputs = "amd64" | "arm64";
export type ArchitectureOutputs = "amd64" | "arm64";
export interface CertRequestSubjectInputs {
    readonly commonName?: pulumi.Input<string>;
    readonly country?: pulumi.Input<string>;
    readonly locality?: pulumi.Input<string>;
    readonly organization?: pulumi.Input<string>;
    readonly organizationalUnit?: pulumi.Input<string>;
    readonly postalCode?: pulumi.Input<string>;
    readonly province?: pulumi.Input<string>;
    readonly serialNumber?: pulumi.Input<string>;
    readonly streetAddresses?: pulumi.Input<pulumi.Input<string>[]>;
}
export interface CertRequestSubjectOutputs {
    readonly commonName?: pulumi.Output<string>;
    readonly country?: pulumi.Output<string>;
    readonly locality?: pulumi.Output<string>;
    readonly organization?: pulumi.Output<string>;
    readonly organizationalUnit?: pulumi.Output<string>;
    readonly postalCode?: pulumi.Output<string>;
    readonly province?: pulumi.Output<string>;
    readonly serialNumber?: pulumi.Output<string>;
    readonly streetAddresses?: pulumi.Output<string[]>;
}
export interface ConnectionInputs {
    readonly agentSocketPath?: pulumi.Input<string>;
    readonly dialErrorLimit?: pulumi.Input<number>;
    readonly host: pulumi.Input<string>;
    readonly password?: pulumi.Input<string>;
    readonly perDialTimeout?: pulumi.Input<number>;
    readonly port?: pulumi.Input<number>;
    readonly privateKey?: pulumi.Input<string>;
    readonly privateKeyPassword?: pulumi.Input<string>;
    readonly user?: pulumi.Input<string>;
}
export interface ConnectionOutputs {
    readonly agentSocketPath?: pulumi.Output<string>;
    readonly dialErrorLimit?: pulumi.Output<number>;
    readonly host: pulumi.Output<string>;
    readonly password?: pulumi.Output<string>;
    readonly perDialTimeout?: pulumi.Output<number>;
    readonly port?: pulumi.Output<number>;
    readonly privateKey?: pulumi.Output<string>;
    readonly privateKeyPassword?: pulumi.Output<string>;
    readonly user?: pulumi.Output<string>;
}
export interface KubeconfigClusterInputs {
    readonly cluster: pulumi.Input<ClusterInputs>;
    readonly name: pulumi.Input<string>;
}
export interface KubeconfigClusterOutputs {
    readonly cluster: pulumi.Output<ClusterOutputs>;
    readonly name: pulumi.Output<string>;
}
export interface ClusterInputs {
    readonly certificateAuthorityData: pulumi.Input<string>;
    readonly server: pulumi.Input<string>;
}
export interface ClusterOutputs {
    readonly certificateAuthorityData: pulumi.Output<string>;
    readonly server: pulumi.Output<string>;
}
export interface ClusterPkiNodeInputs {
    readonly ip?: pulumi.Input<string>;
    readonly role?: pulumi.Input<NodeRoleInputs>;
}
export interface ClusterPkiNodeOutputs {
    readonly ip?: pulumi.Output<string>;
    readonly role?: pulumi.Output<NodeRoleOutputs>;
}
export interface KubeconfigContextInputs {
    readonly context: pulumi.Input<ContextInputs>;
    readonly name: pulumi.Input<string>;
}
export interface KubeconfigContextOutputs {
    readonly context: pulumi.Output<ContextOutputs>;
    readonly name: pulumi.Output<string>;
}
export interface ContextInputs {
    readonly cluster: pulumi.Input<string>;
    readonly user: pulumi.Input<string>;
}
export interface ContextOutputs {
    readonly cluster: pulumi.Output<string>;
    readonly user: pulumi.Output<string>;
}
export type EcdsaCurveInputs = "P224" | "P256" | "P384" | "P521";
export type EcdsaCurveOutputs = "P224" | "P256" | "P384" | "P521";
export interface KeyPairInputs {
    readonly certPem?: pulumi.Input<string>;
    readonly key?: pulumi.Input<tls.index.PrivateKey>;
    readonly privateKeyPem?: pulumi.Input<string>;
    readonly publicKeyPem?: pulumi.Input<string>;
}
export interface KeyPairOutputs {
    readonly certPem?: pulumi.Output<string>;
    readonly key?: pulumi.Output<tls.index.PrivateKey>;
    readonly privateKeyPem?: pulumi.Output<string>;
    readonly publicKeyPem?: pulumi.Output<string>;
}
export interface KubeconfigInputs {
    readonly clusters: pulumi.Input<pulumi.Input<ClusterInputs>[]>;
    readonly contexts: pulumi.Input<pulumi.Input<ContextInputs>[]>;
    readonly users: pulumi.Input<pulumi.Input<UserInputs>[]>;
}
export interface KubeconfigOutputs {
    readonly clusters: pulumi.Output<ClusterOutputs[]>;
    readonly contexts: pulumi.Output<ContextOutputs[]>;
    readonly users: pulumi.Output<UserOutputs[]>;
}
export type KubeconfigTypeInputs = "worker" | "kube-proxy" | "kube-controller-manager" | "kube-scheduler" | "admin";
export type KubeconfigTypeOutputs = "worker" | "kube-proxy" | "kube-controller-manager" | "kube-scheduler" | "admin";
export interface KubeconfigAdminOptionsInputs {
    readonly type: string;
    readonly publicIp?: pulumi.Input<string>;
}
export interface KubeconfigAdminOptionsOutputs {
    readonly type: string;
    readonly publicIp?: pulumi.Output<string>;
}
export interface KubeconfigKubeControllerManagerOptionsInputs {
    readonly type: string;
    readonly publicIp?: pulumi.Input<string>;
}
export interface KubeconfigKubeControllerManagerOptionsOutputs {
    readonly type: string;
    readonly publicIp?: pulumi.Output<string>;
}
export interface KubeconfigKubeProxyOptionsInputs {
    readonly type: string;
    readonly publicIp?: pulumi.Input<string>;
}
export interface KubeconfigKubeProxyOptionsOutputs {
    readonly type: string;
    readonly publicIp?: pulumi.Output<string>;
}
export interface KubeconfigKubeSchedulerOptionsInputs {
    readonly type: string;
    readonly publicIp?: pulumi.Input<string>;
}
export interface KubeconfigKubeSchedulerOptionsOutputs {
    readonly type: string;
    readonly publicIp?: pulumi.Output<string>;
}
export interface KubeconfigWorkerOptionsInputs {
    readonly type?: string;
    readonly name: pulumi.Input<string>;
    readonly publicIp: pulumi.Input<string>;
}
export interface KubeconfigWorkerOptionsOutputs {
    readonly type?: string;
    readonly name: pulumi.Output<string>;
    readonly publicIp: pulumi.Output<string>;
}
export type NodeRoleInputs = "controlplane" | "worker";
export type NodeRoleOutputs = "controlplane" | "worker";
export interface ResourceOptionsInputs {
    readonly parent?: pulumi.Input<any>;
}
export interface ResourceOptionsOutputs {
    readonly parent?: pulumi.Output<any>;
}
export interface KubeconfigUserInputs {
    readonly name: pulumi.Input<string>;
    readonly user: pulumi.Input<UserInputs>;
}
export interface KubeconfigUserOutputs {
    readonly name: pulumi.Output<string>;
    readonly user: pulumi.Output<UserOutputs>;
}
export interface UserInputs {
    readonly clientCertificateData: pulumi.Input<string>;
    readonly clientKeyData: pulumi.Input<string>;
}
export interface UserOutputs {
    readonly clientCertificateData: pulumi.Output<string>;
    readonly clientKeyData: pulumi.Output<string>;
}
export interface ClusterPki_getKubeconfigInputs {
    readonly __self__: pulumi.Input<ClusterPki>;
    readonly options: unknown;
}
export interface ClusterPki_getKubeconfigOutputs {
    readonly result: pulumi.Output<KubeconfigOutputs>;
}
