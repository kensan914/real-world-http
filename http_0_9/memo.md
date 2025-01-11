# HTTP/0.9 概要
- 1990年にドキュメント化、1991年にリリースされたHTTPのプロトタイプ
- HTTP/1.0 がリリースされ話題にされるようになってから、HTTP/0.9 と呼ばれるようになった

## HTTP/0.9 でできないこと
- 0.9で実現されることは、ボディの受信とパスだけ
- HTML 以外のフォーマットを指定できない
- 作成、更新、削除ができない
- ステータスがなかった


# 実装した GO サーバが HTTP/0.9 をサポートしていないので、擬似的に1.0を使用 (P.6)
```shell
curl --http1.0 http://localhost:18888/greeting
```

# クエリパラメータの指定 (P.7)
- --get は -G でも可
- --getがある時、--data-urlencode はクエリパラメータとして動作
```shell
curl --http1.0 --get --data-urlencode "search word" http://localhost:18888/greeting
```






