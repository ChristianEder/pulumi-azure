// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation
{
    public static class GetBoolVariable
    {
        /// <summary>
        /// Use this data source to access information about an existing Automation Bool Variable.
        /// 
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBoolVariableResult> InvokeAsync(GetBoolVariableArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBoolVariableResult>("azure:automation/getBoolVariable:getBoolVariable", args ?? new GetBoolVariableArgs(), options.WithVersion());
    }


    public sealed class GetBoolVariableArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account in which the Automation Variable exists.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Automation Variable.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the automation account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetBoolVariableArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBoolVariableResult
    {
        public readonly string AutomationAccountName;
        /// <summary>
        /// The description of the Automation Variable.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Specifies if the Automation Variable is encrypted. Defaults to `false`.
        /// </summary>
        public readonly bool Encrypted;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The value of the Automation Variable as a `boolean`.
        /// </summary>
        public readonly bool Value;

        [OutputConstructor]
        private GetBoolVariableResult(
            string automationAccountName,

            string description,

            bool encrypted,

            string id,

            string name,

            string resourceGroupName,

            bool value)
        {
            AutomationAccountName = automationAccountName;
            Description = description;
            Encrypted = encrypted;
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Value = value;
        }
    }
}
