# Sokudoku

Spritsやspeedreadのように使える、オープンソースのターミナルアプリケーションです。

このコマンドは、日本語のテキストを読み込んで単語ごとに素早く表示します。

## Install

```
$ export CGO_LDFLAGS="-L/path/to/lib -lmecab -lstdc++"
$ export CGO_CFLAGS="-I/path/to/include"
$ go get github.com/shogo82148/go-mecab
$ go get github.com/umanoda/sokudoku
```

## Basic Example

```
cat cat.txt | sokudoku -w 250
```

デフォルトでは1単語を250ミリ秒ごとに切り替えて表示します。なれるまではこの速度を使うとよいでしょう。
500ミリ秒ごとの表示ならば、かなり簡単に読み取ることが出来ます。

## Controls

コマンドの実行中に以下の操作が行えます

* `[` - 表示速度を10%遅くします
* `]` - 表示速度を10%速くします
* `space` - 停止します (そして、最新の２行を表示します)


## Inspire

[pasky / speedread](https://github.com/pasky/speedread)
