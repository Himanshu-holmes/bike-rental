package models

type RenteeModel struct {
	ID        string       `gorm:"primaryKey;default:gen_random_uuid()"`
	FirstName string
	LastName  string
	NationalIdNumber string
	Phone     string
	Email     string
	HeldBikes     []BikeModel `gorm:"foreignKey:RenteeID;constraint:OnDelete:SET NULL"`
}
