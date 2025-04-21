package models


	type BikeModel struct {
	ID        string `gorm:"primaryKey"`
	OwnerName string  `gorm:"not null"`
	Type      string  
	Make      string  
	Serial    string  
	RenteeID  string  `gorm:"foreignKey:RenteeID;constraint:OnDelete:SET NULL"`
}


