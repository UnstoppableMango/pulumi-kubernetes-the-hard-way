package remote

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateEtcdCluster(commandSpec schema.PackageSpec) schema.ResourceSpec {
	etcdNode := types.LocalType("EtcdNode", "remote")

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
		"configurationDirectory": props.String("The directory to use for etcd configuration."),
		"dataDirectory":          props.String("The directory to use for etcd data."),
		"nodes": {
			Description: "Etcd node configuration. The key should be a name used to identify the node.",
			TypeSpec: schema.TypeSpec{
				Type:                 "object",
				AdditionalProperties: &etcdNode,
				Plain:                true,
			},
		},
		"version": props.String("The version to install."),
	}

	requiredInputs := []string{
		"bundle",
		"nodes",
	}

	etcdInstall := types.LocalResource("EtcdInstall", "remote")
	etcdConfiguration := types.LocalResource("EtcdConfiguration", "remote")
	etcdService := types.LocalResource("EtcdService", "remote")
	startEtcd := types.LocalResource("StartEtcd", "remote")

	outputs := map[string]schema.PropertySpec{
		"configuration": {
			Description: "Map of node name to etcd configuration.",
			TypeSpec: schema.TypeSpec{
				Type:                 "object",
				AdditionalProperties: &etcdConfiguration,
			},
		},
		"install": {
			Description: "Map of node name to etcd install.",
			TypeSpec: schema.TypeSpec{
				Type:                 "object",
				AdditionalProperties: &etcdInstall,
			},
		},
		"service": {
			Description: "Map of node name to etcd systemd service.",
			TypeSpec: schema.TypeSpec{
				Type:                 "object",
				AdditionalProperties: &etcdService,
			},
		},
		"start": {
			Description: "Map of node name to etcd start commands.",
			TypeSpec: schema.TypeSpec{
				Type:                 "object",
				AdditionalProperties: &startEtcd,
			},
		},
	}
	maps.Copy(outputs, inputs)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Creates an etcd cluster from one or more remote systems.",
			Properties:  outputs,
			Required: append(requiredInputs,
				"configuration",
				"install",
				"service",
				"start",
			),
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
