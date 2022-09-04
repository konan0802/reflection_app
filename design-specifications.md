# 仕様書
## FE
* Flutter
    * 基本的には入力端末としてのiOSアプリ
* スプレッドシート
    * 過去の分を含めて一覧で内容を見れるように

## BE（API）
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
* Golang


## Infra
### A案
* GAS（Google Apps Script）

### B案（A案の速度が出なければ）
* App Engine
    * API用
        * Flutterからのリクエストに対してのレスポンスを返す
    * 集計用
        * 日次でスプレッドシートにデータをエクスポート
* Cloud Tasks
    * スプレッドシートへの更新はキューにより非同期で行う
