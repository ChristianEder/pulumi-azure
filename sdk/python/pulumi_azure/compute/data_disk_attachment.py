# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class DataDiskAttachment(pulumi.CustomResource):
    """
    Manages attaching a Disk to a Virtual Machine.
    
    ~> **NOTE:** Data Disks can be attached either directly on the `azurerm_virtual_machine` resource, or using the `azurerm_virtual_machine_data_disk_attachment` resource - but the two cannot be used together. If both are used against the same Virtual Machine, spurious changes will occur.
    
    -> **Please Note:** only Managed Disks are supported via this separate resource, Unmanaged Disks can be attached using the `storage_data_disk` block in the `azurerm_virtual_machine` resource.
    """
    def __init__(__self__, __name__, __opts__=None, caching=None, create_option=None, lun=None, managed_disk_id=None, virtual_machine_id=None, write_accelerator_enabled=None):
        """Create a DataDiskAttachment resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not caching:
            raise TypeError('Missing required property caching')
        elif not isinstance(caching, basestring):
            raise TypeError('Expected property caching to be a basestring')
        __self__.caching = caching
        """
        Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
        """
        __props__['caching'] = caching

        if create_option and not isinstance(create_option, basestring):
            raise TypeError('Expected property create_option to be a basestring')
        __self__.create_option = create_option
        """
        The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
        """
        __props__['createOption'] = create_option

        if not lun:
            raise TypeError('Missing required property lun')
        elif not isinstance(lun, int):
            raise TypeError('Expected property lun to be a int')
        __self__.lun = lun
        """
        The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
        """
        __props__['lun'] = lun

        if not managed_disk_id:
            raise TypeError('Missing required property managed_disk_id')
        elif not isinstance(managed_disk_id, basestring):
            raise TypeError('Expected property managed_disk_id to be a basestring')
        __self__.managed_disk_id = managed_disk_id
        """
        The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
        """
        __props__['managedDiskId'] = managed_disk_id

        if not virtual_machine_id:
            raise TypeError('Missing required property virtual_machine_id')
        elif not isinstance(virtual_machine_id, basestring):
            raise TypeError('Expected property virtual_machine_id to be a basestring')
        __self__.virtual_machine_id = virtual_machine_id
        """
        The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
        """
        __props__['virtualMachineId'] = virtual_machine_id

        if write_accelerator_enabled and not isinstance(write_accelerator_enabled, bool):
            raise TypeError('Expected property write_accelerator_enabled to be a bool')
        __self__.write_accelerator_enabled = write_accelerator_enabled
        """
        Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
        """
        __props__['writeAcceleratorEnabled'] = write_accelerator_enabled

        super(DataDiskAttachment, __self__).__init__(
            'azure:compute/dataDiskAttachment:DataDiskAttachment',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'caching' in outs:
            self.caching = outs['caching']
        if 'createOption' in outs:
            self.create_option = outs['createOption']
        if 'lun' in outs:
            self.lun = outs['lun']
        if 'managedDiskId' in outs:
            self.managed_disk_id = outs['managedDiskId']
        if 'virtualMachineId' in outs:
            self.virtual_machine_id = outs['virtualMachineId']
        if 'writeAcceleratorEnabled' in outs:
            self.write_accelerator_enabled = outs['writeAcceleratorEnabled']