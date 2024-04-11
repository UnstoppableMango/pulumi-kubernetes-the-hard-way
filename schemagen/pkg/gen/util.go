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

func makeExternal(propertySpec map[string]schema.PropertySpec, spec schema.PackageSpec) map[string]schema.PropertySpec {
	prefix := fmt.Sprintf("/%s/%s/schema.json#/", spec.Name, spec.Version)
	return renamePropertiesRefs(propertySpec, "#/", prefix)
}

// func renameComplexRefs(spec schema.ComplexTypeSpec, old, new string) schema.ComplexTypeSpec {
// 	spec.Properties = renamePropertiesRefs(spec.Properties, old, new)
// 	return spec
// }

func renamePropertiesRefs(propertySpec map[string]schema.PropertySpec, old, new string) map[string]schema.PropertySpec {
	properties := map[string]schema.PropertySpec{}
	for k, v := range propertySpec {
		properties[k] = renamePropertyRefs(v, old, new)
	}
	return properties
}

func renamePropertyRefs(propSpec schema.PropertySpec, old, new string) schema.PropertySpec {
	if propSpec.Ref != "" {
		propSpec.Ref = strings.Replace(propSpec.Ref, old, new, 1)
	}
	if propSpec.AdditionalProperties != nil {
		additionalProperties := renameTypeSpecRefs(*propSpec.AdditionalProperties, old, new)
		propSpec.AdditionalProperties = &additionalProperties
	}
	if propSpec.Items != nil {
		items := renameTypeSpecRefs(*propSpec.Items, old, new)
		propSpec.Items = &items
	}
	if propSpec.OneOf != nil {
		propSpec.OneOf = renameTypeSpecsRefs(propSpec.OneOf, old, new)
	}
	return propSpec
}

func renameTypeSpecsRefs(typeSpec []schema.TypeSpec, old, new string) []schema.TypeSpec {
	newSpecs := make([]schema.TypeSpec, len(typeSpec))
	for i, spec := range typeSpec {
		newSpecs[i] = renameTypeSpecRefs(spec, old, new)
	}
	return newSpecs
}

func renameTypeSpecRefs(typeSpec schema.TypeSpec, old, new string) schema.TypeSpec {
	if typeSpec.Ref != "" {
		typeSpec.Ref = strings.Replace(typeSpec.Ref, old, new, 1)
	}
	if typeSpec.AdditionalProperties != nil {
		additionalProperties := renameTypeSpecRefs(*typeSpec.AdditionalProperties, old, new)
		typeSpec.AdditionalProperties = &additionalProperties
	}
	if typeSpec.Items != nil {
		items := renameTypeSpecRefs(*typeSpec.Items, old, new)
		typeSpec.Items = &items
	}
	return typeSpec
}
