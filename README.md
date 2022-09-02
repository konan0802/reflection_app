# Reflection App
チャット形式で日々の振り返りを行うためのiOSアプリ

## DEMO
<img src="ui.png" width="350px">

[Figma](https://www.figma.com/file/Hoqn0h6B3zDAGtcRuOfQGm/YWTReview?node-id=0%3A1)

## Features
* EACNの振り返りフォーマットで日々の振り返り
* iOSアプリのため場所や時間を選ばずに入力が可能
* スプレッドシートに日次でエクスポートすることにより一覧性を担保
* （今後は）対話形式に修正を行い、より”振り返り”の心理的ハードルを下げる

### ◇EACN
* E（経験）
    * 具体的な経験
    * 成功,失敗に関わらず特徴的な経験
* A（抽象化）
    * 具体的な経験を抽象化していく
* C（概念）
    * 「自身の価値観（スタンス）」や「転用可能な法則」を見つける
* N（ネクストアクション）
    * 失敗の経験であれば改善行動
    * 成功の経験であれば転用可能な事柄

※ 開発者が独自に提唱している振り返りフォーマットです

## Architecture
### ◇FE
* Flutter
* Googleスプレッドシート
### ◇BE
* Golang
### ◇Infra
* Lambda
* S3