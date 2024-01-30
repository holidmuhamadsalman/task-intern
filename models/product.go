package models

type Product struct {
	Id        int    `gorm:"primaryKey"`
	Nama      string `gorm:"type:varchar(255)" form:"nama"`
	Harga     int    `form:"harga"`
	Size      string `form:"size"`
	Deskripsi string `gorm:"type:text" form:"deskripsi"`
	Photos    string `form:"photo"`
}