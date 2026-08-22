package repo

import (
	"errors"

	"gorm.io/gorm"
)

// SafeDBMiddleware wraps DB operations to catch nil DB pointer dereferences
// and return a user-friendly error
func SafeDBAction(db *gorm.DB, action func(*gorm.DB) error) error {
	if db == nil {
		return errors.New("database is currently unavailable - using mock/static data mode")
	}
	return action(db)
}

// SafeDBQuery wraps DB query operations that return data
func SafeDBQuery(db *gorm.DB, query func(*gorm.DB) error) error {
	if db == nil {
		return errors.New("database is currently unavailable")
	}
	return query(db)
}

// GetDBErrorMessage returns an appropriate error message when DB is down
func GetDBErrorMessage() string {
	_, errMsg := GetDBStatus()
	if errMsg != "" {
		return "Database unavailable: " + errMsg
	}
	return "Database is currently unavailable"
}
