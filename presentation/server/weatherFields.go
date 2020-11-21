package server

import (
	"damagecalculator/domain/field"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 天候、フィールド一覧の取得
func (s *serverImpl) weatherFieldsList(c *gin.Context) {
	type result struct {
		Weathers []string `json:"weathers"`
		Fields   []string `json:"fields"`
	}

	res := result{
		Weathers: field.WeatherNames(),
		Fields:   field.FieldNames(),
	}

	c.JSON(http.StatusOK, res)
}
