package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateGetKubeletConfiguration() schema.PackageSpec {
	function, resource := generatePseudoFunction(
		"Get the kubelet configuration.",
		schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"anonymous":                props.Boolean(""),
				"webhook":                  props.Boolean(""),
				"clientCAFile":             props.String(""),
				"authorizationMode":        props.String(""),
				"clusterDomain":            props.String(""),
				"clusterDNS":               props.ArrayOf("string", ""),
				"cgroupDriver":             props.String(""),
				"containerRuntimeEndpoint": props.String(""),
				"podCIDR":                  props.String(""),
				"resolvConf":               props.String(""),
				"runtimeRequestTimeout":    props.String(""),
				"tlsCertFile":              props.String(""),
				"tlsPrivateKeyFile":        props.String(""),
			},
			Required: []string{"podCIDR"},
		},
		schema.PropertySpec{
			TypeSpec: types.LocalType("KubeletConfiguration", "config"),
		},
		"yaml", props.String("The yaml representation of the manifest."),
	)

	return schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{},
		Functions: map[string]schema.FunctionSpec{
			name("getKubeletConfiguration"): function,
		},
		Resources: map[string]schema.ResourceSpec{
			name("KubeletConfiguration"): resource,
		},
	}
}
