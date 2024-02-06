# golang-project

## Setup
```
go install github.com/cosmtrek/air@latest // cf. https://github.com/cosmtrek/air?tab=readme-ov-file#via-go-install-recommended
```

## Run

```
docker compose up -d
air
```

## ディレクトリ説明

Serverを書くときは /internal は省略しても全然良い。

```
/cmd                  => 実行エントリーポイント。
/internal/handler/    => ハンドラー置き場。net/httpやAWS Lambdaなどのエントリーポイント。
/internal/model/      => gormのモデルやgormじゃないモデルも。
/internal/pkg/        => 自社サービス関係ない汎用ライブラリ。
/internal/repo/       => クエリ発行場所。gormのラッパー。
/internal/server/     => echoサーバー。
/internal/usecase/    => ユースケース。
/internal/validation/ => validation。
```

### その他
Service (/internal/usecase/taskusecase/service.go)

ここではユースケースから呼び出す最小単位の処理としている。

トランザクションや外部サービスの呼び出しを行う。

必要に応じて作成。
