// Copyright 2023 Cisco Systems, Inc.
//
// Licensed under the MPL License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.mozilla.org/en-US/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	cloudconnectionapi "github.com/cisco-open/appd-cloud-go-client/apis/v1/cloudconnections"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var resourceConnectionAwsRoleAttachmentTest = map[string]string{
	"account_id": os.Getenv("TEST_AWS_ACCOUNT_ID"),
	"role_name":  makeTestVariable(acctest.RandString(5)),
}

var connection_id string

func TestAccAppdynamicscloudConnectionAwsRoleAttachment_Basic(t *testing.T) {
	var connectionAwsRoleAttachment_default cloudconnectionapi.ConnectionResponse
	resourceName := "appdynamicscloud_connection_aws_role_attachment.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		ExternalProviders: awsProvider,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsRoleAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config:      CreateAccConnectionAwsRoleAttachmentWithoutConnectionId(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccConnectionAwsRoleAttachmentWithoutRoleName(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config: CreateAccConnectionAwsRoleAttachmentConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudConnectionAwsRoleAttachmentExists(resourceName, &connectionAwsRoleAttachment_default),
					resource.TestCheckResourceAttr(resourceName, "role_name", resourceConnectionAwsRoleAttachmentTest["role_name"]),
				),
			},
			{
				Config:      CreateAccConnectionAwsRoleAttachmentWithInvalidRoleName("testing-invalid-role"),
				ExpectError: regexp.MustCompile(`Invalid Credentials Error`),
			},
			{
				Config: CreateAccConnectionAwsRoleAttachmentConfig(),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: CreateAccConnectionAwsRoleAttachmentConfig(),
			},
		},
	})
}

func CreateAccConnectionAwsRoleAttachmentWithInvalidRoleName(value string) string {
	var resource string
	resource += aws_role
	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_connection_aws_role_attachment" "test" {
			connection_id = appdynamicscloud_connection_aws.test.id
			role_name = "%v"
		}
	`, value)
	return resource
}

func CreateAccConnectionAwsRoleAttachmentWithoutConnectionId() string {
	var resource string
	resource += aws_role
	//lint:ignore S1039 consistency
	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_connection_aws_role_attachment" "test" {
						role_name = aws_iam_role.role.name
		}`)
	return resource
}
func CreateAccConnectionAwsRoleAttachmentWithoutRoleName() string {
	var resource string
	resource += aws_role
	//lint:ignore S1039 consistency
	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_connection_aws_role_attachment" "test" {
							connection_id = appdynamicscloud_connection_aws.test.id
		}`)
	return resource
}

func CreateAccConnectionAwsRoleAttachmentConfig() string {
	var resource string
	resource += aws_role
	//lint:ignore S1039 consistency
	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_connection_aws_role_attachment" "test" {
			connection_id = appdynamicscloud_connection_aws.test.id
			role_name = aws_iam_role.role.name
		}`)
	return resource
}

func testAccCheckAppdynamicscloudConnectionAwsRoleAttachmentExists(name string, connectionAwsRoleAttachment *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Connection AWS Role Attachment %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AWS Role Attachment connection id was set")
		}

		config := testAccProvider.Meta().(config)
		myctx, _, apiClient := initializeCloudConnectionClient(config)

		resp, _, err := apiClient.ConnectionsApi.GetConnection(myctx, rs.Primary.ID).Execute()
		if err != nil {
			return err
		}

		if resp.GetId() != rs.Primary.ID {
			return fmt.Errorf("Cloud Connection AWS %s not found", rs.Primary.ID)
		}
		connection_id = resp.GetId()
		*connectionAwsRoleAttachment = *resp
		return nil
	}
}

func testAccCheckAppdynamicscloudConnectionAwsRoleAttachmentDestroy(s *terraform.State) error {
	return nil
}

//lint:ignore U1000 might come in handy in the future
func testAccCheckAppdynamicscloudConnectionAwsRoleAttachmentIdEqual(connectionAwsRoleAttachment1, connectionAwsRoleAttachment2 *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if connectionAwsRoleAttachment1.Id != connectionAwsRoleAttachment2.Id {
			return fmt.Errorf("ConnectionAwsRoleAttachment IDs are not equal")
		}
		return nil
	}
}

//lint:ignore U1000 might come in handy in the future
func testAccCheckAppdynamicscloudConnectionAwsRoleAttachmentIdNotEqual(connectionAwsRoleAttachment1, connectionAwsRoleAttachment2 *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if connectionAwsRoleAttachment1.Id == connectionAwsRoleAttachment2.Id {
			return fmt.Errorf("ConnectionAwsRoleAttachment IDs are equal")
		}
		return nil
	}
}

var aws_role string = fmt.Sprintf(`
resource "appdynamicscloud_connection_aws" "test" {
	display_name     = "AWS ROLE TEST"
	description      = "Description for this AWS role delegation connection1"
	connection_details {
	  access_type       = "role_delegation"
	  account_id        = "%v"
	}
	configuration_details{
	  regions=["us-east-1"]
  
		services {
			name = "ec2"
			import_tags {
					enabled = false
					excluded_keys = []
			}
			# tag_filter = "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)"
			polling {
					interval = 5
					unit = "minute"
			}
		  }
		  services {
			name = "elb"
			import_tags {
					enabled = false
					excluded_keys = []
			}
			# tag_filter = "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)"
			polling {
					interval = 5
					unit = "minute"
			}
		  }
	}
  }

resource "aws_iam_role" "role" {
  name = "%v"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
          "AWS": "arn:aws:iam::${appdynamicscloud_connection_aws.test.connection_details[0].appdynamics_aws_account_id}:root"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
          "StringEquals": {
              "sts:ExternalId": "${appdynamicscloud_connection_aws.test.connection_details[0].external_id}"
          }
      }
    }
  ]
}
EOF
}

resource "aws_iam_policy" "policy" {
  name        = "%v"
  description = "A test policy "

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "elasticloadbalancing:DescribeLoadBalancers",
        "ec2:DescribeInstances",
        "cloudwatch:GetMetricData",
        "ec2:DescribeVpcs",
        "ec2:DescribeRegions",
        "ec2:DescribeVolumes",
        "elasticloadbalancing:DescribeTargetHealth",
        "rds:DescribeDBInstances",
        "elasticloadbalancing:DescribeTargetGroups",
        "ec2:DescribeSubnets",
        "cloudwatch:ListMetrics",
        "rds:DescribeDBClusters",
        "tag:GetResources",
        "ecs:ListClusters",
        "ecs:DescribeClusters",
        "ecs:ListServices",
        "ecs:DescribeServices",
        "ecs:ListTasks",
        "ecs:DescribeTasks",
        "ecs:DescribeContainerInstances",
        "ecs:ListTaskDefinitions"
      ],
      "Resource": "*",
      "Effect": "Allow",
      "Sid": "AllowMonitoring"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "test-attach" {
  role       = aws_iam_role.role.name
  policy_arn = aws_iam_policy.policy.arn
}
`, resourceConnectionAwsRoleAttachmentTest["account_id"], resourceConnectionAwsRoleAttachmentTest["role_name"], resourceConnectionAwsRoleAttachmentTest["role_name"])
