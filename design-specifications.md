# 仕様書
## ◇ FE
* Flutter
    * 基本的には入力端末としてのiOSアプリ

## ◇ BE
### API
* **GET /reflection**
    * **Request**
        * **since** yyyy-mm-dd リクエスト開始期間
        * **until** yyyy-mm-dd リクエスト終了期間
    * **Response**
        * **data** object[]
            * **date** yyyy-mm-dd 日付
            * **threads** object[]
                * **type** text メッセージ種別（E,A,C,N）
                * **message** text メッセージ内容
* **POST /reflection**
    * **Request**
        * **date** yyyy-mm-dd 日付
        * **threads** object[]
            * **type** text メッセージ種別（E,A,C,N）
            * **message** text メッセージ内容
    * **Response**
        * **result** boolean 成否

## ◇ DB
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