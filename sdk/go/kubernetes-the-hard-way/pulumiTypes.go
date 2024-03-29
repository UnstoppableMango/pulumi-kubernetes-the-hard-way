// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetesthehardway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

var _ = internal.GetEnvOrDefault

// Polyfill for `tls.CertRequestSubject`.
type CertRequestSubject struct {
	// Distinguished name: CN
	CommonName *string `pulumi:"commonName"`
	// Distinguished name: C
	Country *string `pulumi:"country"`
	// Distinguished name: L
	Locality *string `pulumi:"locality"`
	// Distinguished name: O
	Organization *string `pulumi:"organization"`
	// Distinguished name: OU
	OrganizationalUnit *string `pulumi:"organizationalUnit"`
	// Distinguished name: PC
	PostalCode *string `pulumi:"postalCode"`
	// Distinguished name: ST
	Province *string `pulumi:"province"`
	// Distinguished name: SERIALNUMBER
	SerialNumber *string `pulumi:"serialNumber"`
	// Distinguished name: STREET
	StreetAddresses []string `pulumi:"streetAddresses"`
}

// CertRequestSubjectInput is an input type that accepts CertRequestSubjectArgs and CertRequestSubjectOutput values.
// You can construct a concrete instance of `CertRequestSubjectInput` via:
//
//	CertRequestSubjectArgs{...}
type CertRequestSubjectInput interface {
	pulumi.Input

	ToCertRequestSubjectOutput() CertRequestSubjectOutput
	ToCertRequestSubjectOutputWithContext(context.Context) CertRequestSubjectOutput
}

// Polyfill for `tls.CertRequestSubject`.
type CertRequestSubjectArgs struct {
	// Distinguished name: CN
	CommonName pulumi.StringPtrInput `pulumi:"commonName"`
	// Distinguished name: C
	Country pulumi.StringPtrInput `pulumi:"country"`
	// Distinguished name: L
	Locality pulumi.StringPtrInput `pulumi:"locality"`
	// Distinguished name: O
	Organization pulumi.StringPtrInput `pulumi:"organization"`
	// Distinguished name: OU
	OrganizationalUnit pulumi.StringPtrInput `pulumi:"organizationalUnit"`
	// Distinguished name: PC
	PostalCode pulumi.StringPtrInput `pulumi:"postalCode"`
	// Distinguished name: ST
	Province pulumi.StringPtrInput `pulumi:"province"`
	// Distinguished name: SERIALNUMBER
	SerialNumber pulumi.StringPtrInput `pulumi:"serialNumber"`
	// Distinguished name: STREET
	StreetAddresses pulumi.StringArrayInput `pulumi:"streetAddresses"`
}

func (CertRequestSubjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertRequestSubject)(nil)).Elem()
}

func (i CertRequestSubjectArgs) ToCertRequestSubjectOutput() CertRequestSubjectOutput {
	return i.ToCertRequestSubjectOutputWithContext(context.Background())
}

func (i CertRequestSubjectArgs) ToCertRequestSubjectOutputWithContext(ctx context.Context) CertRequestSubjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertRequestSubjectOutput)
}

func (i CertRequestSubjectArgs) ToCertRequestSubjectPtrOutput() CertRequestSubjectPtrOutput {
	return i.ToCertRequestSubjectPtrOutputWithContext(context.Background())
}

func (i CertRequestSubjectArgs) ToCertRequestSubjectPtrOutputWithContext(ctx context.Context) CertRequestSubjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertRequestSubjectOutput).ToCertRequestSubjectPtrOutputWithContext(ctx)
}

// CertRequestSubjectPtrInput is an input type that accepts CertRequestSubjectArgs, CertRequestSubjectPtr and CertRequestSubjectPtrOutput values.
// You can construct a concrete instance of `CertRequestSubjectPtrInput` via:
//
//	        CertRequestSubjectArgs{...}
//
//	or:
//
//	        nil
type CertRequestSubjectPtrInput interface {
	pulumi.Input

	ToCertRequestSubjectPtrOutput() CertRequestSubjectPtrOutput
	ToCertRequestSubjectPtrOutputWithContext(context.Context) CertRequestSubjectPtrOutput
}

type certRequestSubjectPtrType CertRequestSubjectArgs

func CertRequestSubjectPtr(v *CertRequestSubjectArgs) CertRequestSubjectPtrInput {
	return (*certRequestSubjectPtrType)(v)
}

