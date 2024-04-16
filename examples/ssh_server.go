package examples

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const internalPort = "22/tcp"

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

	var buf bytes.Buffer
	ctxWriter := tar.NewWriter(&buf)
	ctxWriter.AddFS(os.DirFS(cwd))
	if err := ctxWriter.Close(); err != nil {
		return SshServer{}, err
	}

	req := testcontainers.ContainerRequest{
		Name: "node1",
		// Idfk why this is being dumb
		// FromDockerfile: testcontainers.FromDockerfile{
		// 	ContextArchive: bytes.NewReader(buf.Bytes()),
		// },
		Image:        "kthw:dev",
		ExposedPorts: []string{internalPort},
		// Env: map[string]string{
		// 	"PUID":            "1000",
		// 	"PGID":            "1000",
		// 	"TZ":              "America/Chicago",
		// 	"USER_NAME":       options.Username,
		// 	"PASSWORD_ACCESS": strconv.FormatBool(options.Password != ""),
		// 	"USER_PASSWORD":   options.Password,
		// 	"PUBLIC_KEY":      options.PublicKey,
		// 	"SUDO_ACCESS":     "true",
		// },
		Privileged: true,
		HostConfigModifier: func(hc *container.HostConfig) {
			hc.CapAdd = append(hc.CapAdd, "ALL")
		},
		WaitingFor: wait.ForExposedPort(),
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
