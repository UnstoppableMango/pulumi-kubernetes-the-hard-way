package remote

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateStaticPod(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"fileName": {
			Description: "The name of the file on the remote system.",
			TypeSpec:    types.String,
		},
		"pod": {
			Description: "The pod manifest.",
			TypeSpec:    types.LocalType("PodManifest", "config"),
		},
	}

	outputs := map[string]schema.PropertySpec{
		"file": {
			Description: "The remote manifest file.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
		"mkdir": {
			Description: "The mkdir operation to ensure /etc/kubernetes/manifests exists.",
			TypeSpec:    types.LocalResource("Mkdir", "tools"),
		},
		"path": {
			Description: "The path to the manifest on the remote system.",
			TypeSpec:    types.String,
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := []string{
		"connection",
		"file",
		"fileName",
		"mkdir",
		"path",
		"pod",
	}

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Create a static pod manifest on a remote system.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  []string{"connection", "pod"},
	}
}
