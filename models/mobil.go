package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type ModelMobil struct {
  ID int `json:"id_mobil",gorm:"column:id_mobil;AUTO_INCREMENT;primary_key"`
  KodeMobil string `json:"kode_mobil",gorm:"column:kode_mobil;type:varchar(10)"`
  Merk string `json:"merk",gorm:"column:merk;type:varchar(50)"`
  Tipe string `json:"tipe",gorm:"column:tipe;type:varchar(100)"`
  Harga int `json:"harga",gorm:"column:harga;type:int(11)`
  Warna string `json:"warna",gorm:"column:warna;type:varchar(20)"`
  Penggerak string `json:"penggerak",gorm:"column:penggerak;type:varchar(4)"`
  BanyakRoda int `json:"banyak_roda",gorm:"column:banyak_roda;type:int(11)"`
  BanyakBangku int `json:"banyak_bangku",gorm:"column:banyak_bangku;type:int(11)"`
  JenisMesin string `json:"jenis_mesin,gorm:"column:jenis_mesin;type:varchar(20)`
  Mesin string `json:"mesin",gorm:"column:mesin;type:varchar(20)"`
}

func (m ModelMobil) TableName() string {
  return "tb_mobil"
}

func (m Mobil) Get() interface{} {
  db := Connect()
  return db.Find(&ModelMobil)
}
func (m Mobil) Find(id int) interface{} {
  db := Connect()
  return db.First(&ModelMobil, id)
}
