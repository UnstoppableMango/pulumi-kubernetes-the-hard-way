package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func Generate(commandSpec schema.PackageSpec) schema.PackageSpec {
	mkCmd := makeCommand(commandSpec)
	tools := map[string]schema.ResourceSpec{
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
	for tool, typ := range tools {
		c, r := mkCmd(typ)
		types[name(tool)+"Opts"] = c
		resources[name(tool)] = r
	}

	return schema.PackageSpec{
		Functions: map[string]schema.FunctionSpec{},
		Resources: resources,
		Types:     types,
	}
}

func makeCommand(commandSpec schema.PackageSpec) func(r schema.ResourceSpec) (schema.ComplexTypeSpec, schema.ResourceSpec) {
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

	return func(r schema.ResourceSpec) (schema.ComplexTypeSpec, schema.ResourceSpec) {
		c := schema.ComplexTypeSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: r.Description,
				Properties:  r.Properties,
			},
		}

		for propName, prop := range commonProps {
			r.InputProperties[propName] = prop
			r.Properties[propName] = prop
		}
		for name, prop := range commonOutputs {
			r.Properties[name] = prop
		}

		r.IsComponent = true
		r.RequiredInputs = append(r.RequiredInputs,
			"connection",
		)
		r.Required = append(r.Required,
			"binaryPath",
			"command",
			"connection",
			"environment",
			"stderr",
			"stdout",
			"triggers",
		)

		return c, r
	}
}
