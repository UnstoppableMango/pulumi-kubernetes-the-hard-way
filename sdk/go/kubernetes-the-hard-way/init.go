// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetesthehardway

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
	case "kubernetes-the-hard-way:index:Certificate":
		r = &Certificate{}
	case "kubernetes-the-hard-way:index:ClusterPki":
		r = &ClusterPki{}
	case "kubernetes-the-hard-way:index:EncryptionKey":
		r = &EncryptionKey{}
	case "kubernetes-the-hard-way:index:RemoteFile":
		r = &RemoteFile{}
	case "kubernetes-the-hard-way:index:RootCa":
		r = &RootCa{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:kubernetes-the-hard-way" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"kubernetes-the-hard-way",
		"index",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"kubernetes-the-hard-way",
		&pkg{version},
	)
}
