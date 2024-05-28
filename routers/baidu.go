package routers

import (
	"net/http"

	_ "github.com/TaceyWong/HotDaily/models"
	"github.com/labstack/echo/v4"
)

// BaiduHot godoc
//
//	@Summary		百度热搜榜接口
//	@Description	get admin info
//	@Tags			热搜榜
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.ListItem
//	@Failure		400	{object}	HTTPError
//	@Failure		401	{object}	HTTPError
//	@Failure		404	{object}	HTTPError
//	@Failure		500	{object}	HTTPError
//	@Security		ApiKeyAuth
//	@Router			/zhihu [get]
func BaiduHot(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"status": "ok"}, " ")
}
