package gen

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strings"

	"os"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/go/common/util/contract"
)

func GenerateSchema(packageDir string) schema.PackageSpec {
	dependencies := readPackageDependencies(packageDir)
	commandSpec := getPackageSpec("command", dependencies.Command)
	kubernetesSpec := getPackageSpec("kubernetes", dependencies.Kubernetes)
	randomSpec := getPackageSpec("random", dependencies.Random)
	tlsSpec := getPackageSpec("tls", dependencies.Tls)

	packageSpec := schema.PackageSpec{
		Name:              "kubernetes-the-hard-way",
		Description:       "A Pulumi implementation of Kelsey Hightower's Kubernetes the Hard Way",
		Repository:        "https://github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way",
		License:           "Apache-2.0",
		Keywords:          []string{"pulumi", "command", "tls", "kubernetes", "category/utility", "kind/component"},
		PluginDownloadURL: "github://api.github.com/UnstoppableMango",
		Publisher:         "UnstoppableMango",
		Attribution:       "https://github.com/kelseyhightower/kubernetes-the-hard-way",
		Language: map[string]schema.RawMessage{
			"csharp": rawMessage(map[string]interface{}{
				"rootNamespace":                "UnMango",
				"dictionaryConstructors":       true,
				"liftSingleValueMethodReturns": true,
				"packageReferences": map[string]string{
					"Pulumi":            "3.*",
					"Pulumi.Command":    "0.9.*",
					"Pulumi.Kubernetes": "4.*",
					"Pulumi.Random":     "4.*",
					"Pulumi.Tls":        "5.*",
				},
			}),
			"go": rawMessage(map[string]interface{}{
				"generateResourceContainerTypes": true,
				"liftSingleValueMethodReturns":   true,
				"importBasePath":                 "github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way",
				"packageImportAliases": map[string]string{
					"github.com/pulumi/pulumi-command/sdk/go/command/remote":                                    "pulumiCommand",
					"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way": "kthw",
				},
			}),
			"nodejs": rawMessage(map[string]interface{}{
				"packageName":                  "@unmango/pulumi-kubernetes-the-hard-way",
				"liftSingleValueMethodReturns": true,
				"dependencies": map[string]string{
					"@pulumi/pulumi":     "^3.0.0",
					"@pulumi/command":    "^" + dependencies.Command,
					"@pulumi/kubernetes": "^" + dependencies.Kubernetes,
					"@pulumi/random":     "^" + dependencies.Random,
					"@pulumi/tls":        "^" + dependencies.Tls,
				},
				"devDependencies": map[string]string{
					"@types/node": "^18",
					"typescript":  "^4.6.2",
				},
			}),
			"python": rawMessage(map[string]interface{}{
				"liftSingleValueMethodReturns": true,
				"pyproject": map[string]bool{
					"enabled": true,
				},
				"requires": map[string]string{
					"pulumi":            ">=3.91.1,<4.0.0",
					"pulumi-command":    fmt.Sprintf(">=%s,<1.0.0", dependencies.Command),
					"pulumi-kubernetes": fmt.Sprintf(">=%s,<5.0.0", dependencies.Kubernetes),
					"pulumi-random":     fmt.Sprintf(">=%s,<5.0.0", dependencies.Random),
					"pulumi-tls":        fmt.Sprintf(">=%s,<6.0.0", dependencies.Tls),
				},
			}),
			"java": rawMessage(map[string]interface{}{
				"basePackage":                     "com.unmango",
				"buildFiles":                      "gradle",
				"liftSingleValueMethodReturns":    true,
				"gradleNexusPublishPluginVersion": "1.1.0",
				"dependencies": map[string]string{
					"com.google.code.findbugs:jsr305": "3.0.2",
					"com.google.code.gson:gson":       "2.8.9",
					"com.pulumi:command":              dependencies.Command,
					"com.pulumi:kubernetes":           dependencies.Kubernetes,
					"com.pulumi:pulumi":               "0.9.9",
					"com.pulumi:random":               dependencies.Random,
					"com.pulumi:tls":                  dependencies.Tls,
				},
			}),
		},
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{},
		Types:     map[string]schema.ComplexTypeSpec{},
	}

	return extendSchemas(packageSpec,
		generateConfig(kubernetesSpec),
		generateRemote(commandSpec),
		generateTls(randomSpec, tlsSpec),
		generateTools(commandSpec))
}

func extendSchemas(spec schema.PackageSpec, extensions ...schema.PackageSpec) schema.PackageSpec {
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

func rawMessage(v interface{}) schema.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}

func getPackageSpec(name, version string) schema.PackageSpec {
	// If the version has a commit hash, strip it off since they are not used in GitHub URLs by tag.
	urlVersion := version
	if before, _, found := strings.Cut(version, "+"); found {
		urlVersion = before
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/pulumi/pulumi-%s/v%s/provider/cmd/pulumi-resource-%s/schema.json", name, urlVersion, name)
	spec := getSpecFromUrl(url)
	if spec.Version == "" {
		// Version is rarely included, so we'll just add it.
		spec.Version = "v" + version
	}
	return spec
}

func getSpecFromUrl(url string) schema.PackageSpec {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Could not GET %s: %v", url, err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Could not read %s: %v", url, err)
	}
	var spec schema.PackageSpec
	err = json.Unmarshal(body, &spec)
	if err != nil {
		log.Fatalf("Could not parse %s: %v", url, err)
	}

	return spec
}

type Dependencies struct {
	Command    string `json:"@pulumi/command"`
	Kubernetes string `json:"@pulumi/kubernetes"`
	Random     string `json:"@pulumi/random"`
	Tls        string `json:"@pulumi/tls"`
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
