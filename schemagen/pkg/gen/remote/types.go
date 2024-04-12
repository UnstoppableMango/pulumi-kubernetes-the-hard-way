package remote

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func GenerateTypes() map[string]schema.ComplexTypeSpec {
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
					"restartSec": {
						Description: "Configures the time to sleep before restarting a service (as configured with Restart=).",
						TypeSpec:    types.String,
					},
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
					"bindsTo": {
						Description: "Configures requirement dependencies, very similar in style to Requires=.",
						TypeSpec:    types.ArrayOfStrings,
					},
					"description": {
						Description: "A short human readable title of the unit.",
						TypeSpec:    types.String,
					},
					"documentation": {
						Description: "A space-separated list of URIs referencing documentation for this unit or its configuration.",
						TypeSpec:    types.ArrayOfStrings,
					},
					"requires": {
						Description: "Similar to Wants=, but declares a stronger requirement dependency.",
						TypeSpec:    types.ArrayOfStrings,
					},
					"requisite": {
						Description: "Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately.",
						TypeSpec:    types.ArrayOfStrings,
					},
					"wants": {
						Description: "Configures (weak) requirement dependencies on other units.",
						TypeSpec:    types.ArrayOfStrings,
					},
				},
			},
		},
	}
}
