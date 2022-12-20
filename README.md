## Terraform Provider AppDynamics Demo
---
configure vault with required credentials
```shell
vault kv put appd-partner-demo/azure_credentials 
    client_id=adb90c29-204d-43d9-987c-ab406d9199cc 
    client_secret=ZmD8Q~q87S6hbvMVFh5hKS3y8dLA9C1Xfc5jLbjx 
    subscription_id=fca41da2-4908-49e2-b0cb-d3d2080fc5be
    tenant_id=f3db65e4-ce5a-4d7e-a140-255bf017d87f

vault kv put appd-partner-demo/tenant_credentails_test_service_principal  \
    client_id=srv_6i9OwmrLVqaycbz6iflFmQ  \
    client_secret=F19K9M4xkewuG6DLFIIvYLfPtPcrb1oQ_sogLW0CFKE 

vault kv put appd-partner-demo/tenant_credentials_basic_service_principal  \
    client_id=srv_1n3a5qqVougQbxdvnkQPwj  \
    client_secret=Cp_5zt7VEP4nRdJWiGl8d2pxqY_Mf2weMraDIeTLUB0 

vault kv put appd-partner-demo/aws_credentials_access_key  \
    access_key_id=AKIA4Q3VJVOID2MGTYLU  \
    secret_access_key=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 

vault kv put appd-partner-demo/aws_credentials_role_delegation  \
    account_id=860850072464
```
---
cd into vault-workspace, run `terraform init` followed by `terraform apply` to fetch this credentials from hashicorp vault  

cd into appd-workspace, run `terraform init` and `terraform apply`. this will fetch the credentials from the vault workspace and create the resource  

thogh the credentials will be fetched from vault, those will still be visible into terraform state file.  

to secure this, terraform provides a set of [backends](https://developer.hashicorp.com/terraform/language/settings/backends/configuration) that one can use to remotely store the state file on a variety of locations (cloud, cos, kubernetes, consul etc), on top of which access control can be applied to limit the exposure of these credentials. Most of them also supports encryption and versioning.