func (*certRequestSubjectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertRequestSubject)(nil)).Elem()
}

func (i *certRequestSubjectPtrType) ToCertRequestSubjectPtrOutput() CertRequestSubjectPtrOutput {
	return i.ToCertRequestSubjectPtrOutputWithContext(context.Background())
}

func (i *certRequestSubjectPtrType) ToCertRequestSubjectPtrOutputWithContext(ctx context.Context) CertRequestSubjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertRequestSubjectPtrOutput)
}

// Polyfill for `tls.CertRequestSubject`.
type CertRequestSubjectOutput struct{ *pulumi.OutputState }

func (CertRequestSubjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertRequestSubject)(nil)).Elem()
}

func (o CertRequestSubjectOutput) ToCertRequestSubjectOutput() CertRequestSubjectOutput {
	return o
}

func (o CertRequestSubjectOutput) ToCertRequestSubjectOutputWithContext(ctx context.Context) CertRequestSubjectOutput {
	return o
}

func (o CertRequestSubjectOutput) ToCertRequestSubjectPtrOutput() CertRequestSubjectPtrOutput {
	return o.ToCertRequestSubjectPtrOutputWithContext(context.Background())
}

func (o CertRequestSubjectOutput) ToCertRequestSubjectPtrOutputWithContext(ctx context.Context) CertRequestSubjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertRequestSubject) *CertRequestSubject {
		return &v
	}).(CertRequestSubjectPtrOutput)
}

// Distinguished name: CN
func (o CertRequestSubjectOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.CommonName }).(pulumi.StringPtrOutput)
}

// Distinguished name: C
func (o CertRequestSubjectOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.Country }).(pulumi.StringPtrOutput)
}

// Distinguished name: L
func (o CertRequestSubjectOutput) Locality() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.Locality }).(pulumi.StringPtrOutput)
}

// Distinguished name: O
func (o CertRequestSubjectOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

// Distinguished name: OU
func (o CertRequestSubjectOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.OrganizationalUnit }).(pulumi.StringPtrOutput)
}

// Distinguished name: PC
func (o CertRequestSubjectOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.PostalCode }).(pulumi.StringPtrOutput)
}

// Distinguished name: ST
func (o CertRequestSubjectOutput) Province() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.Province }).(pulumi.StringPtrOutput)
}

// Distinguished name: SERIALNUMBER
func (o CertRequestSubjectOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.SerialNumber }).(pulumi.StringPtrOutput)
}

// Distinguished name: STREET
func (o CertRequestSubjectOutput) StreetAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CertRequestSubject) []string { return v.StreetAddresses }).(pulumi.StringArrayOutput)
}

type CertRequestSubjectPtrOutput struct{ *pulumi.OutputState }

func (CertRequestSubjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertRequestSubject)(nil)).Elem()
}

func (o CertRequestSubjectPtrOutput) ToCertRequestSubjectPtrOutput() CertRequestSubjectPtrOutput {
	return o
}

func (o CertRequestSubjectPtrOutput) ToCertRequestSubjectPtrOutputWithContext(ctx context.Context) CertRequestSubjectPtrOutput {
	return o
}

func (o CertRequestSubjectPtrOutput) Elem() CertRequestSubjectOutput {
	return o.ApplyT(func(v *CertRequestSubject) CertRequestSubject {
		if v != nil {
			return *v
		}
		var ret CertRequestSubject
		return ret
	}).(CertRequestSubjectOutput)
}

