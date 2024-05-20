package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateGetKubeProxyConfiguration() schema.PackageSpec {
	function, resource := generatePseudoFunction(
		"kube-proxy configuration.",
		schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"kubeconfig":  props.String("Path to the kubeconfig."),
				"mode":        props.String("TODO"),
				"clusterCIDR": props.String("Cluster CIDR."),
			},
			Required: []string{"kubeconfig", "clusterCIDR"},
		},
		schema.PropertySpec{
			TypeSpec: types.LocalType("KubeProxyConfiguration", "config"),
		},
		"yaml", props.String("The yaml representation of the manifest."),
	)

	return schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{},
		Functions: map[string]schema.FunctionSpec{
			name("getKubeProxyConfiguration"): function,
		},
		Resources: map[string]schema.ResourceSpec{
			name("KubeProxyConfiguration"): resource,
		},
	}
}
