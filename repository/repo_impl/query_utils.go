package repo_impl

import (
	"gorm.io/gorm"
	"hotel-booking-api/db"
	"hotel-booking-api/model/query"
	"hotel-booking-api/utils"
	"reflect"
)

func GenerateQueryGetData(sql *db.Sql, queryModel *query.DataQueryModel, modelStruct any, listIgnoreKey []string) *gorm.DB {
	result := sql.Db
	//query row by conditions
	if queryModel.Search != "" {
		result = result.Where("(unaccent(id) ILIKE CONCAT('%', unaccent(?), '%') ",
			queryModel.Search)
		val := reflect.ValueOf(modelStruct).Elem()
		numberRemainFieldCheck := val.NumField()
		numberIgnoreField := 0
		for i := 0; i < val.NumField(); i++ {
			tempElementJson := utils.GetColumnFieldName(val.Type().Field(i))
			if utils.Contains(listIgnoreKey, tempElementJson) {
				numberIgnoreField++
			}
		}
		for i := 0; i < val.NumField(); i++ {
			tempElementJson := utils.GetColumnFieldName(val.Type().Field(i))
			numberRemainFieldCheck--
			if !utils.Contains(listIgnoreKey, tempElementJson) {
				if numberRemainFieldCheck == numberIgnoreField {
					result = result.Or("unaccent("+
						tempElementJson+"::text) ILIKE CONCAT('%', unaccent(?), '%'))",
						queryModel.Search)
				} else {
					result = result.Or("unaccent("+
						tempElementJson+"::text) ILIKE CONCAT('%', unaccent(?), '%')",
						queryModel.Search)
				}
			} else {
				numberIgnoreField--
			}
		}

	}

	if queryModel.Filter != nil {
		result = result.Where(queryModel.Filter)
	}
	result = result.Where("is_deleted = ?", queryModel.IsShowDeleted)

	//statistic total row, total pages
	var countTotalRows int64
	result.Model(modelStruct).Count(&countTotalRows)
	queryModel.TotalRows = int(countTotalRows)
	countTotalPages := 0
	if queryModel.Limit > 0 && queryModel.TotalRows%queryModel.Limit > 0 {
		countTotalPages = 1
	}
	if queryModel.Limit > 0 {
		countTotalPages += queryModel.TotalRows / queryModel.Limit
	} else {
		countTotalPages = 1
	}
	queryModel.TotalPages = countTotalPages
	//order by
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