// Distinguished name: CN
func (o CertRequestSubjectPtrOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.CommonName
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: C
func (o CertRequestSubjectPtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.Country
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: L
func (o CertRequestSubjectPtrOutput) Locality() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.Locality
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: O
func (o CertRequestSubjectPtrOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.Organization
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: OU
func (o CertRequestSubjectPtrOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationalUnit
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: PC
func (o CertRequestSubjectPtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.PostalCode
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: ST
func (o CertRequestSubjectPtrOutput) Province() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.Province
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: SERIALNUMBER
func (o CertRequestSubjectPtrOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: STREET
func (o CertRequestSubjectPtrOutput) StreetAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertRequestSubject) []string {
		if v == nil {
			return nil
		}
		return v.StreetAddresses
	}).(pulumi.StringArrayOutput)
}

// Node inputs for the PKI.
type ClusterPkiNode struct {
	// The IP address of the node.
	Ip   *string   `pulumi:"ip"`
	Role *NodeRole `pulumi:"role"`
}

// ClusterPkiNodeInput is an input type that accepts ClusterPkiNodeArgs and ClusterPkiNodeOutput values.
// You can construct a concrete instance of `ClusterPkiNodeInput` via:
//
//	ClusterPkiNodeArgs{...}
type ClusterPkiNodeInput interface {
	pulumi.Input

	ToClusterPkiNodeOutput() ClusterPkiNodeOutput
	ToClusterPkiNodeOutputWithContext(context.Context) ClusterPkiNodeOutput
}

// Node inputs for the PKI.
type ClusterPkiNodeArgs struct {
	// The IP address of the node.
	Ip   pulumi.StringPtrInput `pulumi:"ip"`
	Role NodeRolePtrInput      `pulumi:"role"`
}

func (ClusterPkiNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterPkiNode)(nil)).Elem()
}

func (i ClusterPkiNodeArgs) ToClusterPkiNodeOutput() ClusterPkiNodeOutput {
	return i.ToClusterPkiNodeOutputWithContext(context.Background())
}

func (i ClusterPkiNodeArgs) ToClusterPkiNodeOutputWithContext(ctx context.Context) ClusterPkiNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPkiNodeOutput)
}

// ClusterPkiNodeMapInput is an input type that accepts ClusterPkiNodeMap and ClusterPkiNodeMapOutput values.
// You can construct a concrete instance of `ClusterPkiNodeMapInput` via:
//
//	ClusterPkiNodeMap{ "key": ClusterPkiNodeArgs{...} }
type ClusterPkiNodeMapInput interface {
	pulumi.Input

	ToClusterPkiNodeMapOutput() ClusterPkiNodeMapOutput
	ToClusterPkiNodeMapOutputWithContext(context.Context) ClusterPkiNodeMapOutput
}

type ClusterPkiNodeMap map[string]ClusterPkiNodeInput

func (ClusterPkiNodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClusterPkiNode)(nil)).Elem()
}

func (i ClusterPkiNodeMap) ToClusterPkiNodeMapOutput() ClusterPkiNodeMapOutput {
	return i.ToClusterPkiNodeMapOutputWithContext(context.Background())
}

func (i ClusterPkiNodeMap) ToClusterPkiNodeMapOutputWithContext(ctx context.Context) ClusterPkiNodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPkiNodeMapOutput)
}

// Node inputs for the PKI.
type ClusterPkiNodeOutput struct{ *pulumi.OutputState }

func (ClusterPkiNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterPkiNode)(nil)).Elem()
}

func (o ClusterPkiNodeOutput) ToClusterPkiNodeOutput() ClusterPkiNodeOutput {
	return o
}

func (o ClusterPkiNodeOutput) ToClusterPkiNodeOutputWithContext(ctx context.Context) ClusterPkiNodeOutput {
	return o
}

// The IP address of the node.
func (o ClusterPkiNodeOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterPkiNode) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o ClusterPkiNodeOutput) Role() NodeRolePtrOutput {
	return o.ApplyT(func(v ClusterPkiNode) *NodeRole { return v.Role }).(NodeRolePtrOutput)
}

type ClusterPkiNodeMapOutput struct{ *pulumi.OutputState }

func (ClusterPkiNodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClusterPkiNode)(nil)).Elem()
}

func (o ClusterPkiNodeMapOutput) ToClusterPkiNodeMapOutput() ClusterPkiNodeMapOutput {
	return o
}

func (o ClusterPkiNodeMapOutput) ToClusterPkiNodeMapOutputWithContext(ctx context.Context) ClusterPkiNodeMapOutput {
	return o
}

func (o ClusterPkiNodeMapOutput) MapIndex(k pulumi.StringInput) ClusterPkiNodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClusterPkiNode {
		return vs[0].(map[string]ClusterPkiNode)[vs[1].(string)]
	}).(ClusterPkiNodeOutput)
}

