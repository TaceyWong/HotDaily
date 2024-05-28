package routers

import (
	"net/http"

	_ "github.com/TaceyWong/HotDaily/models"
	"github.com/labstack/echo/v4"
)

// TieBaHot godoc
//
//	@Summary		百度贴吧热议榜接口
//	@Description
//	@Tags			热议榜
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.ListItem
//	@Failure		400	{object}	HTTPError
//	@Failure		401	{object}	HTTPError
//	@Failure		404	{object}	HTTPError
//	@Failure		500	{object}	HTTPError
//	@Security		ApiKeyAuth
//	@Router			/zhihu [get]
func TieBaHot(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"status": "ok"}, " ")
}
