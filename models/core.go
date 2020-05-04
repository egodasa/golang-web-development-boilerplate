package models

import (
	"log"
	"strconv"

	sqlQb "github.com/Masterminds/squirrel"
	orm "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
)

type IModels interface {
	GetTableName() string
	GetPkColumn() Column
	GetColumnList() map[string]Column
	GetColumnListAsSql() []string

	All() ([]orm.Params, bool)
	Find(string) (result orm.Params, isError bool)
	Count() (int, bool)

	Select(...string) *Models
	Insert(data map[string]string) *Models
	Update(string, map[string]string) *Models
	Delete(string) *Models
	Get() ([]orm.Params, bool)
	First() (orm.Params, bool)
	Run() bool
}

type Column struct {
	Name          string
	Type          string
	Fillable      bool
	IsPk          bool
	AutoIncrement bool
	DefaultValue  interface{}
}

type Models struct {
	tableName     string
	pkColumn      Column
	columnList    map[string]Column
	currentSql    sqlQb.Sqlizer
	defaultSelect sqlQb.SelectBuilder
	defaultInsert sqlQb.InsertBuilder
	defaultUpdate sqlQb.UpdateBuilder
	defaultDelete sqlQb.DeleteBuilder
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
		tableName:     tableName,
		columnList:    tableStruct,
		defaultSelect: sqlQb.Select("*").From(tableName),
		defaultInsert: sqlQb.Insert(tableName),
		defaultUpdate: sqlQb.Update(tableName),
		defaultDelete: sqlQb.Delete(tableName),
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

func (m *Models) All() ([]orm.Params, bool) {
	m.currentSql = m.defaultSelect
	return m.Get()
}

func (m *Models) Find(id string) (result orm.Params, isError bool) {
	m.currentSql = m.defaultSelect.Where(sqlQb.Eq{m.GetPkColumn().Name: id}).Limit(1)
	return m.First()
}

// Count method untuk menghitung banyak seluruh record pada sebuah models \n
// Hasil return berupa banyak record integer serta error bool
func (m *Models) Count() (count int, isError bool) {
	Db := m.GetDb()
	result := []orm.Params{}
	PkColumn := m.GetPkColumn()
	sqlCount := "SELECT COUNT(" + PkColumn.Name + ") AS id FROM " + m.GetTableName()

	_, err := Db.Raw(sqlCount).Values(&result)

	if err != nil {
		log.Println(err.Error())
		isError = true
		count = 0
	} else {
		totalCount, _ := strconv.Atoi(result[0]["id"].(string))
		count = totalCount
	}

	return count, isError
}

func (m *Models) Select(columns ...string) *Models {
	if len(columns) != 0 {
		m.currentSql = sqlQb.Select(columns...).From(m.GetTableName())
	} else {
		m.currentSql = m.defaultSelect
	}
	return m
}

func (m *Models) Insert(data map[string]string) *Models {
	for key, value := range m.columnList {
		if value.Fillable == true && data[value.Name] != "" {
			m.defaultInsert = m.defaultInsert.Values(key, data[value.Name])
		} else {
			if value.DefaultValue != nil {
				m.defaultInsert = m.defaultInsert.Values(key, value.DefaultValue)
			}
		}
	}
	m.currentSql = m.defaultInsert
	return m
}

func (m *Models) Update(id string, data map[string]string) *Models {
	for _, value := range m.columnList {
		if value.Fillable == true && data[value.Name] != "" {
			m.defaultUpdate = m.defaultUpdate.Set(value.Name, data[value.Name])
		}
	}
	m.currentSql = m.defaultUpdate.Where(sqlQb.Eq{m.GetPkColumn().Name: id})
	return m
}

func (m *Models) Delete(id string) *Models {
	m.currentSql = m.defaultDelete.Where(sqlQb.Eq{m.GetPkColumn().Name: id})
	return m
}

// Get untuk menjalankan select data
// hasil return berupa []map[string]interface{} dan error
func (m *Models) Get() (result []orm.Params, isError bool) {
	Db := m.GetDb()
	sql, args, _ := m.currentSql.ToSql()
	_, err := Db.Raw(sql, args...).Values(&result)
	if err != nil {
		log.Println(err.Error())
		isError = true
	}
	return result, isError
}

// First mengambil satu data dari hasil select
// hasil return berupa map[string]interface{}
func (m *Models) First() (result orm.Params, isError bool) {
	Db := m.GetDb()
	resultTmp := []orm.Params{}
	sql, args, _ := m.currentSql.ToSql()
	_, err := Db.Raw(sql, args...).Values(&resultTmp)
	if err != nil {
		log.Println(err.Error())
		isError = true
	} else {
		if len(resultTmp) != 0 {
			result = resultTmp[0]
		}
	}
	return result, isError
}

// Run method untuk menjalankan insert, update dan delete
// Hasil return berupa error
func (m *Models) Run() (isError bool) {
	sql, args, _ := m.currentSql.ToSql()
	_, err := m.GetDb().Raw(sql, args).Exec()

	if err != nil {
		log.Println(err.Error())
		isError = true
	}
	return isError
}
