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
		name("SystemdInstallSection"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#%5BInstall%5D%20Section%20Options",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"wantedBy": {
						Description: "A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.",
						TypeSpec:    types.ArrayOfStrings,
					},
				},
			},
		},
		name("SystemdServiceExitType"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "Systemd service exit type.",
				Type:        "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "main"},
				{Value: "cgroup"},
			},
		},
		name("SystemdServiceRestart"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "Systemd service restart behavior.",
				Type:        "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "no"},
				{Value: "on-success"},
				{Value: "on-failure"},
				{Value: "on-abnormal"},
				{Value: "on-watchdog"},
				{Value: "on-abort"},
				{Value: "always"},
			},
		},
		name("SystemdServiceType"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "Systemd service type.",
				Type:        "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "simple"},
				{Value: "exec"},
				{Value: "forking"},
				{Value: "oneshot"},
				{Value: "dbus"},
				{Value: "notify"},
				{Value: "notify-reload"},
				{Value: "idle"},
			},
		},
		name("SystemdServiceSection"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html#",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"execStart": {
						Description: "Commands that are executed when this service is started.",
						TypeSpec:    types.String, // This can technically be an array when type is oneshot
					},
					"exitType": {
						Description: "Specifies when the manager should consider the service to be finished.",
						TypeSpec:    types.LocalType("SystemdServiceExitType", "remote"),
					},
					"restart": {
						Description: "Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.",
						TypeSpec:    types.LocalType("SystemdServiceRestart", "remote"),
					},
					"restartSec": props.String("Configures the time to sleep before restarting a service (as configured with Restart=)."),
					"type": {
						Description: "Configures the mechanism via which the service notifies the manager that the service start-up has finished.",
						TypeSpec:    types.LocalType("SystemdServiceType", "remote"),
					},
				},
			},
		},
		name("SystemdUnitSection"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"after":         props.ArrayOf("string", "Those two settings configure ordering dependencies between units."),
					"before":        props.ArrayOf("string", "Those two settings configure ordering dependencies between units."),
					"bindsTo":       props.ArrayOf("string", "Configures requirement dependencies, very similar in style to Requires=."),
					"description":   props.String("A short human readable title of the unit."),
					"documentation": props.ArrayOf("string", "A space-separated list of URIs referencing documentation for this unit or its configuration."),
					"requires":      props.ArrayOf("string", "Similar to Wants=, but declares a stronger requirement dependency."),
					"requisite":     props.ArrayOf("string", "Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately."),
					"wants":         props.ArrayOf("string", "Configures (weak) requirement dependencies on other units."),
				},
			},
		},
	}
}
