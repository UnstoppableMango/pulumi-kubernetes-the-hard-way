package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateGetKubeconfig() schema.PackageSpec {
	return schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{
			name("Cluster"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"certificateAuthorityData": props.String("TODO"),
						"server":                   props.String("TODO"),
					},
					Required: []string{
						"certificateAuthorityData",
						"server",
					},
				},
			},
			name("Context"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"cluster": props.String("TODO"),
						"user":    props.String("TODO"),
					},
					Required: []string{"cluster", "user"},
				},
			},
			name("Kubeconfig"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"clusters": {
							TypeSpec: schema.TypeSpec{
								Type: "array",
								Items: &schema.TypeSpec{
									Ref: types.LocalTypeRef("Cluster", "config"),
								},
							},
						},
						"contexts": {
							TypeSpec: schema.TypeSpec{
								Type: "array",
								Items: &schema.TypeSpec{
									Ref: types.LocalTypeRef("Context", "config"),
								},
							},
						},
						"users": {
							TypeSpec: schema.TypeSpec{
								Type: "array",
								Items: &schema.TypeSpec{
									Ref: types.LocalTypeRef("User", "config"),
								},
							},
						},
					},
					Required: []string{"clusters", "contexts", "users"},
				},
			},
			name("KubeconfigAdminOptions"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"type": {
							TypeSpec: schema.TypeSpec{
								Type:  "string",
								Plain: true,
							},
							Const: "admin",
						},
						"publicIp": props.String("TODO"),
					},
					Required: []string{"type"}, // TODO: Why is public ip not required again?
				},
			},
			name("KubeconfigCluster"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"cluster": {
							TypeSpec: types.LocalType("Cluster", "config"),
						},
						"name": props.String("TODO"),
					},
					Required: []string{"cluster", "name"},
				},
			},
			name("KubeconfigContext"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"context": {
							TypeSpec: types.LocalType("Context", "config"),
						},
						"name": props.String("TODO"),
					},
					Required: []string{"context", "name"},
				},
			},
			name("KubeconfigKubeControllerManagerOptions"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"type": {
							TypeSpec: schema.TypeSpec{
								Type:  "string",
								Plain: true,
							},
							Const: "kube-controller-manager",
						},
						"publicIp": props.String("TODO"),
					},
					Required: []string{"type"}, // TODO: Why is public ip not required again?
				},
			},
			name("KubeconfigKubeProxyOptions"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"type": {
							TypeSpec: schema.TypeSpec{
								Type:  "string",
								Plain: true,
							},
							Const: "kube-proxy",
						},
						"publicIp": props.String("TODO"),
					},
					Required: []string{"type"}, // TODO: Why is public ip not required again?
				},
			},
			name("KubeconfigKubeSchedulerOptions"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"type": {
							TypeSpec: schema.TypeSpec{
								Type:  "string",
								Plain: true,
							},
							Const: "kube-scheduler",
						},
						"publicIp": props.String("TODO"),
					},
					Required: []string{"type"}, // TODO: Why is public ip not required again?
				},
			},
			name("KubeconfigType"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []schema.EnumValueSpec{
					{Value: "worker"},
					{Value: "kube-proxy"},
					{Value: "kube-controller-manager"},
					{Value: "kube-scheduler"},
					{Value: "admin"},
				},
			},
			name("KubeconfigUser"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"name": props.String("TODO"),
						"user": {
							TypeSpec: types.LocalType("User", "config"),
						},
					},
					Required: []string{"name", "user"},
				},
			},
			name("KubeconfigWorkerOptions"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"type": {
							TypeSpec: schema.TypeSpec{
								Type:  "string",
								Plain: true,
							},
							Const: "worker",
						},
						"name":     props.String("TODO"),
						"publicIp": props.String("TODO"),
					},
					Required: []string{"name", "publicIp"}, // TODO: Why is public ip not required again?
				},
			},
			name("User"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"clientCertificateData": props.String("TODO"),
						"clientKeyData":         props.String("TODO"),
					},
					Required: []string{"clientCertificateData", "clientKeyData"},
				},
			},
		},
		Functions: map[string]schema.FunctionSpec{
			name("getKubeconfig"): { // TODO: Need more pems
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
			},
		},
	}
}
