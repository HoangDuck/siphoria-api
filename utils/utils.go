package utils

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"hotel-booking-api/model/query"
	"math"
	"strconv"
	"strings"
)

func GetQueryDataModel(c echo.Context) query.DataQueryModel {
	var model query.DataQueryModel
	//limit item can get
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		model.Limit = math.MinInt64
	}
	model.Limit = limit
	//page index
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		model.Page = 0
	}
	model.Page = page
	//search parameters
	search := c.QueryParam("search")
	model.Search = search
	//sort parameters
	sort := c.QueryParam("sort")
	model.Sort = sort
	//order by parameters
	order := c.QueryParam("order")
	if order == "" {
		model.Order = "DESC"
	} else {
		model.Order = order
	}
	//slice parameters
	start := c.QueryParam("start")
	end := c.QueryParam("end")
	model.Start = start
	model.End = end
	return model
}

func GetFilterQueryDataModel(c echo.Context, queryDataModel query.DataQueryModel) query.DataQueryModel {

	return queryDataModel
}

func GetNewId() (string, error) {
	//Generate UUID
	newId, err := uuid.NewUUID()
	if err != nil {
		logger.Error("Error uuid data", zap.Error(err))
		return "", err
	}
	return newId.String(), nil
}

func DecodeJSONArray(value string) []string {
	var listStringDecode []string
	logger.Info(value)
	value = strings.Replace(value, "\"", "`", -1)
	value = strings.Replace(value, "'", "\"", -1)
	err := json.Unmarshal([]byte(value), &listStringDecode)
	if err != nil {
		logger.Error("Error decode json array data", zap.Error(err))
		return []string{}
	}
	return listStringDecode
}
