module github.com/han/go-gin-example

go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/robfig/cron v1.2.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.9
	github.com/ugorji/go v1.1.13 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/sys v0.0.0-20201101102859-da207088b7d1 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
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
