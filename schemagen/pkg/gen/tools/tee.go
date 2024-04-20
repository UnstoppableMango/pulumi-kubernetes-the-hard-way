package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateTee() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"append":           props.Boolean("Append to the given FILEs, do not overwrite"),
		"files":            props.ArrayOf("string", "Corresponds to the [FILE] argument."),
		"ignoreInterrupts": props.Boolean("Ignore interrupt signals."),
		"outputError": {
			Description: "Set behavior on write error.",
			TypeSpec:    types.LocalType("TeeMode", "tools"),
		},
		"pipe":    props.Boolean("Operate in a more appropriate MODE with pipes."),
		"version": props.Boolean("Output version information and exit."),
	}

	required := []string{
		"files",
		"stdin",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `rm` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"append",
				"ignoreInterrupts",
				"pipe",
				"version",
			),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}
