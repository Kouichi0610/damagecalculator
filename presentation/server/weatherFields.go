package server

import (
	"damagecalculator/domain/field"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 天候、フィールド一覧の取得
func (s *serverImpl) weatherFieldsList(c *gin.Context) {
	type weatherData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	type fieldData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	type result struct {
		Weathers []*weatherData `json:"weathers"`
		Fields   []*fieldData   `json:"fields"`
	}
	fd := make([]*fieldData, 0)
	wd := make([]*weatherData, 0)

	weathers := field.WeatherList()
	fields := field.FieldList()

	for _, w := range weathers {
		wd = append(wd, &weatherData{Name: w.String(), Description: w.Description()})
	}
	for _, f := range fields {
		fd = append(fd, &fieldData{Name: f.String(), Description: f.Description()})
	}
	res := result{
		Weathers: wd,
		Fields:   fd,
	}

	c.JSON(http.StatusOK, res)
}
