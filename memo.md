## AWS CLIからCognito
### サインアップ
```bash
aws cognito-idp sign-up --client-id ${CLIENT_ID} --username ${USER_EMAIL} --password ${PASSWORD} --user-attribute "Name=email,Value=${USER_EMAIL}"
```

### のサインイン
```bash
aws cognito-idp admin-initiate-auth --user-pool-id ${USER_POOL_ID} --client-id ${CLIENT_ID} --auth-flow ADMIN_NO_SRP_AUTH --auth-parameters "USERNAME=${USER_EMAIL},PASSWORD=${PASSWORD}"
```
```bash
aws cognito-idp admin-initiate-auth --user-pool-id ${USER_POOL_ID} --client-id ${CLIENT_ID} --auth-flow ADMIN_USER_PASSWORD_AUTH --auth-parameters USERNAME=${USER_EMAIL},PASSWORD=${PASSWORD} --region ap-northeast-1
```

