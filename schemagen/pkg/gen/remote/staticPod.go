package remote

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateStaticPod(commandSpec schema.PackageSpec) schema.PackageSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": props.Connection(commandSpec),
		"fileName":   props.String("The name of the file on the remote system."),
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
		"path": props.String("The path to the manifest on the remote system."),
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

	return schema.PackageSpec{
		Types:     map[string]schema.ComplexTypeSpec{},
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{
			name("StaticPod"): {
				IsComponent: true,
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "Create a static pod manifest on a remote system.",
					Properties:  outputs,
					Required:    requiredOutputs,
				},
				InputProperties: inputs,
				RequiredInputs:  []string{"connection", "pod"},
			},
		},
	}
}
