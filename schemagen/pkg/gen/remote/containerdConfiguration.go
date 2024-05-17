package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateContainerdConfiguration(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"cri": {
			Description: "The cri configuration.",
			TypeSpec: schema.TypeSpec{
				Plain: true,
				Ref:   types.LocalTypeRef("ContainerdCriPluginConfiguration", "remote"),
			},
		},
	}

	requiredInputs := []string{"connection"}

	outputs := map[string]schema.PropertySpec{
		"file": {
			Description: "The remote configuration file.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(requiredInputs, []string{
		"file",
		"cri",
	})

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "The containerd configuration file.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
