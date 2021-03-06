// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type PolicyCustomRule struct {
	// Type of Actions
	Action string `pulumi:"action"`
	// One or more `matchCondition` block defined below.
	MatchConditions []PolicyCustomRuleMatchCondition `pulumi:"matchConditions"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value
	Priority int `pulumi:"priority"`
	// Describes the type of rule
	RuleType string `pulumi:"ruleType"`
}

// PolicyCustomRuleInput is an input type that accepts PolicyCustomRuleArgs and PolicyCustomRuleOutput values.
// You can construct a concrete instance of `PolicyCustomRuleInput` via:
//
// 		 PolicyCustomRuleArgs{...}
//
type PolicyCustomRuleInput interface {
	pulumi.Input

	ToPolicyCustomRuleOutput() PolicyCustomRuleOutput
	ToPolicyCustomRuleOutputWithContext(context.Context) PolicyCustomRuleOutput
}

type PolicyCustomRuleArgs struct {
	// Type of Actions
	Action pulumi.StringInput `pulumi:"action"`
	// One or more `matchCondition` block defined below.
	MatchConditions PolicyCustomRuleMatchConditionArrayInput `pulumi:"matchConditions"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value
	Priority pulumi.IntInput `pulumi:"priority"`
	// Describes the type of rule
	RuleType pulumi.StringInput `pulumi:"ruleType"`
}

func (PolicyCustomRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyCustomRule)(nil)).Elem()
}

func (i PolicyCustomRuleArgs) ToPolicyCustomRuleOutput() PolicyCustomRuleOutput {
	return i.ToPolicyCustomRuleOutputWithContext(context.Background())
}

func (i PolicyCustomRuleArgs) ToPolicyCustomRuleOutputWithContext(ctx context.Context) PolicyCustomRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCustomRuleOutput)
}

// PolicyCustomRuleArrayInput is an input type that accepts PolicyCustomRuleArray and PolicyCustomRuleArrayOutput values.
// You can construct a concrete instance of `PolicyCustomRuleArrayInput` via:
//
// 		 PolicyCustomRuleArray{ PolicyCustomRuleArgs{...} }
//
type PolicyCustomRuleArrayInput interface {
	pulumi.Input

	ToPolicyCustomRuleArrayOutput() PolicyCustomRuleArrayOutput
	ToPolicyCustomRuleArrayOutputWithContext(context.Context) PolicyCustomRuleArrayOutput
}

type PolicyCustomRuleArray []PolicyCustomRuleInput

func (PolicyCustomRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyCustomRule)(nil)).Elem()
}

func (i PolicyCustomRuleArray) ToPolicyCustomRuleArrayOutput() PolicyCustomRuleArrayOutput {
	return i.ToPolicyCustomRuleArrayOutputWithContext(context.Background())
}

func (i PolicyCustomRuleArray) ToPolicyCustomRuleArrayOutputWithContext(ctx context.Context) PolicyCustomRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCustomRuleArrayOutput)
}

type PolicyCustomRuleOutput struct{ *pulumi.OutputState }

func (PolicyCustomRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyCustomRule)(nil)).Elem()
}

func (o PolicyCustomRuleOutput) ToPolicyCustomRuleOutput() PolicyCustomRuleOutput {
	return o
}

func (o PolicyCustomRuleOutput) ToPolicyCustomRuleOutputWithContext(ctx context.Context) PolicyCustomRuleOutput {
	return o
}

// Type of Actions
func (o PolicyCustomRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyCustomRule) string { return v.Action }).(pulumi.StringOutput)
}

// One or more `matchCondition` block defined below.
func (o PolicyCustomRuleOutput) MatchConditions() PolicyCustomRuleMatchConditionArrayOutput {
	return o.ApplyT(func(v PolicyCustomRule) []PolicyCustomRuleMatchCondition { return v.MatchConditions }).(PolicyCustomRuleMatchConditionArrayOutput)
}

// The name of the policy. Changing this forces a new resource to be created.
func (o PolicyCustomRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyCustomRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value
func (o PolicyCustomRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v PolicyCustomRule) int { return v.Priority }).(pulumi.IntOutput)
}

