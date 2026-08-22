package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/metadata/controller"
	"kisaanSathi/pkg/services/metadata/db"

	"github.com/gin-gonic/gin"
)

type medatahandler struct {	
	controller controller.MetadataController
}

type MetadataHandler interface {
	GetMetadata(*gin.Context)
}
func NewMetadataHandler(controller controller.MetadataController) MetadataHandler{
	return &medatahandler{
		controller: controller,
	}
}
func NewMetaDataController(repo repo.DataObject) controller.MetadataController	{
	store := db.NewMetadataStore(repo.Databases.PgDB)
	return controller.NewMetadataController(store)
}

