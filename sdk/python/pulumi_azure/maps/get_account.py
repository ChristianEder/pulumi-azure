# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetAccountResult:
    """
    A collection of values returned by getAccount.
    """
    def __init__(__self__, name=None, primary_access_key=None, resource_group_name=None, secondary_access_key=None, sku_name=None, tags=None, x_ms_client_id=None, id=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if primary_access_key and not isinstance(primary_access_key, str):
            raise TypeError("Expected argument 'primary_access_key' to be a str")
        __self__.primary_access_key = primary_access_key
        """
        The primary key used to authenticate and authorize access to the Maps REST APIs.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if secondary_access_key and not isinstance(secondary_access_key, str):
            raise TypeError("Expected argument 'secondary_access_key' to be a str")
        __self__.secondary_access_key = secondary_access_key
        """
        The primary key used to authenticate and authorize access to the Maps REST APIs. The second key is given to provide seamless key regeneration.
        """
        if sku_name and not isinstance(sku_name, str):
            raise TypeError("Expected argument 'sku_name' to be a str")
        __self__.sku_name = sku_name
        """
        The sku of the Azure Maps Account.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if x_ms_client_id and not isinstance(x_ms_client_id, str):
            raise TypeError("Expected argument 'x_ms_client_id' to be a str")
        __self__.x_ms_client_id = x_ms_client_id
        """
        A unique identifier for the Maps Account.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetAccountResult(GetAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountResult(
            name=self.name,
            primary_access_key=self.primary_access_key,
            resource_group_name=self.resource_group_name,
            secondary_access_key=self.secondary_access_key,
            sku_name=self.sku_name,
            tags=self.tags,
            x_ms_client_id=self.x_ms_client_id,
            id=self.id)

def get_account(name=None,resource_group_name=None,tags=None,opts=None):
    """
    Use this data source to access information about an existing Azure Maps Account.
    
    :param str name: Specifies the name of the Maps Account.
    :param str resource_group_name: Specifies the name of the Resource Group in which the Maps Account is located.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/maps_account.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:maps/getAccount:getAccount', __args__, opts=opts).value

    return AwaitableGetAccountResult(
        name=__ret__.get('name'),
        primary_access_key=__ret__.get('primaryAccessKey'),
        resource_group_name=__ret__.get('resourceGroupName'),
        secondary_access_key=__ret__.get('secondaryAccessKey'),
        sku_name=__ret__.get('skuName'),
        tags=__ret__.get('tags'),
        x_ms_client_id=__ret__.get('xMsClientId'),
        id=__ret__.get('id'))