// Describes the type of rule
func (o PolicyCustomRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyCustomRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type PolicyCustomRuleArrayOutput struct{ *pulumi.OutputState }

func (PolicyCustomRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyCustomRule)(nil)).Elem()
}

func (o PolicyCustomRuleArrayOutput) ToPolicyCustomRuleArrayOutput() PolicyCustomRuleArrayOutput {
	return o
}

func (o PolicyCustomRuleArrayOutput) ToPolicyCustomRuleArrayOutputWithContext(ctx context.Context) PolicyCustomRuleArrayOutput {
	return o
}

func (o PolicyCustomRuleArrayOutput) Index(i pulumi.IntInput) PolicyCustomRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyCustomRule {
		return vs[0].([]PolicyCustomRule)[vs[1].(int)]
	}).(PolicyCustomRuleOutput)
}

type PolicyCustomRuleMatchCondition struct {
	// Match value
	MatchValues []string `pulumi:"matchValues"`
	// One or more `matchVariable` block defined below.
	MatchVariables []PolicyCustomRuleMatchConditionMatchVariable `pulumi:"matchVariables"`
	// Describes if this is negate condition or not
	NegationCondition *bool `pulumi:"negationCondition"`
	// Describes operator to be matched
	Operator string `pulumi:"operator"`
}

// PolicyCustomRuleMatchConditionInput is an input type that accepts PolicyCustomRuleMatchConditionArgs and PolicyCustomRuleMatchConditionOutput values.
// You can construct a concrete instance of `PolicyCustomRuleMatchConditionInput` via:
//
// 		 PolicyCustomRuleMatchConditionArgs{...}
//
type PolicyCustomRuleMatchConditionInput interface {
	pulumi.Input

	ToPolicyCustomRuleMatchConditionOutput() PolicyCustomRuleMatchConditionOutput
	ToPolicyCustomRuleMatchConditionOutputWithContext(context.Context) PolicyCustomRuleMatchConditionOutput
}

type PolicyCustomRuleMatchConditionArgs struct {
	// Match value
	MatchValues pulumi.StringArrayInput `pulumi:"matchValues"`
	// One or more `matchVariable` block defined below.
	MatchVariables PolicyCustomRuleMatchConditionMatchVariableArrayInput `pulumi:"matchVariables"`
	// Describes if this is negate condition or not
	NegationCondition pulumi.BoolPtrInput `pulumi:"negationCondition"`
	// Describes operator to be matched
	Operator pulumi.StringInput `pulumi:"operator"`
}

func (PolicyCustomRuleMatchConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyCustomRuleMatchCondition)(nil)).Elem()
}

func (i PolicyCustomRuleMatchConditionArgs) ToPolicyCustomRuleMatchConditionOutput() PolicyCustomRuleMatchConditionOutput {
	return i.ToPolicyCustomRuleMatchConditionOutputWithContext(context.Background())
}

func (i PolicyCustomRuleMatchConditionArgs) ToPolicyCustomRuleMatchConditionOutputWithContext(ctx context.Context) PolicyCustomRuleMatchConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCustomRuleMatchConditionOutput)
}

// PolicyCustomRuleMatchConditionArrayInput is an input type that accepts PolicyCustomRuleMatchConditionArray and PolicyCustomRuleMatchConditionArrayOutput values.
// You can construct a concrete instance of `PolicyCustomRuleMatchConditionArrayInput` via:
//
// 		 PolicyCustomRuleMatchConditionArray{ PolicyCustomRuleMatchConditionArgs{...} }
//
type PolicyCustomRuleMatchConditionArrayInput interface {
	pulumi.Input

	ToPolicyCustomRuleMatchConditionArrayOutput() PolicyCustomRuleMatchConditionArrayOutput
	ToPolicyCustomRuleMatchConditionArrayOutputWithContext(context.Context) PolicyCustomRuleMatchConditionArrayOutput
}

type PolicyCustomRuleMatchConditionArray []PolicyCustomRuleMatchConditionInput

func (PolicyCustomRuleMatchConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyCustomRuleMatchCondition)(nil)).Elem()
}

func (i PolicyCustomRuleMatchConditionArray) ToPolicyCustomRuleMatchConditionArrayOutput() PolicyCustomRuleMatchConditionArrayOutput {
	return i.ToPolicyCustomRuleMatchConditionArrayOutputWithContext(context.Background())
}

func (i PolicyCustomRuleMatchConditionArray) ToPolicyCustomRuleMatchConditionArrayOutputWithContext(ctx context.Context) PolicyCustomRuleMatchConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCustomRuleMatchConditionArrayOutput)
}

type PolicyCustomRuleMatchConditionOutput struct{ *pulumi.OutputState }

func (PolicyCustomRuleMatchConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyCustomRuleMatchCondition)(nil)).Elem()
}

func (o PolicyCustomRuleMatchConditionOutput) ToPolicyCustomRuleMatchConditionOutput() PolicyCustomRuleMatchConditionOutput {
	return o
}

func (o PolicyCustomRuleMatchConditionOutput) ToPolicyCustomRuleMatchConditionOutputWithContext(ctx context.Context) PolicyCustomRuleMatchConditionOutput {
	return o
}

// Match value
func (o PolicyCustomRuleMatchConditionOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PolicyCustomRuleMatchCondition) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

// One or more `matchVariable` block defined below.
func (o PolicyCustomRuleMatchConditionOutput) MatchVariables() PolicyCustomRuleMatchConditionMatchVariableArrayOutput {
	return o.ApplyT(func(v PolicyCustomRuleMatchCondition) []PolicyCustomRuleMatchConditionMatchVariable {
		return v.MatchVariables
	}).(PolicyCustomRuleMatchConditionMatchVariableArrayOutput)
}

// Describes if this is negate condition or not
func (o PolicyCustomRuleMatchConditionOutput) NegationCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PolicyCustomRuleMatchCondition) *bool { return v.NegationCondition }).(pulumi.BoolPtrOutput)
}

// Describes operator to be matched
func (o PolicyCustomRuleMatchConditionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyCustomRuleMatchCondition) string { return v.Operator }).(pulumi.StringOutput)
}

type PolicyCustomRuleMatchConditionArrayOutput struct{ *pulumi.OutputState }

func (PolicyCustomRuleMatchConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyCustomRuleMatchCondition)(nil)).Elem()
}

func (o PolicyCustomRuleMatchConditionArrayOutput) ToPolicyCustomRuleMatchConditionArrayOutput() PolicyCustomRuleMatchConditionArrayOutput {
	return o
}

func (o PolicyCustomRuleMatchConditionArrayOutput) ToPolicyCustomRuleMatchConditionArrayOutputWithContext(ctx context.Context) PolicyCustomRuleMatchConditionArrayOutput {
	return o
}

func (o PolicyCustomRuleMatchConditionArrayOutput) Index(i pulumi.IntInput) PolicyCustomRuleMatchConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyCustomRuleMatchCondition {
		return vs[0].([]PolicyCustomRuleMatchCondition)[vs[1].(int)]
	}).(PolicyCustomRuleMatchConditionOutput)
}

type PolicyCustomRuleMatchConditionMatchVariable struct {
	// Describes field of the matchVariable collection
	Selector *string `pulumi:"selector"`
	// The name of the Match Variable
	VariableName string `pulumi:"variableName"`
}

// PolicyCustomRuleMatchConditionMatchVariableInput is an input type that accepts PolicyCustomRuleMatchConditionMatchVariableArgs and PolicyCustomRuleMatchConditionMatchVariableOutput values.
// You can construct a concrete instance of `PolicyCustomRuleMatchConditionMatchVariableInput` via:
//
// 		 PolicyCustomRuleMatchConditionMatchVariableArgs{...}
//
type PolicyCustomRuleMatchConditionMatchVariableInput interface {
	pulumi.Input

	ToPolicyCustomRuleMatchConditionMatchVariableOutput() PolicyCustomRuleMatchConditionMatchVariableOutput
	ToPolicyCustomRuleMatchConditionMatchVariableOutputWithContext(context.Context) PolicyCustomRuleMatchConditionMatchVariableOutput
}

