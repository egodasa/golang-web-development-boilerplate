package models

/*
Select data :
Models.Select().Get()
Models.Select().Where(sqlQb.Eq{"kolom1": 123}).Get()
Models.Select().Where(sqlQb.Eq{"kolom1": 123}).OrderBy("kolom1 ASC").Get()
Models.Select().Where(sqlQb.Eq{"kolom1": 123}).Having(sqlQb.Eq{"kolom1": 123}).OrderBy("kolom1 ASC").Get()
Models.Select("tabel1.*", "tabel2.nm_jenis").Join("tabel2", "tabel1.id_tabel1 = tabel2.id_tabel2").Get()


Insert Data :
Models.SetValue("kolom1", "data 1")
Models.SetValue("Kolom2", 2)
Models.Insert()

Update Data :
Models.SetValue("kolom1", "data 1")
Models.SetValue("Kolom2", 2)
Models.Where(sqlQb.Eq{"kolom1": 123}).Update(1)

Delete Data :
Models.Delete(1)
*/

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
	LeftJoin(string, string) *Models
	Join(string, string) *Models
	RightJoin(string, string) *Models
	OrderBy(...string) *Models
	GroupBy(...string) *Models
	Having(interface{}) *Models
	Where(interface{}) *Models
	Get() ([]orm.Params, bool)
	First() (orm.Params, bool)
	SetValue(column string, data interface{})
	Insert() bool
	Update(string) bool
	Delete(string) bool
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
	tableName      string
	pkColumn       Column
	columnList     map[string]Column
	defaultSelect  sqlQb.SelectBuilder
	defaultInsert  sqlQb.InsertBuilder
	defaultUpdate  sqlQb.UpdateBuilder
	defaultDelete  sqlQb.DeleteBuilder
	Data           map[string]interface{}
	currentSqlType string
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
		Data:          map[string]interface{}{},
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

func (m *Models) ResetDefaultQuery() {
	m.defaultSelect = sqlQb.Select("*").From(m.GetTableName())
	m.defaultInsert = sqlQb.Insert(m.GetTableName())
	m.defaultUpdate = sqlQb.Update(m.GetTableName())
	m.defaultDelete = sqlQb.Delete(m.GetTableName())
}

func (m *Models) All() ([]orm.Params, bool) {
	m.currentSqlType = "SELECT"
	return m.Get()
}

func (m *Models) Find(id string) (result orm.Params, isError bool) {
	m.defaultSelect = m.defaultSelect.Where(sqlQb.Eq{m.GetPkColumn().Name: id}).Limit(1)
	m.currentSqlType = "SELECT"
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
		m.defaultSelect = sqlQb.Select(columns...).From(m.GetTableName())
	}
	m.currentSqlType = "SELECT"
	return m
}

// Save untuk menjalankan query insert dan update
func (m *Models) Insert() bool {
	for keys, value := range m.Data {
		m.defaultInsert = m.defaultInsert.Values(keys, value)
	}
	m.currentSqlType = "INSERT"
	return m.Run()
}

func (m *Models) Update(id string) bool {
	for keys, value := range m.Data {
		m.defaultUpdate = m.defaultUpdate.Set(keys, value)
	}
	m.defaultUpdate = m.defaultUpdate.Where(sqlQb.Eq{m.GetPkColumn().Name: id})
	m.currentSqlType = "UPDATE"
	return m.Run()
}

func (m *Models) Delete(id string) bool {
	m.defaultDelete = m.defaultDelete.Where(sqlQb.Eq{m.GetPkColumn().Name: id})
	m.currentSqlType = "DELETE"
	return m.Run()
}

func (m *Models) LeftJoin(table string, relation string) *Models {
	m.defaultSelect = m.defaultSelect.LeftJoin(table + " ON " + relation)
	m.currentSqlType = "SELECT"
	return m
}

func (m *Models) Join(table string, relation string) *Models {
	m.defaultSelect = m.defaultSelect.Join(table + " ON " + relation)
	m.currentSqlType = "SELECT"
	return m
}

func (m *Models) RightJoin(table string, relation string) *Models {
	m.defaultSelect = m.defaultSelect.RightJoin(table + " ON " + relation)
	m.currentSqlType = "SELECT"
	return m
}

func (m *Models) OrderBy(columns ...string) *Models {
	m.defaultSelect = m.defaultSelect.OrderBy(columns...)
	m.currentSqlType = "SELECT"
	return m
}

func (m *Models) GroupBy(columns ...string) *Models {
	m.defaultSelect = m.defaultSelect.GroupBy(columns...)
	m.currentSqlType = "SELECT"
	return m
}

func (m *Models) Having(condition interface{}) *Models {
	m.defaultSelect = m.defaultSelect.Having(condition)
	m.currentSqlType = "SELECT"
	return m
}

func (m *Models) Where(condition interface{}) *Models {
	switch m.currentSqlType {
	case "SELECT":
		m.defaultSelect = m.defaultSelect.Where(condition)
	case "UPDATE":
		m.defaultUpdate = m.defaultUpdate.Where(condition)
	case "DELETE":
		m.defaultDelete = m.defaultDelete.Where(condition)
	}
	return m
}

func (m *Models) SetValue(column string, val interface{}) {
	for _, value := range m.columnList {
		if value.Fillable == true {
			m.Data[column] = val
		}
	}
}

// Get untuk menjalankan select data
// hasil return berupa []map[string]interface{} dan error
func (m *Models) Get() (result []orm.Params, isError bool) {
	Db := m.GetDb()
	sql, args, _ := m.defaultSelect.ToSql()
	_, err := Db.Raw(sql, args...).Values(&result)
	if err != nil {
		log.Println(err.Error())
		isError = true
	}

	m.ResetDefaultQuery()
	return result, isError
}

// First mengambil satu data dari hasil select
// hasil return berupa map[string]interface{}
func (m *Models) First() (result orm.Params, isError bool) {
	Db := m.GetDb()
	resultTmp := []orm.Params{}
	sql, args, _ := m.defaultSelect.Limit(1).ToSql()
	_, err := Db.Raw(sql, args...).Values(&resultTmp)
	if err != nil {
		log.Println(err.Error())
		isError = true
	} else {
		if len(resultTmp) != 0 {
			result = resultTmp[0]
		}
	}

	m.ResetDefaultQuery()
	return result, isError
}

// Run method untuk menjalankan delete
// Hasil return berupa error
func (m *Models) Run() (isError bool) {
	var sql string
	var args []interface{}

	switch m.currentSqlType {
	case "SELECT":
		sql, args, _ = m.defaultSelect.ToSql()
	case "INSERT":
		sql, args, _ = m.defaultInsert.ToSql()
	case "UPDATE":
		sql, args, _ = m.defaultUpdate.ToSql()
	case "DELETE":
		sql, args, _ = m.defaultDelete.ToSql()
	}
	_, err := m.GetDb().Raw(sql, args).Exec()

	if err != nil {
		log.Println(err.Error())
		isError = true
	}

	m.ResetDefaultQuery()
	m.Data = map[string]interface{}{}
	return isError
}
