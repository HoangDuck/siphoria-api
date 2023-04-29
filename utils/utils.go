package utils

import (
	"github.com/labstack/echo/v4"
	"hotel-booking-api/model/query"
	"math"
	"strconv"
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
