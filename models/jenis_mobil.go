package models

import (
  fmt "fmt"
  orm "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql" // import your required driver
  sqlQb "github.com/Masterminds/squirrel"
)

var JenisMobilColumn = map[string]Column{
  "id_jenis": Column{
    Name: "id_jenis",
    Type: "int",
    Fillable: false,
    IsPk: true,
    AutoIncrement: true,
  },
  "nm_jenis": Column{
    Name: "nm_jenis",
    Type: "varhcar",
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

  sql, args, _ := sqlQb.Select("*").From(jm.tableName).Where(sqlWhere).ToSql();
  _, err := Db.Raw(sql, args).Values(&result);
  
  if err != nil {
    fmt.Println(err.Error());
    isError = true
  }
  
  return result, isError
}

var ModelJenisMobil JenisMobil = JenisMobil{Models: NewModels("tb_jenis_mobil", JenisMobilColumn)}
