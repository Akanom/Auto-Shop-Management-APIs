package models

import (
	"github.com/Akanom/Golang-Auto-Shop-Management-APIs/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Auto struct {
	gorm.Model
	Name            string `json:"name"`
	AutoModel       string `json:"automodel"`
	ManufactureDate string `json:"manufacturedate"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Auto{})
}

func (b *Auto) CreateNewAuto() *Auto {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAuto() []Auto {
	var Autos []Auto
	db.Find(&Autos)
	return Autos
}

func GetAutoById(Id int64) (*Auto, *gorm.DB) {
	var GetAuto Auto
	db := db.Where("ID=?", Id).Find(&GetAuto)
	return &GetAuto, db
}

func DeleteAuto(ID int64) Auto {
	var auto Auto
	db.Where("ID=?", ID).Delete(auto)
	return auto
}
