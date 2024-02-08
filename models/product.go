package models

import "time"

type Product struct {
	Id        int    `gorm:"primaryKey"`
	Nama      string `gorm:"type:varchar(255)" form:"nama"`
	Harga     int    `form:"harga"`
	Size      string `form:"size"`
	Deskripsi string `gorm:"type:text" form:"deskripsi"`
	Photos    string `form:"photo"`
	CreatedAt time.Time
	UpdatedAt time.Time
}