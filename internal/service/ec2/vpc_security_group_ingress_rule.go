package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/experimental/intf"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func init() {
	registerFrameworkResourceFactory(newResourceSecurityGroupIngressRule)
}

// newResourceSecurityGroupIngressRule instantiates a new Resource for the aws_vpc_security_group_ingress_rule resource.
func newResourceSecurityGroupIngressRule(context.Context) (intf.ResourceWithConfigureAndImportState, error) {
	return &resourceSecurityGroupIngressRule{}, nil
}

type resourceSecurityGroupIngressRule struct {
	meta *conns.AWSClient
}

// Metadata should return the full name of the resource, such as
// examplecloud_thing.
func (r *resourceSecurityGroupIngressRule) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = "aws_vpc_security_group_ingress_rule"
}

// GetSchema returns the schema for this resource.
func (r *resourceSecurityGroupIngressRule) GetSchema(context.Context) (tfsdk.Schema, diag.Diagnostics) {
	schema := tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"arn": {
				Type:     types.StringType,
				Computed: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					resource.UseStateForUnknown(),
				},
			},
			"cidr_ipv4": {
				Type:     types.StringType,
				Optional: true,
			},
			"cidr_ipv6": {
				Type:     types.StringType,
				Optional: true,
			},
			"description": {
				Type:     types.StringType,
				Optional: true,
			},
			"from_port": {
				Type:     types.Int64Type,
				Optional: true,
			},
			"id": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
			"ip_protocol": {
				Type:     types.StringType,
				Required: true,
			},
			"tags":     tftags.TagsAttribute(),
			"tags_all": tftags.TagsAttributeComputed(),
			"to_port": {
				Type:     types.Int64Type,
				Optional: true,
			},
		},
	}

	return schema, nil
}

// Configure enables provider-level data or clients to be set in the
// provider-defined Resource type.
func (r *resourceSecurityGroupIngressRule) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if v, ok := request.ProviderData.(*conns.AWSClient); ok {
		r.meta = v
	}
}

// Create is called when the provider must create a new resource.
// Config and planned state values should be read from the CreateRequest and new state values set on the CreateResponse.
func (r *resourceSecurityGroupIngressRule) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var data resourceSecurityGroupIngressRuleData

	response.Diagnostics.Append(request.Plan.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	// conn := r.meta.EC2Conn
	defaultTagsConfig := r.meta.DefaultTagsConfig
	ignoreTagsConfig := r.meta.IgnoreTagsConfig
	tags := defaultTagsConfig.MergeTags(tftags.New(data.Tags))

	data.TagsAll = flex.FlattenFrameworkStringValueMap(ctx, tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map())

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

// Read is called when the provider must read resource values in order to update state.
// Planned state values should be read from the ReadRequest and new state values set on the ReadResponse.
func (r *resourceSecurityGroupIngressRule) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var data resourceSecurityGroupIngressRuleData

	response.Diagnostics.Append(request.State.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	// conn := r.meta.EC2Conn
	// defaultTagsConfig := r.meta.DefaultTagsConfig
	// ignoreTagsConfig := r.meta.IgnoreTagsConfig

	// tags = tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig)
	// data.Tags = flex.FlattenFrameworkStringValueMap(ctx, tags.RemoveDefaultConfig(defaultTagsConfig).Map())
	// data.TagsAll = flex.FlattenFrameworkStringValueMap(ctx, tags.Map())

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

// Update is called to update the state of the resource.
// Config, planned state, and prior state values should be read from the UpdateRequest and new state values set on the UpdateResponse.
func (r *resourceSecurityGroupIngressRule) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var old, new resourceSecurityGroupIngressRuleData

	response.Diagnostics.Append(request.State.Get(ctx, &old)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(request.Plan.Get(ctx, &new)...)

	if response.Diagnostics.HasError() {
		return
	}

	// conn := r.meta.EC2Conn

	if !new.TagsAll.Equal(old.TagsAll) {
	}

	response.Diagnostics.Append(response.State.Set(ctx, &new)...)
}

// Delete is called when the provider must delete the resource.
// Config values may be read from the DeleteRequest.
//
// If execution completes without error, the framework will automatically call DeleteResponse.State.RemoveResource(),
// so it can be omitted from provider logic.
func (r *resourceSecurityGroupIngressRule) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var data resourceSecurityGroupIngressRuleData

	response.Diagnostics.Append(request.State.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	// conn := r.meta.EC2Conn

	tflog.Debug(ctx, "deleting SWF Domain", map[string]interface{}{
		"id": data.ID.Value,
	})
}

// ImportState is called when the provider must import the state of a resource instance.
// This method must return enough state so the Read method can properly refresh the full resource.
//
// If setting an attribute with the import identifier, it is recommended to use the ImportStatePassthroughID() call in this method.
func (r *resourceSecurityGroupIngressRule) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
}

// ModifyPlan is called when the provider has an opportunity to modify
// the plan: once during the plan phase when Terraform is determining
// the diff that should be shown to the user for approval, and once
// during the apply phase with any unknown values from configuration
// filled in with their final values.
//
// The planned new state is represented by
// ModifyPlanResponse.Plan. It must meet the following
// constraints:
// 1. Any non-Computed attribute set in config must preserve the exact
// config value or return the corresponding attribute value from the
// prior state (ModifyPlanRequest.State).
// 2. Any attribute with a known value must not have its value changed
// in subsequent calls to ModifyPlan or Create/Read/Update.
// 3. Any attribute with an unknown value may either remain unknown
// or take on any value of the expected type.
//
// Any errors will prevent further resource-level plan modifications.
func (r *resourceSecurityGroupIngressRule) ModifyPlan(ctx context.Context, request resource.ModifyPlanRequest, response *resource.ModifyPlanResponse) {
	defaultTagsConfig := r.meta.DefaultTagsConfig
	ignoreTagsConfig := r.meta.IgnoreTagsConfig

	var planTags types.Map

	response.Diagnostics.Append(request.Plan.GetAttribute(ctx, path.Root("tags"), &planTags)...)

	if response.Diagnostics.HasError() {
		return
	}

	resourceTags := tftags.New(planTags)

	if defaultTagsConfig.TagsEqual(resourceTags) {
		response.Diagnostics.AddError(
			`"tags" are identical to those in the "default_tags" configuration block of the provider`,
			"please de-duplicate and try again")
	}

	allTags := defaultTagsConfig.MergeTags(resourceTags).IgnoreConfig(ignoreTagsConfig)

	response.Diagnostics.Append(response.Plan.SetAttribute(ctx, path.Root("tags_all"), flex.FlattenFrameworkStringValueMap(ctx, allTags.Map()))...)
}

type resourceSecurityGroupIngressRuleData struct {
	ARN         types.String `tfsdk:"arn"`
	CIDRIPv4    types.String `tfsdk:"cidr_ipv4"`
	CIDRIPv6    types.String `tfsdk:"cidr_ipv6"`
	Description types.String `tfsdk:"description"`
	FromPort    types.Int64  `tfsdk:"from_port"`
	ID          types.String `tfsdk:"id"`
	IPProtocol  types.String `tfsdk:"ip_protocol"`
	Tags        types.Map    `tfsdk:"tags"`
	TagsAll     types.Map    `tfsdk:"tags_all"`
	ToPort      types.Int64  `tfsdk:"to_port"`
}
