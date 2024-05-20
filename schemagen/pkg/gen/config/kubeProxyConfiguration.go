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
		Types: map[string]schema.ComplexTypeSpec{
			name("KubeProxyConfigurationClientConnection"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"kubeconfig": props.String("Path to the kubeconfig."),
					},
					Required: []string{"kubeconfig"},
				},
			},
			name("KubeProxyConfiguration"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"kind": {
							TypeSpec: types.String,
							Const:    "KubeProxyConfiguration",
						},
						"apiVersion": {
							TypeSpec: types.String,
							Const:    "kubeproxy.config.k8s.io/v1alpha1",
						},
						"clientConnection": {
							TypeSpec: types.LocalType("KubeProxyConfigurationClientConnection", "config"),
						},
						"mode":        props.String("TODO"),
						"clusterCIDR": props.String("TODO"),
					},
					Required: []string{"clusterCIDR"},
				},
			},
		},
		Functions: map[string]schema.FunctionSpec{
			name("getKubeProxyConfiguration"): function,
		},
		Resources: map[string]schema.ResourceSpec{
			name("KubeProxyConfiguration"): resource,
		},
	}
}
