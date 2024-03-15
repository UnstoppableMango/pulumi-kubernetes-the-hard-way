package provider

import (
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstallOnArgs struct {
	Connection remote.ConnectionArgs `pulumi:"connection"`
	Name       string                `pulumi:"name"`
	Path       pulumi.StringInput    `pulumi:"path"`
}

type InstallOnResult struct {
	File *RemoteFile `pulumi:"file"`
}
