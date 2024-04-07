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
		tlsMod + "Certificate":   generateCertificate(tlsSpec),
		tlsMod + "ClusterPki":    generateClusterPki(),
		tlsMod + "EncryptionKey": generateEncryptionKey(randomSpec),
		tlsMod + "RootCa":        generateRootCa(tlsSpec),
	}

	return schema.PackageSpec{
		Functions: functions,
		Resources: resources,
		Types:     types,
	}
}

func generateCertificate(tlsSpec schema.PackageSpec) schema.ResourceSpec {
	certRequest := tlsSpec.Resources["tls:index/certRequest:CertRequest"]
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

	// TODO: There are a couple of these that we don't need
	// TODO: Need to rename some refs like `subject`
	inputs := map[string]schema.PropertySpec{}
	maps.Copy(inputs, certRequest.InputProperties)
	maps.Copy(inputs, locallySignedCert.InputProperties)
	maps.Copy(inputs, privateKey.InputProperties)
	maps.Copy(inputs, overrides)

	requiredInputs := []string{
		"algorithm",
		"allowedUses",
		"caCertPem",
		"caPrivateKeyPem",
		"validityPeriodHours",
	}

	outputs := map[string]schema.PropertySpec{
		"cert": {
			Description: "The certificate.",
			TypeSpec: schema.TypeSpec{
				Ref: refResource(tlsSpec, "LocallySignedCert", "index", "locallySignedCert"),
			},
		},
		"csr": {
			Description: "The certificate signing request.",
			TypeSpec: schema.TypeSpec{
				Ref: refResource(tlsSpec, "CertRequest", "index", "certRequest"),
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
	maps.Copy(outputs, certRequest.Properties)
	maps.Copy(outputs, locallySignedCert.Properties)
	maps.Copy(outputs, privateKey.Properties)
	maps.Copy(outputs, overrides)

	requiredOutputs := slices.Concat(
		certRequest.Required,
		locallySignedCert.Required,
		privateKey.Required,
		[]string{"cert", "csr", "key"},
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

func generateClusterPki() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"algorithm": {
			Description: "Name of the algorithm to use when generating the private key.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Algorithm", "tls")},
			Default:     "RSA",
		},
		"clusterName": {
			Description: "A name to use for the cluster",
			TypeSpec:    typeSpecs.String,
		},
		"ecdsaCurve": {
			Description: "When `algorithm` is `ECDSA`, the name of the elliptic curve to use.",
			TypeSpec:    schema.TypeSpec{Ref: localType("EcdsaCurve", "tls")},
		},
		"nodes": {
			Description: "Map of node name to node configuration",
			TypeSpec: schema.TypeSpec{
				Type: "object",
				AdditionalProperties: &schema.TypeSpec{
					Ref: localType("ClusterPkiNode", "tls"),
				},
			},
		},
		"publicIp": {
			Description: "Publicly accessible IP address.",
			TypeSpec:    typeSpecs.String,
		},
		"rsaBits": {
			Description: "When `algorithm` is `RSA`, the size of the generated RSA key, in bits.",
			TypeSpec:    typeSpecs.Integer,
			Default:     2048,
		},
		"validityPeriodHours": {
			Description: "Number of hours, after initial issuing, that the certificate will remain valid.",
			TypeSpec:    typeSpecs.Integer,
			Default:     8076,
		},
	}

	requiredInputs := []string{
		"clusterName",
		"nodes",
		"publicIp",
	}

	outputs := map[string]schema.PropertySpec{
		"admin": {
			Description: "The admin certificate.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Certificate", "tls")},
		},
		"ca": {
			Description: "The cluster certificate authority.",
			TypeSpec:    schema.TypeSpec{Ref: localType("RootCa", "tls")},
		},
		"controllerManager": {
			Description: "The controller manager certificate.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Certificate", "tls")},
		},
		"kubelet": {
			Description: "Map of node name to kubelet certificate.",
			TypeSpec: schema.TypeSpec{
				Type: "object",
				AdditionalProperties: &schema.TypeSpec{
					Ref: localType("Certificate", "tls"),
				},
			},
		},
		"kubeProxy": {
			Description: "The kube proxy certificate.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Certificate", "tls")},
		},
		"kubernetes": {
			Description: "The kubernetes certificate.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Certificate", "tls")},
		},
		"kubeScheduler": {
			Description: "The kube scheduler certificate.",
			TypeSpec:    schema.TypeSpec{Ref: localType("Certificate", "tls")},
		},
		"serviceAccounts": {
			Description: "The service accounts certificate",
			TypeSpec:    schema.TypeSpec{Ref: localType("Certificate", "tls")},
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(
		[]string{
			"admin",
			"ca",
			"controllerManager",
			"kubelet",
			"kubeProxy",
			"kubernetes",
			"kubeScheduler",
			"publicIp",
			"serviceAccounts",
			"validityPeriodHours",
		},
		requiredInputs,
	)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "The private key infrastructure for a cluster",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
		Methods: map[string]string{
			"getKubeconfig": tlsMod + "ClusterPki/getKubeconfig",
		},
	}
}

func generateEncryptionKey(randomSpec schema.PackageSpec) schema.ResourceSpec {
	randomKey := randomSpec.Resources["random:index/randomBytes:RandomBytes"]

	inputs := map[string]schema.PropertySpec{
		// Consider renaming to length
		"bytes": randomKey.InputProperties["length"],
	}

	outputs := map[string]schema.PropertySpec{
		"config": {
			Description: "The generated `v1/EncryptionConfig`.",
			TypeSpec:    typeSpecs.String,
		},
		"key": {
			Description: "The generated random key.",
			TypeSpec:    schema.TypeSpec{Ref: refResource(randomSpec, "index", "randomBytes")},
		},
	}
	maps.Copy(outputs, inputs)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "A cluster encryption key.",
			Properties:  outputs,
			Required:    []string{"bytes", "config", "key"},
		},
		InputProperties: inputs,
		RequiredInputs:  []string{"bytes"},
	}
}

func generateRootCa(tlsSpec schema.PackageSpec) schema.ResourceSpec {
	locallySignedCert := tlsSpec.Resources["tls:index/selfSignedCert:SelfSignedCert"]
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

	// TODO: There are a few of these we don't need.
	// TODO: Need to rename some refs like `subject`
	inputs := map[string]schema.PropertySpec{}
	maps.Copy(inputs, locallySignedCert.InputProperties)
	maps.Copy(inputs, privateKey.InputProperties)
	maps.Copy(inputs, overrides)

	requiredInputs := []string{"validityPeriodHours"}

	outputs := map[string]schema.PropertySpec{
		"cert": {
			Description: "The certificate authority certificate.",
			TypeSpec:    schema.TypeSpec{Ref: refResource(tlsSpec, "SelfSignedCertificate", "index/selfSignedCertificate")},
		},
		"key": {
			Description: "The certificate authority key.",
			TypeSpec:    schema.TypeSpec{Ref: refResource(tlsSpec, "PrivateKey", "index", "privateKey")},
		},
	}
	maps.Copy(outputs, locallySignedCert.Properties)
	maps.Copy(outputs, privateKey.Properties)
	maps.Copy(outputs, inputs)

	requiredOutputs := []string{"cert", "key"}

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Root certificate authority for a cluster.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
