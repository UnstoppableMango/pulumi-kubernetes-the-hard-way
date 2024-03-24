"""A Python Pulumi program"""

import pulumi_kubernetes_the_hard_way as kthw

ca = kthw.RootCa("simple",
                 algorithm=kthw.Algorithm.RSA,
                 validity_period_hours=256)

cert = kthw.Certificate("simple",
                        algorithm=kthw.Algorithm.RSA,
                        allowed_uses=[kthw.AllowedUsage.CERT_SIGNING],
                        ca_cert_pem=ca.cert_pem,
                        ca_private_key_pem=ca.private_key_pem,
                        validity_period_hours=256)
