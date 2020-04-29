package models

import (
  hpl "belajar-ech0-framework/helper"
  "fmt"
  sqlQb "github.com/Masterminds/squirrel"
  orm "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql" // import your required driver
)

type IModels interface {
  GetTableName() string
  GetPkColumn() Column
  GetColumnList() map[string]Column
  GetColumnListAsSql() []string
  Get() ([]orm.Params, bool)
  Find(string) ([]orm.Params, bool)
  Insert(data map[string]string) (string, bool)
  Update(string, map[string]string) bool
  Delete(string) bool
  Count() (int, bool)
}

type Column struct {
  Name          string
  Type          string
  Fillable      bool
  IsPk          bool
  AutoIncrement bool
}

type Models struct {
  tableName  string
  pkColumn   Column
  columnList map[string]Column
}

func (m *Models) GetDb() orm.Ormer {
  return orm.NewOrm()
}

func NewModels(tableName string, tableStruct map[string]Column) *Models {
  var banyakPk int
  for _, value := range tableStruct {
    if value.IsPk == true {
      banyakPk += 1
    }
  }
  if banyakPk != 1 {
    panic("Primary Key berlebih atau belum diset!")
  }

  return &Models{
    tableName:  tableName,
    columnList: tableStruct,
  }
}

func (m *Models) GetTableName() string {
  return m.tableName
}

func (m *Models) GetPkColumn() (pkColumn Column) {
  for index, value := range m.columnList {
    if value.IsPk == true {
      pkColumn = m.columnList[index]
    }
  }
  return pkColumn
}

func (m *Models) GetColumnList() map[string]Column {
  return m.columnList
}

func (m *Models) GetColumnListAsSql() (column []string) {
  for _, value := range m.columnList {
    column = append(column, "`"+m.tableName+"`"+"."+"`"+value.Name+"`")
  }
  return column
}

func (m *Models) Get() (result []orm.Params, isError bool) {
  Db := m.GetDb()
  sql, args, _ := sqlQb.Select("*").From(m.tableName).ToSql()
  _, err := Db.Raw(sql, args).Values(&result)

  if err != nil {
    fmt.Println(err.Error())
    isError = true
  }

  return result, isError
}

func (m *Models) Find(id string) (result []orm.Params, isError bool) {
  Db := m.GetDb()
  PkColumn := m.GetPkColumn()

  sqlWhere := make(sqlQb.Eq)
  sqlWhere[PkColumn.Name] = id

  sql, args, _ := sqlQb.Select("*").From(m.tableName).Where(sqlWhere).ToSql()
  _, err := Db.Raw(sql, args).Values(&result)

  if err != nil {
    fmt.Println(err.Error())
    isError = true
  }

  return result, isError
}

func (m *Models) Insert(data map[string]string) (lastId string, isError bool) {
  Db := m.GetDb()
  dataInsert := map[string]interface{}{}
  PkColumn := m.GetPkColumn()

  for key, value := range m.columnList {
    if value.Fillable == true && data[value.Name] != "" {
      dataInsert[key] = data[value.Name]
    }
  }

  sql, args, _ := hpl.InsertSql(m.GetTableName(), dataInsert)

  _, err := Db.Raw(sql, args).Exec()

  if err != nil {
    fmt.Println(err.Error())
    lastId = "0"
    isError = true
  } else {
    if PkColumn.AutoIncrement == true {
      result := []orm.Params{}
      _, err := Db.Raw("SELECT LAST_INSERT_ID() AS id").Values(&result)
      if err != nil {
        fmt.Println(err.Error())
        lastId = "0"
        isError = true
      } else {
        lastId = result[0]["id"].(string)
      }
    } else {
      lastId = data[PkColumn.Name]
    }
  }
  return lastId, isError
}

func (m *Models) Update(id string, data map[string]string) (isError bool) {
  Db := m.GetDb()
  dataUpdate := make(map[string]interface{})

  PkColumn := m.GetPkColumn()
  sqlWhere := make(map[string]interface{})
  sqlWhere[PkColumn.Name] = id

  for _, value := range m.columnList {
    if value.Fillable == true && data[value.Name] != "" {
      dataUpdate[value.Name] = data[value.Name]
    }
  }

  sql, args, _ := hpl.UpdateSql(m.GetTableName(), dataUpdate, sqlWhere)

  _, err := Db.Raw(sql, args).Exec()

  if err != nil {
    fmt.Println(err.Error())
    isError = true
  }

  return isError
}

func (m *Models) Delete(id string) (isError bool) {
  Db := m.GetDb()
  PkColumn := m.GetPkColumn()

  sqlWhere := make(map[string]interface{})
  sqlWhere[PkColumn.Name] = id

  sql, args, _ := hpl.DeleteSql(m.GetTableName(), sqlWhere)

  _, err := Db.Raw(sql, args).Exec()

  if err != nil {
    fmt.Println(err.Error())
    isError = true
  }

  return isError
}

func (m *Models) Count() (count int, isError bool) {
  Db := m.GetDb()
  result := []orm.Params{}
  sqlCount := "SELECT COUNT(" + m.pkColumn.Name + ") AS id FROM " + m.GetTableName()

  _, err := Db.Raw(sqlCount).Values(&result)

  if err != nil {
    fmt.Println(err.Error())
    isError = true
    count = 0
  } else {
    count = result[0]["id"].(int)
  }

  return count, isError
}
