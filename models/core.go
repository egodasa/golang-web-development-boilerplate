package models

import (
  "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql" // import your required driver
  "fmt"
  sqlQb "github.com/Masterminds/squirrel"
)

type IModels interface {
  Get() (interface{}, bool)
  Find(id interface{}) (interface{}, bool)
  Delete(id string) bool
  GetColumnSql() string
  GetDb() orm.Ormer
  GetPrimaryKey() Column
  GetTableName() string
  Insert(data map[string]interface{}) (string, bool)
  Update(id string, data map[string]interface{}) bool
  Count() (count int, isError bool)
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
  ColumnList map[string]Column
  pkColumn Column
  pKname string
  SelectStatement sqlQb.SelectBuilder
}

func (m *Models) GetDb() orm.Ormer {
  return orm.NewOrm()
}

func NewModels(tableName string, pk string, tableStruct map[string]Column) *Models {
  return &Models{
    tableName: tableName,
    ColumnList: tableStruct,
    pKname: pk,
  }
}

func (m *Models) SetCustomSelect(sql sqlQb.SelectBuilder) {
  m.SelectStatement = sql
}

func (m *Models) GetTableName() string {
  return m.tableName
}

func (m *Models) GetPrimaryKey() Column {
  return m.ColumnList[m.pKname]
}

func (m *Models) GetColumnSql() (column []string) {
  for _, value := range m.ColumnList {
    column = append(column, value.Name);
  }
  return column;
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
  PkColumn := m.GetPrimaryKey()
  
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
  PkColumn := m.GetPrimaryKey()
  
  for _, value := range m.ColumnList {
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
  PkColumn := m.GetPrimaryKey()
  sqlTmp := sqlQb.Update(m.GetTableName());
  
  sqlWhere := make(sqlQb.Eq)
  sqlWhere[PkColumn.Name] = id
  
  for _, value := range m.ColumnList {
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
  PkColumn := m.GetPrimaryKey();
  
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
