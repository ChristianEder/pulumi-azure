// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Action Group within Azure Monitor.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "West US",
 * });
 * const exampleActionGroup = new azure.monitoring.ActionGroup("example", {
 *     armRoleReceivers: [{
 *         name: "armroleaction",
 *         roleId: "de139f84-1756-47ae-9be6-808fbbe84772",
 *         useCommonAlertSchema: true,
 *     }],
 *     automationRunbookReceivers: [{
 *         automationAccountId: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-runbooks/providers/microsoft.automation/automationaccounts/aaa001",
 *         isGlobalRunbook: true,
 *         name: "actionName1",
 *         runbookName: "my runbook",
 *         serviceUri: "https://s13events.azure-automation.net/webhooks?token=randomtoken",
 *         useCommonAlertSchema: true,
 *         webhookResourceId: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-runbooks/providers/microsoft.automation/automationaccounts/aaa001/webhooks/webhook_alert",
 *     }],
 *     azureAppPushReceivers: [{
 *         emailAddress: "admin@contoso.com",
 *         name: "pushtoadmin",
 *     }],
 *     azureFunctionReceivers: [{
 *         functionAppResourceId: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-funcapp/providers/Microsoft.Web/sites/funcapp",
 *         functionName: "myfunc",
 *         httpTriggerUrl: "https://example.com/trigger",
 *         name: "funcaction",
 *         useCommonAlertSchema: true,
 *     }],
 *     emailReceivers: [
 *         {
 *             emailAddress: "admin@contoso.com",
 *             name: "sendtoadmin",
 *         },
 *         {
 *             emailAddress: "devops@contoso.com",
 *             name: "sendtodevops",
 *             useCommonAlertSchema: true,
 *         },
 *     ],
 *     itsmReceivers: [{
 *         connectionId: "53de6956-42b4-41ba-be3c-b154cdf17b13",
 *         name: "createorupdateticket",
 *         region: "southcentralus",
 *         ticketConfiguration: "{}",
 *         workspaceId: "6eee3a18-aac3-40e4-b98e-1f309f329816",
 *     }],
 *     logicAppReceivers: [{
 *         callbackUrl: "https://logicapptriggerurl/...",
 *         name: "logicappaction",
 *         resourceId: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-logicapp/providers/Microsoft.Logic/workflows/logicapp",
 *         useCommonAlertSchema: true,
 *     }],
 *     resourceGroupName: exampleResourceGroup.name,
 *     shortName: "p0action",
 *     smsReceivers: [{
 *         countryCode: "1",
 *         name: "oncallmsg",
 *         phoneNumber: "1231231234",
 *     }],
 *     voiceReceivers: [{
 *         countryCode: "86",
 *         name: "remotesupport",
 *         phoneNumber: "13888888888",
 *     }],
 *     webhookReceivers: [{
 *         name: "callmyapiaswell",
 *         serviceUri: "http://example.com/alert",
 *         useCommonAlertSchema: true,
 *     }],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/monitor_action_group.html.markdown.
 */
export class ActionGroup extends pulumi.CustomResource {
    /**
     * Get an existing ActionGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionGroupState, opts?: pulumi.CustomResourceOptions): ActionGroup {
        return new ActionGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:monitoring/actionGroup:ActionGroup';

    /**
     * Returns true if the given object is an instance of ActionGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionGroup.__pulumiType;
    }

    /**
     * One or more `armRoleReceiver` blocks as defined below.
     */
    public readonly armRoleReceivers!: pulumi.Output<outputs.monitoring.ActionGroupArmRoleReceiver[] | undefined>;
    /**
     * One or more `automationRunbookReceiver` blocks as defined below.
     */
    public readonly automationRunbookReceivers!: pulumi.Output<outputs.monitoring.ActionGroupAutomationRunbookReceiver[] | undefined>;
    /**
     * One or more `azureAppPushReceiver` blocks as defined below.
     */
    public readonly azureAppPushReceivers!: pulumi.Output<outputs.monitoring.ActionGroupAzureAppPushReceiver[] | undefined>;
    /**
     * One or more `azureFunctionReceiver` blocks as defined below.
     */
    public readonly azureFunctionReceivers!: pulumi.Output<outputs.monitoring.ActionGroupAzureFunctionReceiver[] | undefined>;
    /**
     * One or more `emailReceiver` blocks as defined below.
     */
    public readonly emailReceivers!: pulumi.Output<outputs.monitoring.ActionGroupEmailReceiver[] | undefined>;
    /**
     * Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * One or more `itsmReceiver` blocks as defined below.
     */
    public readonly itsmReceivers!: pulumi.Output<outputs.monitoring.ActionGroupItsmReceiver[] | undefined>;
    /**
     * One or more `logicAppReceiver` blocks as defined below.
     */
    public readonly logicAppReceivers!: pulumi.Output<outputs.monitoring.ActionGroupLogicAppReceiver[] | undefined>;
    /**
     * The name of the webhook receiver. Names must be unique (case-insensitive) across all receivers within an action group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Action Group instance.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The short name of the action group. This will be used in SMS messages.
     */
    public readonly shortName!: pulumi.Output<string>;
    /**
     * One or more `smsReceiver` blocks as defined below.
     */
    public readonly smsReceivers!: pulumi.Output<outputs.monitoring.ActionGroupSmsReceiver[] | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * One or more `voiceReceiver` blocks as defined below.
     */
    public readonly voiceReceivers!: pulumi.Output<outputs.monitoring.ActionGroupVoiceReceiver[] | undefined>;
    /**
     * One or more `webhookReceiver` blocks as defined below.
     */
    public readonly webhookReceivers!: pulumi.Output<outputs.monitoring.ActionGroupWebhookReceiver[] | undefined>;

    /**
     * Create a ActionGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionGroupArgs | ActionGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ActionGroupState | undefined;
            inputs["armRoleReceivers"] = state ? state.armRoleReceivers : undefined;
            inputs["automationRunbookReceivers"] = state ? state.automationRunbookReceivers : undefined;
            inputs["azureAppPushReceivers"] = state ? state.azureAppPushReceivers : undefined;
            inputs["azureFunctionReceivers"] = state ? state.azureFunctionReceivers : undefined;
            inputs["emailReceivers"] = state ? state.emailReceivers : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["itsmReceivers"] = state ? state.itsmReceivers : undefined;
            inputs["logicAppReceivers"] = state ? state.logicAppReceivers : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["shortName"] = state ? state.shortName : undefined;
            inputs["smsReceivers"] = state ? state.smsReceivers : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["voiceReceivers"] = state ? state.voiceReceivers : undefined;
            inputs["webhookReceivers"] = state ? state.webhookReceivers : undefined;
        } else {
            const args = argsOrState as ActionGroupArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.shortName === undefined) {
                throw new Error("Missing required property 'shortName'");
            }
            inputs["armRoleReceivers"] = args ? args.armRoleReceivers : undefined;
            inputs["automationRunbookReceivers"] = args ? args.automationRunbookReceivers : undefined;
            inputs["azureAppPushReceivers"] = args ? args.azureAppPushReceivers : undefined;
            inputs["azureFunctionReceivers"] = args ? args.azureFunctionReceivers : undefined;
            inputs["emailReceivers"] = args ? args.emailReceivers : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["itsmReceivers"] = args ? args.itsmReceivers : undefined;
            inputs["logicAppReceivers"] = args ? args.logicAppReceivers : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["shortName"] = args ? args.shortName : undefined;
            inputs["smsReceivers"] = args ? args.smsReceivers : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["voiceReceivers"] = args ? args.voiceReceivers : undefined;
            inputs["webhookReceivers"] = args ? args.webhookReceivers : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ActionGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionGroup resources.
 */
export interface ActionGroupState {
    /**
     * One or more `armRoleReceiver` blocks as defined below.
     */
    readonly armRoleReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupArmRoleReceiver>[]>;
    /**
     * One or more `automationRunbookReceiver` blocks as defined below.
     */
    readonly automationRunbookReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupAutomationRunbookReceiver>[]>;
    /**
     * One or more `azureAppPushReceiver` blocks as defined below.
     */
    readonly azureAppPushReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupAzureAppPushReceiver>[]>;
    /**
     * One or more `azureFunctionReceiver` blocks as defined below.
     */
    readonly azureFunctionReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupAzureFunctionReceiver>[]>;
    /**
     * One or more `emailReceiver` blocks as defined below.
     */
    readonly emailReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupEmailReceiver>[]>;
    /**
     * Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * One or more `itsmReceiver` blocks as defined below.
     */
    readonly itsmReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupItsmReceiver>[]>;
    /**
     * One or more `logicAppReceiver` blocks as defined below.
     */
    readonly logicAppReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupLogicAppReceiver>[]>;
    /**
     * The name of the webhook receiver. Names must be unique (case-insensitive) across all receivers within an action group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Action Group instance.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The short name of the action group. This will be used in SMS messages.
     */
    readonly shortName?: pulumi.Input<string>;
    /**
     * One or more `smsReceiver` blocks as defined below.
     */
    readonly smsReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupSmsReceiver>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * One or more `voiceReceiver` blocks as defined below.
     */
    readonly voiceReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupVoiceReceiver>[]>;
    /**
     * One or more `webhookReceiver` blocks as defined below.
     */
    readonly webhookReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupWebhookReceiver>[]>;
}

/**
 * The set of arguments for constructing a ActionGroup resource.
 */
export interface ActionGroupArgs {
    /**
     * One or more `armRoleReceiver` blocks as defined below.
     */
    readonly armRoleReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupArmRoleReceiver>[]>;
    /**
     * One or more `automationRunbookReceiver` blocks as defined below.
     */
    readonly automationRunbookReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupAutomationRunbookReceiver>[]>;
    /**
     * One or more `azureAppPushReceiver` blocks as defined below.
     */
    readonly azureAppPushReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupAzureAppPushReceiver>[]>;
    /**
     * One or more `azureFunctionReceiver` blocks as defined below.
     */
    readonly azureFunctionReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupAzureFunctionReceiver>[]>;
    /**
     * One or more `emailReceiver` blocks as defined below.
     */
    readonly emailReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupEmailReceiver>[]>;
    /**
     * Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * One or more `itsmReceiver` blocks as defined below.
     */
    readonly itsmReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupItsmReceiver>[]>;
    /**
     * One or more `logicAppReceiver` blocks as defined below.
     */
    readonly logicAppReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupLogicAppReceiver>[]>;
    /**
     * The name of the webhook receiver. Names must be unique (case-insensitive) across all receivers within an action group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Action Group instance.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The short name of the action group. This will be used in SMS messages.
     */
    readonly shortName: pulumi.Input<string>;
    /**
     * One or more `smsReceiver` blocks as defined below.
     */
    readonly smsReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupSmsReceiver>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * One or more `voiceReceiver` blocks as defined below.
     */
    readonly voiceReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupVoiceReceiver>[]>;
    /**
     * One or more `webhookReceiver` blocks as defined below.
     */
    readonly webhookReceivers?: pulumi.Input<pulumi.Input<inputs.monitoring.ActionGroupWebhookReceiver>[]>;
}
