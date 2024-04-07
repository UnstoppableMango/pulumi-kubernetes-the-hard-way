package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func main() {
	printUsage := func() {
		fmt.Printf("Usage: %s <out-dir>\n", os.Args[0])
	}

	args := os.Args[1:]
	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	outDir := args[0]
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	rootDir := filepath.Join(cwd, "..", "..", "..")
	packageDir := filepath.Join(rootDir, "provider", "cmd", "pulumi-resource-kubernetes-the-hard-way")
	spec := gen.GenerateSchema(packageDir)
	mustWritePulumiSchema(spec, outDir)
}

func mustWriteFile(rootDir, filename string, contents []byte) {
	outPath := filepath.Join(rootDir, filename)

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		panic(err)
	}
	err := os.WriteFile(outPath, contents, 0600)
	if err != nil {
		panic(err)
	}
}

func mustWritePulumiSchema(pkgSpec schema.PackageSpec, outdir string) {
	schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		panic(errors.Wrap(err, "marshaling Pulumi schema"))
	}
	mustWriteFile(outdir, "schema.json", schemaJSON)
}
