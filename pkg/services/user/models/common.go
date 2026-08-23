package models

type User struct {
	ID        int     `db:"id" json:"id"`
	FirstName string  `db:"first_name" json:"firstName"`
	LastName  string  `db:"last_name" json:"lastName"`
	Email     string  `db:"email" json:"email"`
	Mobile    string  `db:"phone" json:"mobile"`
	Password  string  `db:"password" json:"-"`
	Role      string  `db:"role" json:"role"`
	Language  string  `db:"language" json:"language"`
	SoilType  string  `db:"soil_type" json:"soilType"`
	State     string  `db:"state" json:"state"`
	District  string  `db:"district" json:"district"`
	Address   string  `db:"address" json:"address"`
	ZipCode   string  `db:"zipcode" json:"zipcode"`
	Lat       float64 `db:"lat" json:"lat"`
	Lng       float64 `db:"lng" json:"lng"`
	Avatar    string  `db:"avatar" json:"avatar"`
	Bio       string  `db:"bio" json:"bio"`
	IsActive  bool    `db:"is_active" json:"isActive"`
	CreatedAt string  `db:"created_at" json:"createdAt"`
	UpdatedAt string  `db:"updated_at" json:"updatedAt"`
}
