package examples

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	cwd, err := os.Getwd()
	if err != nil {
		return SshServer{}, err
	}

	options := SshServerOptions{}
	for _, o := range opts {
		o(&options)
	}

	fmt.Printf("CWD IS: %s\n", cwd)
	fmt.Printf("JOINED IS: %s\n", filepath.Join(cwd, "testdata"))

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "",
			Dockerfile: "Dockerfile",
		},
		ExposedPorts: []string{internalPort},
		Env: map[string]string{
			"PUID":            "1000",
			"PGID":            "1000",
			"TZ":              "America/Chicago",
			"USER_NAME":       options.Username,
			"PASSWORD_ACCESS": strconv.FormatBool(options.Password != ""),
			"USER_PASSWORD":   options.Password,
			"PUBLIC_KEY":      options.PublicKey,
			"SUDO_ACCESS":     "true",
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

func (s *SshServer) ReadFile(ctx context.Context, filePath string) (string, error) {
	reader, err := s.Container.CopyFileFromContainer(ctx, filePath)
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (s *SshServer) CopyFile(ctx context.Context, source, target string) error {
	return s.Container.CopyFileToContainer(ctx, source, target, 666)
}
