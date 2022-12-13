resource "appdynamicscloud_connection_aws_role_delegation_credentials" "example" {
  display_name     = "AWS Dev"
  description      = "Description for this AWS role delegation connection"
  account_id        = "0000-0000-0000"
}


resource "aws_iam_role" "role" {
  name = "test-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
          "AWS": "arn:aws:iam::${appdynamicscloud_connection_aws_role_delegation_credentials.example.appdynamics_aws_account_id}:root"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
          "StringEquals": {
              "sts:ExternalId": "${appdynamicscloud_connection_aws_role_delegation_credentials.example.external_id}"
          }
      }
    }
}
EOF
}

resource "aws_iam_policy" "policy" {
  name        = "test-policy"
  description = "A test policy"

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

resource "appdynamicscloud_connection_aws" "test" {
    connection_id = appdynamicscloud_connection_aws.example.connection_id
    role_name = aws_iam_role.role.role_name
    configuration_id = "000-000-0000"
    state = ""
}