type PolicyCustomRuleMatchConditionMatchVariableArgs struct {
	// Describes field of the matchVariable collection
	Selector pulumi.StringPtrInput `pulumi:"selector"`
	// The name of the Match Variable
	VariableName pulumi.StringInput `pulumi:"variableName"`
}

func (PolicyCustomRuleMatchConditionMatchVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyCustomRuleMatchConditionMatchVariable)(nil)).Elem()
}

func (i PolicyCustomRuleMatchConditionMatchVariableArgs) ToPolicyCustomRuleMatchConditionMatchVariableOutput() PolicyCustomRuleMatchConditionMatchVariableOutput {
	return i.ToPolicyCustomRuleMatchConditionMatchVariableOutputWithContext(context.Background())
}

func (i PolicyCustomRuleMatchConditionMatchVariableArgs) ToPolicyCustomRuleMatchConditionMatchVariableOutputWithContext(ctx context.Context) PolicyCustomRuleMatchConditionMatchVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCustomRuleMatchConditionMatchVariableOutput)
}

// PolicyCustomRuleMatchConditionMatchVariableArrayInput is an input type that accepts PolicyCustomRuleMatchConditionMatchVariableArray and PolicyCustomRuleMatchConditionMatchVariableArrayOutput values.
// You can construct a concrete instance of `PolicyCustomRuleMatchConditionMatchVariableArrayInput` via:
//
// 		 PolicyCustomRuleMatchConditionMatchVariableArray{ PolicyCustomRuleMatchConditionMatchVariableArgs{...} }
//
type PolicyCustomRuleMatchConditionMatchVariableArrayInput interface {
	pulumi.Input

	ToPolicyCustomRuleMatchConditionMatchVariableArrayOutput() PolicyCustomRuleMatchConditionMatchVariableArrayOutput
	ToPolicyCustomRuleMatchConditionMatchVariableArrayOutputWithContext(context.Context) PolicyCustomRuleMatchConditionMatchVariableArrayOutput
}

type PolicyCustomRuleMatchConditionMatchVariableArray []PolicyCustomRuleMatchConditionMatchVariableInput

func (PolicyCustomRuleMatchConditionMatchVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyCustomRuleMatchConditionMatchVariable)(nil)).Elem()
}

func (i PolicyCustomRuleMatchConditionMatchVariableArray) ToPolicyCustomRuleMatchConditionMatchVariableArrayOutput() PolicyCustomRuleMatchConditionMatchVariableArrayOutput {
	return i.ToPolicyCustomRuleMatchConditionMatchVariableArrayOutputWithContext(context.Background())
}

func (i PolicyCustomRuleMatchConditionMatchVariableArray) ToPolicyCustomRuleMatchConditionMatchVariableArrayOutputWithContext(ctx context.Context) PolicyCustomRuleMatchConditionMatchVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCustomRuleMatchConditionMatchVariableArrayOutput)
}

type PolicyCustomRuleMatchConditionMatchVariableOutput struct{ *pulumi.OutputState }

func (PolicyCustomRuleMatchConditionMatchVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyCustomRuleMatchConditionMatchVariable)(nil)).Elem()
}

func (o PolicyCustomRuleMatchConditionMatchVariableOutput) ToPolicyCustomRuleMatchConditionMatchVariableOutput() PolicyCustomRuleMatchConditionMatchVariableOutput {
	return o
}

func (o PolicyCustomRuleMatchConditionMatchVariableOutput) ToPolicyCustomRuleMatchConditionMatchVariableOutputWithContext(ctx context.Context) PolicyCustomRuleMatchConditionMatchVariableOutput {
	return o
}

// Describes field of the matchVariable collection
func (o PolicyCustomRuleMatchConditionMatchVariableOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyCustomRuleMatchConditionMatchVariable) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

