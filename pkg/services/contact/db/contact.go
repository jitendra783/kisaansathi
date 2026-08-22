package db

import "kisaanSathi/pkg/services/contact/models"

func (s *contactStore) SaveContact(req *models.ContactRequest) error {

	query := `
		INSERT INTO contact_messages
			(name, email, phone, subject, message)
		VALUES
			(:name, :email, :phone, :subject, :message)
	`

	_, err := s.db.NamedExec(query, req)
	if err != nil {
		return err
	}

	return nil
}

// GetContactInfo returns contact information
func (s *contactStore) GetContactInfo() (*models.ContactInfo, error) {

	var info models.ContactInfo

	query := `
		SELECT
			email,
			phone,
			address,
			facebook,
			instagram,
			twitter,
			linkedin
		FROM metadata_contact
		LIMIT 1
	`

	if err := s.db.Get(&info, query); err != nil {
		return nil, err
	}

	return &info, nil
}
