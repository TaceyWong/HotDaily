package routers

import (
	"net/http"

	_ "github.com/TaceyWong/HotDaily/models"
	"github.com/labstack/echo/v4"
)

// DouYinHot godoc
//
//	@Summary		抖音热榜接口
//	@Description	实时上升热点
//	@Tags			热榜
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.ListItem
//	@Failure		400	{object}	HTTPError
//	@Failure		401	{object}	HTTPError
//	@Failure		404	{object}	HTTPError
//	@Failure		500	{object}	HTTPError
//	@Security		ApiKeyAuth
//	@Router			/zhihu [get]
func DouYinHot(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"status": "ok"}, " ")
}
