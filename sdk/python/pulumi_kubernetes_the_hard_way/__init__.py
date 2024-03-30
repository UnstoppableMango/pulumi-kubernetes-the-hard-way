# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .install_cert import *
from .install_control_plane import *
from .install_key import *
from .new_certificate import *
from .provider import *
from ._inputs import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_kubernetes_the_hard_way.remote as __remote
    remote = __remote
    import pulumi_kubernetes_the_hard_way.tls as __tls
    tls = __tls
    import pulumi_kubernetes_the_hard_way.tools as __tools
    tools = __tools
else:
    remote = _utilities.lazy_import('pulumi_kubernetes_the_hard_way.remote')
    tls = _utilities.lazy_import('pulumi_kubernetes_the_hard_way.tls')
    tools = _utilities.lazy_import('pulumi_kubernetes_the_hard_way.tools')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "kubernetes-the-hard-way",
  "mod": "remote",
  "fqn": "pulumi_kubernetes_the_hard_way.remote",
  "classes": {
   "kubernetes-the-hard-way:remote:Download": "Download",
   "kubernetes-the-hard-way:remote:File": "File"
  }
 },
 {
  "pkg": "kubernetes-the-hard-way",
  "mod": "tls",
  "fqn": "pulumi_kubernetes_the_hard_way.tls",
  "classes": {
   "kubernetes-the-hard-way:tls:Certificate": "Certificate",
   "kubernetes-the-hard-way:tls:ClusterPki": "ClusterPki",
   "kubernetes-the-hard-way:tls:EncryptionKey": "EncryptionKey",
   "kubernetes-the-hard-way:tls:RootCa": "RootCa"
  }
 },
 {
  "pkg": "kubernetes-the-hard-way",
  "mod": "tools",
  "fqn": "pulumi_kubernetes_the_hard_way.tools",
  "classes": {
   "kubernetes-the-hard-way:tools:Mkdir": "Mkdir",
   "kubernetes-the-hard-way:tools:Tar": "Tar",
   "kubernetes-the-hard-way:tools:Wget": "Wget"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "kubernetes-the-hard-way",
  "token": "pulumi:providers:kubernetes-the-hard-way",
  "fqn": "pulumi_kubernetes_the_hard_way",
  "class": "Provider"
 }
]
"""
)