// The name of the Match Variable
func (o PolicyCustomRuleMatchConditionMatchVariableOutput) VariableName() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyCustomRuleMatchConditionMatchVariable) string { return v.VariableName }).(pulumi.StringOutput)
}

type PolicyCustomRuleMatchConditionMatchVariableArrayOutput struct{ *pulumi.OutputState }

func (PolicyCustomRuleMatchConditionMatchVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyCustomRuleMatchConditionMatchVariable)(nil)).Elem()
}

func (o PolicyCustomRuleMatchConditionMatchVariableArrayOutput) ToPolicyCustomRuleMatchConditionMatchVariableArrayOutput() PolicyCustomRuleMatchConditionMatchVariableArrayOutput {
	return o
}

func (o PolicyCustomRuleMatchConditionMatchVariableArrayOutput) ToPolicyCustomRuleMatchConditionMatchVariableArrayOutputWithContext(ctx context.Context) PolicyCustomRuleMatchConditionMatchVariableArrayOutput {
	return o
}

func (o PolicyCustomRuleMatchConditionMatchVariableArrayOutput) Index(i pulumi.IntInput) PolicyCustomRuleMatchConditionMatchVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyCustomRuleMatchConditionMatchVariable {
		return vs[0].([]PolicyCustomRuleMatchConditionMatchVariable)[vs[1].(int)]
	}).(PolicyCustomRuleMatchConditionMatchVariableOutput)
}

type PolicyPolicySettings struct {
	// Describes if the policy is in enabled state or disabled state Defaults to `Enabled`.
	Enabled *bool `pulumi:"enabled"`
	// Describes if it is in detection mode  or prevention mode at the policy level Defaults to `Prevention`.
	Mode *string `pulumi:"mode"`
}

// PolicyPolicySettingsInput is an input type that accepts PolicyPolicySettingsArgs and PolicyPolicySettingsOutput values.
// You can construct a concrete instance of `PolicyPolicySettingsInput` via:
//
// 		 PolicyPolicySettingsArgs{...}
//
type PolicyPolicySettingsInput interface {
	pulumi.Input

	ToPolicyPolicySettingsOutput() PolicyPolicySettingsOutput
	ToPolicyPolicySettingsOutputWithContext(context.Context) PolicyPolicySettingsOutput
}

type PolicyPolicySettingsArgs struct {
	// Describes if the policy is in enabled state or disabled state Defaults to `Enabled`.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
	// Describes if it is in detection mode  or prevention mode at the policy level Defaults to `Prevention`.
	Mode pulumi.StringPtrInput `pulumi:"mode"`
}

func (PolicyPolicySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyPolicySettings)(nil)).Elem()
}

func (i PolicyPolicySettingsArgs) ToPolicyPolicySettingsOutput() PolicyPolicySettingsOutput {
	return i.ToPolicyPolicySettingsOutputWithContext(context.Background())
}

func (i PolicyPolicySettingsArgs) ToPolicyPolicySettingsOutputWithContext(ctx context.Context) PolicyPolicySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPolicySettingsOutput)
}

func (i PolicyPolicySettingsArgs) ToPolicyPolicySettingsPtrOutput() PolicyPolicySettingsPtrOutput {
	return i.ToPolicyPolicySettingsPtrOutputWithContext(context.Background())
}

func (i PolicyPolicySettingsArgs) ToPolicyPolicySettingsPtrOutputWithContext(ctx context.Context) PolicyPolicySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPolicySettingsOutput).ToPolicyPolicySettingsPtrOutputWithContext(ctx)
}

// PolicyPolicySettingsPtrInput is an input type that accepts PolicyPolicySettingsArgs, PolicyPolicySettingsPtr and PolicyPolicySettingsPtrOutput values.
// You can construct a concrete instance of `PolicyPolicySettingsPtrInput` via:
//
// 		 PolicyPolicySettingsArgs{...}
//
//  or:
//
// 		 nil
//
type PolicyPolicySettingsPtrInput interface {
	pulumi.Input

	ToPolicyPolicySettingsPtrOutput() PolicyPolicySettingsPtrOutput
	ToPolicyPolicySettingsPtrOutputWithContext(context.Context) PolicyPolicySettingsPtrOutput
}

