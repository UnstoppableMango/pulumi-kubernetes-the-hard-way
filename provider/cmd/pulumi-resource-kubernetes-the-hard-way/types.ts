export const Algorithm = {
  RSA: 'RSA',
  ECDSA: 'ECDSA',
  ED25519: 'ED25519',
} as const;

export type Algorithm = (typeof Algorithm)[keyof typeof Algorithm];

export const AllowedUsage = {
  CertSigning: 'cert_signing',
  ClientAuth: 'client_auth',
  CrlSigning: 'crl_signing',
  DigitalSignature: 'digital_signature',
  KeyEncipherment: 'key_encipherment',
  ServerAuth: 'server_auth',
} as const;

export type AllowedUsage = (typeof AllowedUsage)[keyof typeof AllowedUsage];

export const EcdsaCurve = {
  P224: 'P224',
  P256: 'P256',
  P384: 'P384',
  P521: 'P521',
} as const;

export type EcdsaCurve = (typeof EcdsaCurve)[keyof typeof EcdsaCurve];

export const KubeconfigType = {
  admin: 'admin',
  kubeControllerManager: 'kube-controller-manager',
  kubeProxy: 'kube-proxy',
  kubeScheduler: 'kube-scheduler',
  worker: 'worker',
} as const;

export type KubeconfigType = (typeof KubeconfigType)[keyof typeof KubeconfigType];
