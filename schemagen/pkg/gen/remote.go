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
	}

	functions := map[string]schema.FunctionSpec{}

	resources := map[string]schema.ResourceSpec{
		remoteMod + "CniPluginsInstall": generateArchiveInstall(commandSpec, "Installs cni-plugins on a remote system",
			"bandwidth", "bridge", "dhcp", "dummy", "firewall", "hostDevice",
			"hostLocal", "ipvlan", "loopback", "macvlan", "portmap", "ptp", "sbr",
			"static", "tap", "tuning", "vlan", "vrf"),
		remoteMod + "ContainerdInstall":     generateArchiveInstall(commandSpec, "Installs containerd on a remote system", "containerd"),
		remoteMod + "CrictlInstall":         generateArchiveInstall(commandSpec, "Installs crictl on a remote system", "crictl"),
		remoteMod + "Download":              generateDownload(commandSpec),
		remoteMod + "EtcdInstall":           generateArchiveInstall(commandSpec, "Installs etcd on a remote system", "etcd", "etcdctl"),
		remoteMod + "KubeApiServerInstall":  generateBinaryInstall(commandSpec, "Installs kube-apiserver on a remote system."),
		remoteMod + "KubeControllerManager": generateBinaryInstall(commandSpec, "Installs kube-controller-manager on a remote system."),
		remoteMod + "KubectlInstall":        generateBinaryInstall(commandSpec, "Installs kubectl on a remote system."),
		remoteMod + "KubeletInstall":        generateBinaryInstall(commandSpec, "Installs kubelet on a remote system."),
		remoteMod + "KubeProxyInstall":      generateBinaryInstall(commandSpec, "Installs kube-proxy on a remote system."),
		remoteMod + "KubeSchedulerInstall":  generateBinaryInstall(commandSpec, "Installs kube-scheduler on a remote system."),
		remoteMod + "RuncInstall":           generateBinaryInstall(commandSpec, "Installs runc on a remote system."),
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
			TypeSpec:    schema.TypeSpec{Ref: localType("Download", "remote")},
		},
		"mkdir": {
			Description: "The mkdir operation.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Mkdir", "tools")},
		},
		"mktemp": {
			Description: "The mktemp operation.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Mktemp", "tools")},
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
			TypeSpec:    schema.TypeSpec{Ref: localType("Download", "remote")},
		},
		"mkdir": {
			Description: "The mkdir operation.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Mkdir", "tools")},
		},
		"mktemp": {
			Description: "The mktemp operation.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Mktemp", "tools")},
		},
		"mv": {
			Description: "The mv operation.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Mv", "tools")},
		},
		"path": {
			Description: "The path to the installed binary.",
			TypeSpec:    typeSpecs.String,
		},
		"rm": {
			Description: "The rm operation.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Rm", "tools")},
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
