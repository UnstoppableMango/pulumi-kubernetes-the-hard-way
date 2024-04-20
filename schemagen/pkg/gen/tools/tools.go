package tools

import (
	"fmt"
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

type tool struct {
	optsType schema.ComplexTypeSpec
	types    map[string]schema.ComplexTypeSpec
}

func Generate(commandSpec schema.PackageSpec) schema.PackageSpec {
	tools := map[string]tool{
		"Chmod":       generateChmod(),
		"Curl":        generateCurl(),
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

	types := generateTypes()
	resources := map[string]schema.ResourceSpec{}

	// TODO: This gets confusing fast
	for name, tool := range tools {
		for typeName, typ := range tool.types {
			types[qualifyName(name)+typeName] = typ
		}
		types[qualifyName(name)+"Opts"] = tool.optsType
		resources[qualifyName(name)] = tool.resourceSpec(
			commandSpec,
			optsType(name),
		)
	}

	return schema.PackageSpec{
		Functions: map[string]schema.FunctionSpec{},
		Resources: resources,
		Types:     types,
	}
}

func qualifyName(x string) string {
	return fmt.Sprintf("kubernetes-the-hard-way:tools:%s", x)
}

func optsType(x string) schema.TypeSpec {
	typ := types.LocalType(x+"Opts", "tools")
	typ.Plain = true
	return typ
}

func (tool tool) resourceSpec(commandSpec schema.PackageSpec, optsType schema.TypeSpec) schema.ResourceSpec {
	command := commandSpec.Resources["command:remote:Command"]

	inputs := map[string]schema.PropertySpec{
		"binaryPath": props.String("Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH"),
		"create": {
			Description: command.InputProperties["create"].Description,
			TypeSpec:    optsType,
		},
		"connection": {
			Description: "Connection details for the remote system",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"delete": {
			Description: command.InputProperties["delete"].Description,
			TypeSpec:    optsType,
		},
		"environment": props.StringMap("Environment variables"),
		"stdin":       props.String("TODO"),
		"triggers": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Type: "array",
				Items: &schema.TypeSpec{
					Ref: "pulumi.json#/Any",
				},
			},
		},
		"update": {
			Description: command.InputProperties["update"].Description,
			TypeSpec:    optsType,
		},
	}

	outputs := map[string]schema.PropertySpec{
		"command": {
			Description: "The underlying command",
			TypeSpec:    types.ExtResource(commandSpec, "Command", "remote"),
		},
		"stderr": props.String("TODO"),
		"stdin":  props.String("TODO"),
		"stdout": props.String("TODO"),
	}
	maps.Copy(outputs, inputs)

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: tool.optsType.Description,
			Properties:  outputs,
			Required: []string{
				"binaryPath",
				"command",
				"connection",
				"environment",
				"stderr",
				"stdout",
				"triggers",
			},
		},
		InputProperties: inputs,
		RequiredInputs:  []string{"connection"},
	}
}
