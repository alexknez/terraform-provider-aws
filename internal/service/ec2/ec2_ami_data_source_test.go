// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"testing"

	"github.com/YakDriver/regexache"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccEC2AMIDataSource_natInstance(t *testing.T) {
	ctx := acctest.Context(t)
	datasourceName := "data.aws_ami.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAMIDataSourceConfig_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check attributes. Some attributes are tough to test - any not contained here should not be considered
					// stable and should not be used in interpolation. Exception to block_device_mappings which should both
					// show up consistently and break if certain references are not available. However modification of the
					// snapshot ID which is bound to happen on the NAT AMIs will cause testing to break consistently, so
					// deep inspection is not included, simply the count is checked.
					// Tags and product codes may need more testing, but I'm having a hard time finding images with
					// these attributes set.
					resource.TestCheckResourceAttr(datasourceName, "architecture", "x86_64"),
					acctest.MatchResourceAttrRegionalARNNoAccount(datasourceName, names.AttrARN, "ec2", regexache.MustCompile(`image/ami-.+`)),
					resource.TestCheckResourceAttr(datasourceName, "block_device_mappings.#", acctest.CtOne),
					resource.TestMatchResourceAttr(datasourceName, names.AttrCreationDate, regexache.MustCompile("^20[0-9]{2}-")),
					resource.TestMatchResourceAttr(datasourceName, "deprecation_time", regexache.MustCompile("^20[0-9]{2}-")),
					resource.TestMatchResourceAttr(datasourceName, names.AttrDescription, regexache.MustCompile("^Amazon Linux AMI")),
					resource.TestCheckResourceAttr(datasourceName, "ena_support", "true"),
					resource.TestCheckResourceAttr(datasourceName, "hypervisor", "xen"),
					resource.TestMatchResourceAttr(datasourceName, "image_id", regexache.MustCompile("^ami-")),
					resource.TestMatchResourceAttr(datasourceName, "image_location", regexache.MustCompile("^amazon/")),
					resource.TestCheckResourceAttr(datasourceName, "image_owner_alias", "amazon"),
					resource.TestCheckResourceAttr(datasourceName, "image_type", "machine"),
					resource.TestCheckResourceAttr(datasourceName, "imds_support", ""),
					resource.TestCheckResourceAttr(datasourceName, names.AttrMostRecent, "true"),
					resource.TestMatchResourceAttr(datasourceName, names.AttrName, regexache.MustCompile("^amzn-ami-vpc-nat")),
					acctest.MatchResourceAttrAccountID(datasourceName, names.AttrOwnerID),
					resource.TestCheckResourceAttr(datasourceName, "platform_details", "Linux/UNIX"),
					resource.TestCheckResourceAttr(datasourceName, "product_codes.#", acctest.CtZero),
					resource.TestCheckResourceAttr(datasourceName, "public", "true"),
					resource.TestCheckResourceAttr(datasourceName, "root_device_name", "/dev/xvda"),
					resource.TestCheckResourceAttr(datasourceName, "root_device_type", "ebs"),
					resource.TestMatchResourceAttr(datasourceName, "root_snapshot_id", regexache.MustCompile("^snap-")),
					resource.TestCheckResourceAttr(datasourceName, "sriov_net_support", "simple"),
					resource.TestCheckResourceAttr(datasourceName, names.AttrState, "available"),
					resource.TestCheckResourceAttr(datasourceName, "state_reason.code", "UNSET"),
					resource.TestCheckResourceAttr(datasourceName, "state_reason.message", "UNSET"),
					resource.TestCheckResourceAttr(datasourceName, acctest.CtTagsPercent, acctest.CtZero),
					resource.TestCheckResourceAttr(datasourceName, "usage_operation", "RunInstances"),
					resource.TestCheckResourceAttr(datasourceName, "virtualization_type", "hvm"),
				),
			},
		},
	})
}

