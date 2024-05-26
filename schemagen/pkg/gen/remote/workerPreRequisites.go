package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateWorkerPreRequisites(commandSpec schema.PackageSpec) schema.PackageSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": props.Connection(commandSpec),
		"triggers": {
			Description: "Trigger recheck on changes to this input.",
			TypeSpec: schema.TypeSpec{
				Type: "array",
				Items: &schema.TypeSpec{
					Ref: "pulumi.json#/Any",
				},
			},
		},
	}

	requiredInputs := []string{
		"connection",
	}

	outputs := map[string]schema.PropertySpec{
		"conntrack": {
			Description: "Verifies that the conntrack binary exists.",
			TypeSpec:    types.ExtResource(commandSpec, "Command", "remote"),
		},
		"ipset": {
			Description: "Verifies that the ipset binary exists.",
			TypeSpec:    types.ExtResource(commandSpec, "Command", "remote"),
		},
		"socat": {
			Description: "Verifies that the socat binary exists.",
			TypeSpec:    types.ExtResource(commandSpec, "Command", "remote"),
		},
		"swap": {
			Description: "Verifies that swap is disabled.",
			TypeSpec:    types.ExtResource(commandSpec, "Command", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(requiredInputs, []string{
		"conntrack",
		"ipset",
		"socat",
		"swap",
	})

	return schema.PackageSpec{
		Types:     map[string]schema.ComplexTypeSpec{},
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{
			name("WorkerPreRequisites"): {
				IsComponent: true,
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "Verifies that all worker node pre-requisites have been met.",
					Properties:  outputs,
					Required:    requiredOutputs,
				},
				InputProperties: inputs,
				RequiredInputs:  requiredInputs,
			},
		},
	}
}
