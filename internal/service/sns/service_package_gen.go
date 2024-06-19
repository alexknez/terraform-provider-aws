// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package sns

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	sns_sdkv2 "github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
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
			Factory:  dataSourceTopic,
			TypeName: "aws_sns_topic",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourcePlatformApplication,
			TypeName: "aws_sns_platform_application",
		},
		{
			Factory:  resourceSMSPreferences,
			TypeName: "aws_sns_sms_preferences",
		},
		{
			Factory:  resourceTopic,
			TypeName: "aws_sns_topic",
			Name:     "Topic",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTopicDataProtectionPolicy,
			TypeName: "aws_sns_topic_data_protection_policy",
		},
		{
			Factory:  resourceTopicPolicy,
			TypeName: "aws_sns_topic_policy",
		},
		{
			Factory:  resourceTopicSubscription,
			TypeName: "aws_sns_topic_subscription",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SNS
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*sns_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return sns_sdkv2.NewFromConfig(cfg,
		sns_sdkv2.WithEndpointResolverV2(newEndpointResolver()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func withBaseEndpoint(endpoint string) func(*sns_sdkv2.Options) {
	return func(o *sns_sdkv2.Options) {
		if endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
