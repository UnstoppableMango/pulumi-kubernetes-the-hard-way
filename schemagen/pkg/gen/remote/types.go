package remote

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateTypes() map[string]schema.ComplexTypeSpec {
	return map[string]schema.ComplexTypeSpec{
		name("Architecture"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "CPU architecture",
				Type:        "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "amd64"},
				{Value: "arm64"},
			},
		},
		name("CniBridgeIpam"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The CNI plugins IPAM",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"type": props.String("CNI bridge IPAM type"),
					"ranges": {
						Description: "IPAM ranges.",
						TypeSpec: schema.TypeSpec{
							Type: "array",
							Items: &schema.TypeSpec{
								Type:                 "object",
								AdditionalProperties: &types.String,
							},
						},
					},
					"routes": {
						Description: "IPAM routes.",
						TypeSpec: schema.TypeSpec{
							Type: "array",
							Items: &schema.TypeSpec{
								Type:                 "object",
								AdditionalProperties: &types.String,
							},
						},
					},
				},
			},
		},
		name("ContainerdCriPluginConfiguration"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "containerd cri plugin configuration.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"containerd": {
						Description: "containerd configuration.",
						TypeSpec: schema.TypeSpec{
							Plain: true,
							Ref:   types.LocalTypeRef("ContainerdCriPluginConfigurationContainerd", "remote"),
						},
					},
					"cni": {
						Description: "cni configuration.",
						TypeSpec: schema.TypeSpec{
							Plain: true,
							Ref:   types.LocalTypeRef("ContainerdCriPluginConfigurationCni", "remote"),
						},
					},
				},
				Required: []string{"containerd", "cni"},
			},
		},
		name("ContainerdCriPluginConfigurationContainerd"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "containerd cri plugin configuration.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"defaultRuntimeName": props.String("default_runtime_name"),
					"snapshotter":        props.String("snapshotter"),
					"runtimes": { // TODO: This doesn't correspond 1:1 with the TOML file
						Description: "The containerd runtime configuration.",
						TypeSpec: schema.TypeSpec{
							Plain: true,
							Ref:   types.LocalTypeRef("ContainerdCriPluginConfigurationContainerdRunc", "remote"),
						},
					},
				},
			},
		},
		name("ContainerdCriPluginConfigurationContainerdRunc"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "containerd cri runc plugin configuration.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"runtimeType": props.String("runtime_type"),
					"options": {
						Description: "runc options.",
						TypeSpec: schema.TypeSpec{
							Plain: true,
							Ref:   types.LocalTypeRef("ContainerdCriPluginConfigurationContainerdRuncOptions", "remote"),
						},
					},
				},
				Required: []string{"options"},
			},
		},
		name("ContainerdCriPluginConfigurationContainerdRuncOptions"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "containerd cri runc plugin configuration.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"systemdCgroup": props.Boolean("SystemdCgroup"),
				},
			},
		},
		name("ContainerdCriPluginConfigurationCni"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "containerd cri plugin configuration.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"binDir":  props.String("bin_dir"),
					"confDir": props.String("conf_dir"),
				},
			},
		},
	}
}
