# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .certificate import *
from .cluster_pki import *
from .encryption_key import *
from .install_cert import *
from .install_control_plane import *
from .install_key import *
from .new_certificate import *
from .provider import *
from .remote_file import *
from .root_ca import *
from ._inputs import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_kubernetes_the_hard_way.tools as __tools
    tools = __tools
else:
    tools = _utilities.lazy_import('pulumi_kubernetes_the_hard_way.tools')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "kubernetes-the-hard-way",
  "mod": "index",
  "fqn": "pulumi_kubernetes_the_hard_way",
  "classes": {
   "kubernetes-the-hard-way:index:Certificate": "Certificate",
   "kubernetes-the-hard-way:index:ClusterPki": "ClusterPki",
   "kubernetes-the-hard-way:index:EncryptionKey": "EncryptionKey",
   "kubernetes-the-hard-way:index:RemoteFile": "RemoteFile",
   "kubernetes-the-hard-way:index:RootCa": "RootCa"
  }
 },
 {
  "pkg": "kubernetes-the-hard-way",
  "mod": "tools",
  "fqn": "pulumi_kubernetes_the_hard_way.tools",
  "classes": {
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
