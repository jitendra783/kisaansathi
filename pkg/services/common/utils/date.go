package utils

import (
	"kisaanSathi/pkg/logger"
	"time"
)

var TimeNow = time.Now

type Date string

const DateLayout = "02/01/2006"

func (d Date) Date() (time.Time, error) {
	layout := DateLayout
	if len(d) > 10 {
		layout = time.RFC3339
	}
	parsedTime, err := time.Parse(layout, string(d))
	if err != nil {
		logger.Log().Error("Date Parsing Error : " + err.Error())
		return parsedTime, err
	}
	return parsedTime.Local(), err
}

func (d Date) String() string {
	return string(d)
}

func (d Date) IsBefore(anotherTime time.Time) (bool, error) {
	parsedDate, err := d.Date()
	if err != nil {
		return false, err
	}
	return parsedDate.Before(anotherTime.Truncate(24 * time.Hour)), nil
}

func (d Date) IsBeforeOrSame(anotherTime time.Time) (bool, error) {
	firstDate, err := d.Date()
	if err != nil {
		return false, err
	}
	secondDate := anotherTime.Truncate(24 * time.Hour)
	return firstDate.Before(secondDate) || firstDate.Equal(secondDate), nil
}

func (d Date) IsAfter(anotherTime time.Time) (bool, error) {
	parsedDate, err := d.Date()
	if err != nil {
		return false, err
	}
	return parsedDate.After(anotherTime.Truncate(24 * time.Hour)), nil
}

func (d Date) IsAfterOrSame(anotherTime time.Time) (bool, error) {
	firstDate, err := d.Date()
	if err != nil {
		return false, err
	}
	secondDate := anotherTime.Truncate(24 * time.Hour)
	return firstDate.After(secondDate) || firstDate.Equal(secondDate), nil
}

func (d Date) IsEqual(anotherTime time.Time) (bool, error) {
	firstDate, err := d.Date()
	if err != nil {
		return false, err
	}
	secondDate := anotherTime.Truncate(24 * time.Hour)
	return firstDate.Equal(secondDate), nil
}
