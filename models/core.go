package models

import (
  "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql" // import your required driver
  "strings"
  "fmt"
  sqlQb "github.com/Masterminds/squirrel"
)

type Column struct {
  Name string
  Type string
  Fillable bool
  IsPk bool
}

type Models struct {
  tableName string
  sql string
  ColumnList []Column
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

func (m *Models) logSql() {
  fmt.Println(m.sql);
}

func (m *Models) setTableStruct(tableStruct []Column) {
  m.ColumnList = tableStruct
}

func (m *Models) GetTableName() string {
  return m.tableName
}

func (m *Models) GetPrimaryKey() string {
  for _, value := range m.ColumnList {
    if value.IsPk == true {
      return value.Name
    }
  }
  return "id"
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

func (m *Models) Get() interface{} {
  Db := m.GetDb()
  result := []orm.Params{}
  sqlTmp := sqlQb.Select("*");
  sqlTmp = sqlTmp.From(m.tableName);
  sql, args, _ := sqlTmp.ToSql();
  m.sql = sql;
  num, err := Db.Raw(sql, args).Values(&result);
  
  if err != nil {
    panic(err.Error());
  }
  
  if num > 0 {
    return result
  } else {
    return struct{}{}
  }
}

func (m *Models) Find(id interface{}) interface{} {
  Db := m.GetDb()
  result := []orm.Params{}
  
  sqlWhere := make(sqlQb.Eq)
  sqlWhere[m.GetPrimaryKey()] = id.(string)
  
  sqlTmp := sqlQb.Select("*");
  sqlTmp = sqlTmp.From(m.tableName);
  sqlTmp = sqlTmp.Where(sqlWhere);
  
  sql, args, _ := sqlTmp.ToSql();
  num, err := Db.Raw(sql, args).Values(&result);
  
  if err != nil {
    panic(err.Error());
  }
  
  if num > 0 {
    return result[0]
  } else {
    return struct{}{}
  }
}

func (m *Models) Insert(data map[string]interface{}) bool {
  Db := m.GetDb()
  
  columns := []string{}
  values := []interface{}{}
  
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
    return false
  }
  
  return true
}

func (m *Models) Update(id string, data map[string]interface{}) bool {
  Db := m.GetDb()
  
  sqlTmp := sqlQb.Update(m.GetTableName());
  
  sqlWhere := make(sqlQb.Eq)
  sqlWhere[m.GetPrimaryKey()] = id
  
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
  sqlWhere := make(sqlQb.Eq)
  sqlWhere[m.GetPrimaryKey()] = id
  
  sql, args, _ := sqlQb.Delete(m.GetTableName()).Where(sqlWhere).ToSql();
  
  _, err := Db.Raw(sql, args).Exec();
  
  if err != nil {
    fmt.Println(err.Error());
    return false
  }
  
  return true
}
