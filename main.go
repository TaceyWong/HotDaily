package hotdaily

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/TaceyWong/HotDaily/docs"
	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
	echo_middle "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//go:generate go run meta/meta.go routers

//	go:generate /Users/tacey/go/bin/swag init

// CustomValidator 自定义请求字段验证器(github.com/go-playground/validator)
type CustomValidator struct {
	validator *validator.Validate
}

// Validate 请求字段验证器接口方法
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// SwaggerRedirect SwaggerRedirect
func SwaggerRedirect(c echo.Context) error {
	return c.Redirect(http.StatusPermanentRedirect, "/api-docs/index.html")
}

// SetupRouter 设置路由
func SetupRouter(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.GET("/api-docs/*", echoSwagger.WrapHandler)
	e.GET("/api-docs", SwaggerRedirect)
	e.GET("/api-docs/", SwaggerRedirect)
	e.GET("", SwaggerRedirect)

}

// @title HotDaily API
// @version 1.0
// @description 聚合新闻热榜数据的API接口.

// @contact.name 问题追踪
// @contact.url https://github.com/TaceyWong/hot-daily/issues
// @contact.email xinyong.wang@qq.com

// @license.name MIT
// @license.url https://www.tldrlegal.com/license/mit-license

// @host 127.0.0.1:8181
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name HOT-DAILY
// @description 用于限定接口访问权限以及用户追踪

// @schemes http https
func Server(host string, port int) {
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.Validator = &CustomValidator{validator: validator.New()}
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Use(echo_middle.Recover())
	e.Use(echo_middle.Logger())
	SetupRouter(e)
	if e.Debug {
		e.Logger.SetLevel(log.DEBUG)
	}
	e.Logger.Info(fmt.Sprintf("🔥 DailyHot API 成功在端口%d运行", port))
	e.Logger.Info(fmt.Sprintf("🔗 Local: 👉 http://%s:%d", host, port))
	e.Logger.Fatal(e.StartServer(server))
}
