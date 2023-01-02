data "aws_secretsmanager_secret" "example" {
  name = "appd/partner-demo"
}

data "aws_secretsmanager_secret_version" "example" {
  secret_id = data.aws_secretsmanager_secret.example.id
}

output "credentials" {
  value = jsondecode(data.aws_secretsmanager_secret_version.example.secret_string)
  sensitive = true
}