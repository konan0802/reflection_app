# 仕様書
## ◇ Cloud Firestore
### Usersテーブル
| PK          | -      | -     | -      | -    | -               |
| ----------- | ------ | ----- | ------ | ---- | --------------- |
| CognitoName | UserId | Email | Status | Type | TypeDescription |

・Type: やったこと, 分かったこと, 次にやること<br>
・TypeDescription: 

### Reflsテーブル
| PK     | SK   | -     | -    | -    |
| ------ | ---- | ----- | ---- | ---- |
| UserId | Date | Order | Type | Text |

### Transitionsテーブル
| PK     | SK   | -       |
| ------ | ---- | ------- |
| UserId | Date | WordCnt |

## 命名
* Refl：[Type + Text] ※メッセージ種別と内容の固まり
* Thread：[Date + Refls] ※特定日のRefの固まり