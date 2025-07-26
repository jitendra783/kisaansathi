package utils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestStrings_IsBackOfficeUser(t *testing.T) {
	userId := "system"
	isBackOfficeUser := Strings(userId).IsBackOfficeUser()
	assert.True(t, isBackOfficeUser)

	userId = "testuser"
	isBackOfficeUser = Strings(userId).IsBackOfficeUser()
	assert.False(t, isBackOfficeUser)
}

func TestStrings_IsAgent(t *testing.T) {
	userId := "Y"
	isAgent := Strings(userId).IsAgent()
	assert.True(t, isAgent)

	userId = "N"
	isAgent = Strings(userId).IsAgent()
	assert.False(t, isAgent)
}

func TestStrings_IsBusinessPartner(t *testing.T) {
	userId := "#test"
	isBusinessPartner := Strings(userId).IsBusinessPartner()
	assert.True(t, isBusinessPartner)

	userId = "test"
	isBusinessPartner = Strings(userId).IsBusinessPartner()
	assert.False(t, isBusinessPartner)
}

func TestStrings_IsDummyFolio(t *testing.T) {
	folio := "_D"
	isDummyFolio := Strings(folio).IsDummyFolio()
	assert.True(t, isDummyFolio)

	folio = "123123"
	isDummyFolio = Strings(folio).IsDummyFolio()
	assert.False(t, isDummyFolio)
}

func TestStrings_IsYes(t *testing.T) {
	flag := "Y"
	isYes := Strings(flag).IsYes()
	assert.True(t, isYes)

	flag = "N"
	isYes = Strings(flag).IsYes()
	assert.False(t, isYes)
}

func TestStrings_IsNo(t *testing.T) {
	flag := "N"
	isNo := Strings(flag).IsNo()
	assert.True(t, isNo)

	flag = "Y"
	isNo = Strings(flag).IsNo()
	assert.False(t, isNo)
}

func TestStrings_IsDemat(t *testing.T) {
	flag := "D"
	isDemat := Strings(flag).IsDemat()
	assert.True(t, isDemat)

	flag = "N"
	isDemat = Strings(flag).IsDemat()
	assert.False(t, isDemat)
}

func TestStrings_String(t *testing.T) {
	flag := "D"
	stringValue := Strings(flag)
	assert.True(t, stringValue.String() == flag)
	assert.True(t, reflect.TypeOf(stringValue).String() == "utils.Strings")
	assert.True(t, reflect.TypeOf(stringValue.String()).String() == "string")
}
