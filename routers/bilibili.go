package routers

import (
	"net/http"

	_ "github.com/TaceyWong/HotDaily/models"
	"github.com/labstack/echo/v4"
)

// BilibiliHot godoc
//
//	@Summary		哔哩哔哩排行榜接口
//	@Description	get admin info
//	@Tags			热门榜
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.ListItem
//	@Failure		400	{object}	HTTPError
//	@Failure		401	{object}	HTTPError
//	@Failure		404	{object}	HTTPError
//	@Failure		500	{object}	HTTPError
//	@Security		ApiKeyAuth
//	@Router			/bilibili [get]
func BilibiliHot(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"status": "ok"}, " ")
}
