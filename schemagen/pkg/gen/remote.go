package gen

import (
	"fmt"
	"maps"
	"slices"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

const remoteMod = "kubernetes-the-hard-way:remote:"

func generateRemote(commandSpec schema.PackageSpec) schema.PackageSpec {
	types := map[string]schema.ComplexTypeSpec{
		remoteMod + "Architecture": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "CPU architecture",
				Type:        "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "amd64"},
				{Value: "arm64"},
			},
		},
		remoteMod + "SystemdInstallSection": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#%5BInstall%5D%20Section%20Options",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"wantedBy": {
						Description: "A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.",
						TypeSpec:    typeSpecs.ArrayOfStrings,
					},
				},
			},
		},
		remoteMod + "SystemdServiceExitType": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "Systemd service exit type.",
				Type:        "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "main"},
				{Value: "cgroup"},
			},
		},
		remoteMod + "SystemdServiceRestart": {
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
		remoteMod + "SystemdServiceType": {
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
		remoteMod + "SystemdServiceSection": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html#",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"execStart": {
						Description: "Commands that are executed when this service is started.",
						TypeSpec:    typeSpecs.String, // This can technically be an array when type is oneshot
					},
					"exitType": {
						Description: "Specifies when the manager should consider the service to be finished.",
						TypeSpec:    schema.TypeSpec{Ref: localType("SystemdServiceExitType", "remote")},
					},
					"restart": {
						Description: "Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.",
						TypeSpec:    schema.TypeSpec{Ref: localType("SystemdServiceRestart", "remote")},
					},
					"restartSec": {
						Description: "Configures the time to sleep before restarting a service (as configured with Restart=).",
						TypeSpec:    typeSpecs.String,
					},
					"type": {
						Description: "Configures the mechanism via which the service notifies the manager that the service start-up has finished.",
						TypeSpec:    schema.TypeSpec{Ref: localType("SystemdServiceType", "remote")},
					},
				},
			},
		},
		remoteMod + "SystemdUnitSection": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"bindsTo": {
						Description: "Configures requirement dependencies, very similar in style to Requires=.",
						TypeSpec:    typeSpecs.ArrayOfStrings,
					},
					"description": {
						Description: "A short human readable title of the unit.",
						TypeSpec:    typeSpecs.String,
					},
					"documentation": {
						Description: "A space-separated list of URIs referencing documentation for this unit or its configuration.",
						TypeSpec:    typeSpecs.ArrayOfStrings,
					},
					"requires": {
						Description: "Similar to Wants=, but declares a stronger requirement dependency.",
						TypeSpec:    typeSpecs.ArrayOfStrings,
					},
					"requisite": {
						Description: "Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately.",
						TypeSpec:    typeSpecs.ArrayOfStrings,
					},
					"wants": {
						Description: "Configures (weak) requirement dependencies on other units.",
						TypeSpec:    typeSpecs.ArrayOfStrings,
					},
				},
			},
		},
	}

	functions := map[string]schema.FunctionSpec{}

	resources := map[string]schema.ResourceSpec{
		remoteMod + "CniPluginsInstall": generateArchiveInstall(commandSpec, "Installs cni-plugins on a remote system",
			"bandwidth", "bridge", "dhcp", "dummy", "firewall", "hostDevice",
			"hostLocal", "ipvlan", "loopback", "macvlan", "portmap", "ptp", "sbr",
			"static", "tap", "tuning", "vlan", "vrf"),
		remoteMod + "ContainerdInstall":            generateArchiveInstall(commandSpec, "Installs containerd on a remote system", "containerd"),
		remoteMod + "CrictlInstall":                generateArchiveInstall(commandSpec, "Installs crictl on a remote system", "crictl"),
		remoteMod + "Download":                     generateDownload(commandSpec),
		remoteMod + "EtcdConfiguration":            generateEtcdConfiguration(commandSpec),
		remoteMod + "EtcdInstall":                  generateArchiveInstall(commandSpec, "Installs etcd on a remote system", "etcd", "etcdctl"),
		remoteMod + "File":                         generateFile(commandSpec),
		remoteMod + "KubeApiServerInstall":         generateBinaryInstall(commandSpec, "Installs kube-apiserver on a remote system."),
		remoteMod + "KubeControllerManagerInstall": generateBinaryInstall(commandSpec, "Installs kube-controller-manager on a remote system."),
		remoteMod + "KubectlInstall":               generateBinaryInstall(commandSpec, "Installs kubectl on a remote system."),
		remoteMod + "KubeletInstall":               generateBinaryInstall(commandSpec, "Installs kubelet on a remote system."),
		remoteMod + "KubeProxyInstall":             generateBinaryInstall(commandSpec, "Installs kube-proxy on a remote system."),
		remoteMod + "KubeSchedulerInstall":         generateBinaryInstall(commandSpec, "Installs kube-scheduler on a remote system."),
		remoteMod + "RuncInstall":                  generateBinaryInstall(commandSpec, "Installs runc on a remote system."),
		remoteMod + "SystemdService":               generateSystemdService(commandSpec),
	}

	return schema.PackageSpec{
		Functions: functions,
		Resources: resources,
		Types:     types,
	}
}

