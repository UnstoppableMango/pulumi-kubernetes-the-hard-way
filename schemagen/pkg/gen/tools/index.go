package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func Generate(commandSpec schema.PackageSpec) schema.PackageSpec {
	commonProps := map[string]schema.PropertySpec{
		"binaryPath": props.String("Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH"),
		"connection": {
			Description: "Connection details for the remote system",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"environment": props.StringMap("Environment variables"),
		"lifecycle": {
			Description: "At what stage(s) in the resource lifecycle should the command be run",
			TypeSpec: schema.TypeSpec{
				Ref:   types.LocalType("CommandLifecycle", "tools").Ref,
				Plain: true,
			},
		},
		"stdin": props.String("TODO"),
		"triggers": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Type: "array",
				Items: &schema.TypeSpec{
					Ref: "pulumi.json#/Any",
				},
			},
		},
	}

	commonOutputs := map[string]schema.PropertySpec{
		"command": {
			Description: "The underlying command",
			TypeSpec:    types.ExtResource(commandSpec, "Command", "remote"),
		},
		"stderr": props.String("TODO"),
		"stdin":  props.String("TODO"),
		"stdout": props.String("TODO"),
	}

	tools := map[string]schema.ResourceSpec{
		"Chmod":       generateChmod(),
		"Etcdctl":     generateEtcdctl(),
		"Hostnamectl": generateHostnamectl(),
		"Mkdir":       generateMkdir(),
		"Mktemp":      generateMktemp(),
		"Mv":          generateMv(),
		"Rm":          generateRm(),
		"Sed":         generateSed(),
		"Systemctl":   generateSystemctl(),
		"Tar":         generateTar(),
		"Tee":         generateTee(),
		"Wget":        generateWget(),
	}

	resources := map[string]schema.ResourceSpec{}
	for tool, resource := range tools {
		for propName, prop := range commonProps {
			resource.InputProperties[propName] = prop
			resource.Properties[propName] = prop
		}
		for name, prop := range commonOutputs {
			resource.Properties[name] = prop
		}

		resource.IsComponent = true
		resource.RequiredInputs = append(resource.RequiredInputs,
			"connection",
		)
		resource.Required = append(resource.Required,
			"binaryPath",
			"command",
			"connection",
			"environment",
			"stderr",
			"stdout",
			"triggers",
		)

		resources[name(tool)] = resource
	}

	return schema.PackageSpec{
		Functions: map[string]schema.FunctionSpec{},
		Resources: resources,
		Types:     generateTypes(),
	}
}
