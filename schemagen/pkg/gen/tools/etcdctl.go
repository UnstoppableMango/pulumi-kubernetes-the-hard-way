package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateEtcdctl() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"caCert": props.String("TODO"),
		"cert":   props.String("TODO"),
		"commands": {
			Description: "TODO",
			TypeSpec:    types.LocalType("EtcdctlCommand", "tools"),
		},
		"endpoints": props.String("TODO"),
		"key":       props.String("TODO"),
	}

	required := []string{
		"caCert",
		"cert",
		"commands",
		"endpoints",
		"key",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `etcdctl` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required:    required,
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}
