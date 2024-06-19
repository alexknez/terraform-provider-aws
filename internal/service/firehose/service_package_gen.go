// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package firehose

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	firehose_sdkv2 "github.com/aws/aws-sdk-go-v2/service/firehose"
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
			Factory:  dataSourceDeliveryStream,
			TypeName: "aws_kinesis_firehose_delivery_stream",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceDeliveryStream,
			TypeName: "aws_kinesis_firehose_delivery_stream",
			Name:     "Delivery Stream",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrName,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Firehose
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*firehose_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return firehose_sdkv2.NewFromConfig(cfg,
		firehose_sdkv2.WithEndpointResolverV2(newEndpointResolver()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func withBaseEndpoint(endpoint string) func(*firehose_sdkv2.Options) {
	return func(o *firehose_sdkv2.Options) {
		if endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
