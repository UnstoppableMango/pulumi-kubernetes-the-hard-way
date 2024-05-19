package gen

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

var typeSpecs = struct {
	ArrayOfStrings   schema.TypeSpec
	Boolean          schema.TypeSpec
	Integer          schema.TypeSpec
	OneOrMoreStrings schema.TypeSpec
	String           schema.TypeSpec
	StringMap        schema.TypeSpec
}{
	ArrayOfStrings: schema.TypeSpec{
		Type: "array",
		Items: &schema.TypeSpec{
			Type: "string",
		},
	},
	Boolean: schema.TypeSpec{Type: "boolean"},
	Integer: schema.TypeSpec{Type: "integer"},
	OneOrMoreStrings: schema.TypeSpec{
		OneOf: []schema.TypeSpec{
			{Type: "string"},
			{Type: "array", Items: &schema.TypeSpec{
				Type: "string",
			}},
		},
	},
	String: schema.TypeSpec{Type: "string"},
	StringMap: schema.TypeSpec{
		Type:                 "object",
		AdditionalProperties: &schema.TypeSpec{Type: "string"},
	},
}

func implicitOutputs(inputs, outputs map[string]schema.PropertySpec) map[string]schema.PropertySpec {
	for k, v := range inputs {
		outputs[k] = v
	}

	return outputs
}

func localResource(name string, modules ...string) string {
	return fmt.Sprintf("#/resources/kubernetes-the-hard-way:%s:%s",
		strings.Join(modules, "/"),
		name,
	)
}

func localType(name string, modules ...string) string {
	return fmt.Sprintf("#/types/kubernetes-the-hard-way:%s:%s",
		strings.Join(modules, "/"),
		name,
	)
}

func refResource(spec schema.PackageSpec, name string, modules ...string) string {
	return fmt.Sprintf("/%s/%s/schema.json#/resources/%s:%s:%s",
		spec.Name,
		spec.Version,
		spec.Name,
		strings.Join(modules, "/"),
		name,
	)
}

func refType(spec schema.PackageSpec, name string, modules ...string) string {
	return fmt.Sprintf("/%s/%s/schema.json#/types/%s:%s:%s",
		spec.Name,
		spec.Version,
		spec.Name,
		strings.Join(modules, "/"),
		name,
	)
}

// func renameComplexRefs(spec schema.ComplexTypeSpec, old, new string) schema.ComplexTypeSpec {
// 	spec.Properties = renamePropertiesRefs(spec.Properties, old, new)
// 	return spec
// }
