package types

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

var ArrayOfStrings = schema.TypeSpec{
	Type: "array",
	Items: &schema.TypeSpec{
		Type: "string",
	},
}

var Boolean = schema.TypeSpec{Type: "boolean"}

var Integer = schema.TypeSpec{Type: "integer"}

var OneOrMoreStrings = schema.TypeSpec{
	OneOf: []schema.TypeSpec{
		{Type: "string"},
		{Type: "array", Items: &schema.TypeSpec{
			Type: "string",
		}},
	},
}

var String = schema.TypeSpec{Type: "string"}

var StringMap = schema.TypeSpec{
	Type:                 "object",
	AdditionalProperties: &schema.TypeSpec{Type: "string"},
}

func LocalResource(name string, modules ...string) schema.TypeSpec {
	return schema.TypeSpec{
		Ref: localResource(name, modules...),
	}
}

func localResource(name string, modules ...string) string {
	return fmt.Sprintf("#/resources/kubernetes-the-hard-way:%s:%s",
		strings.Join(modules, "/"),
		name,
	)
}

func LocalType(name string, modules ...string) schema.TypeSpec {
	return schema.TypeSpec{
		Ref: localType(name, modules...),
	}
}

func localType(name string, modules ...string) string {
	return fmt.Sprintf("#/types/kubernetes-the-hard-way:%s:%s",
		strings.Join(modules, "/"),
		name,
	)
}

func ExtResource(spec schema.PackageSpec, name string, modules ...string) schema.TypeSpec {
	return schema.TypeSpec{
		Ref: refResource(spec, name, modules...),
	}
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

func ExtType(spec schema.PackageSpec, name string, modules ...string) schema.TypeSpec {
	return schema.TypeSpec{
		Ref: refType(spec, name, modules...),
	}
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
