package provider

import (
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstallOnArgs struct {
	Connection remote.ConnectionArgs
	Name       string
	Path       pulumi.StringInput
}
