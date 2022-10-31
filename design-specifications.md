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
* **PUT reflection/put**
    * **Request**
        * **date** yyyy-mm-dd 日付
        * **threads** object[]
            * **type** text メッセージ種別（E,A,C,N）
            * **text** text メッセージ内容
    * **Response**
        * **result** boolean 成否

## ◇ DynamoDB
| パーティションキー | ソートキー |    属性1      |    属性2     |
| --------------- | -------- | ------------ | ------------ |
| 対象日           | 順序      | メッセージ種別 | メッセージ内容 |

※ SQLの場合は
```sql
CREATE TABLE `content` (
  `id` bigint unsigned NOT NULL COMMENT 'ID',
  `type` smallint unsigned NOT NULL DEFAULT 0 COMMENT 'メッセージ種別',
  `text` text COMMENT 'メッセージ内容',
  `sort` smallint unsigned NOT NULL DEFAULT 0 COMMENT '1日ごとのメッセージの順番',
  `date` date NOT NULL DEFAULT '2006-01-01' COMMENT '対象日',
  PRIMARY KEY (`id`, `date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```