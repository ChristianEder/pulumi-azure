// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.RecoveryServices
{
    /// <summary>
    /// &gt; **NOTE:** This resource has been deprecated in favour of the `azure.backup.PolicyVM` resource and will be removed in the next major version of the AzureRM Provider. The new resource shares the same fields as this one, and information on migrating across can be found in this guide.
    /// 
    /// Manages an Recovery Services VM Protection Policy.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/recovery_services_protection_policy_vm.html.markdown.
    /// </summary>
    public partial class ProtectionPolicyVM : Pulumi.CustomResource
    {
        /// <summary>
        /// Configures the Policy backup frequecent, times &amp; days as documented in the `backup` block below. 
        /// </summary>
        [Output("backup")]
        public Output<Outputs.ProtectionPolicyVMBackup> Backup { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Recovery Services Vault Policy. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        /// </summary>
        [Output("recoveryVaultName")]
        public Output<string> RecoveryVaultName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Configures the policy daily retention as documented in the `retention_daily` block below. Required when backup frequency is `Daily`.
        /// </summary>
        [Output("retentionDaily")]
        public Output<Outputs.ProtectionPolicyVMRetentionDaily?> RetentionDaily { get; private set; } = null!;

        /// <summary>
        /// Configures the policy monthly retention as documented in the `retention_monthly` block below.
        /// </summary>
        [Output("retentionMonthly")]
        public Output<Outputs.ProtectionPolicyVMRetentionMonthly?> RetentionMonthly { get; private set; } = null!;

        /// <summary>
        /// Configures the policy weekly retention as documented in the `retention_weekly` block below. Required when backup frequency is `Weekly`.
        /// </summary>
        [Output("retentionWeekly")]
        public Output<Outputs.ProtectionPolicyVMRetentionWeekly?> RetentionWeekly { get; private set; } = null!;

        /// <summary>
        /// Configures the policy yearly retention as documented in the `retention_yearly` block below.
        /// </summary>
        [Output("retentionYearly")]
        public Output<Outputs.ProtectionPolicyVMRetentionYearly?> RetentionYearly { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the timezone. Defaults to `UTC`
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;


        /// <summary>
        /// Create a ProtectionPolicyVM resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProtectionPolicyVM(string name, ProtectionPolicyVMArgs args, CustomResourceOptions? options = null)
            : base("azure:recoveryservices/protectionPolicyVM:ProtectionPolicyVM", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ProtectionPolicyVM(string name, Input<string> id, ProtectionPolicyVMState? state = null, CustomResourceOptions? options = null)
            : base("azure:recoveryservices/protectionPolicyVM:ProtectionPolicyVM", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProtectionPolicyVM resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProtectionPolicyVM Get(string name, Input<string> id, ProtectionPolicyVMState? state = null, CustomResourceOptions? options = null)
        {
            return new ProtectionPolicyVM(name, id, state, options);
        }
    }

    public sealed class ProtectionPolicyVMArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configures the Policy backup frequecent, times &amp; days as documented in the `backup` block below. 
        /// </summary>
        [Input("backup", required: true)]
        public Input<Inputs.ProtectionPolicyVMBackupArgs> Backup { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Recovery Services Vault Policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        /// </summary>
        [Input("recoveryVaultName", required: true)]
        public Input<string> RecoveryVaultName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Configures the policy daily retention as documented in the `retention_daily` block below. Required when backup frequency is `Daily`.
        /// </summary>
        [Input("retentionDaily")]
        public Input<Inputs.ProtectionPolicyVMRetentionDailyArgs>? RetentionDaily { get; set; }

        /// <summary>
        /// Configures the policy monthly retention as documented in the `retention_monthly` block below.
        /// </summary>
        [Input("retentionMonthly")]
        public Input<Inputs.ProtectionPolicyVMRetentionMonthlyArgs>? RetentionMonthly { get; set; }

        /// <summary>
        /// Configures the policy weekly retention as documented in the `retention_weekly` block below. Required when backup frequency is `Weekly`.
        /// </summary>
        [Input("retentionWeekly")]
        public Input<Inputs.ProtectionPolicyVMRetentionWeeklyArgs>? RetentionWeekly { get; set; }

        /// <summary>
        /// Configures the policy yearly retention as documented in the `retention_yearly` block below.
        /// </summary>
        [Input("retentionYearly")]
        public Input<Inputs.ProtectionPolicyVMRetentionYearlyArgs>? RetentionYearly { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the timezone. Defaults to `UTC`
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public ProtectionPolicyVMArgs()
        {
        }
    }

    public sealed class ProtectionPolicyVMState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configures the Policy backup frequecent, times &amp; days as documented in the `backup` block below. 
        /// </summary>
        [Input("backup")]
        public Input<Inputs.ProtectionPolicyVMBackupGetArgs>? Backup { get; set; }

        /// <summary>
        /// Specifies the name of the Recovery Services Vault Policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        /// </summary>
        [Input("recoveryVaultName")]
        public Input<string>? RecoveryVaultName { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Configures the policy daily retention as documented in the `retention_daily` block below. Required when backup frequency is `Daily`.
        /// </summary>
        [Input("retentionDaily")]
        public Input<Inputs.ProtectionPolicyVMRetentionDailyGetArgs>? RetentionDaily { get; set; }

        /// <summary>
        /// Configures the policy monthly retention as documented in the `retention_monthly` block below.
        /// </summary>
        [Input("retentionMonthly")]
        public Input<Inputs.ProtectionPolicyVMRetentionMonthlyGetArgs>? RetentionMonthly { get; set; }

        /// <summary>
        /// Configures the policy weekly retention as documented in the `retention_weekly` block below. Required when backup frequency is `Weekly`.
        /// </summary>
        [Input("retentionWeekly")]
        public Input<Inputs.ProtectionPolicyVMRetentionWeeklyGetArgs>? RetentionWeekly { get; set; }

        /// <summary>
        /// Configures the policy yearly retention as documented in the `retention_yearly` block below.
        /// </summary>
        [Input("retentionYearly")]
        public Input<Inputs.ProtectionPolicyVMRetentionYearlyGetArgs>? RetentionYearly { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the timezone. Defaults to `UTC`
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public ProtectionPolicyVMState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ProtectionPolicyVMBackupArgs : Pulumi.ResourceArgs
    {
        [Input("frequency", required: true)]
        public Input<string> Frequency { get; set; } = null!;

        [Input("time", required: true)]
        public Input<string> Time { get; set; } = null!;

        [Input("weekdays")]
        private InputList<string>? _weekdays;
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        public ProtectionPolicyVMBackupArgs()
        {
        }
    }

    public sealed class ProtectionPolicyVMBackupGetArgs : Pulumi.ResourceArgs
    {
        [Input("frequency", required: true)]
        public Input<string> Frequency { get; set; } = null!;

        [Input("time", required: true)]
        public Input<string> Time { get; set; } = null!;

        [Input("weekdays")]
        private InputList<string>? _weekdays;
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        public ProtectionPolicyVMBackupGetArgs()
        {
        }
    }

    public sealed class ProtectionPolicyVMRetentionDailyArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        public ProtectionPolicyVMRetentionDailyArgs()
        {
        }
    }

    public sealed class ProtectionPolicyVMRetentionDailyGetArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        public ProtectionPolicyVMRetentionDailyGetArgs()
        {
        }
    }

    public sealed class ProtectionPolicyVMRetentionMonthlyArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("weekdays", required: true)]
        private InputList<string>? _weekdays;
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        [Input("weeks", required: true)]
        private InputList<string>? _weeks;
        public InputList<string> Weeks
        {
            get => _weeks ?? (_weeks = new InputList<string>());
            set => _weeks = value;
        }

        public ProtectionPolicyVMRetentionMonthlyArgs()
        {
        }
    }

    public sealed class ProtectionPolicyVMRetentionMonthlyGetArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("weekdays", required: true)]
        private InputList<string>? _weekdays;
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        [Input("weeks", required: true)]
        private InputList<string>? _weeks;
        public InputList<string> Weeks
        {
            get => _weeks ?? (_weeks = new InputList<string>());
            set => _weeks = value;
        }

        public ProtectionPolicyVMRetentionMonthlyGetArgs()
        {
        }
    }

    public sealed class ProtectionPolicyVMRetentionWeeklyArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("weekdays", required: true)]
        private InputList<string>? _weekdays;
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        public ProtectionPolicyVMRetentionWeeklyArgs()
        {
        }
    }

    public sealed class ProtectionPolicyVMRetentionWeeklyGetArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("weekdays", required: true)]
        private InputList<string>? _weekdays;
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        public ProtectionPolicyVMRetentionWeeklyGetArgs()
        {
        }
    }

    public sealed class ProtectionPolicyVMRetentionYearlyArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("months", required: true)]
        private InputList<string>? _months;
        public InputList<string> Months
        {
            get => _months ?? (_months = new InputList<string>());
            set => _months = value;
        }

        [Input("weekdays", required: true)]
        private InputList<string>? _weekdays;
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        [Input("weeks", required: true)]
        private InputList<string>? _weeks;
        public InputList<string> Weeks
        {
            get => _weeks ?? (_weeks = new InputList<string>());
            set => _weeks = value;
        }

        public ProtectionPolicyVMRetentionYearlyArgs()
        {
        }
    }

    public sealed class ProtectionPolicyVMRetentionYearlyGetArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("months", required: true)]
        private InputList<string>? _months;
        public InputList<string> Months
        {
            get => _months ?? (_months = new InputList<string>());
            set => _months = value;
        }

        [Input("weekdays", required: true)]
        private InputList<string>? _weekdays;
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        [Input("weeks", required: true)]
        private InputList<string>? _weeks;
        public InputList<string> Weeks
        {
            get => _weeks ?? (_weeks = new InputList<string>());
            set => _weeks = value;
        }

        public ProtectionPolicyVMRetentionYearlyGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ProtectionPolicyVMBackup
    {
        public readonly string Frequency;
        public readonly string Time;
        public readonly ImmutableArray<string> Weekdays;

        [OutputConstructor]
        private ProtectionPolicyVMBackup(
            string frequency,
            string time,
            ImmutableArray<string> weekdays)
        {
            Frequency = frequency;
            Time = time;
            Weekdays = weekdays;
        }
    }

    [OutputType]
    public sealed class ProtectionPolicyVMRetentionDaily
    {
        public readonly int Count;

        [OutputConstructor]
        private ProtectionPolicyVMRetentionDaily(int count)
        {
            Count = count;
        }
    }

    [OutputType]
    public sealed class ProtectionPolicyVMRetentionMonthly
    {
        public readonly int Count;
        public readonly ImmutableArray<string> Weekdays;
        public readonly ImmutableArray<string> Weeks;

        [OutputConstructor]
        private ProtectionPolicyVMRetentionMonthly(
            int count,
            ImmutableArray<string> weekdays,
            ImmutableArray<string> weeks)
        {
            Count = count;
            Weekdays = weekdays;
            Weeks = weeks;
        }
    }

    [OutputType]
    public sealed class ProtectionPolicyVMRetentionWeekly
    {
        public readonly int Count;
        public readonly ImmutableArray<string> Weekdays;

        [OutputConstructor]
        private ProtectionPolicyVMRetentionWeekly(
            int count,
            ImmutableArray<string> weekdays)
        {
            Count = count;
            Weekdays = weekdays;
        }
    }

    [OutputType]
    public sealed class ProtectionPolicyVMRetentionYearly
    {
        public readonly int Count;
        public readonly ImmutableArray<string> Months;
        public readonly ImmutableArray<string> Weekdays;
        public readonly ImmutableArray<string> Weeks;

        [OutputConstructor]
        private ProtectionPolicyVMRetentionYearly(
            int count,
            ImmutableArray<string> months,
            ImmutableArray<string> weekdays,
            ImmutableArray<string> weeks)
        {
            Count = count;
            Months = months;
            Weekdays = weekdays;
            Weeks = weeks;
        }
    }
    }
}
