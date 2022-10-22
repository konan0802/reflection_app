# 仕様書
## ◇ FE
* Flutter
    * 基本的には入力端末としてのiOSアプリ
* スプレッドシート
    * 過去の分を含めて一覧で内容を見れるように

## ◇ BE
### API
* GET /reflection
    * date yyyy-mm-dd 日付
    * threads object[] 振り返り内容
        * type text タイプ（E,A,C,N）
        * message text 内容
* POST /reflection
    * date yyyy-mm-dd 日付
    * threads object[] 振り返り内容
        * type text タイプ（E,A,C,N）
        * message text 内容

### バッチ（Artisan）
* export-spread-sheet：日次でDBの内容をスプレッドシートにエクスポート

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