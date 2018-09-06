# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Hub(pulumi.CustomResource):
    """
    Manages a Notification Hub within a Notification Hub Namespace.
    """
    def __init__(__self__, __name__, __opts__=None, apns_credential=None, gcm_credential=None, location=None, name=None, namespace_name=None, resource_group_name=None):
        """Create a Hub resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if apns_credential and not isinstance(apns_credential, dict):
            raise TypeError('Expected property apns_credential to be a dict')
        __self__.apns_credential = apns_credential
        """
        A `apns_credential` block as defined below.
        """
        __props__['apnsCredential'] = apns_credential

        if gcm_credential and not isinstance(gcm_credential, dict):
            raise TypeError('Expected property gcm_credential to be a dict')
        __self__.gcm_credential = gcm_credential
        """
        A `gcm_credential` block as defined below.
        """
        __props__['gcmCredential'] = gcm_credential

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
        """
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name to use for this Notification Hub. Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if not namespace_name:
            raise TypeError('Missing required property namespace_name')
        elif not isinstance(namespace_name, basestring):
            raise TypeError('Expected property namespace_name to be a basestring')
        __self__.namespace_name = namespace_name
        """
        The name of the Notification Hub Namespace in which to create this Notification Hub. Changing this forces a new resource to be created.
        """
        __props__['namespaceName'] = namespace_name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        super(Hub, __self__).__init__(
            'azure:notificationhub/hub:Hub',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'apnsCredential' in outs:
            self.apns_credential = outs['apnsCredential']
        if 'gcmCredential' in outs:
            self.gcm_credential = outs['gcmCredential']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'namespaceName' in outs:
            self.namespace_name = outs['namespaceName']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']