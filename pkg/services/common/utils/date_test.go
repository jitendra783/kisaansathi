package utils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestDate_Date(t *testing.T) {
	parsedDate, err := Date("31/08/2003").Date()
	assert.NoError(t, err)
	assert.Equal(t, parsedDate.Day(), 31)
	assert.Equal(t, parsedDate.Month(), time.August)
	assert.Equal(t, parsedDate.Year(), 2003)

	parsedDate, err = Date("2043-08-18T00:00:00Z").Date()
	assert.NoError(t, err)
	assert.Equal(t, parsedDate.Day(), 18)
	assert.Equal(t, parsedDate.Month(), time.August)
	assert.Equal(t, parsedDate.Year(), 2043)
}

func TestDate_String(t *testing.T) {
	dateString := "31/08/2003"
	dateValue := Date(dateString)
	assert.True(t, dateValue.String() == dateString)
	assert.True(t, reflect.TypeOf(dateValue).String() == "utils.Date")
	assert.True(t, reflect.TypeOf(dateValue.String()).String() == "string")
}

func TestDate_IsBefore(t *testing.T) {
	firstDate := Date("30/08/2003")
	secondDate, err := time.Parse("02/01/2006 15:04:05", "31/08/2003 11:21:29.000")
	assert.NoError(t, err)

	ok, err := firstDate.IsBefore(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, true)

	firstDate = "31/08/2003"
	ok, err = firstDate.IsBefore(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, false)

	firstDate = "01/09/2003"
	ok, err = firstDate.IsBefore(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, false)
}

func TestDate_IsBeforeOrSame(t *testing.T) {
	firstDate := Date("30/08/2003")
	secondDate, err := time.Parse("02/01/2006 15:04:05", "31/08/2003 11:21:29.000")
	assert.NoError(t, err)

	ok, err := firstDate.IsBeforeOrSame(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, true)

	firstDate = "31/08/2003"
	ok, err = firstDate.IsBeforeOrSame(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, true)

	firstDate = "01/09/2003"
	ok, err = firstDate.IsBeforeOrSame(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, false)
}

func TestDate_IsAfter(t *testing.T) {
	firstDate := Date("30/08/2003")
	secondDate, err := time.Parse("02/01/2006 15:04:05", "31/08/2003 11:21:29.000")
	assert.NoError(t, err)

	ok, err := firstDate.IsAfter(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, false)

	firstDate = "31/08/2003"
	ok, err = firstDate.IsAfter(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, false)

	firstDate = "01/09/2003"
	ok, err = firstDate.IsAfter(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, true)
}

func TestDate_IsAfterOrSame(t *testing.T) {
	firstDate := Date("30/08/2003")
	secondDate, err := time.Parse("02/01/2006 15:04:05", "31/08/2003 11:21:29.000")
	assert.NoError(t, err)

	ok, err := firstDate.IsAfterOrSame(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, false)

	firstDate = "31/08/2003"
	ok, err = firstDate.IsAfterOrSame(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, true)

	firstDate = "01/09/2003"
	ok, err = firstDate.IsAfterOrSame(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, true)
}

func TestDate_IsEqual(t *testing.T) {
	firstDate := Date("30/08/2003")
	secondDate, err := time.Parse("02/01/2006 15:04:05", "31/08/2003 11:21:29.000")
	assert.NoError(t, err)

	ok, err := firstDate.IsEqual(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, false)

	firstDate = "31/08/2003"
	ok, err = firstDate.IsEqual(secondDate)
	assert.NoError(t, err)
	assert.Equal(t, ok, true)
}