func generateArchiveInstall(
	commandSpec schema.PackageSpec,
	description string,
	files ...string,
) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"architecture": {
			Description: "The CPU architecture to install.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Architecture", "remote")},
		},
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    schema.TypeSpec{Ref: refType(commandSpec, "Connection", "remote")},
		},
		"directory": {
			Description: "The directory to install the binary to.",
			TypeSpec:    typeSpecs.String,
			Default:     "/usr/local/bin",
		},
		"version": {
			Description: "The version to install.",
			TypeSpec:    typeSpecs.String,
		},
	}

	requiredInputs := []string{"connection"}

	outputs := map[string]schema.PropertySpec{
		"archiveName": {
			Description: "The name of the downloaded archive.",
			TypeSpec:    typeSpecs.String,
		},
		"download": {
			Description: "The download operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Download", "remote")},
		},
		"mkdir": {
			Description: "The mkdir operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Mkdir", "tools")},
		},
		"mktemp": {
			Description: "The mktemp operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Mktemp", "tools")},
		},
		"path": {
			Description: "The path to the installed binary.",
			TypeSpec:    typeSpecs.String,
		},
		"rm": {
			Description: "The rm operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Rm", "tools")},
		},
		"tar": {
			Description: "The tar operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Tar", "tools")},
		},
		"url": {
			Description: "The url used to download the binary.",
			TypeSpec:    typeSpecs.String,
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := []string{
		"architecture",
		"archiveName",
		"connection",
		"download",
		"directory",
		"mkdir",
		"mktemp",
		"rm",
		"tar",
		"url",
		"version",
	}

	for _, f := range files {
		mvProp := f + "Mv"
		outputs[mvProp] = schema.PropertySpec{
			Description: fmt.Sprintf("The %s mv operation.", f),
			TypeSpec:    schema.TypeSpec{Ref: localResource("Mv", "tools")},
		}
		pathProp := f + "Path"
		outputs[pathProp] = schema.PropertySpec{
			Description: fmt.Sprintf("The %s path on the remote system", f),
			TypeSpec:    typeSpecs.String,
		}
		requiredOutputs = append(requiredOutputs, mvProp, pathProp)
	}

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: description,
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}

func generateBinaryInstall(
	commandSpec schema.PackageSpec,
	description string,
) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"architecture": {
			Description: "The CPU architecture to install.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Architecture", "remote")},
		},
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    schema.TypeSpec{Ref: refType(commandSpec, "Connection", "remote")},
		},
		"directory": {
			Description: "The directory to install the binary to.",
			TypeSpec:    typeSpecs.String,
			Default:     "/usr/local/bin",
		},
		"version": {
			Description: "The version to install.",
			TypeSpec:    typeSpecs.String,
		},
	}

	requiredInputs := []string{"connection"}

	outputs := map[string]schema.PropertySpec{
		"binName": {
			Description: "The name of the installed binary.",
			TypeSpec:    typeSpecs.String,
		},
		"download": {
			Description: "The download operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Download", "remote")},
		},
		"mkdir": {
			Description: "The mkdir operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Mkdir", "tools")},
		},
		"mktemp": {
			Description: "The mktemp operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Mktemp", "tools")},
		},
		"mv": {
			Description: "The mv operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Mv", "tools")},
		},
		"path": {
			Description: "The path to the installed binary.",
			TypeSpec:    typeSpecs.String,
		},
		"rm": {
			Description: "The rm operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Rm", "tools")},
		},
		"url": {
			Description: "The url used to download the binary.",
			TypeSpec:    typeSpecs.String,
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(
		requiredInputs,
		[]string{
			"architecture",
			"directory",
			"download",
			"mkdir",
			"mktemp",
			"mv",
			"path",
			"rm",
			"url",
			"version",
		},
	)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: description,
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}

func generateDownload(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    schema.TypeSpec{Ref: refType(commandSpec, "Connection", "remote")},
		},
		"destination": {
			Description: "The fully qualified path on the remote system where the file should be downloaded to.",
			TypeSpec:    typeSpecs.String,
		},
		"removeOnDelete": {
			Description: "Remove the downloaded fiel when the resource is deleted.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"url": {
			Description: "The URL of the file to be downloaded.",
			TypeSpec:    typeSpecs.String,
		},
	}

	requiredInputs := []string{
		"connection",
		"destination",
		"url",
	}

	outputs := map[string]schema.PropertySpec{
		"mkdir": {
			Description: "The mkdir operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Mkdir", "tools")},
		},
		"wget": {
			Description: "The wget operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Wget", "tools")},
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := []string{
		"connection",
		"destination",
		"mkdir",
		"removeOnDelete",
		"url",
		"wget",
	}

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Downloads the file specified by `url` onto a remote system.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}

