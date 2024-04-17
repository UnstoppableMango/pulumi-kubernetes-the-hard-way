//go:build nodejs || all

package examples

import (
	"context"
	_ "embed"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
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

	ctx := context.Background()
	server, err := StartSshServer(ctx,
		WithSshUsername(username),
		WithSshPassword(password))
	assert.NoError(t, err)

	port, err := server.Port(ctx)
	assert.NoError(t, err)

	defer StopSshServer(ctx, server) // TODO: Error handling?

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "remote-install-ts"),
			Quick:         true,
			SkipRefresh:   true,
			RunUpdateTest: false,
			Config: map[string]string{
				"host":     "localhost",
				"port":     port,
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
	const (
		username = "root"
		password = "root"
		content  = "Some content idk"
	)

	ctx := context.Background()
	server, err := StartSshServer(ctx,
		WithSshUsername(username),
		WithSshPassword(password))
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	port, err := server.Port(ctx)
	assert.NoError(t, err)

	defer StopSshServer(ctx, server) // TODO: Error handling?

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
				"port":     port,
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

				data, err := server.ReadFile(ctx, "/config/index.html")
				assert.NoError(t, err)
				assert.Equal(t, exampleComHtml, data)

				// assert.Empty(t, stack.Outputs["tarStdout"])
				// assert.Empty(t, stack.Outputs["tarStderr"])

				assert.NotEmpty(t, stack.Outputs["mktemp"])
			},
		})

	integration.ProgramTest(t, &test)
}

func TestTlsRootCaTs(t *testing.T) {
	validateSimple := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)
		allowedUses, ok := res.Outputs["allowedUses"]
		assert.True(t, ok, "Outputs did not contain `allowedUses`")
		assert.ElementsMatch(t,
			[]string{"cert_signing", "key_encipherment", "server_auth", "client_auth"},
			allowedUses,
		)

		algorithm, ok := res.Outputs["algorithm"]
		assert.True(t, ok, "Outputs did not contain `algorithm`")
		assert.Equal(t, "RSA", algorithm)

		validityPeriodHours, ok := res.Outputs["validityPeriodHours"]
		assert.True(t, ok, "Outputs did not contain `validityPeriodHours`")
		assert.Equal(t, 256., validityPeriodHours)

		assert.Contains(t, res.Outputs, "cert")
		assert.Contains(t, res.Outputs, "certPem")
		assert.Contains(t, res.Outputs, "key")
		assert.Contains(t, res.Outputs, "privateKeyPem")
		assert.Contains(t, res.Outputs, "publicKeyPem")
	}

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:   path.Join(getCwd(t), "tls", "root-ca-ts"),
			Quick: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				validatedResources := []string{}
				for _, res := range stack.Deployment.Resources {
					if res.Type != "kubernetes-the-hard-way:tls:RootCa" {
						continue
					}
					switch res.URN.Name() {
					case "simple":
						validateSimple(t, res)
						validatedResources = append(validatedResources, "simple")
					}
				}

				assert.Equal(t, []string{"simple"}, validatedResources, "Not all resources were validated")
			},
		})

	integration.ProgramTest(t, &test)
}

func TestTlsTs(t *testing.T) {
	validateRootCa := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)
		assert.ElementsMatch(t,
			[]string{"cert_signing", "key_encipherment", "server_auth", "client_auth"},
			res.Outputs["allowedUses"],
		)

		assert.Equal(t, "RSA", res.Outputs["algorithm"])

		if res.Parent.Type() == "kubernetes-the-hard-way:tls:ClusterPki" {
			assert.Equal(t, 8076., res.Outputs["validityPeriodHours"])
		} else {
			assert.Equal(t, 256., res.Outputs["validityPeriodHours"])
		}

		assert.NotNil(t, res.Outputs["cert"])
		assert.NotNil(t, res.Outputs["certPem"])
		assert.NotNil(t, res.Outputs["key"])
		assert.NotNil(t, res.Outputs["privateKeyPem"])
		assert.NotNil(t, res.Outputs["publicKeyPem"])
	}

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:   path.Join(getCwd(t), "tls-ts"),
			Quick: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				for _, res := range stack.Deployment.Resources {
					switch res.Type {
					case "kubernetes-the-hard-way:tls:RootCa":
						validateRootCa(t, res)
					}
				}
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
