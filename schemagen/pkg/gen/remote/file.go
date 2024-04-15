package remote

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateFile(commandSpec schema.PackageSpec) schema.ResourceSpec {
	command := commandSpec.Resources["command:remote:Command"]

	inputs := map[string]schema.PropertySpec{
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"content": {
			Description: "The content of the file.",
			TypeSpec:    types.String,
		},
		"path": {
			Description: "The path to the file on the remote host.",
			TypeSpec:    types.String,
		},
	}

	requiredInputs := []string{
		"connection",
		"content",
		"path",
	}

	outputs := map[string]schema.PropertySpec{
		"command": {
			Description: "The executed command.",
			TypeSpec:    types.ExtResource(commandSpec, "Command", "remote"),
		},
		"stderr": command.Properties["stderr"],
		"stdin":  command.Properties["stdin"],
		"stdout": command.Properties["stdout"],
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := []string{
		"command",
		"connection",
		"content",
		"path",
		"stderr",
		"stdin",
		"stdout",
	}

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