func generateEtcdConfiguration(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"caPem": {
			Description: "The PEM encoded certificate authority data.",
			TypeSpec:    typeSpecs.String,
		},
		"certPem": {
			Description: "The PEM encoded certificate data.",
			TypeSpec:    typeSpecs.String,
		},
		"configurationDirectory": {
			Description: "The directory to store etcd configuration.",
			TypeSpec:    typeSpecs.String,
			Default:     "/etc/etcd",
		},
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    schema.TypeSpec{Ref: refType(commandSpec, "Connection", "remote")},
		},
		"dataDirectory": {
			Description: "The directory etcd will store its data.",
			TypeSpec:    typeSpecs.String,
		},
		"etcdPath": {
			Description: "The path to the `etcd` binary.",
			TypeSpec:    typeSpecs.String,
		},
		"internalIp": {
			Description: "The IP used to serve client requests and communicate with etcd peers.",
			TypeSpec:    typeSpecs.String,
		},
		"keyPem": {
			Description: "The PEM encoded key data.",
			TypeSpec:    typeSpecs.String,
		},
	}

	requiredInputs := []string{
		"caPem",
		"certPem",
		"connection",
		"etcdPath",
		"internalIp",
		"keyPem",
	}

	outputs := map[string]schema.PropertySpec{
		"caFile": {
			Description: "The remote certificate authority file.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("File", "remote")},
		},
		"certFile": {
			Description: "The remote certificate file.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("File", "remote")},
		},
		"configurationMkdir": {
			Description: "The configuration mkdir operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Mkdir", "tools")},
		},
		"dataMkdir": {
			Description: "The data mkdir operation.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("Mkdir", "tools")},
		},
		"keyFile": {
			Description: "The remote key file.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("File", "remote")},
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(
		requiredInputs,
		[]string{
			"caFile",
			"certFile",
			"configurationMkdir",
			"dataMkdir",
			"keyFile",
		},
	)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Configures etcd on a remote system.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}

func generateFile(commandSpec schema.PackageSpec) schema.ResourceSpec {
	command := commandSpec.Resources["command:remote:Command"]

	inputs := map[string]schema.PropertySpec{
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    schema.TypeSpec{Ref: refType(commandSpec, "Connection", "remote")},
		},
		"content": {
			Description: "The content of the file.",
			TypeSpec:    typeSpecs.String,
		},
		"path": {
			Description: "The path to the file on the remote host.",
			TypeSpec:    typeSpecs.String,
		},
	}

	requiredInputs := []string{
		"connection",
		"content",
		"path",
	}

	outputs := map[string]schema.PropertySpec{
		"command": {
			Description: "The executed command.",
			TypeSpec:    schema.TypeSpec{Ref: refResource(commandSpec, "Command", "remote")},
		},
		"stderr": command.Properties["stderr"],
		"stdin":  command.Properties["stdin"],
		"stdout": command.Properties["stdout"],
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := []string{
		"command",
		"connection",
		"content",
		"path",
		"stderr",
		"stdin",
		"stdout",
	}

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}

func generateSystemdService(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    schema.TypeSpec{Ref: refType(commandSpec, "Connection", "remote")},
		},
		"directory": {
			Description: "The location to create the service file.",
			TypeSpec:    typeSpecs.String,
			Default:     "/etc/systemd/system",
		},
		"install": {
			Description: "Describes the [Install] section of a systemd service file.",
			TypeSpec:    schema.TypeSpec{Ref: localType("SystemdInstallSection", "remote")},
		},
		"service": {
			Description: "Describes the [Service] section of a systemd service file.",
			TypeSpec:    schema.TypeSpec{Ref: localType("SystemdServiceSection", "remote")},
		},
		"unit": {
			Description: "Describes the [Unit] section of a systemd service file.",
			TypeSpec:    schema.TypeSpec{Ref: localType("SystemdUnitSection", "remote")},
		},
	}

	requiredInputs := []string{
		"connection",
		"service",
	}

	outputs := map[string]schema.PropertySpec{
		"file": {
			Description: "The service file on the remote machine.",
			TypeSpec:    schema.TypeSpec{Ref: localResource("File", "remote")},
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := []string{
		"connection",
		"directory",
		"file",
		"service",
	}

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "A systemd service on a remote system.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
