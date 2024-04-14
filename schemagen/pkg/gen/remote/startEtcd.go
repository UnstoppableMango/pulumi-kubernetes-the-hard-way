package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateStartEtcd() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{}
	requiredInputs := []string{}

	outputs := map[string]schema.PropertySpec{
		"daemonReload": {
			Description: "The daemon-reload command.",
			TypeSpec:    types.LocalResource("Systemctl", "tools"),
		},
		"enable": {
			Description: "The enable command.",
			TypeSpec:    types.LocalResource("Systemctl", "tools"),
		},
		"start": {
			Description: "The start command.",
			TypeSpec:    types.LocalResource("Systemctl", "tools"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(
		requiredInputs,
		[]string{
			"daemonReload",
			"enable",
			"start",
		},
	)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Starts etcd on a remote system.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