type policyPolicySettingsPtrType PolicyPolicySettingsArgs

func PolicyPolicySettingsPtr(v *PolicyPolicySettingsArgs) PolicyPolicySettingsPtrInput {
	return (*policyPolicySettingsPtrType)(v)
}

func (*policyPolicySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyPolicySettings)(nil)).Elem()
}

func (i *policyPolicySettingsPtrType) ToPolicyPolicySettingsPtrOutput() PolicyPolicySettingsPtrOutput {
	return i.ToPolicyPolicySettingsPtrOutputWithContext(context.Background())
}

func (i *policyPolicySettingsPtrType) ToPolicyPolicySettingsPtrOutputWithContext(ctx context.Context) PolicyPolicySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPolicySettingsPtrOutput)
}

type PolicyPolicySettingsOutput struct{ *pulumi.OutputState }

func (PolicyPolicySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyPolicySettings)(nil)).Elem()
}

func (o PolicyPolicySettingsOutput) ToPolicyPolicySettingsOutput() PolicyPolicySettingsOutput {
	return o
}

func (o PolicyPolicySettingsOutput) ToPolicyPolicySettingsOutputWithContext(ctx context.Context) PolicyPolicySettingsOutput {
	return o
}

func (o PolicyPolicySettingsOutput) ToPolicyPolicySettingsPtrOutput() PolicyPolicySettingsPtrOutput {
	return o.ToPolicyPolicySettingsPtrOutputWithContext(context.Background())
}

func (o PolicyPolicySettingsOutput) ToPolicyPolicySettingsPtrOutputWithContext(ctx context.Context) PolicyPolicySettingsPtrOutput {
	return o.ApplyT(func(v PolicyPolicySettings) *PolicyPolicySettings {
		return &v
	}).(PolicyPolicySettingsPtrOutput)
}

// Describes if the policy is in enabled state or disabled state Defaults to `Enabled`.
func (o PolicyPolicySettingsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PolicyPolicySettings) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Describes if it is in detection mode  or prevention mode at the policy level Defaults to `Prevention`.
func (o PolicyPolicySettingsOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyPolicySettings) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type PolicyPolicySettingsPtrOutput struct{ *pulumi.OutputState }

func (PolicyPolicySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyPolicySettings)(nil)).Elem()
}

func (o PolicyPolicySettingsPtrOutput) ToPolicyPolicySettingsPtrOutput() PolicyPolicySettingsPtrOutput {
	return o
}

func (o PolicyPolicySettingsPtrOutput) ToPolicyPolicySettingsPtrOutputWithContext(ctx context.Context) PolicyPolicySettingsPtrOutput {
	return o
}

func (o PolicyPolicySettingsPtrOutput) Elem() PolicyPolicySettingsOutput {
	return o.ApplyT(func(v *PolicyPolicySettings) PolicyPolicySettings { return *v }).(PolicyPolicySettingsOutput)
}

// Describes if the policy is in enabled state or disabled state Defaults to `Enabled`.
func (o PolicyPolicySettingsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyPolicySettings) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

// Describes if it is in detection mode  or prevention mode at the policy level Defaults to `Prevention`.
func (o PolicyPolicySettingsPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyPolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyCustomRuleOutput{})
	pulumi.RegisterOutputType(PolicyCustomRuleArrayOutput{})
	pulumi.RegisterOutputType(PolicyCustomRuleMatchConditionOutput{})
	pulumi.RegisterOutputType(PolicyCustomRuleMatchConditionArrayOutput{})
	pulumi.RegisterOutputType(PolicyCustomRuleMatchConditionMatchVariableOutput{})
	pulumi.RegisterOutputType(PolicyCustomRuleMatchConditionMatchVariableArrayOutput{})
	pulumi.RegisterOutputType(PolicyPolicySettingsOutput{})
	pulumi.RegisterOutputType(PolicyPolicySettingsPtrOutput{})
}
