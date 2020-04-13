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

func (m *Models) getTableName() string {
  return m.tableName
}

func (m *Models) getPrimaryKey() string {
  for _, value := range m.ColumnList {
    if value.IsPk == true {
      return value.Name
    }
  }
  return "id"
}

func (m *Models) getColumnSql() string {
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
  
  m.logSql();
  
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
  sqlWhere[m.getPrimaryKey()] = id.(string)
  
  sqlTmp := sqlQb.Select("*");
  sqlTmp = sqlTmp.From(m.tableName);
  sqlTmp = sqlTmp.Where(sqlWhere);
  
  sql, args, _ := sqlTmp.ToSql();
  m.sql = sql;
  num, err := Db.Raw(sql, args).Values(&result);
  
  m.logSql();
  
  if err != nil {
    panic(err.Error());
  }
  
  if num > 0 {
    return result[0]
  } else {
    return struct{}{}
  }
}
