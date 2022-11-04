# 仕様書
## ◇ FE
* Flutter
    * iOSアプリ
        * iPhone SE3
        * iPad mini6
    * Webアプリ

## ◇ BE
### API
* **GET reflection/get**
    * **Request**
        * **since** yyyy-mm-dd リクエスト開始期間
        * **until** yyyy-mm-dd リクエスト終了期間
    * **Response**
        * **data** object[]
            * **date** yyyy-mm-dd 日付
            * **threads** object[]
                * **type** text メッセージ種別（E,A,C,N）
                * **text** text メッセージ内容
    * **Header**
        * **Authorization** ID_TOKEN（Cognito ID Token）
* **PUT reflection/put**
    * **Request**
        * **date** yyyy-mm-dd 日付
        * **threads** object[]
            * **type** text メッセージ種別（E,A,C,N）
            * **text** text メッセージ内容
    * **Response**
        * **result** boolean 成否
    * **Header**
        * **Authorization** ID_TOKEN（Cognito ID Token）

## ◇ DynamoDB
### Usersテーブル
| PK     | -           | -     | -      |
| ------ | ----------- | ----- | ------ |
| UserId | CognitoName | Email | Status |
### Reflectionテーブル
| PK     | SK   | -     | -    | -    |
| ------ | ---- | ----- | ---- | ---- |
| UserId | Date | Order | Type | Text |
