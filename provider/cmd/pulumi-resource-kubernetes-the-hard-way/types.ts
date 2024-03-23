export const Algorithm = {
  RSA: 'RSA',
  ECDSA: 'ECDSA',
  ED25519: 'ED25519',
} as const;

export type Algorithm = (typeof Algorithm)[keyof typeof Algorithm];

export const AllowedUsage = {
  Cert_signing: 'cert_signing',
  Client_auth: 'client_auth',
  Crl_signing: 'crl_signing',
  Digital_signature: 'digital_signature',
  Key_encipherment: 'key_encipherment',
  Server_auth: 'server_auth',
} as const;

export type AllowedUsage = (typeof AllowedUsage)[keyof typeof AllowedUsage];

export const EcdsaCurve = {
  P224: 'P224',
  P256: 'P256',
  P384: 'P384',
  P521: 'P521',
} as const;

export type EcdsaCurve = (typeof EcdsaCurve)[keyof typeof EcdsaCurve];
