package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateProvisionEtcd(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"architecture": {
			Description: "TODO",
			TypeSpec:    types.LocalType("Architecture", "remote"),
		},
		"binaryDirectory": props.String("TODO"),
		"bundle": {
			Description: "The TLS bundle.",
			TypeSpec:    types.LocalType("Bundle", "tls"),
		},
		"connection":             props.Connection(commandSpec),
		"internalIp":             props.String("The internal IP of the etcd node"),
		"configurationDirectory": props.String("The directory to use for etcd configuration."),
		"dataDirectory":          props.String("The directory to use for etcd data."),
		"version":                props.String("The version to install."),
	}
	requiredInputs := []string{"bundle", "connection", "internalIp"}

	outputs := map[string]schema.PropertySpec{
		"configuration": {
			Description: "Etcd configuration.",
			TypeSpec:    types.LocalResource("EtcdConfiguration", "remote"),
		},
		"install": {
			Description: "Install etcd.",
			TypeSpec:    types.LocalResource("EtcdInstall", "remote"),
		},
		"service": {
			Description: "Systemd service.",
			TypeSpec:    types.LocalResource("SystemdService", "remote"),
		},
		"start": {
			Description: "Start etcd",
			TypeSpec:    types.LocalResource("StartEtcd", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(
		requiredInputs,
		[]string{
			"configuration",
			"install",
			"service",
			"start",
		},
	)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Starts etcd on a remote system.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
