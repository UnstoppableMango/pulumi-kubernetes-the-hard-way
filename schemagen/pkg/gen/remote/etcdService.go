package remote

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateEtcdService(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": connection(commandSpec),
	}

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
			Properties: outputs,
		},
		InputProperties: inputs,
	}
}
