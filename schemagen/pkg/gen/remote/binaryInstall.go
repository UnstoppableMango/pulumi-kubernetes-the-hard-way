package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateBinaryInstall(
	commandSpec, commandxSpec schema.PackageSpec,
	typ, description string,
) schema.PackageSpec {
	inputs := map[string]schema.PropertySpec{
		"architecture": {
			Description: "The CPU architecture to install.",
			TypeSpec:    types.LocalType("Architecture", "remote"),
		},
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"directory": {
			Description: "The directory to install the binary to.",
			TypeSpec:    types.String,
			Default:     "/usr/local/bin",
		},
		"version": {
			Description: "The version to install.",
			TypeSpec:    types.String,
		},
	}

	requiredInputs := []string{"connection"}

	outputs := map[string]schema.PropertySpec{
		"binName": {
			Description: "The name of the installed binary.",
			TypeSpec:    types.String,
		},
		"download": {
			Description: "The download operation.",
			TypeSpec:    types.LocalResource("Download", "remote"),
		},
		"mkdir": {
			Description: "The mkdir operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mkdir", "remote"),
		},
		"mktemp": {
			Description: "The mktemp operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mktemp", "remote"),
		},
		"mv": {
			Description: "The mv operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mv", "remote"),
		},
		"path": props.String("The path to the installed binary."),
		"rm": {
			Description: "The rm operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Rm", "remote"),
		},
		"url": props.String("The url used to download the binary."),
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(
		requiredInputs,
		[]string{
			"architecture",
			"directory",
			"download",
			"mkdir",
			"mktemp",
			"mv",
			"path",
			"rm",
			"url",
			"version",
		},
	)

	return schema.PackageSpec{
		Types:     map[string]schema.ComplexTypeSpec{},
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{
			name(typ): {
				IsComponent: true,
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: description,
					Properties:  outputs,
					Required:    requiredOutputs,
				},
				InputProperties: inputs,
				RequiredInputs:  requiredInputs,
			},
		},
	}
}
