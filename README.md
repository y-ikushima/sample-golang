# sample-golang

## golang インストール

色々
brew でもhttps://go.dev/dl/ でも

```
go version
```

sqlc を使っているので version は最低でも 1.21 は必須

> FROM golang:1.23-alpine

とりあえず最新でやる

## プロジェクトとか依存関係とか（構築時のメモ、動かすならすっ飛ばして OK）

### プロジェクト作成

go.mod と go.sum ができる

```
go mod init sample-golang
```

### パッケージに追加

```
go get github.com/gin-gonic/gin
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go get github.com/jackc/pgx/v5/pgxpool
```

### sqlc.yaml に sqlc の設定周りを追加

今回は横着して init.sql を schema 設定している

### コードの生成

```
sqlc generate
```

db/sqlc 配下に sqlc.yaml で読んだファイルでコードが自動生成される  
ここらへんを見て  
https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html

## 起動

`.env.temp` を `.env.local` へリネーム

```
docker-compose build
docker-compose up
docker-compose down
```

init.sql でコンテナ起動時にテーブル定義とデータ投入

psotman あたりで api をコール

```
http://localhost:8080/user?id=1
```

sql は random で同じパラメでも異なる値が帰るようにしている

### DB 最大接続数

環境変数

```
POSTGRES_MAX_CONNECTIONS
```

で設定、CloudSQL とかならインフラ側か

### デプロイ

dockerfile は適当
