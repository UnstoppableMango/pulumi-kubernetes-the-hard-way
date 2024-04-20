package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateTar() tool {
	inputs := map[string]schema.PropertySpec{
		"archive":   props.String("Corresponds to the [ARCHIVE] argument."),
		"directory": props.String("Corresponds to the `--directory` option."),
		"extract":   props.Boolean("Corresponds to the `--extract` option."),
		"files":     props.ArrayOf("string", "Corresponds to the [FILE] argument."),
		"gzip":      props.Boolean("Corresponds to the `--gzip` option."),
		// TODO: Reconsider this SOLID violation
		"onDelete":        props.Boolean("Whether rm should be run when the resource is created or deleted."),
		"recursive":       props.Boolean("Corresponds to the `--recursive` option."),
		"stripComponents": props.Integer("Corresponds to the `--strip-components` option."),
	}

	required := []string{"archive"}

	typ := schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `tar` utility on a remote system.",
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
// 			"extract",
// 			"files",
// 		),
// 	},
// 	InputProperties: inputs,
// 	RequiredInputs:  required,
// }
