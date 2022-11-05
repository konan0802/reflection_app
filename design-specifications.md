# 仕様書
## ◇ Cloud Firestore
### Usersテーブル
| PK          | -      | -     | -      | -     |
| ----------- | ------ | ----- | ------ | ----- |
| CognitoName | UserId | Email | Status | Types |
### Reflsテーブル
| PK     | SK   | -     | -    | -    |
| ------ | ---- | ----- | ---- | ---- |
| UserId | Date | Order | Type | Text |

## 命名
* Refl：[Type + Text] ※メッセージ種別と内容の固まり
* Thread：[Date + Refls] ※特定日のRefの固まり