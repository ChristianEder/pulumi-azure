// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an App Service Slot (within an App Service).
 * 
 * > **Note:** When using Slots - the `appSettings`, `connectionString` and `siteConfig` blocks on the `azure.appservice.AppService` resource will be overwritten when promoting a Slot using the `azure.appservice.ActiveSlot` resource.
 * 
 * 
 * ## Example Usage (.net 4.x)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as random from "@pulumi/random";
 * 
 * const server = new random.RandomId("server", {
 *     byteLength: 8,
 *     keepers: {
 *         azi_id: 1,
 *     },
 * });
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "West Europe",
 * });
 * const examplePlan = new azure.appservice.Plan("example", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: {
 *         size: "S1",
 *         tier: "Standard",
 *     },
 * });
 * const exampleAppService = new azure.appservice.AppService("example", {
 *     appServicePlanId: examplePlan.id,
 *     appSettings: {
 *         SOME_KEY: "some-value",
 *     },
 *     connectionStrings: [{
 *         name: "Database",
 *         type: "SQLServer",
 *         value: "Server=some-server.mydomain.com;Integrated Security=SSPI",
 *     }],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     siteConfig: {
 *         dotnetFrameworkVersion: "v4.0",
 *     },
 * });
 * const exampleSlot = new azure.appservice.Slot("example", {
 *     appServiceName: exampleAppService.name,
 *     appServicePlanId: examplePlan.id,
 *     appSettings: {
 *         SOME_KEY: "some-value",
 *     },
 *     connectionStrings: [{
 *         name: "Database",
 *         type: "SQLServer",
 *         value: "Server=some-server.mydomain.com;Integrated Security=SSPI",
 *     }],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     siteConfig: {
 *         dotnetFrameworkVersion: "v4.0",
 *     },
 * });
 * ```
 * 
 * ## Example Usage (Java 1.8)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as random from "@pulumi/random";
 * 
 * const server = new random.RandomId("server", {
 *     byteLength: 8,
 *     keepers: {
 *         azi_id: 1,
 *     },
 * });
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "West Europe",
 * });
 * const examplePlan = new azure.appservice.Plan("example", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: {
 *         size: "S1",
 *         tier: "Standard",
 *     },
 * });
 * const exampleAppService = new azure.appservice.AppService("example", {
 *     appServicePlanId: examplePlan.id,
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     siteConfig: {
 *         javaContainer: "JETTY",
 *         javaContainerVersion: "9.3",
 *         javaVersion: "1.8",
 *     },
 * });
 * const exampleSlot = new azure.appservice.Slot("example", {
 *     appServiceName: exampleAppService.name,
 *     appServicePlanId: examplePlan.id,
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     siteConfig: {
 *         javaContainer: "JETTY",
 *         javaContainerVersion: "9.3",
 *         javaVersion: "1.8",
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/app_service_slot.html.markdown.
 */
export class Slot extends pulumi.CustomResource {
    /**
     * Get an existing Slot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SlotState, opts?: pulumi.CustomResourceOptions): Slot {
        return new Slot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appservice/slot:Slot';

    /**
     * Returns true if the given object is an instance of Slot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Slot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Slot.__pulumiType;
    }

    /**
     * The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
     */
    public readonly appServiceName!: pulumi.Output<string>;
    /**
     * The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
     */
    public readonly appServicePlanId!: pulumi.Output<string>;
    /**
     * A key-value pair of App Settings.
     */
    public readonly appSettings!: pulumi.Output<{[key: string]: string}>;
    /**
     * A `authSettings` block as defined below.
     */
    public readonly authSettings!: pulumi.Output<outputs.appservice.SlotAuthSettings>;
    /**
     * Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
     */
    public readonly clientAffinityEnabled!: pulumi.Output<boolean>;
    /**
     * An `connectionString` block as defined below.
     */
    public readonly connectionStrings!: pulumi.Output<outputs.appservice.SlotConnectionString[]>;
    /**
     * The Default Hostname associated with the App Service Slot - such as `mysite.azurewebsites.net`
     */
    public /*out*/ readonly defaultSiteHostname!: pulumi.Output<string>;
    /**
     * Is the App Service Slot Enabled?
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
     */
    public readonly httpsOnly!: pulumi.Output<boolean | undefined>;
    /**
     * A Managed Service Identity block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.appservice.SlotIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    public readonly logs!: pulumi.Output<outputs.appservice.SlotLogs>;
    /**
     * The name of the Connection String.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the App Service Slot component.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    public readonly siteConfig!: pulumi.Output<outputs.appservice.SlotSiteConfig>;
    /**
     * A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
     */
    public /*out*/ readonly siteCredential!: pulumi.Output<outputs.appservice.SlotSiteCredential>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Slot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SlotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SlotArgs | SlotState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SlotState | undefined;
            inputs["appServiceName"] = state ? state.appServiceName : undefined;
            inputs["appServicePlanId"] = state ? state.appServicePlanId : undefined;
            inputs["appSettings"] = state ? state.appSettings : undefined;
            inputs["authSettings"] = state ? state.authSettings : undefined;
            inputs["clientAffinityEnabled"] = state ? state.clientAffinityEnabled : undefined;
            inputs["connectionStrings"] = state ? state.connectionStrings : undefined;
            inputs["defaultSiteHostname"] = state ? state.defaultSiteHostname : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["httpsOnly"] = state ? state.httpsOnly : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["logs"] = state ? state.logs : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["siteConfig"] = state ? state.siteConfig : undefined;
            inputs["siteCredential"] = state ? state.siteCredential : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SlotArgs | undefined;
            if (!args || args.appServiceName === undefined) {
                throw new Error("Missing required property 'appServiceName'");
            }
            if (!args || args.appServicePlanId === undefined) {
                throw new Error("Missing required property 'appServicePlanId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["appServiceName"] = args ? args.appServiceName : undefined;
            inputs["appServicePlanId"] = args ? args.appServicePlanId : undefined;
            inputs["appSettings"] = args ? args.appSettings : undefined;
            inputs["authSettings"] = args ? args.authSettings : undefined;
            inputs["clientAffinityEnabled"] = args ? args.clientAffinityEnabled : undefined;
            inputs["connectionStrings"] = args ? args.connectionStrings : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["httpsOnly"] = args ? args.httpsOnly : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["logs"] = args ? args.logs : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["siteConfig"] = args ? args.siteConfig : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["defaultSiteHostname"] = undefined /*out*/;
            inputs["siteCredential"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Slot.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Slot resources.
 */
export interface SlotState {
    /**
     * The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
     */
    readonly appServiceName?: pulumi.Input<string>;
    /**
     * The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
     */
    readonly appServicePlanId?: pulumi.Input<string>;
    /**
     * A key-value pair of App Settings.
     */
    readonly appSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `authSettings` block as defined below.
     */
    readonly authSettings?: pulumi.Input<inputs.appservice.SlotAuthSettings>;
    /**
     * Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
     */
    readonly clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * An `connectionString` block as defined below.
     */
    readonly connectionStrings?: pulumi.Input<pulumi.Input<inputs.appservice.SlotConnectionString>[]>;
    /**
     * The Default Hostname associated with the App Service Slot - such as `mysite.azurewebsites.net`
     */
    readonly defaultSiteHostname?: pulumi.Input<string>;
    /**
     * Is the App Service Slot Enabled?
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
     */
    readonly httpsOnly?: pulumi.Input<boolean>;
    /**
     * A Managed Service Identity block as defined below.
     */
    readonly identity?: pulumi.Input<inputs.appservice.SlotIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    readonly logs?: pulumi.Input<inputs.appservice.SlotLogs>;
    /**
     * The name of the Connection String.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the App Service Slot component.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    readonly siteConfig?: pulumi.Input<inputs.appservice.SlotSiteConfig>;
    /**
     * A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
     */
    readonly siteCredential?: pulumi.Input<inputs.appservice.SlotSiteCredential>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Slot resource.
 */
export interface SlotArgs {
    /**
     * The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
     */
    readonly appServiceName: pulumi.Input<string>;
    /**
     * The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
     */
    readonly appServicePlanId: pulumi.Input<string>;
    /**
     * A key-value pair of App Settings.
     */
    readonly appSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `authSettings` block as defined below.
     */
    readonly authSettings?: pulumi.Input<inputs.appservice.SlotAuthSettings>;
    /**
     * Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
     */
    readonly clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * An `connectionString` block as defined below.
     */
    readonly connectionStrings?: pulumi.Input<pulumi.Input<inputs.appservice.SlotConnectionString>[]>;
    /**
     * Is the App Service Slot Enabled?
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
     */
    readonly httpsOnly?: pulumi.Input<boolean>;
    /**
     * A Managed Service Identity block as defined below.
     */
    readonly identity?: pulumi.Input<inputs.appservice.SlotIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    readonly logs?: pulumi.Input<inputs.appservice.SlotLogs>;
    /**
     * The name of the Connection String.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the App Service Slot component.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    readonly siteConfig?: pulumi.Input<inputs.appservice.SlotSiteConfig>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
