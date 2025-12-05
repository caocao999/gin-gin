# GIN　
## 1. go.mod は最初に必ず作る（必須）

プロジェクトディレクトリを作成して、その中で以下を実行：
```
go mod init example.com/ginapi
```
## 2.Gin をインストール（go.mod の後）

```
go get github.com/gin-gonic/gin
```

# go.mod に JWT ライブラリを追加
```
go get github.com/golang-jwt/jwt/v5
```


# DB GORM 設定　の手順

## 1 GORM + MySQL ドライバのインストール
```
go get gorm.io/gorm
go get gorm.io/driver/mysql
```

## 2 MySQL の DSN を決める

基本的に .envで情報は持つ。

```
ユーザー名:パスワード@tcp(ホスト:ポート)/データベース名?charset=utf8mb4&parseTime=True&loc=Local
root:secret@tcp(127.0.0.1:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local
```





## JWT　参考ドキュメント

https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Keyfunc


https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Parse
