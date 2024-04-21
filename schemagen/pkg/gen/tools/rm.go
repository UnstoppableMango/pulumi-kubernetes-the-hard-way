package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateRm() tool {
	inputs := map[string]schema.PropertySpec{
		"dir":   props.Boolean("Corresponds to the `--dir` option."),
		"files": props.ArrayOf("string", "Corresponds to the [FILE] argument."),
		"force": props.Boolean("Corresponds to the `--force` option."),
		// TODO: Reconsider this SOLID violation
		"onDelete":  props.Boolean("Whether rm should be run when the resource is created or deleted."),
		"recursive": props.Boolean("Corresponds to the `--recursive` option."),
		"verbose":   props.Boolean("Corresponds to the `--verbose` option."),
	}

	required := []string{"files"}

	typ := schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `rm` utility on a remote system.",
			Type:        "object",
			Properties:  inputs,
			Required:    required,
		},
	}

	return tool{optsType: typ, types: map[string]schema.ComplexTypeSpec{}}
}

// If we ever get a way to add the "required outputs" logic around a complexType
// resource := schema.ResourceSpec{
// 	ObjectTypeSpec: schema.ObjectTypeSpec{
// 		Description: "Abstraction over the `rm` utility on a remote system.",
// 		Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
// 		Required: append(required,
// 			"dir",
// 			"force",
// 			"onDelete",
// 			"recursive",
// 			"verbose",
// 		),
// 	},
// 	InputProperties: inputs,
// 	RequiredInputs:  required,
// }
