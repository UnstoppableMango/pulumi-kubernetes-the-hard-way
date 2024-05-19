package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateGetKubeconfig() schema.FunctionSpec {
	return schema.FunctionSpec{ // TODO: Need more pems
		Inputs: &schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"caPem": props.String("Certificate authority data."),
				"options": {
					Description: "Options for creating the kubeconfig.",
					TypeSpec: schema.TypeSpec{
						Plain: true,
						OneOf: []schema.TypeSpec{
							types.LocalType("KubeconfigAdminOptions", "config"),
							types.LocalType("KubeconfigKubeControllerManagerOptions", "config"),
							types.LocalType("KubeconfigKubeProxyOptions", "config"),
							types.LocalType("KubeconfigKubeSchedulerOptions", "config"),
							types.LocalType("KubeconfigWorkerOptions", "config"),
						},
						Discriminator: &schema.DiscriminatorSpec{
							PropertyName: "type",
							Mapping: map[string]string{
								"admin":                   types.LocalTypeRef("KubeconfigAdminOptions", "config"),
								"kube-controller-manager": types.LocalTypeRef("KubeconfigKubeControllerManagerOptions", "config"),
								"kube-proxy":              types.LocalTypeRef("KubeconfigKubeProxyOptions", "config"),
								"kube-scheduler":          types.LocalTypeRef("KubeconfigKubeSchedulerOptions", "config"),
								"worker":                  types.LocalTypeRef("KubeconfigWorkerOptions", "config"),
							},
						},
					},
				},
			},
			Required: []string{"caPem", "options"},
		},
		Outputs: &schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"result": {
					TypeSpec: types.LocalType("Kubeconfig", "config"),
				},
			},
			Required: []string{"result"},
		},
	}
}
