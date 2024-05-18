//go:build nodejs || all

package examples

import (
	_ "embed"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/example.com.html
var exampleComHtml string

func TestSimpleTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "simple-ts"),
			Quick:         true,
			SkipRefresh:   true,
			RunUpdateTest: false,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				assert.NotEmpty(t, stack.Outputs["certPem"])
				assert.NotEmpty(t, stack.Outputs["keyPem"])
				assert.NotEmpty(t, stack.Outputs["kubeconfigJson"])
				assert.NotEmpty(t, stack.Outputs["kubeVipYaml"])
			},
		})

	integration.ProgramTest(t, &test)
}

func TestRemoteInstallTs(t *testing.T) {
	skipIfShort(t)
	skipIfCi(t)

	const (
		username = "root"
		password = "root"
	)

	node := newNode(t,
		WithSshUsername(username),
		WithSshPassword(password),
	)

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "remote-install-ts"),
			Quick:         true,
			SkipRefresh:   true,
			RunUpdateTest: false,
			Config: map[string]string{
				"host":     "localhost",
				"port":     node.Port,
				"user":     username,
				"password": password,
				"basePath": "/config",
			},
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@unmango/pulumi-kubernetes-the-hard-way",
		},
	})

	return baseJS
}
