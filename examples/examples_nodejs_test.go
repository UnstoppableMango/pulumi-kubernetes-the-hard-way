//go:build nodejs || all

package examples

import (
	"context"
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
)

// func TestSimpleTs(t *testing.T) {
// 	test := getJSBaseOptions(t).
// 		With(integration.ProgramTestOptions{
// 			Dir:           path.Join(getCwd(t), "simple-ts"),
// 			Quick:         true,
// 			SkipRefresh:   true,
// 			RunUpdateTest: false,
// 		})

// 	integration.ProgramTest(t, &test)
// }

func TestRemoteFileTs(t *testing.T) {
	ctx := context.Background()
	program := auto.Program(func(ctx *pulumi.Context) error {
		image, err := docker.NewRemoteImage(ctx, "node", &docker.RemoteImageArgs{
			Name: pulumi.String("lscr.io/linuxserver/openssh-server:latest"),
		})
		if err != nil {
			return err
		}

		_, err = docker.NewContainer(ctx, "node", &docker.ContainerArgs{
			Image: image.ImageId,
			Envs: pulumi.StringArray{
				pulumi.String("PUID=1000"),
				pulumi.String("PGID=1000"),
				pulumi.String("TZ=America/Chicago"),
			},
			Ports: docker.ContainerPortArray{docker.ContainerPortArgs{
				External: pulumi.IntPtr(2222),
				Internal: pulumi.Int(2222),
			}},
		})
		if err != nil {
			return err
		}

		return nil
	})

	workspace, err := auto.NewLocalWorkspace(ctx, program)
	assert.NoError(t, err)

	stackName := fmt.Sprintf("UnstoppableMango/kubernetes-the-hard-way/test-%d", time.Now().Unix())
	err = createOrSelectStack(ctx, workspace, stackName)
	assert.NoError(t, err)

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "remote-file"),
			Quick:         true,
			SkipRefresh:   true,
			RunUpdateTest: false,
			Config: map[string]string{
				"host":    "localhost",
				"content": "Some test content idk",
				"path":    "/home/testuser/remote-file-test",
			},
		})

	integration.ProgramTest(t, &test)
	err = workspace.RemoveStack(ctx, stackName)
	assert.NoError(t, err)
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
