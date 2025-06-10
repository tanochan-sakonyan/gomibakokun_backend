<div id="top"></div>

## 使用技術一覧

<!-- シールド一覧 -->
<!-- 該当するプロジェクトの中から任意のものを選ぶ-->
<p style="display: inline">
  <!-- バックエンドのフレームワーク一覧 -->
  <img src="https://img.shields.io/badge/-Echo-black.svg?logo=echo&style=for-the-badge">
  <!-- バックエンドの言語一覧 -->
  <img src="https://img.shields.io/badge/-Go-black.svg?logo=go&style=for-the-badge">
  <!-- インフラ一覧 -->
  <img src="https://img.shields.io/badge/-Firestore-black.svg?&style=for-the-badge">
  <img src="https://img.shields.io/badge/-Heroku-black.svg?logo=heroku&style=for-the-badge">
</p>

## 目次

1. 概要
2. アプリの起動
3. エンドポイント
4. ディレクトリ構成


## ごみばこくんバックエンド

ユーザーが街中のゴミ箱情報を投稿すると、それがアプリ内のマップに反映されるユーザー協力型のゴミ箱マップ作成アプリのバックエンド開発です。

## アプリの起動
### 環境
| 言語・フレームワーク  | バージョン |
| --------------------- | ---------- |
| Go                | 1.24.0   windows/amd64   |
パッケージは[go.mod](go.mod)を確認

### 環境変数
`GOOGLE_CREDENTIALS_JSON_BASE64`：firestore使用時に必要 \
取得方法：[unagami](https://github.com/UnagamiKanta)に連絡

### 起動
```
$ go run main.go
```
BaseURL :[`http://localhost:8080`](http://localhost:8080)

### エンドポイント

|URL|メソッド|内容|リクエストボディ|
|----|----|----|----|
|/trashcan|POST|ゴミ箱情報を保存|`{"latitude":"float", "longitude":float, "image":"json", "trashType":"[]string", "nearetBuildint":"string", "note":"string", "SelectedButton":"string"}`|
|/trashcan?latitude=hoge&longitude=huga | GET| 半径5km内のすべてのゴミ箱情報を取得(現在はデモのため半径無限) | なし|
|/trashcan?id=hoge | DELETE| 指定したゴミ箱情報を削除 |なし|

## ディレクトリ構成

<!-- Treeコマンドを使ってディレクトリ構成を記載 -->
```
rootdir│  .gitignore
│  go.mod
│  go.sum
│  main.go
│  Readme.md
│
├─domain
│  │  error.go
│  │  helper.go
│  │  trashcan.go
│  │  trashcan_config.go
│  │
│  └─repository
│          trashcan_repository.go
│
├─infrastructure
│  │  firestore.go
│  │
│  └─persistence
│          trashcan.go
│
├─interfaces
│  └─handler
│          trashcan.go
│
└─usecase
        trashcan.go
```


