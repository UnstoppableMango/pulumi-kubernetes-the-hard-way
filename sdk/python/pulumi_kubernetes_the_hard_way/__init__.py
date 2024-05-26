# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_kubernetes_the_hard_way.config as __config
    config = __config
    import pulumi_kubernetes_the_hard_way.remote as __remote
    remote = __remote
    import pulumi_kubernetes_the_hard_way.tls as __tls
    tls = __tls
    import pulumi_kubernetes_the_hard_way.tools as __tools
    tools = __tools
else:
    config = _utilities.lazy_import('pulumi_kubernetes_the_hard_way.config')
    remote = _utilities.lazy_import('pulumi_kubernetes_the_hard_way.remote')
    tls = _utilities.lazy_import('pulumi_kubernetes_the_hard_way.tls')
    tools = _utilities.lazy_import('pulumi_kubernetes_the_hard_way.tools')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "kubernetes-the-hard-way",
  "mod": "config",
  "fqn": "pulumi_kubernetes_the_hard_way.config",
  "classes": {
   "kubernetes-the-hard-way:config:CniBridgePluginConfiguration": "CniBridgePluginConfiguration",
   "kubernetes-the-hard-way:config:CniLoopbackPluginConfiguration": "CniLoopbackPluginConfiguration",
   "kubernetes-the-hard-way:config:ContainerdConfiguration": "ContainerdConfiguration",
   "kubernetes-the-hard-way:config:KubeProxyConfiguration": "KubeProxyConfiguration",
   "kubernetes-the-hard-way:config:KubeVipManifest": "KubeVipManifest",
   "kubernetes-the-hard-way:config:KubeletConfiguration": "KubeletConfiguration"
  }
 },
 {
  "pkg": "kubernetes-the-hard-way",
  "mod": "remote",
  "fqn": "pulumi_kubernetes_the_hard_way.remote",
  "classes": {
   "kubernetes-the-hard-way:remote:CniPluginsInstall": "CniPluginsInstall",
   "kubernetes-the-hard-way:remote:ContainerdInstall": "ContainerdInstall",
   "kubernetes-the-hard-way:remote:ContainerdService": "ContainerdService",
   "kubernetes-the-hard-way:remote:ControlPlaneNode": "ControlPlaneNode",
   "kubernetes-the-hard-way:remote:CrictlInstall": "CrictlInstall",
   "kubernetes-the-hard-way:remote:Download": "Download",
   "kubernetes-the-hard-way:remote:EtcdCluster": "EtcdCluster",
   "kubernetes-the-hard-way:remote:EtcdConfiguration": "EtcdConfiguration",
   "kubernetes-the-hard-way:remote:EtcdInstall": "EtcdInstall",
   "kubernetes-the-hard-way:remote:EtcdService": "EtcdService",
   "kubernetes-the-hard-way:remote:File": "File",
   "kubernetes-the-hard-way:remote:KubeApiServerInstall": "KubeApiServerInstall",
   "kubernetes-the-hard-way:remote:KubeControllerManagerInstall": "KubeControllerManagerInstall",
   "kubernetes-the-hard-way:remote:KubeProxyInstall": "KubeProxyInstall",
   "kubernetes-the-hard-way:remote:KubeProxyService": "KubeProxyService",
   "kubernetes-the-hard-way:remote:KubeSchedulerInstall": "KubeSchedulerInstall",
   "kubernetes-the-hard-way:remote:KubectlInstall": "KubectlInstall",
   "kubernetes-the-hard-way:remote:KubeletInstall": "KubeletInstall",
   "kubernetes-the-hard-way:remote:KubeletService": "KubeletService",
   "kubernetes-the-hard-way:remote:ProvisionEtcd": "ProvisionEtcd",
   "kubernetes-the-hard-way:remote:RuncInstall": "RuncInstall",
   "kubernetes-the-hard-way:remote:StartContainerd": "StartContainerd",
   "kubernetes-the-hard-way:remote:StartEtcd": "StartEtcd",
   "kubernetes-the-hard-way:remote:StartKubeProxy": "StartKubeProxy",
   "kubernetes-the-hard-way:remote:StartKubelet": "StartKubelet",
   "kubernetes-the-hard-way:remote:StaticPod": "StaticPod",
   "kubernetes-the-hard-way:remote:SystemdService": "SystemdService",
   "kubernetes-the-hard-way:remote:WorkerNode": "WorkerNode",
   "kubernetes-the-hard-way:remote:WorkerPreRequisites": "WorkerPreRequisites"
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
   "kubernetes-the-hard-way:tools:Chmod": "Chmod",
   "kubernetes-the-hard-way:tools:Etcdctl": "Etcdctl",
   "kubernetes-the-hard-way:tools:Hostnamectl": "Hostnamectl",
   "kubernetes-the-hard-way:tools:Mkdir": "Mkdir",
   "kubernetes-the-hard-way:tools:Mktemp": "Mktemp",
   "kubernetes-the-hard-way:tools:Mv": "Mv",
   "kubernetes-the-hard-way:tools:Rm": "Rm",
   "kubernetes-the-hard-way:tools:Sed": "Sed",
   "kubernetes-the-hard-way:tools:Systemctl": "Systemctl",
   "kubernetes-the-hard-way:tools:Tar": "Tar",
   "kubernetes-the-hard-way:tools:Tee": "Tee",
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
