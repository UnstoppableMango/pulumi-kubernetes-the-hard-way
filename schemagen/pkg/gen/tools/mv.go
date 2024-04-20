package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateMv() tool {
	inputs := map[string]schema.PropertySpec{
		"backup": {
			Description: "Corresponds to the `-b` and `--backup` options depending on whether [CONTROL] is supplied.",
			TypeSpec: schema.TypeSpec{
				Type:  "boolean",
				Plain: true,
			},
		},
		"context":              props.Boolean("Corresponds to the `--context` option."),
		"control":              props.Boolean("Corresponds to the [CONTROL] argument for the `--backup` option."),
		"dest":                 props.String("Corresponds to the [DEST] argument."),
		"directory":            props.String("Corresponds to the [DIRECTORY] argument."),
		"force":                props.Boolean("Corresponds to the `--force` option."),
		"noClobber":            props.Boolean("Corresponds to the `--no-clobber` option."),
		"noTargetDirectory":    props.Boolean("Corresponds to the `--no-target-directory` option."),
		"source":               props.ArrayOf("string", "Corresponds to the [SOURCE] argument."),
		"stripTrailingSlashes": props.Boolean("Corresponds to the `--strip-trailing-slashes` option."),
		"suffix":               props.String("Corresponds to the `--suffix` option."),
		"targetDirectory":      props.Boolean("Corresponds to the `--target-directory` option."),
		"update":               props.Boolean("Corresponds to the `--update` option."),
		"verbose":              props.Boolean("Corresponds to the `--verbose` option."),
	}

	required := []string{"source"}

	typ := schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `mv` utility on a remote system.",
			Properties:  inputs,
			Required:    required,
		},
	}

	return tool{optsType: typ, types: map[string]schema.ComplexTypeSpec{}}
}

// If we ever get a way to add the "required outputs" logic around a complexType
// resource := schema.ResourceSpec{
// 	ObjectTypeSpec: schema.ObjectTypeSpec{
// 		Description: "Abstraction over the `mv` utility on a remote system.",
// 		Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
// 		Required: append(required,
// 			"backup",
// 			"context",
// 			"force",
// 			"noClobber",
// 			"noTargetDirectory",
// 			"source",
// 			"stripTrailingSlashes",
// 			"update",
// 			"verbose",
// 		),
// 	},
// 	InputProperties: inputs,
// 	RequiredInputs:  required,
// }
