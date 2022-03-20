package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Car struct {
	ID    uint   `json:"ID" gorm:"primary_key;auto_increment"`
	Model string `json:"model" gorm:"size:255;not null"`
	Color string `json:"color" gorm:"size:255;not null"`
	Brand string `json:"brand" gorm:"size:255;not null"`
}

func (c *Car) CreateCar(db *gorm.DB) (*Car, error) {
	err := db.Debug().Model(&Car{}).Create(&c).Error
	if err != nil {
		return &Car{}, err
	}
	return c, nil
}

func (c *Car) GetAllCars(db *gorm.DB) (*[]Car, error) {
	carModels := []Car{}
	err := db.Debug().Model(&Car{}).Limit(100).Find(&carModels).Error
	if err != nil {
		return nil, err
	}
	return &carModels, nil
}

func (c *Car) GetCarByID(db *gorm.DB, uid uint64) (*Car, error) {
	err := db.Debug().Model(Car{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("car not found")
	}
	return c, err
}

func (c *Car) UpdateCar(db *gorm.DB, uid uint64) (*Car, error) {
	db = db.Debug().Model(&Car{}).Where("id = ?", uid).Take(&Car{}).UpdateColumns(
		map[string]interface{}{
			"model": c.Model,
			"color": c.Color,
			"brand": c.Brand,
		},
	)
	if db.Error != nil {
		return &Car{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&Car{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return &Car{}, err
	}
	return c, nil
}
