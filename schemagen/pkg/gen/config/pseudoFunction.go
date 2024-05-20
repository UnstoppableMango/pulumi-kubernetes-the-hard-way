package config

import (
	"maps"
	"slices"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

// Generates "pseudo functions" to work around https://github.com/pulumi/pulumi/issues/7583

func generatePseudoFunction(
	description string,
	spec schema.ObjectTypeSpec,
	result schema.PropertySpec,
	extraOutputs ...interface{},
) (schema.FunctionSpec, schema.ResourceSpec) {
	function := schema.FunctionSpec{
		Description: description,
		Inputs:      &spec,
		Outputs: &schema.ObjectTypeSpec{
			Description: description,
			Properties: map[string]schema.PropertySpec{
				"result": result,
			},
			Required: []string{"result"},
		},
	}

	outputs := *function.Outputs
	outputs.Properties = maps.Clone(outputs.Properties)
	outputs.Required = slices.Clone(outputs.Required)

	var key string
	for i, x := range extraOutputs {
		if i%2 == 0 {
			parsed, ok := x.(string)
			if !ok {
				panic("Invalid input")
			}

			key = parsed
		} else {
			parsed, ok := x.(schema.PropertySpec)
			if !ok {
				panic("Invalid input")
			}

			outputs.Properties[key] = parsed
			outputs.Required = append(outputs.Required, key)
		}
	}

	resource := schema.ResourceSpec{
		IsComponent:     true,
		ObjectTypeSpec:  outputs,
		InputProperties: function.Inputs.Properties,
		RequiredInputs:  function.Inputs.Required,
	}

	return function, resource
}
