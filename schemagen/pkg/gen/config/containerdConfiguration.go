package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateContainerdConfiguration() schema.PackageSpec {
	function, resource := generatePseudoFunction(
		"Get the containerd configuration.",
		schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"cri": {
					Description: "The cri configuration.",
					TypeSpec: schema.TypeSpec{
						Plain: true,
						Ref:   types.LocalTypeRef("ContainerdCriPluginConfiguration", "config"),
					},
				},
			},
			Required: []string{},
		},
		schema.PropertySpec{
			TypeSpec: types.LocalType("ContainerdConfiguration", "config"),
		},
		"toml", props.String("The toml representation of the configuration."),
	)

	return schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{
			name("ContainerdCriPluginConfiguration"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "containerd cri plugin configuration.",
					Type:        "object",
					Properties: map[string]schema.PropertySpec{
						"containerd": {
							Description: "containerd configuration.",
							TypeSpec: schema.TypeSpec{
								Plain: true,
								Ref:   types.LocalTypeRef("ContainerdCriPluginConfigurationContainerd", "config"),
							},
						},
						"cni": {
							Description: "cni configuration.",
							TypeSpec: schema.TypeSpec{
								Plain: true,
								Ref:   types.LocalTypeRef("ContainerdCriPluginConfigurationCni", "config"),
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
								Ref:   types.LocalTypeRef("ContainerdCriPluginConfigurationContainerdRunc", "config"),
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
								Ref:   types.LocalTypeRef("ContainerdCriPluginConfigurationContainerdRuncOptions", "config"),
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
			name("ContainerdConfiguration"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "TODO",
					Type:        "object",
					Properties:  map[string]schema.PropertySpec{},
					Required:    []string{},
				},
			},
		},
		Functions: map[string]schema.FunctionSpec{
			name("getContainerdConfiguration"): function,
		},
		Resources: map[string]schema.ResourceSpec{
			name("ContainerdConfiguration"): resource,
		},
	}
}
