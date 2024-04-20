package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateMktemp() tool {
	inputs := map[string]schema.PropertySpec{
		"directory": props.Boolean("Corresponds to the `--directory` option."),
		"dryRun":    props.Boolean("Corresponds to the `--dry-run` option."),
		"quiet":     props.Boolean("Corresponds to the `--quiet` option."),
		"suffix":    props.String("Corresponds to the `--suffix` option."),
		"template":  props.String("Corresponds to the [TEMPLATE] argument."),
		"tmpdir":    props.String("Corresponds to the `--tmpdir` option."),
	}

	typ := schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `mktemp` utility on a remote system.",
			Properties:  inputs,
			Required:    []string{},
		},
	}

	return tool{optsType: typ, types: map[string]schema.ComplexTypeSpec{}}
}

// If we ever get a way to add the "required outputs" logic around a complexType
// resource := schema.ResourceSpec{
// 	ObjectTypeSpec: schema.ObjectTypeSpec{
// 		Description: "Abstraction over the `mkdir` utility on a remote system.",
// 		Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
// 		Required: []string{
// 			"dryRun",
// 			"quiet",
// 		},
// 	},
// 	InputProperties: inputs,
// 	RequiredInputs:  []string{},
// }
