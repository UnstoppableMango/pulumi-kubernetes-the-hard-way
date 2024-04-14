package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateProvisionEtcd(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"configuration": {
			Description: "Etcd configuration.",
			TypeSpec:    types.LocalResource("EtcdConfiguration", "remote"),
		},
		"connection": connection(commandSpec),
	}
	requiredInputs := []string{"connection"}

	outputs := map[string]schema.PropertySpec{
		"configuration": {
			Description: "Etcd configuration.",
			TypeSpec:    types.LocalResource("EtcdConfiguration", "remote"),
		},
		"install": {
			Description: "Install etcd.",
			TypeSpec:    types.LocalResource("EtcdInstall", "remote"),
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
