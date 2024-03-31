// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetesthehardway

import (
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

// Polyfill for `pulumi.ComponentResourceOptions`.
type ResourceOptions struct {
	Parent interface{} `pulumi:"parent"`
}

func init() {
}
