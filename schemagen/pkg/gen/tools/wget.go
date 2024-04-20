package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateWget() tool {
	inputs := map[string]schema.PropertySpec{
		"directoryPrefix": props.String("The  directory prefix is the directory where all other files and subdirectories will be saved to, i.e. the top of the retrieval tree.  The default is . (the current directory)."),
		"httpsOnly":       props.Boolean("When in recursive mode, only HTTPS links are followed."),
		"noVerbose":       props.Boolean("Turn off verbose without being completely quiet (use -q for that), which means that error messages and basic information still get printed."),
		"outputDocument":  props.String("The  documents  will  not  be  written  to the appropriate files, but all will be concatenated together and written to file."),
		"quiet":           props.Boolean("Turn off Wget's output."),
		"timestamping":    props.Boolean("Turn on time-stamping."),
		"url":             props.ArrayOf("string", "Corresponds to the [URL...] argument."),
	}

	required := []string{"url"}

	typ := schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `wget` utility on a remote system.",
			Properties:  inputs,
			Required:    required,
		},
	}

	return tool{optsType: typ, types: map[string]schema.ComplexTypeSpec{}}
}

// If we ever get a way to add the "required outputs" logic around a complexType
// resource := schema.ResourceSpec{
// 	ObjectTypeSpec: schema.ObjectTypeSpec{
// 		Description: "Abstraction over the `wget` utility on a remote system.",
// 		Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
// 		Required: append(required,
// 			"httpsOnly",
// 			"quiet",
// 			"noVerbose",
// 			"timestamping",
// 		),
// 	},
// 	InputProperties: inputs,
// 	RequiredInputs:  required,
// }
