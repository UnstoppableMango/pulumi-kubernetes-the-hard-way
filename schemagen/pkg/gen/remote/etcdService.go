package remote

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateEtcdService(commandSpec schema.PackageSpec) schema.ResourceSpec {
	peerItems := types.LocalType("EtcdConfigurationProps", "remote")

	inputs := map[string]schema.PropertySpec{
		"configuration": {
			Description: "Etcd configuration.",
			TypeSpec:    types.LocalType("EtcdConfigurationProps", "remote"),
		},
		"connection":    connection(commandSpec),
		"description":   props.String("Optional systemd unit description."),
		"directory":     props.String("The location to create the service file."),
		"documentation": props.String("Optional systemd unit documentation"),
		"peers": {
			Description: "Etcd peer configuration.",
			TypeSpec: schema.TypeSpec{
				Type:  "array",
				Items: &peerItems,
			},
		},
		"restart": {
			Description: "Optionally override the systemd service restart behaviour. Defaults to `on-failure`.",
			TypeSpec:    types.LocalType("SystemdServiceRestart", "remote"),
		},
		"restartSec": props.String("Optionally override the systemd service RestartSec. Defaults to `5`."),
		"wantedBy":   props.String("Optionally override the systemd service wanted-by. Defaults to `multi-user.target`."),
	}

	requiredInputs := []string{"configuration", "connection"}

	outputs := map[string]schema.PropertySpec{
		"service": {
			Description: "The remote systemd service.",
			TypeSpec:    types.LocalResource("SystemdService", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Etcd systemd service file. Will likely get replaced with a static function when https://github.com/pulumi/pulumi/issues/7583 gets resolved.",
			Properties:  outputs,
			Required:    append(requiredInputs, "peers", "service"),
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
