package remote

import (
	"fmt"
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateArchiveInstall(
	commandSpec, commandxSpec schema.PackageSpec,
	typ, description string,
	files ...string,
) schema.PackageSpec {
	// Hacky... but better than refactoring this whole thing AGAIN
	var defaultDir string
	if typ == "CniPluginsInstall" {
		defaultDir = "/opt/cni/bin"
	} else {
		defaultDir = "/usr/local/bin"
	}

	inputs := map[string]schema.PropertySpec{
		"architecture": {
			Description: "The CPU architecture to install.",
			TypeSpec:    types.LocalType("Architecture", "remote"),
		},
		"connection": props.Connection(commandSpec),
		"directory": {
			Description: "The directory to install the binary to.",
			TypeSpec:    types.String,
			Default:     defaultDir,
		},
		"version": props.String("The version to install."),
	}

	requiredInputs := []string{"connection"}

	outputs := map[string]schema.PropertySpec{
		"archiveName": props.String("The name of the downloaded archive."),
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
		"path": props.String("The path to the installed binary."),
		"rm": {
			Description: "The rm operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Rm", "remote"),
		},
		"tar": {
			Description: "The tar operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Tar", "remote"),
		},
		"url": props.String("The url used to download the binary."),
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := []string{
		"architecture",
		"archiveName",
		"connection",
		"download",
		"directory",
		"mkdir",
		"mktemp",
		"rm",
		"tar",
		"url",
		"version",
	}

	for _, f := range files {
		mvProp := f + "Mv"
		outputs[mvProp] = schema.PropertySpec{
			Description: fmt.Sprintf("The %s mv operation.", f),
			TypeSpec:    types.ExtResource(commandxSpec, "Mv", "remote"),
		}
		pathProp := f + "Path"
		outputs[pathProp] = schema.PropertySpec{
			Description: fmt.Sprintf("The %s path on the remote system", f),
			TypeSpec:    types.String,
		}
		requiredOutputs = append(requiredOutputs, mvProp, pathProp)
	}

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
