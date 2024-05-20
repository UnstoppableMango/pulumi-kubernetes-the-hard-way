package remote

import (
	"fmt"
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateStartSystemdService(commandSpec schema.PackageSpec, name string) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": props.Connection(commandSpec),
	}
	requiredInputs := []string{"connection"}

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
			Description: fmt.Sprintf("Starts `%s` on a remote system", name),
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
