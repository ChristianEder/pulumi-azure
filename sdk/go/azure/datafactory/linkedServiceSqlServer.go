// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manage a Linked Service (connection) between a SQL Server and Azure Data Factory.
// 
// > **Note:** All arguments including the client secret will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
type LinkedServiceSqlServer struct {
	s *pulumi.ResourceState
}

// NewLinkedServiceSqlServer registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceSqlServer(ctx *pulumi.Context,
	name string, args *LinkedServiceSqlServerArgs, opts ...pulumi.ResourceOpt) (*LinkedServiceSqlServer, error) {
	if args == nil || args.ConnectionString == nil {
		return nil, errors.New("missing required argument 'ConnectionString'")
	}
	if args == nil || args.DataFactoryName == nil {
		return nil, errors.New("missing required argument 'DataFactoryName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["additionalProperties"] = nil
		inputs["annotations"] = nil
		inputs["connectionString"] = nil
		inputs["dataFactoryName"] = nil
		inputs["description"] = nil
		inputs["integrationRuntimeName"] = nil
		inputs["name"] = nil
		inputs["parameters"] = nil
		inputs["resourceGroupName"] = nil
	} else {
		inputs["additionalProperties"] = args.AdditionalProperties
		inputs["annotations"] = args.Annotations
		inputs["connectionString"] = args.ConnectionString
		inputs["dataFactoryName"] = args.DataFactoryName
		inputs["description"] = args.Description
		inputs["integrationRuntimeName"] = args.IntegrationRuntimeName
		inputs["name"] = args.Name
		inputs["parameters"] = args.Parameters
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	s, err := ctx.RegisterResource("azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LinkedServiceSqlServer{s: s}, nil
}

// GetLinkedServiceSqlServer gets an existing LinkedServiceSqlServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceSqlServer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LinkedServiceSqlServerState, opts ...pulumi.ResourceOpt) (*LinkedServiceSqlServer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["additionalProperties"] = state.AdditionalProperties
		inputs["annotations"] = state.Annotations
		inputs["connectionString"] = state.ConnectionString
		inputs["dataFactoryName"] = state.DataFactoryName
		inputs["description"] = state.Description
		inputs["integrationRuntimeName"] = state.IntegrationRuntimeName
		inputs["name"] = state.Name
		inputs["parameters"] = state.Parameters
		inputs["resourceGroupName"] = state.ResourceGroupName
	}
	s, err := ctx.ReadResource("azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LinkedServiceSqlServer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LinkedServiceSqlServer) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LinkedServiceSqlServer) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
func (r *LinkedServiceSqlServer) AdditionalProperties() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["additionalProperties"])
}

// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
func (r *LinkedServiceSqlServer) Annotations() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["annotations"])
}

// The connection string in which to authenticate with the SQL Server.
func (r *LinkedServiceSqlServer) ConnectionString() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["connectionString"])
}

// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
func (r *LinkedServiceSqlServer) DataFactoryName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dataFactoryName"])
}

// The description for the Data Factory Linked Service SQL Server.
func (r *LinkedServiceSqlServer) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
func (r *LinkedServiceSqlServer) IntegrationRuntimeName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["integrationRuntimeName"])
}

// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
func (r *LinkedServiceSqlServer) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A map of parameters to associate with the Data Factory Linked Service SQL Server.
func (r *LinkedServiceSqlServer) Parameters() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["parameters"])
}

// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
func (r *LinkedServiceSqlServer) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Input properties used for looking up and filtering LinkedServiceSqlServer resources.
type LinkedServiceSqlServerState struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties interface{}
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations interface{}
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString interface{}
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName interface{}
	// The description for the Data Factory Linked Service SQL Server.
	Description interface{}
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName interface{}
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name interface{}
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters interface{}
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName interface{}
}

// The set of arguments for constructing a LinkedServiceSqlServer resource.
type LinkedServiceSqlServerArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties interface{}
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations interface{}
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString interface{}
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName interface{}
	// The description for the Data Factory Linked Service SQL Server.
	Description interface{}
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName interface{}
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name interface{}
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters interface{}
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName interface{}
}