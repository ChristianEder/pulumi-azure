// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package policy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Policy Remediation at the specified Scope.
type Remediation struct {
	pulumi.CustomResourceState

	// A list of the resource locations that will be remediated.
	LocationFilters pulumi.StringArrayOutput `pulumi:"locationFilters"`
	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId pulumi.StringOutput `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrOutput `pulumi:"policyDefinitionReferenceId"`
	// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
	Scope pulumi.StringOutput `pulumi:"scope"`
}

// NewRemediation registers a new resource with the given unique name, arguments, and options.
func NewRemediation(ctx *pulumi.Context,
	name string, args *RemediationArgs, opts ...pulumi.ResourceOption) (*Remediation, error) {
	if args == nil || args.PolicyAssignmentId == nil {
		return nil, errors.New("missing required argument 'PolicyAssignmentId'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &RemediationArgs{}
	}
	var resource Remediation
	err := ctx.RegisterResource("azure:policy/remediation:Remediation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemediation gets an existing Remediation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemediation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationState, opts ...pulumi.ResourceOption) (*Remediation, error) {
	var resource Remediation
	err := ctx.ReadResource("azure:policy/remediation:Remediation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Remediation resources.
type remediationState struct {
	// A list of the resource locations that will be remediated.
	LocationFilters []string `pulumi:"locationFilters"`
	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId *string `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
	Scope *string `pulumi:"scope"`
}

type RemediationState struct {
	// A list of the resource locations that will be remediated.
	LocationFilters pulumi.StringArrayInput
	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId pulumi.StringPtrInput
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
	Scope pulumi.StringPtrInput
}

func (RemediationState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationState)(nil)).Elem()
}

type remediationArgs struct {
	// A list of the resource locations that will be remediated.
	LocationFilters []string `pulumi:"locationFilters"`
	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId string `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a Remediation resource.
type RemediationArgs struct {
	// A list of the resource locations that will be remediated.
	LocationFilters pulumi.StringArrayInput
	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId pulumi.StringInput
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
	Scope pulumi.StringInput
}

func (RemediationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationArgs)(nil)).Elem()
}
