package remote

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateSystemdService(commandSpec schema.PackageSpec) schema.PackageSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"directory": {
			Description: "The location to create the service file.",
			TypeSpec:    types.String,
			Default:     "/etc/systemd/system",
		},
		"install": {
			Description: "Describes the [Install] section of a systemd service file.",
			TypeSpec:    types.LocalType("SystemdInstallSection", "remote"),
		},
		"service": {
			Description: "Describes the [Service] section of a systemd service file.",
			TypeSpec:    types.LocalType("SystemdServiceSection", "remote"),
		},
		"unit": {
			Description: "Describes the [Unit] section of a systemd service file.",
			TypeSpec:    types.LocalType("SystemdUnitSection", "remote"),
		},
		"unitName": props.String("Name of the systemd unit."),
	}

	requiredInputs := []string{
		"connection",
		"service",
	}

	outputs := map[string]schema.PropertySpec{
		"file": {
			Description: "The service file on the remote machine.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := []string{
		"connection",
		"directory",
		"file",
		"service",
	}

	return schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{
			name("SystemdDelegate"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "https://www.man7.org/linux/man-pages/man5/systemd.resource-control.5.html",
					Type:        "string",
				},
				Enum: []schema.EnumValueSpec{
					{Value: "yes"},
					{Value: "no"},
					{Value: "cpu"},
					{Value: "cpuacct"},
					{Value: "cpuset"},
					{Value: "io"},
					{Value: "blkio"},
					{Value: "memory"},
					{Value: "devices"},
					{Value: "pids"},
					{Value: "bpf-firewall"},
					{Value: "bpf-devices"},
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
			name("SystemdKillMode"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "https://www.freedesktop.org/software/systemd/man/latest/systemd.kill.html#Description",
					Type:        "string",
				},
				Enum: []schema.EnumValueSpec{
					{Value: "control-group"},
					{Value: "mixed"},
					{Value: "process"},
					{Value: "none"},
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
						"delegate": {
							Description: "Turns on delegation of further resource control partitioning to processes of the unit.",
							TypeSpec:    types.LocalType("SystemdDelegate", "remote"),
						},
						"execStart":    props.String("Commands that are executed when this service is started."), // This can technically be an array when type is oneshot
						"execStartPre": props.String("Additional commands that are executed before the command in ExecStart=."),
						"exitType": {
							Description: "Specifies when the manager should consider the service to be finished.",
							TypeSpec:    types.LocalType("SystemdServiceExitType", "remote"),
						},
						"killMode": {
							Description: "Specifies how processes of this unit shall be killed.",
							TypeSpec:    types.LocalType("SystemdKillMode", "remote"),
						},
						"limitNoFile":    props.Integer("https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties"),
						"limitNProc":     props.String("https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties"),
						"limitCore":      props.String("https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties"),
						"oomScoreAdjust": props.Integer("https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#OOMScoreAdjust="),
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
		},
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{
			name("SystemdService"): {
				IsComponent: true,
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "A systemd service on a remote system.",
					Properties:  outputs,
					Required:    requiredOutputs,
				},
				InputProperties: inputs,
				RequiredInputs:  requiredInputs,
			},
		},
	}
}
