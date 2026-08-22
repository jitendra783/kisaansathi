package controller

import (
	"kisaanSathi/pkg/services/metadata/models"
)

func (c *controller) GetMetadata() (*models.MetadataResponse, error) {

	records, err := c.store.GetMetadata()
	if err != nil {
		return nil, err
	}

	response := &models.MetadataResponse{}

	for _, item := range records {

		switch item.Key {

		case "farmers":
			response.Stats.Farmers = item.Value

		case "experts":
			response.Stats.Experts = item.Value

		case "districts":
			response.Stats.Districts = item.Value

		case "support":
			response.Stats.Support = item.Value
		}
	}

	return response, nil
}
