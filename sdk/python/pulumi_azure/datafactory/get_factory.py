# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetFactoryResult:
    """
    A collection of values returned by getFactory.
    """
    def __init__(__self__, github_configurations=None, id=None, identities=None, location=None, name=None, resource_group_name=None, tags=None, vsts_configurations=None):
        if github_configurations and not isinstance(github_configurations, list):
            raise TypeError("Expected argument 'github_configurations' to be a list")
        __self__.github_configurations = github_configurations
        """
        A `github_configuration` block as defined below.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if identities and not isinstance(identities, list):
            raise TypeError("Expected argument 'identities' to be a list")
        __self__.identities = identities
        """
        An `identity` block as defined below.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The Azure location where the resource exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the resource.
        ---
        """
        if vsts_configurations and not isinstance(vsts_configurations, list):
            raise TypeError("Expected argument 'vsts_configurations' to be a list")
        __self__.vsts_configurations = vsts_configurations
        """
        A `vsts_configuration` block as defined below.
        """
class AwaitableGetFactoryResult(GetFactoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFactoryResult(
            github_configurations=self.github_configurations,
            id=self.id,
            identities=self.identities,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags,
            vsts_configurations=self.vsts_configurations)

def get_factory(name=None,resource_group_name=None,opts=None):
    """
    Use this data source to access information about an existing Azure Data Factory (Version 2).




    :param str name: Specifies the name of the Data Factory to retrieve information about. 
    :param str resource_group_name: The name of the resource group where the Data Factory exists.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:datafactory/getFactory:getFactory', __args__, opts=opts).value

    return AwaitableGetFactoryResult(
        github_configurations=__ret__.get('githubConfigurations'),
        id=__ret__.get('id'),
        identities=__ret__.get('identities'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'),
        tags=__ret__.get('tags'),
        vsts_configurations=__ret__.get('vstsConfigurations'))
