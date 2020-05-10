package models

import (
	"fmt"
	"reflect"

	sqlQb "github.com/Masterminds/squirrel"
	orm "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
)

var JenisMobilColumn = []Column{
	{
		Name:          "id_jenis",
		Type:          reflect.Int,
		Fillable:      false,
		IsPk:          true,
		AutoIncrement: true,
	},
	{
		Name:     "nm_jenis",
		Type:     reflect.String,
		Fillable: true,
	},
}

// inisaliasi model perusahaan
// nanti variabel ini akan digunakan di controller
type JenisMobil struct {
	*Models
}

func (jm JenisMobil) Cari(keyword string) (result []orm.Params, isError bool) {
	Db := jm.GetDb()

	sqlWhere := make(sqlQb.Eq)
	sqlWhere["nm_jenis"] = keyword

	sql, args, _ := sqlQb.Select("*").From(jm.tableName).Where(sqlWhere).ToSql()
	_, err := Db.Raw(sql, args).Values(&result)

	if err != nil {
		fmt.Println(err.Error())
		isError = true
	}

	return result, isError
}

var ModelJenisMobil JenisMobil = JenisMobil{Models: NewModels("tb_jenis_mobil", JenisMobilColumn)}
