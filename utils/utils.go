package utils

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hotel-booking-api/db"
	"hotel-booking-api/logger"
	"hotel-booking-api/model/query"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func GetQueryDataModel(c echo.Context, listIgnoreColumns []string, modelStruct any) query.DataQueryModel {
	var model query.DataQueryModel
	modelFilter := GetFilterQueryDataModel(c, modelStruct, listIgnoreColumns)
	//limit item can get
	tempValueLimit := c.QueryParam("offset")
	limit, err := strconv.ParseInt(tempValueLimit, 10, 32)
	if err != nil {
		model.Limit = math.MaxInt32
	}
	//page index
	tempValuePage := c.QueryParam("page")
	page, err := strconv.ParseInt(tempValuePage, 10, 32)

	if page > 0 {
		model.PageViewIndex = int(page)
		page = limit*page - limit
	}
	if err != nil {
		page = 0
	}
	model.Limit = int(limit)
	model.Filter = modelFilter
	//logger.Error(string(modelFilter))
	model.Page = int(page)
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
	model.ListIgnoreColumns = listIgnoreColumns
	return model
}

func GetFilterQueryDataModel(c echo.Context, modelStruct any, listIgnoreKey []string) map[string]interface{} {
	modelFilter := map[string]interface{}{}
	val := reflect.ValueOf(modelStruct).Elem()
	for i := 0; i < val.NumField(); i++ {
		tempElementJson := GetColumnFieldName(val.Type().Field(i))
		if !Contains(listIgnoreKey, tempElementJson) && c.QueryParam(tempElementJson) != "" {
			modelFilter[tempElementJson] = c.QueryParam(tempElementJson)
		}
	}
	return modelFilter
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
	logger.Info(value)
	err := json.Unmarshal([]byte(value), &listStringDecode)
	if err != nil {
		logger.Error("Error decode json array data", zap.Error(err))
		return []string{}
	}
	return listStringDecode
}

func GetColumnFieldName(f reflect.StructField) (name string) {
	tag := f.Tag.Get("json")
	if tag == "" {
		return f.Name
	}
	if tag == "-" {
		return ""
	}
	if i := strings.Index(tag, ","); i != -1 {
		if i == 0 {
			return f.Name
		} else {
			return tag[:i]
		}
	}
	return tag
}

func Contains(array []string, valueCheck string) bool {
	for _, valueIndex := range array {
		if valueIndex == valueCheck {
			return true
		}
	}
	return false
}

func ConvertStructToMap(modelStruct any) map[string]interface{} {
	modelStructJson, _ := json.Marshal(modelStruct)
	var mapFromJson map[string]interface{}
	_ = json.Unmarshal(modelStructJson, &mapFromJson)
	delete(mapFromJson, "hotel_type")
	delete(mapFromJson, "hotel_facility")
	delete(mapFromJson, "hotel")
	delete(mapFromJson, "room_type_facility")
	delete(mapFromJson, "room_type_views")
	delete(mapFromJson, "room_nights")
	delete(mapFromJson, "rate_plans")
	delete(mapFromJson, "kitchen_tool")
	delete(mapFromJson, "-")
	return mapFromJson
}

func GreaterThanEqual(value float32, column string, db *db.Sql) *gorm.DB {
	return db.Db.Where(column+" >= ?", value)
}

func GreaterThan(value float32, column string, db *db.Sql) *gorm.DB {
	return db.Db.Where(column+" > ?", value)
}

func LessThanEqual(value float32, column string, db *db.Sql) *gorm.DB {
	return db.Db.Where(column+" <= ?", value)
}

func LessThan(value float32, column string, db *db.Sql) *gorm.DB {
	return db.Db.Where(column+" < ?", value)
}