// Instructions for how to connect to a remote endpoint. Polyfill for `command.ConnectionArgs`.
type Connection struct {
	// SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
	AgentSocketPath *string `pulumi:"agentSocketPath"`
	// Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
	DialErrorLimit *int `pulumi:"dialErrorLimit"`
	// The address of the resource to connect to.
	Host string `pulumi:"host"`
	// The password we should use for the connection.
	Password *string `pulumi:"password"`
	// Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
	PerDialTimeout *int `pulumi:"perDialTimeout"`
	// The port to connect to.
	Port *int `pulumi:"port"`
	// The contents of an SSH key to use for the connection. This takes preference over the password if provided.
	PrivateKey *string `pulumi:"privateKey"`
	// The password to use in case the private key is encrypted.
	PrivateKeyPassword *string `pulumi:"privateKeyPassword"`
	// The user that we should use for the connection.
	User *string `pulumi:"user"`
}

// Defaults sets the appropriate defaults for Connection
func (val *Connection) Defaults() *Connection {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DialErrorLimit == nil {
		dialErrorLimit_ := 10
		tmp.DialErrorLimit = &dialErrorLimit_
	}
	if tmp.PerDialTimeout == nil {
		perDialTimeout_ := 15
		tmp.PerDialTimeout = &perDialTimeout_
	}
	if tmp.Port == nil {
		port_ := 22
		tmp.Port = &port_
	}
	if tmp.User == nil {
		user_ := "root"
		tmp.User = &user_
	}
	return &tmp
}

// ConnectionInput is an input type that accepts ConnectionArgs and ConnectionOutput values.
// You can construct a concrete instance of `ConnectionInput` via:
//
//	ConnectionArgs{...}
type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(context.Context) ConnectionOutput
}

// Instructions for how to connect to a remote endpoint. Polyfill for `command.ConnectionArgs`.
type ConnectionArgs struct {
	// SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
	AgentSocketPath pulumi.StringPtrInput `pulumi:"agentSocketPath"`
	// Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
	DialErrorLimit pulumi.IntPtrInput `pulumi:"dialErrorLimit"`
	// The address of the resource to connect to.
	Host pulumi.StringInput `pulumi:"host"`
	// The password we should use for the connection.
	Password pulumi.StringPtrInput `pulumi:"password"`
	// Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
	PerDialTimeout pulumi.IntPtrInput `pulumi:"perDialTimeout"`
	// The port to connect to.
	Port pulumi.IntPtrInput `pulumi:"port"`
	// The contents of an SSH key to use for the connection. This takes preference over the password if provided.
	PrivateKey pulumi.StringPtrInput `pulumi:"privateKey"`
	// The password to use in case the private key is encrypted.
	PrivateKeyPassword pulumi.StringPtrInput `pulumi:"privateKeyPassword"`
	// The user that we should use for the connection.
	User pulumi.StringPtrInput `pulumi:"user"`
}

// Defaults sets the appropriate defaults for ConnectionArgs
func (val *ConnectionArgs) Defaults() *ConnectionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DialErrorLimit == nil {
		tmp.DialErrorLimit = pulumi.IntPtr(10)
	}
	if tmp.PerDialTimeout == nil {
		tmp.PerDialTimeout = pulumi.IntPtr(15)
	}
	if tmp.Port == nil {
		tmp.Port = pulumi.IntPtr(22)
	}
	if tmp.User == nil {
		tmp.User = pulumi.StringPtr("root")
	}
	return &tmp
}
func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil)).Elem()
}

func (i ConnectionArgs) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i ConnectionArgs) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

// Instructions for how to connect to a remote endpoint. Polyfill for `command.ConnectionArgs`.
type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

// SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
func (o ConnectionOutput) AgentSocketPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.AgentSocketPath }).(pulumi.StringPtrOutput)
}

// Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
func (o ConnectionOutput) DialErrorLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Connection) *int { return v.DialErrorLimit }).(pulumi.IntPtrOutput)
}

// The address of the resource to connect to.
func (o ConnectionOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v Connection) string { return v.Host }).(pulumi.StringOutput)
}

// The password we should use for the connection.
func (o ConnectionOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.Password }).(pulumi.StringPtrOutput)
}

// Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
func (o ConnectionOutput) PerDialTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Connection) *int { return v.PerDialTimeout }).(pulumi.IntPtrOutput)
}

// The port to connect to.
func (o ConnectionOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Connection) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// The contents of an SSH key to use for the connection. This takes preference over the password if provided.
func (o ConnectionOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

// The password to use in case the private key is encrypted.
func (o ConnectionOutput) PrivateKeyPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.PrivateKeyPassword }).(pulumi.StringPtrOutput)
}

