package models

import (
	"fmt"
	hpl "golang-web-development/helper"

	orm "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// MobilColumn berisi struktur tabel tb_mobil
// MobilColumn berguna untuk menggambarkan struktur tabel tb_mobil
var MobilColumn = map[string]Column{
	"id_mobil": Column{
		Name:          "id_mobil",
		Type:          "int",
		Fillable:      false,
		IsPk:          true,
		AutoIncrement: true,
	},
	"id_perusahaan": Column{
		Name:     "id_perusahaan",
		Type:     "int",
		Fillable: true,
	},
	"nm_mobil": Column{
		Name:     "nm_mobil",
		Type:     "varchar",
		Fillable: true,
	},
	"jenis_penggerak": Column{
		Name:     "jenis_penggerak",
		Type:     "varchar",
		Fillable: true,
	},
	"banyak_roda": Column{
		Name:     "banyak_roda",
		Type:     "int",
		Fillable: true,
	},
	"id_jenis": Column{
		Name:     "id_jenis",
		Type:     "int",
		Fillable: true,
	},
	"harga": Column{
		Name:     "harga",
		Type:     "int",
		Fillable: true,
	},
}

// Mobil struct turunan dari Models
// Struct ini berguna untuk manajemen model Mobil
type Mobil struct {
	*Models
}

// GetMobil parameternya adalah map[string]interface{}
// Method ini berguna untuk mengambil data mobil
// Dari tabel mobil dengan parameternya adalah kondisi WHERE
// Dengan format map[string]interface{}
func (m *Mobil) GetMobil(condition map[string]interface{}) (result []orm.Params, isError bool) {
	Db := m.GetDb()
	sqlColumn := []string{
		"tb_mobil.*",
		"tb_jenis_mobil.nm_jenis",
		"tb_perusahaan.nm_perusahaan",
		"tb_perusahaan.alamat",
	}
	defaultCondition := map[string]interface{}{
		"[><] tb_jenis_mobil": "tb_mobil.id_jenis = tb_jenis_mobil.id_jenis",
		"[><] tb_perusahaan":  "tb_mobil.id_perusahaan = tb_perusahaan.id_perusahaan",
	}

	if condition != nil {
		for keys, val := range condition {
			defaultCondition[keys] = val
		}
	}

	sql, args, _ := hpl.SelectSql(m.GetTableName(), sqlColumn, defaultCondition)
	_, err := Db.Raw(sql, args).Values(&result)

	if err != nil {
		fmt.Println(err.Error())
		isError = true
	}

	return result, isError
}

var ModelMobil Mobil = Mobil{NewModels("tb_mobil", MobilColumn)}
