package config

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateKubeletConfiguration(getKubeletConfiguration schema.FunctionSpec) schema.ResourceSpec {
	outputs := maps.Clone(getKubeletConfiguration.Outputs.Properties)
	outputs["yaml"] = props.String("The yaml representation of the manifest")

	requiredOutputs := append(
		getKubeletConfiguration.Outputs.Required,
		"yaml",
	)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Properties: outputs,
			Required:   requiredOutputs,
		},
		InputProperties: getKubeletConfiguration.Inputs.Properties,
		RequiredInputs:  getKubeletConfiguration.Inputs.Required,
	}
}
