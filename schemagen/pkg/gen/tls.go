package gen

import (
	"maps"
	"slices"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

const tlsMod = "kubernetes-the-hard-way:tls:"

func generateTls(randomSpec, tlsSpec schema.PackageSpec) schema.PackageSpec {
	types := map[string]schema.ComplexTypeSpec{
		tlsMod + "Algorithm": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "Private key algorithm.",
				Type:        "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "RSA"},
				{Value: "ECDSA"},
				{Value: "ED25519"},
			},
		},
		tlsMod + "AllowedUsage": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "Certificate allowed usage",
				Type:        "string",
			},
			Enum: []schema.EnumValueSpec{
				{Name: "certSigning", Value: "cert_signing"},
				{Name: "clientAuth", Value: "client_auth"},
				{Name: "crlSigning", Value: "crl_signing"},
				{Name: "digitalSignature", Value: "digital_signature"},
				{Name: "keyEncipherment", Value: "key_encipherment"},
				{Name: "serverAuth", Value: "server_auth"},
			},
		},
		tlsMod + "Bundle": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "A CA + Cert + Key bundle",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"caPem": {
						Description: "The PEM encoded certificate authority data.",
						TypeSpec:    typeSpecs.String,
					},
					"certPem": {
						Description: "The PEM encoded certificate data.",
						TypeSpec:    typeSpecs.String,
					},
					"keyPem": { // TODO: Verify public/private
						Description: "The PEM encoded private key data",
						TypeSpec:    typeSpecs.String,
					},
				},
				Required: []string{
					"caPem",
					"certPem",
					"keyPem",
				},
			},
		},
		tlsMod + "ClusterPkiNode": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "TODO",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"ip": {
						Description: "The IP address of the node",
						TypeSpec:    typeSpecs.String,
					},
					"role": {
						Description: "The role a node should be configured for",
						TypeSpec: schema.TypeSpec{
							Ref: localType("NodeRole", "tls"),
						},
					},
				},
			},
		},
		tlsMod + "EcdsaCurve": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "ECDSA algorithm curve",
				Type:        "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "P224"},
				{Value: "P256"},
				{Value: "P384"},
				{Value: "P521"},
			},
		},
		tlsMod + "KeyPair": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "A certificate and key pair",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"cert": {
						Description: "The certificate resource.",
						TypeSpec: schema.TypeSpec{
							OneOf: []schema.TypeSpec{
								{Ref: refResource(tlsSpec, "LocallySignedCert", "index", "locallySignedCert")},
								{Ref: refResource(tlsSpec, "SelfSignedCert", "index", "selfSignedCert")},
							},
						},
					},
					"certPem": {
						Description: "The PEM encoded certificate data.",
						TypeSpec:    typeSpecs.String,
					},
					"key": {
						Description: "The private key resource.",
						TypeSpec: schema.TypeSpec{
							Ref: refResource(tlsSpec, "PrivateKey", "index", "privateKey"),
						},
					},
					"privateKeyPem": {
						Description: "The PEM encoded private key data.",
						TypeSpec:    typeSpecs.String,
					},
					"publicKeyPem": {
						Description: "The PEM encoded public key data.",
						TypeSpec:    typeSpecs.String,
					},
				},
				Required: []string{
					"cert",
					"certPem",
					"key",
					"privateKeyPem",
					"publicKeyPem",
				},
			},
		},
		tlsMod + "NodeRole": { // TODO: Does this belong in the TLS module?
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The role a node will play in the final cluster",
				Type:        "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "controlplane"},
				{Value: "worker"},
			},
		},
	}

	functions := map[string]schema.FunctionSpec{}
	resources := map[string]schema.ResourceSpec{
		tlsMod + "Certificate": generateCertificate(tlsSpec),
	}

	return schema.PackageSpec{
		Functions: functions,
		Resources: resources,
		Types:     types,
	}
}

func generateCertificate(tlsSpec schema.PackageSpec) schema.ResourceSpec {
	locallySignedCert := tlsSpec.Resources["tls:index/locallySignedCert:LocallySignedCert"]
	privateKey := tlsSpec.Resources["tls:index/privateKey:PrivateKey"]

	overrides := map[string]schema.PropertySpec{
		"algorithm": {
			Description: "Name of the algorithm to use when generating the private key.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Algorithm", "tls")},
		},
		"allowedUses": {
			Description: "List of key usages allowed for the issued certificate.",
			TypeSpec: schema.TypeSpec{
				Type:  "array",
				Items: &schema.TypeSpec{Ref: localType("AllowedUsage", "tls")},
			},
		},
	}

	inputs := map[string]schema.PropertySpec{}
	maps.Copy(inputs, locallySignedCert.InputProperties)
	maps.Copy(inputs, privateKey.InputProperties)
	maps.Copy(inputs, overrides)

	requiredInputs := slices.Concat(
		locallySignedCert.RequiredInputs,
		privateKey.RequiredInputs,
	)

	outputs := map[string]schema.PropertySpec{
		"cert": {
			Description: "The certificate",
			TypeSpec: schema.TypeSpec{
				Ref: refResource(tlsSpec, "LocallySignedCert", "index", "locallySignedCert"),
			},
		},
		"key": {
			Description: "The private key",
			TypeSpec: schema.TypeSpec{
				Ref: refResource(tlsSpec, "PrivateKey", "index", "privateKey"),
			},
		},
	}
	maps.Copy(outputs, inputs)
	maps.Copy(outputs, locallySignedCert.Properties)
	maps.Copy(outputs, privateKey.Properties)
	maps.Copy(outputs, overrides)

	requiredOutputs := slices.Concat(
		locallySignedCert.Required,
		privateKey.Required,
		[]string{"cert", "key"},
	)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "A certificate key pair.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
