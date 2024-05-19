package internal

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func MakeExternal(propertySpec map[string]schema.PropertySpec, spec schema.PackageSpec) map[string]schema.PropertySpec {
	prefix := fmt.Sprintf("/%s/%s/schema.json#/", spec.Name, spec.Version)
	return renamePropertiesRefs(propertySpec, "#/", prefix)
}

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
