package helper

import (
	sq "github.com/Masterminds/squirrel"
	str "strings"
)

func InsertSql(table string, data map[string]interface{}) (string, []interface{}, error) {
	sqlResult := sq.Insert(table)

	for keys, value := range data {
		sqlResult = sqlResult.Values(keys, value)
	}
	sql, args, err := sqlResult.ToSql()
	return sql, args, err
}

func OperatorSql(column string, value interface{}) (result sq.Sqlizer) {
	op := str.Split(column, " ")
	// operator sama dengan
	if len(op) < 2 {
		result = sq.Eq{column: value}
	} else if len(op) == 2 {
		// op[0] = nama kolom
		// op[1] = operator
		if op[1] != "OR" { // kolom selain OR akan bernilai tunggal
			switch op[1] {
			case "[>]":
				result = sq.Gt{op[0]: value}
			case "[<]":
				result = sq.Lt{op[0]: value}
			case "[>=]":
				result = sq.GtOrEq{op[0]: value}
			case "[<=]":
				result = sq.LtOrEq{op[0]: value}
			case "[!=]":
				result = sq.NotEq{op[0]: value}
			case "[~]":
				result = sq.Like{op[0]: value}
			default:
				result = sq.Eq{column: value}
			}
		} else if op[1] == "OR" { // isi dari kolom OR (value) adalah map[string]interface{}
			result := sq.Or{}
			for keys, val := range value.(map[string]interface{}) {
				result = append(result, OperatorSql(keys, val)) // semua isi kolom OR diproses rekursif oleh fungsi ini
			}
		}

	}
	return result
}

func UpdateSql(table string, data map[string]interface{}, where map[string]interface{}) (string, []interface{}, error) {
	sqlResult := sq.Update(table)
	for keys, value := range data {
		sqlResult = sqlResult.Set(keys, value)
	}

	for keys, value := range where {
		sqlResult = sqlResult.Where(OperatorSql(keys, value))
	}

	sql, args, err := sqlResult.ToSql()
	return sql, args, err
}

func DeleteSql(table string, where map[string]interface{}) (string, []interface{}, error) {
	sqlResult := sq.Delete(table)
	for keys, value := range where {
		sqlResult = sqlResult.Where(OperatorSql(keys, value))
	}

	sql, args, err := sqlResult.ToSql()
	return sql, args, err
}
