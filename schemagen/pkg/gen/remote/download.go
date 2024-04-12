package remote

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateDownload(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"destination": {
			Description: "The fully qualified path on the remote system where the file should be downloaded to.",
			TypeSpec:    types.String,
		},
		"removeOnDelete": {
			Description: "Remove the downloaded fiel when the resource is deleted.",
			TypeSpec:    types.Boolean,
		},
		"url": {
			Description: "The URL of the file to be downloaded.",
			TypeSpec:    types.String,
		},
	}

	requiredInputs := []string{
		"connection",
		"destination",
		"url",
	}

	outputs := map[string]schema.PropertySpec{
		"mkdir": {
			Description: "The mkdir operation.",
			TypeSpec:    types.LocalResource("Mkdir", "tools"),
		},
		"wget": {
			Description: "The wget operation.",
			TypeSpec:    types.LocalResource("Wget", "tools"),
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

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Downloads the file specified by `url` onto a remote system.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
