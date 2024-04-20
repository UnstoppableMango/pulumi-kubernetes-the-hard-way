package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateChmod() tool {
	types := map[string]schema.ComplexTypeSpec{}

	inputs := map[string]schema.PropertySpec{
		"changes":        props.Boolean("Like verbose but report only when a change is made."),
		"files":          props.ArrayOf("string", "Corresponds to the [FILE] argument."),
		"help":           props.Boolean("Display help and exit."),
		"mode":           props.String("Modes may be absolute or symbolic. An absolute mode is an octal number..."),
		"noPreserveRoot": props.Boolean("Do not treat '/' specially (the default)."),
		"preserveRoot":   props.Boolean("Fail to operate recursively on '/'."),
		"quiet":          props.Boolean("Suppress most error messages. Same as `silent`."),
		"recursive":      props.Boolean("Change files and directories recursively."),
		"reference":      props.String("Use RFILE's mode instead of specifying MODE values. RFILE is always dereferenced if a symbolic link."),
		"silent":         props.Boolean("Suppress most error messages. Same as `quiet`."),
		"version":        props.Boolean("Output version information and exit."),
	}

	required := []string{"files", "mode"}

	typ := schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `chmod` utility on a remote system.",
			Type:        "object",
			Properties:  inputs,
			Required:    required,
		},
	}

	return tool{optsType: typ, types: types}
}

// If we ever get a way to add the "required outputs" logic around a complexType
// resource := schema.ResourceSpec{
// 	ObjectTypeSpec: schema.ObjectTypeSpec{
// 		Description: "Abstraction over the `chmod` utility on a remote system.",
// 		Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
// 		Required: append(required,
// 			"changes",
// 			"help",
// 			"noPreserveRoot",
// 			"preserveRoot",
// 			"quiet",
// 			"recursive",
// 			"silent",
// 			"version",
// 		),
// 	},
// 	InputProperties: inputs,
// 	RequiredInputs:  required,
// }
