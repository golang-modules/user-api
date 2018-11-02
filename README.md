# API Layer for UserModule

## Depandency Packages
```go

go get github.com/go-sql-driver/mysql
go get golang.org/x/crypto/bcrypt
go get github.com/golang-modules/users

go get github.com/gin-gonic/gin
go get github.com/gin-contrib/sessions
go get github.com/gin-contrib/sessions/memstore
go get github.com/gin-contrib/cors

```

## Environmental variables

```bash

export MYSQL_DB_URL="mysqluser:mysqlpwd@tcp(127.0.0.1:3306)/main"

```

## Running 
go run server.go
