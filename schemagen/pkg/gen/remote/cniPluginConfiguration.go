package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateCniPluginConfiguration(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"bridge": {
			Description: "The CNI bridge configuration.",
			TypeSpec:    types.LocalType("CniBridgePluginConfiguration", "remote"),
		},
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"loopback": {
			Description: "The CNI loopback configuration.",
			TypeSpec:    types.LocalType("CniLoopbackPluginConfiguration", "remote"),
		},
	}

	requiredInputs := []string{"connection"}

	outputs := map[string]schema.PropertySpec{}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(requiredInputs, []string{
		"bridge",
		"loopback",
	})

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "CNI plugin configuration.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
