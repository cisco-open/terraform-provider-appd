# Example for Role attachment of AWS Connection 

resource "appdynamicscloud_connection_aws_role_attachment" "example" {
  connection_id = appdynamicscloud_connection_aws.test.id
  role_name     = aws_iam_role.role.name
}

# Example for Creating Role at AWS

resource "aws_iam_role" "role" {
  name = "test-role-1"

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

# Example for Creating Policy at AWS

resource "aws_iam_policy" "policy" {
  name        = "test-policy-1"
  description = "A test policy 1"

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

# Example for Role Policy Attachment at AWS

resource "aws_iam_role_policy_attachment" "test-attach" {
  role       = aws_iam_role.role.name
  policy_arn = aws_iam_policy.policy.arn
}