// The user that we should use for the connection.
func (o ConnectionOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.User }).(pulumi.StringPtrOutput)
}

// A certificate and key pair.
type KeyPair struct {
	// The PEM encoded certificate data
	CertPem *string `pulumi:"certPem"`
	// The private key.
	Key *tls.PrivateKey `pulumi:"key"`
	// The PEM encoded private key data.
	PrivateKeyPem *string `pulumi:"privateKeyPem"`
	// The PEM encoded public key data.
	PublicKeyPem *string `pulumi:"publicKeyPem"`
}

// KeyPairInput is an input type that accepts KeyPairArgs and KeyPairOutput values.
// You can construct a concrete instance of `KeyPairInput` via:
//
//	KeyPairArgs{...}
type KeyPairInput interface {
	pulumi.Input

	ToKeyPairOutput() KeyPairOutput
	ToKeyPairOutputWithContext(context.Context) KeyPairOutput
}

// A certificate and key pair.
type KeyPairArgs struct {
	// The PEM encoded certificate data
	CertPem pulumi.StringPtrInput `pulumi:"certPem"`
	// The private key.
	Key tls.PrivateKeyInput `pulumi:"key"`
	// The PEM encoded private key data.
	PrivateKeyPem pulumi.StringPtrInput `pulumi:"privateKeyPem"`
	// The PEM encoded public key data.
	PublicKeyPem pulumi.StringPtrInput `pulumi:"publicKeyPem"`
}

func (KeyPairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPair)(nil)).Elem()
}

func (i KeyPairArgs) ToKeyPairOutput() KeyPairOutput {
	return i.ToKeyPairOutputWithContext(context.Background())
}

func (i KeyPairArgs) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairOutput)
}

// A certificate and key pair.
type KeyPairOutput struct{ *pulumi.OutputState }

func (KeyPairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPair)(nil)).Elem()
}

func (o KeyPairOutput) ToKeyPairOutput() KeyPairOutput {
	return o
}

func (o KeyPairOutput) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return o
}

// The PEM encoded certificate data
func (o KeyPairOutput) CertPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyPair) *string { return v.CertPem }).(pulumi.StringPtrOutput)
}

// The private key.
func (o KeyPairOutput) Key() tls.PrivateKeyOutput {
	return o.ApplyT(func(v KeyPair) *tls.PrivateKey { return v.Key }).(tls.PrivateKeyOutput)
}

// The PEM encoded private key data.
func (o KeyPairOutput) PrivateKeyPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyPair) *string { return v.PrivateKeyPem }).(pulumi.StringPtrOutput)
}

// The PEM encoded public key data.
func (o KeyPairOutput) PublicKeyPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyPair) *string { return v.PublicKeyPem }).(pulumi.StringPtrOutput)
}

// Polyfill for `pulumi.ComponentResourceOptions`.
type ResourceOptions struct {
	Parent interface{} `pulumi:"parent"`
}

// ResourceOptionsInput is an input type that accepts ResourceOptionsArgs and ResourceOptionsOutput values.
// You can construct a concrete instance of `ResourceOptionsInput` via:
//
//	ResourceOptionsArgs{...}
type ResourceOptionsInput interface {
	pulumi.Input

	ToResourceOptionsOutput() ResourceOptionsOutput
	ToResourceOptionsOutputWithContext(context.Context) ResourceOptionsOutput
}

// Polyfill for `pulumi.ComponentResourceOptions`.
type ResourceOptionsArgs struct {
	Parent pulumi.Input `pulumi:"parent"`
}

func (ResourceOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceOptions)(nil)).Elem()
}

func (i ResourceOptionsArgs) ToResourceOptionsOutput() ResourceOptionsOutput {
	return i.ToResourceOptionsOutputWithContext(context.Background())
}

func (i ResourceOptionsArgs) ToResourceOptionsOutputWithContext(ctx context.Context) ResourceOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOptionsOutput)
}

func (i ResourceOptionsArgs) ToResourceOptionsPtrOutput() ResourceOptionsPtrOutput {
	return i.ToResourceOptionsPtrOutputWithContext(context.Background())
}

func (i ResourceOptionsArgs) ToResourceOptionsPtrOutputWithContext(ctx context.Context) ResourceOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOptionsOutput).ToResourceOptionsPtrOutputWithContext(ctx)
}

