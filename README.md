# Reflection App（Bloomer）
チャット形式で日々の振り返りを行うためのiOSアプリ

## DEMO
<img src="ui.png" width="350px">

[Figma](https://www.figma.com/file/Hoqn0h6B3zDAGtcRuOfQGm/YWTReview?node-id=0%3A1)

## Features
* 「EASN」という日々の経験から自己を振り返るフォーマットを利用
* iOS(iPhone SE3, iPad mini6)とWebに対応しているため場所や時間を選ばずに入力が可能
* （今後は）対話形式のUIに変更して、より”振り返り”の心理的ハードルを下げる

### ◇EASN
* E（Eperience：経験）
    * 具体的な経験
    * 成功,失敗に関わらず特徴的な経験
* A（Action：行動）
    * 行動レベルの振り返り
    * その経験に際してどのような行動を取っていたか？
    * これまではどのような行動を取っていたか？
* I（Inside：内面）
    * スタンスレベルの振り返り
    * その行動の前提としてある自身のスタンスは何か？
    * 何を感じて、考えて、その行動を取っていることが多いか？
* N（NextAction：ネクストアクション）
    * スタンスを鑑みた上での次回の打開行動
    * スタンスは自身の本質でもあるため蔑ろや否定は行わない
    * 失敗の経験であれば改善行動、成功の経験であれば転用可能な事柄

※ 開発者が独自に提唱している振り返りフォーマットです

## Architecture
### ◇FE
* Flutter
### ◇DB
* Cloud Firestore
### ◇Infra
* Firebase Authentication
* Firebase Hosting

<img src="drawio.png" width="500px">

## Ref
* [Flutterでエラーが発生した時の解決方法](https://qiita.com/717natsuki/items/ddb4adf13aec95e5f2e9)
* [#0 Flutter の設計を決める](https://wasabeef.medium.com/0-flutter-%E3%81%AE%E8%A8%AD%E8%A8%88%E3%82%92%E6%B1%BA%E3%82%81%E3%82%8B-4c6df9a77d67)
* [【Flutter】アプリ全体のアーキテクチャを0から考えて作り直した話](https://zenn.dev/chooyan/articles/eefc76dbd2ba25)
* [バックエンドがFirebaseだけでiOSアプリは作れるのか？](https://qiita.com/jumbOrNot/items/646e0c6b72ab47f452f5)
* [Firebaseでアプリを開発するならClient Side Joinを前提にすること](https://qiita.com/1amageek/items/afc1c0ceb15ffc2372fd)
* [Firebase Realtime DBを実践投入するにあたって考えたこと](https://qiita.com/1amageek/items/64bf85ec2cf1613cf507)
* [Providerで状態管理](https://www.flutter-study.dev/firebase-app/provider)