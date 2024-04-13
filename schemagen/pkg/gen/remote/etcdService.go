package remote

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateEtcdService(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"connection":    connection(commandSpec),
		"description":   props.String("Optional systemd unit description."),
		"documentation": props.String("Optional systemd unit documentation"),
		"restart": {
			Description: "Optionally override the systemd service restart behaviour. Defaults to `on-failure`.",
			TypeSpec:    types.LocalType("SystemdServiceRestart"),
		},
		"restartSec": {
			Description: "Optionally override the systemd service RestartSec. Defaults to `5`.",
		},
		"wantedBy": props.String("Optionally override the systemd service wanted-by. Defaults to `multi-user.target`."),
	}

	requiredInputs := []string{"connection"}

	outputs := map[string]schema.PropertySpec{
		"ip": {
			Description: "The IP address of the node.",
			TypeSpec:    types.String,
		},
		"service": {
			Description: "The remote systemd service.",
			TypeSpec:    types.LocalResource("SystemdService", "remote"),
		},
	}

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Etcd systemd service file. Will likely get replaced with a static function when https://github.com/pulumi/pulumi/issues/7583 gets resolved.",
			Properties:  outputs,
			Required:    append(requiredInputs, "ip", "service"),
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
