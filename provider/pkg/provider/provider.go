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
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

func construct(ctx *pulumi.Context, typ, name string, inputs provider.ConstructInputs,
	options pulumi.ResourceOption) (*provider.ConstructResult, error) {
	switch typ {
	case "kubernetes-the-hard-way:index:Certificate":
		return constructCertificate(ctx, name, inputs, options)
	case "kubernetes-the-hard-way:index:RemoteFile":
		return constructRemoteFile(ctx, name, inputs, options)
	case "kubernetes-the-hard-way:index:RootCa":
		return constructRootCa(ctx, name, inputs, options)
	//case "kubernetes-the-hard-way:index:test":
	//	return constructGeneric(ctx, name, inputs, &RootCaArgs{}, NewRootCa, options)
	default:
		return nil, errors.Errorf("unknown resource type %s", typ)
	}
}

//type builder func(ctx *pulumi.Context, name string, args interface{}, opts ...pulumi.ResourceOption)
//
//func constructGeneric[T any](ctx *pulumi.Context, name string, inputs provider.ConstructInputs,
//	args *T,
//	construct func(ctx *pulumi.Context, name string, args interface{}, opts ...pulumi.ResourceOption) (pulumi.ComponentResource, error),
//	opts pulumi.ResourceOption) (*provider.ConstructResult, error) {
//	if err := inputs.CopyTo(args); err != nil {
//		return nil, errors.Wrap(err, "setting args")
//	}
//
//	component, err := construct(ctx, name, args, opts)
//	if err != nil {
//		return nil, errors.Wrap(err, "creating component")
//	}
//
//	return provider.NewConstructResult(component)
//}

func constructCertificate(ctx *pulumi.Context, name string, inputs provider.ConstructInputs,
	options pulumi.ResourceOption) (*provider.ConstructResult, error) {
	args := &CertificateArgs{}
	if err := inputs.CopyTo(args); err != nil {
		return nil, errors.Wrap(err, "setting args")
	}

	component, err := NewCertificate(ctx, name, args, options)
	if err != nil {
		return nil, errors.Wrap(err, "creating component")
	}

	return provider.NewConstructResult(component)
}

func constructRemoteFile(ctx *pulumi.Context, name string, inputs provider.ConstructInputs,
	options pulumi.ResourceOption) (*provider.ConstructResult, error) {
	args := &RemoteFileArgs{}
	if err := inputs.CopyTo(args); err != nil {
		return nil, errors.Wrap(err, "setting args")
	}

	component, err := NewRemoteFile(ctx, name, args, options)
	if err != nil {
		return nil, errors.Wrap(err, "creating component")
	}

	return provider.NewConstructResult(component)
}

func constructRootCa(ctx *pulumi.Context, name string, inputs provider.ConstructInputs,
	options pulumi.ResourceOption) (*provider.ConstructResult, error) {
	args := &RootCaArgs{}
	if err := inputs.CopyTo(args); err != nil {
		return nil, errors.Wrap(err, "setting args")
	}

	component, err := NewRootCa(ctx, name, args, options)
	if err != nil {
		return nil, errors.Wrap(err, "creating component")
	}

	return provider.NewConstructResult(component)
}
