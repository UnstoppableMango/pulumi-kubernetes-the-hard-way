package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateRm() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"dir":   props.Boolean("Corresponds to the `--dir` option."),
		"files": props.OneOrMoreStrings("Corresponds to the [FILE] argument."),
		"force": props.Boolean("Corresponds to the `--force` option."),
		// TODO: Reconsider this SOLID violation
		"onDelete":  props.Boolean("Whether rm should be run when the resource is created or deleted."),
		"recursive": props.Boolean("Corresponds to the `--recursive` option."),
		"verbose":   props.Boolean("Corresponds to the `--verbose` option."),
	}

	required := []string{
		"files",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `rm` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"dir",
				"force",
				"onDelete",
				"recursive",
				"verbose",
			),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}