func TestAccEC2AMIDataSource_windowsInstance(t *testing.T) {
	ctx := acctest.Context(t)
	datasourceName := "data.aws_ami.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAMIDataSourceConfig_windows,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "architecture", "x86_64"),
					resource.TestCheckResourceAttr(datasourceName, "block_device_mappings.#", "27"),
					resource.TestMatchResourceAttr(datasourceName, names.AttrCreationDate, regexache.MustCompile("^20[0-9]{2}-")),
					resource.TestMatchResourceAttr(datasourceName, "deprecation_time", regexache.MustCompile("^20[0-9]{2}-")),
					resource.TestMatchResourceAttr(datasourceName, names.AttrDescription, regexache.MustCompile("^Microsoft Windows Server")),
					resource.TestCheckResourceAttr(datasourceName, "ena_support", "true"),
					resource.TestCheckResourceAttr(datasourceName, "hypervisor", "xen"),
					resource.TestMatchResourceAttr(datasourceName, "image_id", regexache.MustCompile("^ami-")),
					resource.TestMatchResourceAttr(datasourceName, "image_location", regexache.MustCompile("^amazon/")),
					resource.TestCheckResourceAttr(datasourceName, "image_owner_alias", "amazon"),
					resource.TestCheckResourceAttr(datasourceName, "image_type", "machine"),
					resource.TestCheckResourceAttr(datasourceName, names.AttrMostRecent, "true"),
					resource.TestMatchResourceAttr(datasourceName, names.AttrName, regexache.MustCompile("^Windows_Server-2012-R2")),
					acctest.MatchResourceAttrAccountID(datasourceName, names.AttrOwnerID),
					resource.TestCheckResourceAttr(datasourceName, "platform", "windows"),
					resource.TestMatchResourceAttr(datasourceName, "platform_details", regexache.MustCompile(`Windows`)),
					resource.TestCheckResourceAttr(datasourceName, "product_codes.#", acctest.CtZero),
					resource.TestCheckResourceAttr(datasourceName, "public", "true"),
					resource.TestCheckResourceAttr(datasourceName, "root_device_name", "/dev/sda1"),
					resource.TestCheckResourceAttr(datasourceName, "root_device_type", "ebs"),
					resource.TestMatchResourceAttr(datasourceName, "root_snapshot_id", regexache.MustCompile("^snap-")),
					resource.TestCheckResourceAttr(datasourceName, "sriov_net_support", "simple"),
					resource.TestCheckResourceAttr(datasourceName, names.AttrState, "available"),
					resource.TestCheckResourceAttr(datasourceName, "state_reason.code", "UNSET"),
					resource.TestCheckResourceAttr(datasourceName, "state_reason.message", "UNSET"),
					resource.TestCheckResourceAttr(datasourceName, acctest.CtTagsPercent, acctest.CtZero),
					resource.TestCheckResourceAttr(datasourceName, "tpm_support", ""),
					resource.TestMatchResourceAttr(datasourceName, "usage_operation", regexache.MustCompile(`^RunInstances`)),
					resource.TestCheckResourceAttr(datasourceName, "virtualization_type", "hvm"),
				),
			},
		},
	})
}

func TestAccEC2AMIDataSource_instanceStore(t *testing.T) {
	ctx := acctest.Context(t)
	datasourceName := "data.aws_ami.ubuntu-bionic-ami-hvm-instance-store"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAMIDataSourceConfig_latestUbuntuBionicHVMInstanceStore(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "architecture", "x86_64"),
					resource.TestCheckResourceAttr(datasourceName, "block_device_mappings.#", acctest.CtZero),
					resource.TestMatchResourceAttr(datasourceName, names.AttrCreationDate, regexache.MustCompile("^20[0-9]{2}-")),
					resource.TestMatchResourceAttr(datasourceName, "deprecation_time", regexache.MustCompile("^20[0-9]{2}-")),
					resource.TestCheckResourceAttr(datasourceName, "ena_support", "true"),
					resource.TestCheckResourceAttr(datasourceName, "hypervisor", "xen"),
					resource.TestMatchResourceAttr(datasourceName, "image_id", regexache.MustCompile("^ami-")),
					resource.TestMatchResourceAttr(datasourceName, "image_location", regexache.MustCompile(`ubuntu-images-.*-release/.*/.*/hvm/instance-store`)),
					resource.TestCheckResourceAttr(datasourceName, "image_type", "machine"),
					resource.TestCheckResourceAttr(datasourceName, names.AttrMostRecent, "true"),
					resource.TestMatchResourceAttr(datasourceName, names.AttrName, regexache.MustCompile(`ubuntu/images/hvm-instance/.*`)),
					acctest.MatchResourceAttrAccountID(datasourceName, names.AttrOwnerID),
					resource.TestCheckResourceAttr(datasourceName, "platform_details", "Linux/UNIX"),
					resource.TestCheckResourceAttr(datasourceName, "product_codes.#", acctest.CtZero),
					resource.TestCheckResourceAttr(datasourceName, "public", "true"),
					resource.TestCheckResourceAttr(datasourceName, "root_device_type", "instance-store"),
					resource.TestCheckResourceAttr(datasourceName, "root_snapshot_id", ""),
					resource.TestCheckResourceAttr(datasourceName, "sriov_net_support", "simple"),
					resource.TestCheckResourceAttr(datasourceName, names.AttrState, "available"),
					resource.TestCheckResourceAttr(datasourceName, "state_reason.code", "UNSET"),
					resource.TestCheckResourceAttr(datasourceName, "state_reason.message", "UNSET"),
					resource.TestCheckResourceAttr(datasourceName, acctest.CtTagsPercent, acctest.CtZero),
					resource.TestCheckResourceAttr(datasourceName, "tpm_support", ""),
					resource.TestCheckResourceAttr(datasourceName, "usage_operation", "RunInstances"),
					resource.TestCheckResourceAttr(datasourceName, "virtualization_type", "hvm"),
				),
			},
		},
	})
}

