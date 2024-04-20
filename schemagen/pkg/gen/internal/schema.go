package internal

import (
	"log"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func EmptyPackage() schema.PackageSpec {
	return schema.PackageSpec{
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{},
		Types:     map[string]schema.ComplexTypeSpec{},
	}
}

func ExtendSchemas(spec schema.PackageSpec, extensions ...schema.PackageSpec) schema.PackageSpec {
	for _, extension := range extensions {
		for k, v := range extension.Resources {
			if _, found := spec.Resources[k]; found {
				log.Fatalf("resource already defined %q", k)
			}
			spec.Resources[k] = v
		}

		for k, v := range extension.Types {
			if _, found := spec.Types[k]; found {
				log.Fatalf("type already defined %q", k)
			}
			spec.Types[k] = v
		}

		for k, v := range extension.Functions {
			if _, found := spec.Functions[k]; found {
				log.Fatalf("function already defined %q", k)
			}
			spec.Functions[k] = v
		}
	}

	return spec
}
