package remote

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateSystemdService(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"directory": {
			Description: "The location to create the service file.",
			TypeSpec:    types.String,
			Default:     "/etc/systemd/system",
		},
		"install": {
			Description: "Describes the [Install] section of a systemd service file.",
			TypeSpec:    types.LocalType("SystemdInstallSection", "remote"),
		},
		"service": {
			Description: "Describes the [Service] section of a systemd service file.",
			TypeSpec:    types.LocalType("SystemdServiceSection", "remote"),
		},
		"unit": {
			Description: "Describes the [Unit] section of a systemd service file.",
			TypeSpec:    types.LocalType("SystemdUnitSection", "remote"),
		},
	}

	requiredInputs := []string{
		"connection",
		"service",
	}

	outputs := map[string]schema.PropertySpec{
		"file": {
			Description: "The service file on the remote machine.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := []string{
		"connection",
		"directory",
		"file",
		"service",
	}

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "A systemd service on a remote system.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
