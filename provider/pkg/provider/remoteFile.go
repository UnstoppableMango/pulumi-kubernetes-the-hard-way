// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The RemoteFileArgs for creating a RemoteFile component resource.
type RemoteFileArgs struct {
	Connection remote.ConnectionArgs `pulumi:"connection"`
	Content    pulumi.StringInput    `pulumi:"content"`
	Path       pulumi.StringInput    `pulumi:"path"`
}

// The RemoteFile component resource.
type RemoteFile struct {
	pulumi.ResourceState

	Command *remote.Command `pulumi:"command"`
}

// NewRemoteFile creates a new RemoteFile component resource.
func NewRemoteFile(ctx *pulumi.Context,
	name string, args *RemoteFileArgs, opts ...pulumi.ResourceOption) (*RemoteFile, error) {
	if args == nil {
		args = &RemoteFileArgs{}
	}

	component := &RemoteFile{}
	err := ctx.RegisterComponentResource("kubernetes-the-hard-way:index:RemoteFile", name, component, opts...)
	if err != nil {
		return nil, err
	}

	command, err := remote.NewCommand(ctx, name, &remote.CommandArgs{
		Connection: args.Connection,
		Stdin:      args.Content,
		Create:     pulumi.Sprintf("tee %s", args.Path),
		Delete:     pulumi.Sprintf("rm %s", args.Path),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	component.Command = command

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"command": command,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
