package db

import "gorm.io/gorm"

func ListRecords(db *gorm.DB, model interface{}) ([]interface{}, error) {
	var records []interface{}

	err := db.Find(&records).Error
	if err != nil {
		return nil, err
	}

	return records, nil
}
