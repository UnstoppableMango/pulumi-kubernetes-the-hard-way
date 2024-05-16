package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateCniBridgePluginConfiguration(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"bridge":     props.String("Bridge name."),
		"cniVersion": props.String("CNI version."),
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"isGateway": props.Boolean("Is gateway."),
		"ipMasq":    props.Boolean("IP masq."),
		"ipam": {
			Description: "IPAM",
			TypeSpec:    types.LocalType("CniBridgeIpam", "remote"),
		},
		"name":   props.String("CNI plugin name."),
		"path":   props.String("Path to put the configuration file on the remote system"),
		"subnet": props.String("The subnet to use."),
		"type":   props.String("CNI plugin type."),
	}

	requiredInputs := []string{"connection"}

	outputs := map[string]schema.PropertySpec{
		"file": {
			Description: "The file on the remote system.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(requiredInputs, []string{
		"bridge",
		"cniVersion",
		"isGateway",
		"ipMasq",
		"ipam",
		"name",
		"path",
		"type",
	})

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "The CNI bridge plugin configuration.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
