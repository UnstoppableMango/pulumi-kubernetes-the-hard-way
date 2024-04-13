package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateSystemctl() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"command": { // The common output will overwrite this, that's fine.
			Description: "Corresponds to the COMMAND argument.",
			TypeSpec: schema.TypeSpec{
				Ref:   types.LocalType("SystemctlCommand", "tools").Ref,
				Plain: true,
			},
		},
		"pattern": props.String("Corresponds to the [PATTERN] argument"),
		"unit":    props.String("Corresponds to the [UNIT...] argument."),
	}

	required := []string{
		"command",
		"unit",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `systemctl` utility on a remote system.",
			Properties: implicitOutputs(inputs, map[string]schema.PropertySpec{
				"systemctlCommand": inputs["command"],
			}),
			Required: []string{
				"systemctlCommand",
				"unit",
			},
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}
