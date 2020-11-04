module github.com/han/go-gin-example

go 1.15

require (
	github.com/astaxie/beego v1.12.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/jinzhu/gorm v1.9.16
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace (
	github.com/han/go-gin-example/conf => ./go-gin-example/pkg/conf
	github.com/han/go-gin-example/middleware/jwt => ./go-gin-example/middleware/jwt
	github.com/han/go-gin-example/models => ./go-gin-example/models
	github.com/han/go-gin-example/pkg/e => ./go-gin-example/pkg/e
	github.com/han/go-gin-example/pkg/setting => ./go-gin-example/pkg/setting
	github.com/han/go-gin-example/pkg/util => ./go-gin-example/pkg/util
	github.com/han/go-gin-example/routers => ./go-gin-example/routers
	github.com/han/go-gin-example/routers/api => ./go-gin-example/routers/api
)
