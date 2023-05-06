package repo_impl

import (
	"gorm.io/gorm"
	"hotel-booking-api/db"
	"hotel-booking-api/model/query"
	"hotel-booking-api/utils"
	"reflect"
)

func GenerateQueryGetData(sql *db.Sql, queryModel query.DataQueryModel, modelStruct any, listIgnoreKey []string) *gorm.DB {
	result := sql.Db
	if queryModel.Search != "" {
		result = result.Where("(fn_convertCoDauToKhongDau(email) LIKE ('%' || fn_convertCoDauToKhongDau(?) || '%')",
			queryModel.Search)
		val := reflect.ValueOf(modelStruct).Elem()
		numberField := val.NumField()
		for i := 0; i < val.NumField(); i++ {
			tempElementJson := utils.GetColumnFieldName(val.Type().Field(i))
			if !utils.Contains(listIgnoreKey, tempElementJson) {
				if i == numberField-1-(val.NumField()-numberField) {
					result = result.Or("fn_convertCoDauToKhongDau("+
						tempElementJson+"::text) LIKE ('%' || fn_convertCoDauToKhongDau(?::text) || '%'))",
						queryModel.Search)
				} else {
					result = result.Or("fn_convertCoDauToKhongDau("+
						tempElementJson+"::text) LIKE ('%' || fn_convertCoDauToKhongDau(?::text) || '%')",
						queryModel.Search)
				}
			} else {
				numberField--
			}
		}

	}
	if queryModel.Filter != nil {
		result = result.Where(queryModel.Filter)
	}
	result = result.Where("is_deleted = ?", queryModel.IsShowDeleted)
	if queryModel.Sort != "" {
		result = result.Order(queryModel.Sort + " " + queryModel.Order)
	} else {
		queryModel.Sort = "id"
		result = result.Order(queryModel.Sort + " " + queryModel.Order)
	}
	if queryModel.Limit > 0 {
		result = result.Limit(queryModel.Limit)
		result = result.Offset(queryModel.Page)
	}
	return result
}