func TestAccEC2AMIDataSource_localNameFilter(t *testing.T) {
	ctx := acctest.Context(t)
	datasourceName := "data.aws_ami.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAMIDataSourceConfig_nameRegex,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(datasourceName, "image_id", regexache.MustCompile("^ami-")),
				),
			},
		},
	})
}

func TestAccEC2AMIDataSource_gp3BlockDevice(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ami.test"
	datasourceName := "data.aws_ami.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAMIDataSourceConfig_gp3BlockDevice(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "architecture", resourceName, "architecture"),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrARN, resourceName, names.AttrARN),
					resource.TestCheckResourceAttrPair(datasourceName, "block_device_mappings.#", resourceName, "ebs_block_device.#"),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrDescription, resourceName, names.AttrDescription),
					resource.TestCheckResourceAttrPair(datasourceName, "image_id", resourceName, names.AttrID),
					acctest.CheckResourceAttrAccountID(datasourceName, names.AttrOwnerID),
					resource.TestCheckResourceAttrPair(datasourceName, "root_device_name", resourceName, "root_device_name"),
					resource.TestCheckResourceAttr(datasourceName, "root_device_type", "ebs"),
					resource.TestCheckResourceAttrPair(datasourceName, "root_snapshot_id", resourceName, "root_snapshot_id"),
					resource.TestCheckResourceAttrPair(datasourceName, "sriov_net_support", resourceName, "sriov_net_support"),
					resource.TestCheckResourceAttrPair(datasourceName, acctest.CtTagsPercent, resourceName, acctest.CtTagsPercent),
					resource.TestCheckResourceAttrPair(datasourceName, "virtualization_type", resourceName, "virtualization_type"),
				),
			},
		},
	})
}

// testAccAMIDataSourceConfig_latestUbuntuBionicHVMInstanceStore returns the configuration for a data source that
// describes the latest Ubuntu 18.04 AMI using HVM virtualization and an instance store root device.
// The data source is named 'ubuntu-bionic-ami-hvm-instance-store'.
func testAccAMIDataSourceConfig_latestUbuntuBionicHVMInstanceStore() string {
	return `
data "aws_ami" "ubuntu-bionic-ami-hvm-instance-store" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-instance/ubuntu-bionic-18.04-amd64-server-*"]
  }

  filter {
    name   = "root-device-type"
    values = ["instance-store"]
  }
}
`
}

// Using NAT AMIs for testing - I would expect with NAT gateways now a thing,
// that this will possibly be deprecated at some point in time. Other candidates
// for testing this after that may be Ubuntu's AMI's, or Amazon's regular
// Amazon Linux AMIs.
const testAccAMIDataSourceConfig_basic = `
data "aws_ami" "test" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn-ami-vpc-nat*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  filter {
    name   = "root-device-type"
    values = ["ebs"]
  }

  filter {
    name   = "block-device-mapping.volume-type"
    values = ["standard"]
  }
}
`

// Windows image test.
const testAccAMIDataSourceConfig_windows = `
data "aws_ami" "test" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["Windows_Server-2012-R2*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  filter {
    name   = "root-device-type"
    values = ["ebs"]
  }

  filter {
    name   = "block-device-mapping.volume-type"
    values = ["gp2"]
  }
}
`

// Testing name_regex parameter
const testAccAMIDataSourceConfig_nameRegex = `
data "aws_ami" "test" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn-ami-*"]
  }

  name_regex = "^amzn-ami-min[a-z]{4}-hvm"
}
`

func testAccAMIDataSourceConfig_gp3BlockDevice(rName string) string {
	return acctest.ConfigCompose(
		testAccAMIConfig_gp3BlockDevice(rName),
		`
data "aws_caller_identity" "current" {}

data "aws_ami" "test" {
  owners = [data.aws_caller_identity.current.account_id]

  filter {
    name   = "image-id"
    values = [aws_ami.test.id]
  }
}
`)
}
