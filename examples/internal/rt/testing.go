package rt

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

type ResourceValidator func(t *testing.T, res apitype.ResourceV3)

type ResourceContext struct {
	tokens map[validatorKey]ResourceValidator
}

type validatorKey struct {
	Type tokens.Type
	Name string
}

func Validate(ctx *ResourceContext, typ tokens.Type, name string, validator ResourceValidator) {
	ctx.tokens[validatorKey{Type: typ, Name: name}] = validator
}

func ResourceTest(t *testing.T, project string, baseOptions integration.ProgramTestOptions, validation func(ctx *ResourceContext)) {
	if _, ok := baseOptions.Config["test:container"]; !ok {
		baseOptions = baseOptions.With(SingleContainerSetup(t))
	}

	test := baseOptions.With(integration.ProgramTestOptions{
		Dir: path.Join(getCwd(t), project),
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			ctx := ResourceContext{tokens: map[validatorKey]ResourceValidator{}}
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
