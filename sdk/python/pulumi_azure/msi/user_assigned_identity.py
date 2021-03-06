# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

warnings.warn("azure.UserAssignedIdentity has been deprecated in favour of azure.UserAssignedIdentity", DeprecationWarning)
class UserAssignedIdentity(pulumi.CustomResource):
    client_id: pulumi.Output[str]
    """
    Client ID associated with the user assigned identity.
    """
    location: pulumi.Output[str]
    """
    The location/region where the user assigned identity is
    created.
    """
    name: pulumi.Output[str]
    """
    The name of the user assigned identity. Changing this forces a
    new identity to be created.
    """
    principal_id: pulumi.Output[str]
    """
    Service Principal ID associated with the user assigned identity.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to
    create the user assigned identity.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    warnings.warn("azure.UserAssignedIdentity has been deprecated in favour of azure.UserAssignedIdentity", DeprecationWarning)
    def __init__(__self__, resource_name, opts=None, location=None, name=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a user assigned identity.



        Deprecated: azure.UserAssignedIdentity has been deprecated in favour of azure.UserAssignedIdentity

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location/region where the user assigned identity is
               created.
        :param pulumi.Input[str] name: The name of the user assigned identity. Changing this forces a
               new identity to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the user assigned identity.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        pulumi.log.warn("UserAssignedIdentity is deprecated: azure.UserAssignedIdentity has been deprecated in favour of azure.UserAssignedIdentity")
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['client_id'] = None
            __props__['principal_id'] = None
        super(UserAssignedIdentity, __self__).__init__(
            'azure:msi/userAssignedIdentity:UserAssignedIdentity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, client_id=None, location=None, name=None, principal_id=None, resource_group_name=None, tags=None):
        """
        Get an existing UserAssignedIdentity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: Client ID associated with the user assigned identity.
        :param pulumi.Input[str] location: The location/region where the user assigned identity is
               created.
        :param pulumi.Input[str] name: The name of the user assigned identity. Changing this forces a
               new identity to be created.
        :param pulumi.Input[str] principal_id: Service Principal ID associated with the user assigned identity.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the user assigned identity.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["client_id"] = client_id
        __props__["location"] = location
        __props__["name"] = name
        __props__["principal_id"] = principal_id
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return UserAssignedIdentity(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

