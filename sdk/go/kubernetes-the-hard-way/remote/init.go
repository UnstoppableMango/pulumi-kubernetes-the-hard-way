// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes-the-hard-way:remote:CniPluginsInstall":
		r = &CniPluginsInstall{}
	case "kubernetes-the-hard-way:remote:ContainerdInstall":
		r = &ContainerdInstall{}
	case "kubernetes-the-hard-way:remote:CrictlInstall":
		r = &CrictlInstall{}
	case "kubernetes-the-hard-way:remote:Download":
		r = &Download{}
	case "kubernetes-the-hard-way:remote:EtcdCluster":
		r = &EtcdCluster{}
	case "kubernetes-the-hard-way:remote:EtcdConfiguration":
		r = &EtcdConfiguration{}
	case "kubernetes-the-hard-way:remote:EtcdInstall":
		r = &EtcdInstall{}
	case "kubernetes-the-hard-way:remote:EtcdService":
		r = &EtcdService{}
	case "kubernetes-the-hard-way:remote:File":
		r = &File{}
	case "kubernetes-the-hard-way:remote:KubeApiServerConfiguration":
		r = &KubeApiServerConfiguration{}
	case "kubernetes-the-hard-way:remote:KubeApiServerInstall":
		r = &KubeApiServerInstall{}
	case "kubernetes-the-hard-way:remote:KubeApiServerService":
		r = &KubeApiServerService{}
	case "kubernetes-the-hard-way:remote:KubeControllerManagerInstall":
		r = &KubeControllerManagerInstall{}
	case "kubernetes-the-hard-way:remote:KubeProxyInstall":
		r = &KubeProxyInstall{}
	case "kubernetes-the-hard-way:remote:KubeSchedulerInstall":
		r = &KubeSchedulerInstall{}
	case "kubernetes-the-hard-way:remote:KubectlInstall":
		r = &KubectlInstall{}
	case "kubernetes-the-hard-way:remote:KubeletInstall":
		r = &KubeletInstall{}
	case "kubernetes-the-hard-way:remote:ProvisionEtcd":
		r = &ProvisionEtcd{}
	case "kubernetes-the-hard-way:remote:RuncInstall":
		r = &RuncInstall{}
	case "kubernetes-the-hard-way:remote:StartEtcd":
		r = &StartEtcd{}
	case "kubernetes-the-hard-way:remote:StaticPod":
		r = &StaticPod{}
	case "kubernetes-the-hard-way:remote:SystemdService":
		r = &SystemdService{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"kubernetes-the-hard-way",
		"remote",
		&module{version},
	)
}
