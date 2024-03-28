package examples

import (
	"context"
	"strconv"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const internalPort = "2222"

type SshServer struct {
	Container testcontainers.Container
}

type SshServerOptions struct {
	Password  string
	PublicKey string
	Username  string
}

type SshServerOption func(o *SshServerOptions)

func WithSshPassword(password string) SshServerOption {
	return func(o *SshServerOptions) {
		o.Password = password
	}
}

func WithSshPublicKey(key string) SshServerOption {
	return func(o *SshServerOptions) {
		o.PublicKey = key
	}
}

func WithSshUsername(username string) SshServerOption {
	return func(o *SshServerOptions) {
		o.Username = username
	}
}

func StartSshServer(ctx context.Context, opts ...SshServerOption) (SshServer, error) {
	options := SshServerOptions{}
	for _, o := range opts {
		o(&options)
	}

	req := testcontainers.ContainerRequest{
		Image:        "lscr.io/linuxserver/openssh-server:latest",
		ExposedPorts: []string{internalPort},
		Env: map[string]string{
			"PUID":            "1000",
			"PGID":            "1000",
			"TZ":              "America/Chicago",
			"PUBLIC_KEY":      options.PublicKey,
			"PASSWORD_ACCESS": strconv.FormatBool(options.Password != ""),
			"USER_NAME":       options.Username,
			"USER_PASSWORD":   options.Password,
		},
		WaitingFor: wait.ForExposedPort(), // TODO: What's a good condition
	}

	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	return SshServer{Container: c}, err
}

func StopSshServer(ctx context.Context, server SshServer) error {
	return server.Container.Terminate(ctx)
}

func (s *SshServer) Port(ctx context.Context) (string, error) {
	port, err := s.Container.MappedPort(ctx, nat.Port(internalPort))

	return port.Port(), err
}
