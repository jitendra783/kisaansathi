package utils

import (
	"kisaanSathi/pkg/logger"
	"time"
)

func GetDateSomeTimeAgo(initialdate time.Time, years int, months int, days int, hours int) time.Time {
	t1 := time.Now()
	if !initialdate.IsZero() {
		t1 = initialdate
	}
	t1 = t1.AddDate(-years, -months, -days)
	t1 = t1.Add(-time.Hour * time.Duration(hours))
	return t1
}

func GetDaysDiff(targetTime string) (int, error) {
	targetTime = targetTime[:10]
	parsedTime, err := time.Parse("2006-01-02", targetTime)
	if err != nil {
		return 0, err
	}
	t1 := time.Now()

	return int(t1.Sub(parsedTime).Hours() / 24), nil
}

func DateFormat(targetTime string) string {
	parsedTime, err := time.Parse(time.RFC3339, targetTime)

	// Parse the date string
	if err != nil {
		logger.Log().Error(err.Error())
		return ""
	}

	// Format the parsed time into the desired format
	outputDate := parsedTime.Format("02-Jan-2006")
	return outputDate
}
