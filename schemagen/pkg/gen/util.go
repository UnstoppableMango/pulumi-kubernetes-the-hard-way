package gen

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

var arrayOfStrings = schema.TypeSpec{
	Type: "array",
	Items: &schema.TypeSpec{
		Type: "string",
	},
}

func implicitOutputs(inputs, outputs map[string]schema.PropertySpec) map[string]schema.PropertySpec {
	for k, v := range inputs {
		outputs[k] = v
	}

	return outputs
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
