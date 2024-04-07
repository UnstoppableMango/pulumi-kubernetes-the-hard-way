package gen

import "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

const configMod = "kubernetes-the-hard-way:config:"

func generateConfig() schema.PackageSpec {
	types := map[string]schema.ComplexTypeSpec{
		configMod + "Cluster": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"certificateAuthorityData": {TypeSpec: typeSpecs.String},
					"server":                   {TypeSpec: typeSpecs.String},
				},
				Required: []string{"certificateAuthorityData", "server"},
			},
		},
		configMod + "Context": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"cluster": {TypeSpec: typeSpecs.String},
					"user":    {TypeSpec: typeSpecs.String},
				},
				Required: []string{"cluster", "user"},
			},
		},
		configMod + "Kubeconfig": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"clusters": {
						TypeSpec: schema.TypeSpec{
							Type: "array",
							AdditionalProperties: &schema.TypeSpec{
								Ref: localResource("Cluster", "config"),
							},
						},
					},
					"contexts": {
						TypeSpec: schema.TypeSpec{
							Type: "array",
							AdditionalProperties: &schema.TypeSpec{
								Ref: localResource("Context", "config"),
							},
						},
					},
					"users": {
						TypeSpec: schema.TypeSpec{
							Type: "array",
							AdditionalProperties: &schema.TypeSpec{
								Ref: localResource("User", "config"),
							},
						},
					},
				},
				Required: []string{"clusters", "contexts", "users"},
			},
		},
		configMod + "KubeconfigAdminOptions": {
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
					"publicIp": {TypeSpec: typeSpecs.String},
				},
				Required: []string{"type"}, // TODO: Why is public ip not required again?
			},
		},
		configMod + "KubeconfigCluster": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"cluster": {
						TypeSpec: schema.TypeSpec{Ref: localType("Cluster", "config")},
					},
					"name": {TypeSpec: typeSpecs.String},
				},
				Required: []string{"cluster", "name"},
			},
		},
		configMod + "KubeconfigContext": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"context": {
						TypeSpec: schema.TypeSpec{Ref: localType("Context", "config")},
					},
					"name": {TypeSpec: typeSpecs.String},
				},
				Required: []string{"context", "name"},
			},
		},
		configMod + "KubeconfigKubeControllerManagerOptions": {
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
					"publicIp": {TypeSpec: typeSpecs.String},
				},
				Required: []string{"type"}, // TODO: Why is public ip not required again?
			},
		},
		configMod + "KubeconfigKubeProxyOptions": {
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
					"publicIp": {TypeSpec: typeSpecs.String},
				},
				Required: []string{"type"}, // TODO: Why is public ip not required again?
			},
		},
		configMod + "KubeconfigKubeSchedulerOptions": {
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
					"publicIp": {TypeSpec: typeSpecs.String},
				},
				Required: []string{"type"}, // TODO: Why is public ip not required again?
			},
		},
		configMod + "KubeconfigType": {
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
		configMod + "KubeconfigUser": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"name": {TypeSpec: typeSpecs.String},
					"user": {
						TypeSpec: schema.TypeSpec{Ref: localType("User", "config")},
					},
				},
				Required: []string{"name", "user"},
			},
		},
		configMod + "KubeconfigWorkerOptions": {
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
					"name":     {TypeSpec: typeSpecs.String},
					"publicIp": {TypeSpec: typeSpecs.String},
				},
				Required: []string{"name", "publicIp"}, // TODO: Why is public ip not required again?
			},
		},
		configMod + "User": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"clientCertificateData": {TypeSpec: typeSpecs.String},
					"clientKeyData":         {TypeSpec: typeSpecs.String},
				},
				Required: []string{"clientCertificateData", "clientKeyData"},
			},
		},
	}

	return schema.PackageSpec{
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{},
		Types:     types,
	}
}