// ResourceOptionsPtrInput is an input type that accepts ResourceOptionsArgs, ResourceOptionsPtr and ResourceOptionsPtrOutput values.
// You can construct a concrete instance of `ResourceOptionsPtrInput` via:
//
//	        ResourceOptionsArgs{...}
//
//	or:
//
//	        nil
type ResourceOptionsPtrInput interface {
	pulumi.Input

	ToResourceOptionsPtrOutput() ResourceOptionsPtrOutput
	ToResourceOptionsPtrOutputWithContext(context.Context) ResourceOptionsPtrOutput
}

type resourceOptionsPtrType ResourceOptionsArgs

func ResourceOptionsPtr(v *ResourceOptionsArgs) ResourceOptionsPtrInput {
	return (*resourceOptionsPtrType)(v)
}

func (*resourceOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceOptions)(nil)).Elem()
}

func (i *resourceOptionsPtrType) ToResourceOptionsPtrOutput() ResourceOptionsPtrOutput {
	return i.ToResourceOptionsPtrOutputWithContext(context.Background())
}

func (i *resourceOptionsPtrType) ToResourceOptionsPtrOutputWithContext(ctx context.Context) ResourceOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOptionsPtrOutput)
}

// Polyfill for `pulumi.ComponentResourceOptions`.
type ResourceOptionsOutput struct{ *pulumi.OutputState }

func (ResourceOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceOptions)(nil)).Elem()
}

func (o ResourceOptionsOutput) ToResourceOptionsOutput() ResourceOptionsOutput {
	return o
}

func (o ResourceOptionsOutput) ToResourceOptionsOutputWithContext(ctx context.Context) ResourceOptionsOutput {
	return o
}

func (o ResourceOptionsOutput) ToResourceOptionsPtrOutput() ResourceOptionsPtrOutput {
	return o.ToResourceOptionsPtrOutputWithContext(context.Background())
}

func (o ResourceOptionsOutput) ToResourceOptionsPtrOutputWithContext(ctx context.Context) ResourceOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceOptions) *ResourceOptions {
		return &v
	}).(ResourceOptionsPtrOutput)
}

func (o ResourceOptionsOutput) Parent() pulumi.AnyOutput {
	return o.ApplyT(func(v ResourceOptions) interface{} { return v.Parent }).(pulumi.AnyOutput)
}

type ResourceOptionsPtrOutput struct{ *pulumi.OutputState }

func (ResourceOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceOptions)(nil)).Elem()
}

func (o ResourceOptionsPtrOutput) ToResourceOptionsPtrOutput() ResourceOptionsPtrOutput {
	return o
}

func (o ResourceOptionsPtrOutput) ToResourceOptionsPtrOutputWithContext(ctx context.Context) ResourceOptionsPtrOutput {
	return o
}

func (o ResourceOptionsPtrOutput) Elem() ResourceOptionsOutput {
	return o.ApplyT(func(v *ResourceOptions) ResourceOptions {
		if v != nil {
			return *v
		}
		var ret ResourceOptions
		return ret
	}).(ResourceOptionsOutput)
}

func (o ResourceOptionsPtrOutput) Parent() pulumi.AnyOutput {
	return o.ApplyT(func(v *ResourceOptions) interface{} {
		if v == nil {
			return nil
		}
		return v.Parent
	}).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertRequestSubjectInput)(nil)).Elem(), CertRequestSubjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertRequestSubjectPtrInput)(nil)).Elem(), CertRequestSubjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPkiNodeInput)(nil)).Elem(), ClusterPkiNodeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPkiNodeMapInput)(nil)).Elem(), ClusterPkiNodeMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), ConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairInput)(nil)).Elem(), KeyPairArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceOptionsInput)(nil)).Elem(), ResourceOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceOptionsPtrInput)(nil)).Elem(), ResourceOptionsArgs{})
	pulumi.RegisterOutputType(CertRequestSubjectOutput{})
	pulumi.RegisterOutputType(CertRequestSubjectPtrOutput{})
	pulumi.RegisterOutputType(ClusterPkiNodeOutput{})
	pulumi.RegisterOutputType(ClusterPkiNodeMapOutput{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(KeyPairOutput{})
	pulumi.RegisterOutputType(ResourceOptionsOutput{})
	pulumi.RegisterOutputType(ResourceOptionsPtrOutput{})
}
