// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Links a Log Analytics (formally Operational Insights) Workspace to another resource. The (currently) only linkable service is an Azure Automation Account.
 * 
 * > **NOTE:** This resource has been deprecated in favour of the `azure.loganalytics.LinkedService` resource and will be removed in the next major version of the AzureRM Provider. The new resource shares the same fields as this one, and information on migrating across can be found in this guide.
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "West Europe",
 * });
 * const exampleAccount = new azure.automation.Account("example", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: {
 *         name: "Basic",
 *     },
 *     tags: {
 *         environment: "development",
 *     },
 * });
 * const exampleAnalyticsWorkspace = new azure.operationalinsights.AnalyticsWorkspace("example", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     retentionInDays: 30,
 *     sku: "PerGB2018",
 * });
 * const exampleAnalyticsWorkspaceLinkedService = new azure.operationalinsights.AnalyticsWorkspaceLinkedService("example", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     resourceId: exampleAccount.id,
 *     workspaceName: exampleAnalyticsWorkspace.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/log_analytics_workspace_linked_service.html.markdown.
 */
export class AnalyticsWorkspaceLinkedService extends pulumi.CustomResource {
    /**
     * Get an existing AnalyticsWorkspaceLinkedService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnalyticsWorkspaceLinkedServiceState, opts?: pulumi.CustomResourceOptions): AnalyticsWorkspaceLinkedService {
        return new AnalyticsWorkspaceLinkedService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:operationalinsights/analyticsWorkspaceLinkedService:AnalyticsWorkspaceLinkedService';

    /**
     * Returns true if the given object is an instance of AnalyticsWorkspaceLinkedService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnalyticsWorkspaceLinkedService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnalyticsWorkspaceLinkedService.__pulumiType;
    }

    /**
     * Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspaceName`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
     */
    public readonly linkedServiceName!: pulumi.Output<string | undefined>;
    /**
     * A `linkedServiceProperties` block as defined below.
     */
    public readonly linkedServiceProperties!: pulumi.Output<outputs.operationalinsights.AnalyticsWorkspaceLinkedServiceLinkedServiceProperties>;
    /**
     * The automatically generated name of the Linked Service. This cannot be specified. The format is always `<workspace_name>/<linked_service_name>` e.g. `workspace1/Automation`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The resource id of the resource that will be linked to the workspace. This field has been deprecated in favour of the top-level `resourceId` field and will be removed in v2.0 of the AzureRM Provider.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
     */
    public readonly workspaceName!: pulumi.Output<string>;

    /**
     * Create a AnalyticsWorkspaceLinkedService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnalyticsWorkspaceLinkedServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AnalyticsWorkspaceLinkedServiceArgs | AnalyticsWorkspaceLinkedServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AnalyticsWorkspaceLinkedServiceState | undefined;
            inputs["linkedServiceName"] = state ? state.linkedServiceName : undefined;
            inputs["linkedServiceProperties"] = state ? state.linkedServiceProperties : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["workspaceName"] = state ? state.workspaceName : undefined;
        } else {
            const args = argsOrState as AnalyticsWorkspaceLinkedServiceArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["linkedServiceName"] = args ? args.linkedServiceName : undefined;
            inputs["linkedServiceProperties"] = args ? args.linkedServiceProperties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AnalyticsWorkspaceLinkedService.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AnalyticsWorkspaceLinkedService resources.
 */
export interface AnalyticsWorkspaceLinkedServiceState {
    /**
     * Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspaceName`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
     */
    readonly linkedServiceName?: pulumi.Input<string>;
    /**
     * A `linkedServiceProperties` block as defined below.
     */
    readonly linkedServiceProperties?: pulumi.Input<inputs.operationalinsights.AnalyticsWorkspaceLinkedServiceLinkedServiceProperties>;
    /**
     * The automatically generated name of the Linked Service. This cannot be specified. The format is always `<workspace_name>/<linked_service_name>` e.g. `workspace1/Automation`
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The resource id of the resource that will be linked to the workspace. This field has been deprecated in favour of the top-level `resourceId` field and will be removed in v2.0 of the AzureRM Provider.
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
     */
    readonly workspaceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AnalyticsWorkspaceLinkedService resource.
 */
export interface AnalyticsWorkspaceLinkedServiceArgs {
    /**
     * Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspaceName`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
     */
    readonly linkedServiceName?: pulumi.Input<string>;
    /**
     * A `linkedServiceProperties` block as defined below.
     */
    readonly linkedServiceProperties?: pulumi.Input<inputs.operationalinsights.AnalyticsWorkspaceLinkedServiceLinkedServiceProperties>;
    /**
     * The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource id of the resource that will be linked to the workspace. This field has been deprecated in favour of the top-level `resourceId` field and will be removed in v2.0 of the AzureRM Provider.
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
     */
    readonly workspaceName: pulumi.Input<string>;
}
