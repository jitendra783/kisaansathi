package repo

import (
	"sync"
)

// DBStatus holds the current database health status
type DBStatus struct {
	mu        sync.RWMutex
	IsHealthy bool
	ErrorMsg  string
}

var dbStatus = &DBStatus{
	IsHealthy: true,
	ErrorMsg:  "",
}

// SetDBStatus sets the database health status
func SetDBStatus(healthy bool, errMsg string) {
	dbStatus.mu.Lock()
	defer dbStatus.mu.Unlock()
	dbStatus.IsHealthy = healthy
	dbStatus.ErrorMsg = errMsg
}

// GetDBStatus retrieves the current database health status
func GetDBStatus() (bool, string) {
	dbStatus.mu.RLock()
	defer dbStatus.mu.RUnlock()
	return dbStatus.IsHealthy, dbStatus.ErrorMsg
}

// IsDBAvailable checks if database is available
func IsDBAvailable() bool {
	healthy, _ := GetDBStatus()
	return healthy
}
