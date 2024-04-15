package remote

import (
	"fmt"
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func name(x string) string {
	return fmt.Sprintf("kubernetes-the-hard-way:remote:%s", x)
}

func generateArchiveInstall(
	commandSpec schema.PackageSpec,
	description string,
	files ...string,
) schema.ResourceSpec {
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
		"archiveName": {
			Description: "The name of the downloaded archive.",
			TypeSpec:    types.String,
		},
		"download": {
			Description: "The download operation.",
			TypeSpec:    types.LocalResource("Download", "remote"),
		},
		"mkdir": {
			Description: "The mkdir operation.",
			TypeSpec:    types.LocalResource("Mkdir", "tools"),
		},
		"mktemp": {
			Description: "The mktemp operation.",
			TypeSpec:    types.LocalResource("Mktemp", "tools"),
		},
		"path": {
			Description: "The path to the installed binary.",
			TypeSpec:    types.String,
		},
		"rm": {
			Description: "The rm operation.",
			TypeSpec:    types.LocalResource("Rm", "tools"),
		},
		"tar": {
			Description: "The tar operation.",
			TypeSpec:    types.LocalResource("Tar", "tools"),
		},
		"url": {
			Description: "The url used to download the binary.",
			TypeSpec:    types.String,
		},
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
			TypeSpec:    types.LocalResource("Mv", "tools"),
		}
		pathProp := f + "Path"
		outputs[pathProp] = schema.PropertySpec{
			Description: fmt.Sprintf("The %s path on the remote system", f),
			TypeSpec:    types.String,
		}
		requiredOutputs = append(requiredOutputs, mvProp, pathProp)
	}

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: description,
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}

func generateBinaryInstall(
	commandSpec schema.PackageSpec,
	description string,
) schema.ResourceSpec {
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
			TypeSpec:    types.LocalResource("Mkdir", "tools"),
		},
		"mktemp": {
			Description: "The mktemp operation.",
			TypeSpec:    types.LocalResource("Mktemp", "tools"),
		},
		"mv": {
			Description: "The mv operation.",
			TypeSpec:    types.LocalResource("Mv", "tools"),
		},
		"path": {
			Description: "The path to the installed binary.",
			TypeSpec:    types.String,
		},
		"rm": {
			Description: "The rm operation.",
			TypeSpec:    types.LocalResource("Rm", "tools"),
		},
		"url": {
			Description: "The url used to download the binary.",
			TypeSpec:    types.String,
		},
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

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: description,
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
