package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateTar() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"archive":   props.String("Corresponds to the [ARCHIVE] argument."),
		"directory": props.String("Corresponds to the `--directory` option."),
		"extract":   props.Boolean("Corresponds to the `--extract` option."),
		"files":     props.OneOrMoreStrings("Corresponds to the [FILE] argument."),
		"gzip":      props.Boolean("Corresponds to the `--gzip` option."),
		// TODO: Reconsider this SOLID violation
		"onDelete":        props.Boolean("Whether rm should be run when the resource is created or deleted."),
		"recursive":       props.Boolean("Corresponds to the `--recursive` option."),
		"stripComponents": props.Integer("Corresponds to the `--strip-components` option."),
	}

	required := []string{
		"archive",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `rm` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"extract",
				"files",
			),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}
