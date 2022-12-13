resource "appdynamicscloud_connection_aws" "example" {
  display_name     = "AWS Dev"
  description      = "Description for this AWS role delegation connection"
  configuration_id = appdynamicscloud_connection_configuration_aws.example
  state            = "ACTIVE"

  details {
    access_type       = "role_delegation"
    account_id        = "0000-0000-0000"
  }
}


resource "aws_iam_role" "role" {
  name = "test-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
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
        "ec2:Describe*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "test-attach" {
  role       = aws_iam_role.role.name
  policy_arn = aws_iam_policy.policy.arn
}

resource "appdynamicscloud_connection_aws_role_attachment" "test" {
    connection_id = appdynamicscloud_connection_aws.example.connection_id
    role_name = aws_iam_role.role.role_name
}