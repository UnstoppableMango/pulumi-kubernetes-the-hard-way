package examples

import (
	"fmt"
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
	skipIfShort(t)

	const (
		username = "root"
		password = "root"
	)

	node := newNode(t,
		WithSshUsername(username),
		WithSshPassword(password),
	)

	test := baseOptions.With(integration.ProgramTestOptions{
		Dir: path.Join(getCwd(t), project),
		Config: map[string]string{
			"host":     "localhost",
			"port":     node.Port,
			"user":     username,
			"password": password,
		},
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			ctx := ResourceContext{tokens: map[validatorKey]ResourceValidator{}}
			validation(&ctx)
			lookup := ctx.tokens

			validated := []validatorKey{}
			for _, res := range stack.Deployment.Resources {
				key := validatorKey{Type: res.Type, Name: res.URN.Name()}
				testName := fmt.Sprintf("Testing resource %s", res.URN.Name())
				if validator, ok := lookup[key]; ok && t.Run(testName, func(t *testing.T) {
					validator(t, res)
				}) {
					validated = append(validated, key)
				}
			}

			assert.ElementsMatch(t, maps.Keys(lookup), validated)
		},
	})

	integration.ProgramTest(t, &test)
}
