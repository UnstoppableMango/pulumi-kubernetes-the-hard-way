package gen

import (
	"encoding/json"
	"log"
	"path"

	"os"

	"github.com/pulumi/pulumi/pkg/codegen/schema"
)

func GenerateSchema(packageDir string) schema.PackageSpec {
	dependencies := readPackageDependencies(packageDir)
	packageSpec := schema.PackageSpec{}

	return packageSpec
}

type Dependencies struct {
	Command string `json:"@pulumi/command"`
	Random  string `json:"@pulumi/random"`
	Tls     string `json:"@pulumi/tls"`
}

type PackageJson struct {
	Dependencies Dependencies
}

func readPackageDependencies(packageDir string) Dependencies {
	content, err := os.ReadFile(path.Join(packageDir, "package.json"))
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload PackageJson
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload.Dependencies
}
