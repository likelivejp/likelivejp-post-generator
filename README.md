# postgen

likelivejpで空のブログ用mdを生成するスクリプト

## 使い方

第一引数にファイル名を入力します。

```
$ go run main.go foobar
```

すると、以下のようにファイルを生成してくれます。

```
.
├── content
│   └── posts
│       └── 2016
│           └── 01
│               └── 02
│                   └── foobar.md
```

## 開発

```
$ go get github.com/urfave/cli
```

## ビルド

```
$ go build -o postgen main.go
```

ビルドしたものがlikelivejp/likelivejpのbin/に入っています。