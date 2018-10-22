# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Policy(pulumi.CustomResource):
    """
    Manages a Policy within a Dev Test Policy Set.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, evaluator_type=None, fact_data=None, lab_name=None, name=None, policy_set_name=None, resource_group_name=None, tags=None, threshold=None):
        """Create a Policy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        A description for the Policy.
        """
        __props__['description'] = description

        if not evaluator_type:
            raise TypeError('Missing required property evaluator_type')
        elif not isinstance(evaluator_type, basestring):
            raise TypeError('Expected property evaluator_type to be a basestring')
        __self__.evaluator_type = evaluator_type
        """
        The Evaluation Type used for this Policy. Possible values include: 'AllowedValuesPolicy', 'MaxValuePolicy'. Changing this forces a new resource to be created.
        """
        __props__['evaluatorType'] = evaluator_type

        if fact_data and not isinstance(fact_data, basestring):
            raise TypeError('Expected property fact_data to be a basestring')
        __self__.fact_data = fact_data
        """
        The Fact Data for this Policy.
        """
        __props__['factData'] = fact_data

        if not lab_name:
            raise TypeError('Missing required property lab_name')
        elif not isinstance(lab_name, basestring):
            raise TypeError('Expected property lab_name to be a basestring')
        __self__.lab_name = lab_name
        """
        Specifies the name of the Dev Test Lab in which the Policy should be created. Changing this forces a new resource to be created.
        """
        __props__['labName'] = lab_name

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        Specifies the name of the Dev Test Policy. Possible values are `GalleryImage`, `LabPremiumVmCount`, `LabTargetCost`, `LabVmCount`, `LabVmSize`, `UserOwnedLabPremiumVmCount`, `UserOwnedLabVmCount` and `UserOwnedLabVmCountInSubnet`. Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if not policy_set_name:
            raise TypeError('Missing required property policy_set_name')
        elif not isinstance(policy_set_name, basestring):
            raise TypeError('Expected property policy_set_name to be a basestring')
        __self__.policy_set_name = policy_set_name
        """
        Specifies the name of the Policy Set within the Dev Test Lab where this policy should be created. Changing this forces a new resource to be created.
        """
        __props__['policySetName'] = policy_set_name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        __props__['tags'] = tags

        if not threshold:
            raise TypeError('Missing required property threshold')
        elif not isinstance(threshold, basestring):
            raise TypeError('Expected property threshold to be a basestring')
        __self__.threshold = threshold
        """
        The Threshold for this Policy.
        """
        __props__['threshold'] = threshold

        super(Policy, __self__).__init__(
            'azure:devtest/policy:Policy',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'description' in outs:
            self.description = outs['description']
        if 'evaluatorType' in outs:
            self.evaluator_type = outs['evaluatorType']
        if 'factData' in outs:
            self.fact_data = outs['factData']
        if 'labName' in outs:
            self.lab_name = outs['labName']
        if 'name' in outs:
            self.name = outs['name']
        if 'policySetName' in outs:
            self.policy_set_name = outs['policySetName']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'threshold' in outs:
            self.threshold = outs['threshold']