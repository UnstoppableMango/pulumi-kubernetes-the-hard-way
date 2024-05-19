package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateGetKubeProxyConfiguration() pseudoFunction {
	return generatePseudoFunction(
		"kube-proxy configuration.",
		schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"kubeconfig":  props.String("Path to the kubeconfig."),
				"mode":        props.String("TODO"),
				"clusterCIDR": props.String("Cluster CIDR."),
			},
			Required: []string{"kubeconfig"},
		},
		schema.PropertySpec{
			TypeSpec: types.LocalType("KubeProxyConfiguration", "config"),
		},
		"yaml", props.String("The yaml representation of the manifest."),
	)
}
