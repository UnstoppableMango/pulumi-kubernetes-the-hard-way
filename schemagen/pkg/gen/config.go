package gen

import "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

const configMod = "kubernetes-the-hard-way:config:"

func generateConfig(kubernetesSpec schema.PackageSpec) schema.PackageSpec {
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
							Items: &schema.TypeSpec{
								Ref: localType("Cluster", "config"),
							},
						},
					},
					"contexts": {
						TypeSpec: schema.TypeSpec{
							Type: "array",
							Items: &schema.TypeSpec{
								Ref: localType("Context", "config"),
							},
						},
					},
					"users": {
						TypeSpec: schema.TypeSpec{
							Type: "array",
							Items: &schema.TypeSpec{
								Ref: localType("User", "config"),
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
		configMod + "PodManifest": generatePodManifest(kubernetesSpec),
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

	functions := map[string]schema.FunctionSpec{
		configMod + "getKubeconfig": { // TODO: Need more pems
			Inputs: &schema.ObjectTypeSpec{
				Properties: map[string]schema.PropertySpec{
					"caPem": {
						Description: "Certificate authority data.",
						TypeSpec:    typeSpecs.String,
					},
					"options": {
						Description: "Options for creating the kubeconfig.",
						TypeSpec: schema.TypeSpec{
							Plain: true,
							OneOf: []schema.TypeSpec{
								{Ref: localType("KubeconfigAdminOptions", "config")},
								{Ref: localType("KubeconfigKubeControllerManagerOptions", "config")},
								{Ref: localType("KubeconfigKubeProxyOptions", "config")},
								{Ref: localType("KubeconfigKubeSchedulerOptions", "config")},
								{Ref: localType("KubeconfigWorkerOptions", "config")},
							},
							Discriminator: &schema.DiscriminatorSpec{
								PropertyName: "type",
								Mapping: map[string]string{
									"admin":                   localType("KubeconfigAdminOptions", "config"),
									"kube-controller-manager": localType("KubeconfigKubeControllerManagerOptions", "config"),
									"kube-proxy":              localType("KubeconfigKubeProxyOptions", "config"),
									"kube-scheduler":          localType("KubeconfigKubeSchedulerOptions", "config"),
									"worker":                  localType("KubeconfigWorkerOptions", "config"),
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
						TypeSpec: schema.TypeSpec{Ref: localType("Kubeconfig", "config")},
					},
				},
				Required: []string{"result"},
			},
		},
		configMod + "getKubeVipManifest": {
			Description: "Gets the static pod manifests for KubeVip.",
			Inputs: &schema.ObjectTypeSpec{
				Properties: map[string]schema.PropertySpec{
					"address": {
						Description: "TODO",
						TypeSpec:    typeSpecs.String,
					},
					"cpEnable": {
						Description: "TODO",
						TypeSpec:    typeSpecs.Boolean,
					},
					"cpNamespace": {
						Description: "TODO",
						TypeSpec:    typeSpecs.String,
					},
					"image": {
						Description: "Override the kube-vip image.",
						TypeSpec:    typeSpecs.String,
					},
					"kubeconfigPath": {
						Description: "Path to the kubeconfig on the remote host.",
						TypeSpec:    typeSpecs.String,
					},
					"port": {
						Description: "TODO",
						TypeSpec:    typeSpecs.Integer,
						Default:     6443,
					},
					"svcEnable": {
						Description: "TODO",
						TypeSpec:    typeSpecs.Boolean,
					},
					"version": {
						Description: "Version of kube-vip to use.",
						TypeSpec:    typeSpecs.String,
					},
					"vipArp": {
						Description: "TODO",
						TypeSpec:    typeSpecs.Boolean,
					},
					"vipCidr": {
						Description: "TODO",
						TypeSpec:    typeSpecs.Integer,
					},
					"vipDdns": {
						Description: "TODO",
						TypeSpec:    typeSpecs.Boolean,
					},
					"vipInterface": {
						Description: "TODO",
						TypeSpec:    typeSpecs.String,
					},
					"vipLeaderElection": {
						Description: "TODO",
						TypeSpec:    typeSpecs.Boolean,
					},
					"vipLeaseDuration": {
						Description: "TODO",
						TypeSpec:    typeSpecs.Integer,
					},
					"vipRenewDeadline": {
						Description: "TODO",
						TypeSpec:    typeSpecs.Integer,
					},
					"vipRetryPeriod": {
						Description: "TODO",
						TypeSpec:    typeSpecs.Integer,
					},
				},
				Required: []string{"address", "kubeconfigPath", "vipCidr"},
			},
			Outputs: &schema.ObjectTypeSpec{
				Properties: map[string]schema.PropertySpec{
					"result": {
						TypeSpec: schema.TypeSpec{
							Ref: localType("PodManifest", "config"),
						},
					},
				},
				Required: []string{"result"},
			},
		},
	}

	return schema.PackageSpec{
		Functions: functions,
		Resources: map[string]schema.ResourceSpec{},
		Types:     types,
	}
}

func generatePodManifest(kubernetesSpec schema.PackageSpec) schema.ComplexTypeSpec {
	pod := kubernetesSpec.Resources["kubernetes:core/v1:Pod"]

	return schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: pod.Description,
			Type:        "object",
			Properties:  makeExternal(pod.Properties, kubernetesSpec),
		},
	}
}
