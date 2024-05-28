package rt

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

type ResourceContext struct {
	Stack   integration.RuntimeValidationStackInfo
	options integration.ProgramTestOptions
	tokens  map[validatorKey]ResourceValidator
}

func ForOptions(options integration.ProgramTestOptions) ResourceContext {
	return ResourceContext{
		options: options,
		tokens:  map[validatorKey]ResourceValidator{},
	}
}

func ResourceTest(t *testing.T, project string, options integration.ProgramTestOptions, validation func(ctx *ResourceContext)) {
	if _, ok := options.Config["test:container"]; !ok {
		options = options.With(SingleContainerSetup(t))
	}

	test := options.With(integration.ProgramTestOptions{
		Dir: path.Join(getCwd(t), project),
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			ctx := ForOptions(options)
			validation(&ctx)
			lookup := ctx.tokens

			validated := []string{}
			for _, res := range stack.Deployment.Resources {
				key := validatorKey{Type: res.Type, Name: res.URN.Name()}
				testName := fmt.Sprintf("%s::%s", res.Type, res.URN.Name())
				if validator, ok := lookup[key]; ok && t.Run(testName, func(t *testing.T) {
					validator(t, res)
				}) {
					validated = append(validated, fmt.Sprintf("%s:%s", key.Name, key.Type))
				}
			}

			expected := []string{}
			for _, key := range maps.Keys(lookup) {
				expected = append(expected, fmt.Sprintf("%s:%s", key.Name, key.Type))
			}

			assert.ElementsMatch(t, expected, validated, "Not all resources were validated")
		},
	})

	integration.ProgramTest(t, &test)
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}
