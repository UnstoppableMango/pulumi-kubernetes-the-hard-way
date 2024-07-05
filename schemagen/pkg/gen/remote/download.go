package remote

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateDownload(commandSpec, commandxSpec schema.PackageSpec) schema.PackageSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"destination": props.String("The fully qualified path on the remote system where the file should be downloaded to."),
		"removeOnDelete": {
			Description: "Remove the downloaded fiel when the resource is deleted.",
			TypeSpec:    schema.TypeSpec{Type: "boolean", Plain: true},
		},
		"url": props.String("The URL of the file to be downloaded."),
	}

	requiredInputs := []string{
		"connection",
		"destination",
		"url",
	}

	outputs := map[string]schema.PropertySpec{
		"mkdir": {
			Description: "The mkdir operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mkdir", "remote"),
		},
		"wget": {
			Description: "The wget operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Wget", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := []string{
		"connection",
		"destination",
		"mkdir",
		"removeOnDelete",
		"url",
		"wget",
	}

	return schema.PackageSpec{
		Types:     map[string]schema.ComplexTypeSpec{},
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{
			name("Download"): {
				IsComponent: true,
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "Downloads the file specified by `url` onto a remote system.",
					Properties:  outputs,
					Required:    requiredOutputs,
				},
				InputProperties: inputs,
				RequiredInputs:  requiredInputs,
			},
		},
	}
}
