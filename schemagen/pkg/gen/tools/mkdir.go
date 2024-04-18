package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateMkdir() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"directory": props.String("The fully qualified path of the directory on the remote system."),
		"parents":   props.Boolean("Corresponds to the `--parents` option."),
		// TODO: Reconsider this SOLID violation
		"removeOnDelete": props.Boolean("Remove the created directory when the `Mkdir` resource is deleted or updated."),
	}

	required := []string{
		"directory",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `mkdir` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"parents",
				"removeOnDelete"),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}