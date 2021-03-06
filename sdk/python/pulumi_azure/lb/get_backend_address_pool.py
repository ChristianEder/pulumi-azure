# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetBackendAddressPoolResult:
    """
    A collection of values returned by getBackendAddressPool.
    """
    def __init__(__self__, backend_ip_configurations=None, id=None, loadbalancer_id=None, name=None):
        if backend_ip_configurations and not isinstance(backend_ip_configurations, list):
            raise TypeError("Expected argument 'backend_ip_configurations' to be a list")
        __self__.backend_ip_configurations = backend_ip_configurations
        """
        An array of references to IP addresses defined in network interfaces.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if loadbalancer_id and not isinstance(loadbalancer_id, str):
            raise TypeError("Expected argument 'loadbalancer_id' to be a str")
        __self__.loadbalancer_id = loadbalancer_id
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the Backend Address Pool.
        """
class AwaitableGetBackendAddressPoolResult(GetBackendAddressPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackendAddressPoolResult(
            backend_ip_configurations=self.backend_ip_configurations,
            id=self.id,
            loadbalancer_id=self.loadbalancer_id,
            name=self.name)

def get_backend_address_pool(loadbalancer_id=None,name=None,opts=None):
    """
    Use this data source to access information about an existing Load Balancer's Backend Address Pool.




    :param str loadbalancer_id: The ID of the Load Balancer in which the Backend Address Pool exists.
    :param str name: Specifies the name of the Backend Address Pool.
    """
    __args__ = dict()


    __args__['loadbalancerId'] = loadbalancer_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:lb/getBackendAddressPool:getBackendAddressPool', __args__, opts=opts).value

    return AwaitableGetBackendAddressPoolResult(
        backend_ip_configurations=__ret__.get('backendIpConfigurations'),
        id=__ret__.get('id'),
        loadbalancer_id=__ret__.get('loadbalancerId'),
        name=__ret__.get('name'))
