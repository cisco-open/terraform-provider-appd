
## AWS Secret Manager Configuration
- sign in to AWS console, search for AWS Secret Manager service.
- select the secret type as `"other type of secret"`
- in `key/value pairs` section, open the `plaintext` tab and paste the given value 
- click next, set the name as `appd/partner-demo`
```
{"client_id":"srv_6i9OwmrLVqaycbz6iflFmQ","client_secret":"7JqM3auWPz_dmv4_hkbe_JyfauGt9A21ZI5HfWpPMks","azure_client_id":"adb90c29-204d-43d9-987c-ab406d9199cc","azure_client_secret":"ZmD8Q~q87S6hbvMVFh5hKS3y8dLA9C1Xfc5jLbjx","azure_subscription_id":"fca41da2-4908-49e2-b0cb-d3d2080fc5be","azure_tenant_id":"f3db65e4-ce5a-4d7e-a140-255bf017d87f","aws_access_key_id":"AKIA4Q3VJVOID2MGTYLU","aws_secret_access_key":"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","aws_account_id":"860850072464"}
```
  
## Vault Configuration
- install and start the vault development server with
```
vault server -dev
```

- follow the instructions. export the `VAULT_ADDR` environment variable, open `http://127.0.0.1:8200` in a browser, use the root token from the output to log in to vault development server.

- create a kv secret engine named `appd-partner-demo` of type `kv` and run the following commands in a separate terminal.


```shell
vault kv put appd-partner-demo/azure_credentials  \
    client_id=adb90c29-204d-43d9-987c-ab406d9199cc  \
    client_secret=ZmD8Q~q87S6hbvMVFh5hKS3y8dLA9C1Xfc5jLbjx  \
    subscription_id=fca41da2-4908-49e2-b0cb-d3d2080fc5be  \
    tenant_id=f3db65e4-ce5a-4d7e-a140-255bf017d87f

vault kv put appd-partner-demo/tenant_credentails_test_service_principal  \
    client_id=srv_6i9OwmrLVqaycbz6iflFmQ  \
    client_secret=F19K9M4xkewuG6DLFIIvYLfPtPcrb1oQ_sogLW0CFKE 

vault kv put appd-partner-demo/tenant_credentials_basic_service_principal  \
    client_id=srv_1n3a5qqVougQbxdvnkQPwj  \
    client_secret=Cp_5zt7VEP4nRdJWiGl8d2pxqY_Mf2weMraDIeTLUB0 

vault kv put appd-partner-demo/aws_credentials \
    access_key_id=AKIA4Q3VJVOID2MGTYLU  \
    secret_access_key=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  \
    account_id=860850072464
```
---

All the demos are designed in such a way to illustrate the use of:
- [x] Credentials passed as hard-coded string
- [x] Credentials passed through terraform variables -- terraform variables can be set and used as environment variables as well (https://developer.hashicorp.com/terraform/language/values/variables)
- [x] credentials fetched using Vault
- [x] credentials fetched using AWS Secret Manager

In addition to this, the Azure demo has backend configured. uncomment it to use the AWS S3 bucket to encrypt and securely store the state on the cloud. A bucket with the name as specified in the backend config must be present on the AWS.

> **NOTE:** To create AWS resources (role and policy for connection type aws, role delegation) and s3 backend, one must have AWS env variables set or authenticated using CLI.

export the following variables using your credentials:
```
export AWS_ACCESS_KEY_ID=AKIA4Q3VJVOID2MGTYLU
export AWS_SECRET_ACCESS_KEY=fELEQROV7T7p4ajIJ5pO3d8xmz849R8nJgxazUdE
export AWS_REGION="us-west-2"
```

export the following environment variables for appdynamics provider configuration
```
export APPDYNAMICS_TENANT_NAME=partner-demo
export APPDYNAMICS_LOGIN_MODE=service_principal
export APPDYNAMICS_SAVE_TOKEN=true

export APPDYNAMICS_CLIENT_ID=srv_6i9OwmrLVqaycbz6iflFmQ
export APPDYNAMICS_CLIENT_SECRET=7JqM3auWPz_dmv4_hkbe_JyfauGt9A21ZI5HfWpPMks
export APPDYNAMICS_USERNAME=aniket.kariya@crestdatasys.com
export APPDYNAMICS_PASSWORD=SuperStrongPassword@123
```

putting these here just in case
```
unset APPDYNAMICS_TENANT_NAME
unset APPDYNAMICS_LOGIN_MODE
unset APPDYNAMICS_SAVE_TOKEN

unset APPDYNAMICS_CLIENT_ID
unset APPDYNAMICS_CLIENT_SECRET
unset APPDYNAMICS_USERNAME
unset APPDYNAMICS_PASSWORD
```