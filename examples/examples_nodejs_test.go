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
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			},
		})

	integration.ProgramTest(t, &test)
}

func TestRemoteTs(t *testing.T) {
	t.Skip("Disabling in favor of future, more specific, tests")

	const (
		username = "root"
		password = "root"
		content  = "Some content idk"
	)

	node := newNode(t,
		WithSshUsername(username),
		WithSshPassword(password),
	)

	// err = server.CopyFile(ctx,
	// 	path.Join(getCwd(t), "testdata", "text-file.tar.gz"),
	// 	path.Join("/config", "text-file.tar.gz"))
	// assert.NoError(t, err)

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "remote-ts"),
			Quick:         true,
			SkipRefresh:   true,
			RunUpdateTest: false,
			Config: map[string]string{
				"host":     "localhost",
				"port":     node.Port,
				"user":     username,
				"password": password,
				"content":  content,
				"basePath": "/config",
			},
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				assert.Equal(t, content, stack.Outputs["fileStdout"])
				assert.Empty(t, stack.Outputs["fileStderr"])

				// Wget outputs progress to stderr
				assert.Empty(t, stack.Outputs["wgetStdout"])
				assert.NotEmpty(t, stack.Outputs["wgetStderr"])

				data := node.ReadFile(t, "/config/index.html")
				assert.Equal(t, exampleComHtml, data)

				// assert.Empty(t, stack.Outputs["tarStdout"])
				// assert.Empty(t, stack.Outputs["tarStderr"])

				assert.NotEmpty(t, stack.Outputs["mktemp"])
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
