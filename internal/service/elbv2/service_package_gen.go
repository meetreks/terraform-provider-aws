// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package elbv2

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceLoadBalancer,
			TypeName: "aws_alb",
		},
		{
			Factory:  DataSourceListener,
			TypeName: "aws_alb_listener",
		},
		{
			Factory:  DataSourceTargetGroup,
			TypeName: "aws_alb_target_group",
		},
		{
			Factory:  DataSourceLoadBalancer,
			TypeName: "aws_lb",
		},
		{
			Factory:  DataSourceHostedZoneID,
			TypeName: "aws_lb_hosted_zone_id",
		},
		{
			Factory:  DataSourceListener,
			TypeName: "aws_lb_listener",
		},
		{
			Factory:  DataSourceTargetGroup,
			TypeName: "aws_lb_target_group",
		},
		{
			Factory:  DataSourceLoadBalancers,
			TypeName: "aws_lbs",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceLoadBalancer,
			TypeName: "aws_alb",
		},
		{
			Factory:  ResourceListener,
			TypeName: "aws_alb_listener",
		},
		{
			Factory:  ResourceListenerCertificate,
			TypeName: "aws_alb_listener_certificate",
		},
		{
			Factory:  ResourceListenerRule,
			TypeName: "aws_alb_listener_rule",
		},
		{
			Factory:  ResourceTargetGroup,
			TypeName: "aws_alb_target_group",
		},
		{
			Factory:  ResourceTargetGroupAttachment,
			TypeName: "aws_alb_target_group_attachment",
		},
		{
			Factory:  ResourceLoadBalancer,
			TypeName: "aws_lb",
		},
		{
			Factory:  ResourceListener,
			TypeName: "aws_lb_listener",
		},
		{
			Factory:  ResourceListenerCertificate,
			TypeName: "aws_lb_listener_certificate",
		},
		{
			Factory:  ResourceListenerRule,
			TypeName: "aws_lb_listener_rule",
		},
		{
			Factory:  ResourceTargetGroup,
			TypeName: "aws_lb_target_group",
		},
		{
			Factory:  ResourceTargetGroupAttachment,
			TypeName: "aws_lb_target_group_attachment",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ELBV2
}

var ServicePackage = &servicePackage{}