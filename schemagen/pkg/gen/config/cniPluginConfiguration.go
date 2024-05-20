package config

import (
	"fmt"
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateGetCniBridgePluginConfiguration() schema.PackageSpec {
	function, resource := generateCniPluginConfiguration(
		"CniBridgePluginConfiguration",
		"bridge",
		map[string]schema.PropertySpec{
			"bridge":    props.String("Bridge name."),
			"subnet":    props.String("The subnet to use."),
			"isGateway": props.Boolean("Is gateway."),
			"ipMasq":    props.Boolean("IP masq."),
			"ipam": {
				Description: "IPAM",
				TypeSpec:    types.LocalType("CniBridgeIpam", "remote"),
			},
		},
		[]string{"subnet"},
	)

	return schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{},
		Functions: map[string]schema.FunctionSpec{
			name("getCniBridgePluginConfiguration"): function,
		},
		Resources: map[string]schema.ResourceSpec{
			name("CniBridgePluginConfiguration"): resource,
		},
	}
}

func generateGetCniLoopbackPluginConfiguration() schema.PackageSpec {
	function, resource := generateCniPluginConfiguration(
		"CniLoopbackPluginConfiguration",
		"loopback",
		map[string]schema.PropertySpec{},
		[]string{},
	)

	return schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{},
		Functions: map[string]schema.FunctionSpec{
			name("getCniLoopbackPluginConfiguration"): function,
		},
		Resources: map[string]schema.ResourceSpec{
			name("CniLoopbackPluginConfiguration"): resource,
		},
	}
}

func generateCniPluginConfiguration(
	typ, name string,
	properties map[string]schema.PropertySpec,
	required []string,
) (schema.FunctionSpec, schema.ResourceSpec) {
	maps.Copy(properties, map[string]schema.PropertySpec{
		"cniVersion": props.String("CNI version."),
		"name":       props.String("CNI plugin name."),
		"path":       props.String("Path to put the configuration file on the remote system"),
		"type":       props.String("CNI plugin type."),
	})

	return generatePseudoFunction(
		fmt.Sprintf("Get the `%s` configuration.", name),
		schema.ObjectTypeSpec{
			Properties: properties,
			Required:   required,
		},
		schema.PropertySpec{
			TypeSpec: types.LocalType(typ, "config"),
		},
		"yaml", props.String("The yaml representation of the manifest."),
	)
}
