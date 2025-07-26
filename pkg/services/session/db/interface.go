package db

import (
	"context"

	"gorm.io/gorm"
)

type dbSt struct {
	oracle *gorm.DB
}

type DBLayer interface {
	Session
}
type Session interface {
	ValidateSession(context.Context, string, string) (bool, error)
}

func NewDBObject(oracle *gorm.DB) DBLayer {
	return &dbSt{oracle: oracle}
}
