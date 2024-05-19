package config

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateKubeVipManifest(getKubeVipManifest schema.FunctionSpec) schema.ResourceSpec {
	outputs := maps.Clone(getKubeVipManifest.Outputs.Properties)
	outputs["yaml"] = props.String("The yaml representation of the manifest")

	requiredOutputs := append(
		getKubeVipManifest.Outputs.Required,
		"yaml",
	)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Pseudo resource for generating the kube-vip manifest.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: getKubeVipManifest.Inputs.Properties,
		RequiredInputs:  getKubeVipManifest.Inputs.Required,
	}
}
