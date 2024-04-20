package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateEtcdctl() tool {
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

	required := []string{"commands"}

	typ := schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `etcdctl` utility on a remote system.",
			Properties:  inputs,
			Required:    required,
		},
	}

	return tool{optsType: typ, types: map[string]schema.ComplexTypeSpec{}}
}

// If we ever get a way to add the "required outputs" logic around a complexType
// resource := schema.ResourceSpec{
// 	ObjectTypeSpec: schema.ObjectTypeSpec{
// 		Description: "Abstraction over the `etcdctl` utility on a remote system.",
// 		Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
// 		Required:    required,
// 	},
// 	InputProperties: inputs,
// 	RequiredInputs:  required,
// }
