package controller

import "kisaanSathi/pkg/services/contact/models"

func (c *contactController) ContactUs(req *models.ContactRequest) (*models.ContactResponse, error) {

	if err := c.store.SaveContact(req); err != nil {
		return nil, err
	}

	return &models.ContactResponse{
		Message: "Your message has been submitted successfully.",
	}, nil
}

// GetContactInfo
func (c *contactController) GetContactInfo() (*models.ContactInfoResponse, error) {

	info, err := c.store.GetContactInfo()
	if err != nil {
		return nil, err
	}

	return &models.ContactInfoResponse{
		Data: info,
	}, nil
}
