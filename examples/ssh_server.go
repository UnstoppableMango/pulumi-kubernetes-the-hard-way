package examples

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type SshServer struct {
	Container testcontainers.Container
}

func startSshServer(ctx context.Context) (SshServer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "lscr.io/linuxserver/openssh-server:latest",
		ExposedPorts: []string{"2222"},
		Env: map[string]string{
			"PUID": "1000",
			"PGID": "1000",
			"TZ":   "America/Chicago",
		},
		WaitingFor: wait.ForExposedPort(), // TODO: What's a good condition
	}

	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	return SshServer{Container: c}, err
}

func stopSshServer(ctx context.Context, server SshServer) error {
	return server.Container.Terminate(ctx)
}
