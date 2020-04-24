package models

import (
  orm "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql" // import your required driver
  "fmt"
  sqlQb "github.com/Masterminds/squirrel"
)

type IModels interface {
  GetTableName() string
  GetPkColumn() Column
  GetColumnList() map[string]Column
}

type Column struct {
  Name string
  Type string
  Fillable bool
  IsPk bool
  AutoIncrement bool
}

type Models struct {
  tableName string
  pkColumn Column
  columnList map[string]Column
  relatedModels []IModels
}

func (m *Models) GetDb() orm.Ormer {
  return orm.NewOrm()
}

func NewModels(tableName string, tableStruct map[string]Column, relatedModels []IModels) *Models {
  var banyakPk int
  for _, value := range tableStruct {
    if value.IsPk == true {
      banyakPk += 1
    }
  }
  if banyakPk != 1 {
    panic("Primary Key berlebih atau belum diset!");
  }
  
  return &Models{
    tableName: tableName,
    columnList: tableStruct,
    relatedModels: relatedModels,
  }
}

func (m *Models) GetRelatedModels() []IModels {
  return m.relatedModels
}

func (m *Models) AddRelatedModels(model IModels) {
  m.relatedModels = append(m.relatedModels, model);
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

func (m *Models) Get() (result []orm.Params, isError bool) {
  Db := m.GetDb()
  sql, args, _ := sqlQb.Select("*").From(m.tableName).ToSql();
  _, err := Db.Raw(sql, args).Values(&result);
  
  if err != nil {
    fmt.Println(err.Error());
    isError = true
  }
  
  return result, isError
}

func (m *Models) Find(id string) (result []orm.Params, isError bool) {
  Db := m.GetDb()
  PkColumn := m.GetPkColumn()
  
  sqlWhere := make(sqlQb.Eq)
  sqlWhere[PkColumn.Name] = id

  sql, args, _ := sqlQb.Select("*").From(m.tableName).Where(sqlWhere).ToSql();
  _, err := Db.Raw(sql, args).Values(&result);
  
  if err != nil {
    fmt.Println(err.Error());
    isError = true
  }
  
  return result, isError
}

func (m *Models) Insert(data map[string]string) (lastId string, isError bool) {
  Db := m.GetDb()
  columns := []string{}
  values := []interface{}{}
  PkColumn := m.GetPkColumn()
  
  for _, value := range m.columnList {
    if value.Fillable == true && data[value.Name] != "" {
      columns = append(columns, value.Name);
      values = append(values, data[value.Name]);
    }
  }
  
  sql, args, _ := sqlQb.Insert(m.GetTableName()).Columns(columns...).Values(values...).ToSql()
  
  _, err := Db.Raw(sql, args).Exec();
  
  if err != nil {
    fmt.Println(err.Error());
    lastId = "0"
    isError = true
  } else {
      if PkColumn.AutoIncrement == true {
        result := []orm.Params{}
        _, err := Db.Raw("SELECT LAST_INSERT_ID() AS id").Values(&result);
        if err != nil {
          fmt.Println(err.Error());
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
  PkColumn := m.GetPkColumn()
  sqlTmp := sqlQb.Update(m.GetTableName());
  
  sqlWhere := make(sqlQb.Eq)
  sqlWhere[PkColumn.Name] = id
  
  for _, value := range m.columnList {
    if value.Fillable == true && data[value.Name] != "" {
      sqlTmp = sqlTmp.Set(value.Name, data[value.Name]);
    }
  }
  
  sql, args, _ := sqlTmp.Where(sqlWhere).ToSql()
  
  _, err := Db.Raw(sql, args).Exec();
  
  if err != nil {
    fmt.Println(err.Error());
    isError = true
  }
  
  return isError
}

func (m *Models) Delete(id string) (isError bool) {
  Db := m.GetDb()
  PkColumn := m.GetPkColumn();
  
  sqlWhere := make(sqlQb.Eq)
  sqlWhere[PkColumn.Name] = id
  
  sql, args, _ := sqlQb.Delete(m.GetTableName()).Where(sqlWhere).ToSql();
  
  _, err := Db.Raw(sql, args).Exec();
  
  if err != nil {
    fmt.Println(err.Error());
    isError = true
  }
  
  return isError
}

func (m *Models) Count() (count int, isError bool) {
  Db := m.GetDb()
  result := []orm.Params{};
  sqlCount := "SELECT COUNT(" + m.pkColumn.Name + ") AS id FROM " + m.GetTableName();
  
  _, err := Db.Raw(sqlCount).Values(&result);
  
  if err != nil {
    fmt.Println(err.Error());
    isError = true
    count = 0
  } else {
      count = result[0]["id"].(int)
  }
  
  return count, isError
}
