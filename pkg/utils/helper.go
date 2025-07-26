package utils

import (
	"kisaanSathi/pkg/logger"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetFloat64FromInterface(v interface{}) (f float64) {
	switch t := v.(type) {
	case float64:
		if v != nil {
			f = v.(float64)
		}
	default:
		logger.Log().Info("This type is not  handled", zap.Any("type", t))
	}

	return
}

func GetSliceFromStringBySeparator(str string, sep string) []string {
	var (
		data []string
	)
	for _, v := range strings.Split(str, sep) {
		v = strings.TrimSpace(v)
		if v != "" {
			data = append(data, v)
		}
	}
	return data
}

func IsOpen(start_date, end_date string) bool {
	start_date_diff_current, err := GetDaysDiff(start_date)
	if err != nil {
		logger.Log().Error("error while getting difference between todays and start date", zap.Error(err))
	}
	end_date_diff_current, err := GetDaysDiff(end_date)
	if err != nil {
		logger.Log().Error("error while getting difference between todays and close date", zap.Error(err))
	}
	if end_date_diff_current <= 0 && start_date_diff_current >= 0 {
		return true
	}
	return false
}

func PrepareNameString(names string, delimeter string) string {
	splitNames := strings.Split(names, delimeter)
	noOfNames := len(splitNames)
	if noOfNames == 1 {
		return names
	}
	return strings.Join(splitNames[:noOfNames-1], ",") + " and" + splitNames[noOfNames-1]

}

/*
Paginate is a helper function for implementing pagination in all fundlist apis
*/
func Paginate(c *gin.Context, page string, limit string) func(db *gorm.DB) *gorm.DB {
	logger.Log(c).Debug("START")
	defer logger.Log(c).Debug("END")
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(limit)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
