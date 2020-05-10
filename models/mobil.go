package models

import (
	"fmt"
	hpl "golang-web-development/helper"
	"reflect"

	orm "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Mobil berisi struktur tabel tb_mobil
// Mobil berguna untuk menggambarkan struktur tabel tb_mobil
var MobilColumn = []Column{
	{
		Name:          "id_mobil",
		Type:          reflect.Int,
		Fillable:      false,
		IsPk:          true,
		AutoIncrement: true,
	},
	{
		Name:     "id_perusahaan",
		Type:     reflect.Int,
		Fillable: true,
	},
	{
		Name:     "nm_mobil",
		Type:     reflect.String,
		Fillable: true,
	},
	{
		Name:     "jenis_penggerak",
		Type:     reflect.String,
		Fillable: true,
	},
	{
		Name:     "banyak_roda",
		Type:     reflect.Int,
		Fillable: true,
	},
	{
		Name:     "id_jenis",
		Type:     reflect.Int,
		Fillable: true,
	},
	{
		Name:     "harga",
		Type:     reflect.Int,
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
