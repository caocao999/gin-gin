# 1. go.mod は最初に必ず作る（必須）

プロジェクトディレクトリを作成して、その中で以下を実行：
```
go mod init example.com/ginapi
```
# 2.Gin をインストール（go.mod の後）

```
go get github.com/gin-gonic/gin
```

# 3 go.mod に JWT ライブラリを追加
```
go get github.com/golang-jwt/jwt/v5
```





## JWT　参考ドキュメント

https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Keyfunc


https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Parse
