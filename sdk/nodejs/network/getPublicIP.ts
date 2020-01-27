// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Public IP Address.
 * 
 * ## Example Usage (reference an existing)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const example = azure.network.getPublicIP({
 *     name: "nameOfPublicIp",
 *     resourceGroupName: "nameOfResourceGroup",
 * });
 * 
 * export const domainNameLabel = example.domainNameLabel;
 * export const publicIpAddress = example.ipAddress;
 * ```
 * 
 * ## Example Usage (Retrieve the Dynamic Public IP of a new VM)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "West US 2",
 * });
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("example", {
 *     addressSpaces: ["10.0.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("example", {
 *     addressPrefix: "10.0.2.0/24",
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 * });
 * const examplePublicIp = new azure.network.PublicIp("example", {
 *     allocationMethod: "Dynamic",
 *     idleTimeoutInMinutes: 30,
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     tags: {
 *         environment: "test",
 *     },
 * });
 * const exampleNetworkInterface = new azure.network.NetworkInterface("example", {
 *     ipConfigurations: [{
 *         name: "testconfiguration1",
 *         privateIpAddress: "10.0.2.5",
 *         privateIpAddressAllocation: "Static",
 *         publicIpAddressId: examplePublicIp.id,
 *         subnetId: exampleSubnet.id,
 *     }],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleVirtualMachine = new azure.compute.VirtualMachine("example", {
 *     location: exampleResourceGroup.location,
 *     networkInterfaceIds: [exampleNetworkInterface.id],
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const examplePublicIP = pulumi.all([examplePublicIp.name, exampleVirtualMachine.resourceGroupName]).apply(([name, resourceGroupName]) => azure.network.getPublicIP({
 *     name: name,
 *     resourceGroupName: resourceGroupName,
 * }));
 * 
 * export const publicIpAddress = examplePublicIP.ipAddress;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/public_ip.html.markdown.
 */
export function getPublicIP(args: GetPublicIPArgs, opts?: pulumi.InvokeOptions): Promise<GetPublicIPResult> & GetPublicIPResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetPublicIPResult> = pulumi.runtime.invoke("azure:network/getPublicIP:getPublicIP", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "tags": args.tags,
        "zones": args.zones,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getPublicIP.
 */
export interface GetPublicIPArgs {
    /**
     * Specifies the name of the public IP address.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group.
     */
    readonly resourceGroupName: string;
    readonly tags?: {[key: string]: string};
    readonly zones?: string[];
}

/**
 * A collection of values returned by getPublicIP.
 */
export interface GetPublicIPResult {
    readonly allocationMethod: string;
    /**
     * The label for the Domain Name.
     */
    readonly domainNameLabel: string;
    /**
     * Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
     */
    readonly fqdn: string;
    /**
     * Specifies the timeout for the TCP idle connection.
     */
    readonly idleTimeoutInMinutes: number;
    /**
     * The IP address value that was allocated.
     */
    readonly ipAddress: string;
    /**
     * The IP version being used, for example `IPv4` or `IPv6`.
     */
    readonly ipVersion: string;
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    readonly reverseFqdn: string;
    readonly sku: string;
    /**
     * A mapping of tags to assigned to the resource.
     */
    readonly tags: {[key: string]: string};
    readonly zones: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
