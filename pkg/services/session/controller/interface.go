package controller

import (
	"context"
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/session/db"
)

type repoObject struct {
	db db.DBLayer
	// cache cache.CacheLayer
}

type RepoLayer interface {
	Validate(context.Context, string, string) (bool, error)
}

func NewRepoLayerObject(repo repo.DataObject) RepoLayer {
	return &repoObject{
		db: db.NewDBObject(repo.Databases.PgDB),
	}
}
