package models

import (
  "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql" // import your required driver
  "strings"
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
  ColumnList []Column
  PkColumn Column
}

func (m *Models) GetDb() orm.Ormer {
  return orm.NewOrm()
}

func NewModels(tableName string, tableStruct []Column) *Models {
  return &Models{
    tableName: tableName,
    ColumnList: tableStruct,
  }
}

func (m *Models) setTableStruct(tableStruct []Column) {
  m.ColumnList = tableStruct
}

func (m *Models) GetTableName() string {
  return m.tableName
}

func (m *Models) GetPrimaryKey() Column {
  for i, value := range m.ColumnList {
    if value.IsPk == true {
      return m.ColumnList[i]
    }
  }
  return m.ColumnList[0]
}

func (m *Models) GetColumnSql() string {
  listColumn := []string{}
  for _, value := range m.ColumnList {
    if value.Fillable == true {
      listColumn = append(listColumn, value.Name);
    }
  }
  return strings.Join(listColumn, ",");
}

func (m *Models) Get() (interface{}, bool) {
  Db := m.GetDb()
  result := []orm.Params{}
  sqlTmp := sqlQb.Select("*");
  sqlTmp = sqlTmp.From(m.tableName);
  sql, args, _ := sqlTmp.ToSql();
  num, err := Db.Raw(sql, args).Values(&result);
  
  if err != nil {
    fmt.Println(err.Error());
    return struct{}{}, true
  }
  
  if num > 0 {
    return result, false
  } else {
    return struct{}{}, false
  }
}

func (m *Models) Find(id interface{}) (interface{}, bool) {
  Db := m.GetDb()
  result := []orm.Params{}
  PkColumn := m.GetPrimaryKey()
  
  sqlWhere := make(sqlQb.Eq)
  sqlWhere[PkColumn.Name] = id.(string)
  
  sqlTmp := sqlQb.Select("*");
  sqlTmp = sqlTmp.From(m.tableName);
  sqlTmp = sqlTmp.Where(sqlWhere);
  
  sql, args, _ := sqlTmp.ToSql();
  num, err := Db.Raw(sql, args).Values(&result);
  
  if err != nil {
    fmt.Println(err.Error());
    return struct{}{}, true
  }
  
  if num > 0 {
    return result[0], false
  } else {
    return struct{}{}, false
  }
}

func (m *Models) Insert(data map[string]interface{}) (string, bool) {
  Db := m.GetDb()
  columns := []string{}
  values := []interface{}{}
  PkColumn := m.GetPrimaryKey()
  
  for _, value := range m.ColumnList {
    if value.Fillable == true {
      columns = append(columns, value.Name);
      values = append(values, data[value.Name]);
    }
  }
  
  sql, args, _ := sqlQb.Insert(m.GetTableName()).Columns(columns...).Values(values...).ToSql()
  
  _, err := Db.Raw(sql, args).Exec();
  
  if err != nil {
    fmt.Println(err.Error());
    return "0", true
  }

  if PkColumn.AutoIncrement == true {
    result := []orm.Params{}
    sqlLastID := "SELECT LAST_INSERT_ID() AS id"
    _, err := Db.Raw(sqlLastID).Values(&result);
    if err != nil {
      fmt.Println(err.Error());
      return "0", true
    }
    return result[0]["id"].(string), false
  } else {
    return data[PkColumn.Name].(string), false
  }
}

func (m *Models) Update(id string, data map[string]interface{}) bool {
  Db := m.GetDb()
  PkColumn := m.GetPrimaryKey()
  sqlTmp := sqlQb.Update(m.GetTableName());
  
  sqlWhere := make(sqlQb.Eq)
  sqlWhere[PkColumn.Name] = id
  
  for _, value := range m.ColumnList {
    if value.Fillable == true {
      sqlTmp = sqlTmp.Set(value.Name, data[value.Name]);
    }
  }
  
  sql, args, _ := sqlTmp.Where(sqlWhere).ToSql()
  
  _, err := Db.Raw(sql, args).Exec();
  
  if err != nil {
    fmt.Println(err.Error());
    return false
  }
  
  return true
}

func (m *Models) Delete(id string) bool {
  Db := m.GetDb()
  PkColumn := m.GetPrimaryKey();
  
  sqlWhere := make(sqlQb.Eq)
  sqlWhere[PkColumn.Name] = id
  
  sql, args, _ := sqlQb.Delete(m.GetTableName()).Where(sqlWhere).ToSql();
  
  _, err := Db.Raw(sql, args).Exec();
  
  if err != nil {
    fmt.Println(err.Error());
    return false
  }
  
  return true
}

func (m *Models) Count() (int, bool) {
  Db := m.GetDb()
  result := []orm.Params{}
  
  sqlCount := "SELECT COUNT(" + m.PkColumn.Name + ") AS id FROM " + m.GetTableName();
  
  _, err := Db.Raw(sqlCount).Values(&result);
  
  if err != nil {
    fmt.Println(err.Error());
    return 0, true
  }
  
  return result[0]["id"].(int), false
}
