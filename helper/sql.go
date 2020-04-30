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

func SelectSql(table string, column []string, whereList map[string]interface{}) (string, []interface{}, error) {
	result := sq.Select(column...).From(table)
	for keys, val := range whereList {
		// cek apakah ada join atau tidak
		// [>] == LEFT JOIN
		// [<] == RIGH JOIN
		// [<>] == FULL JOIN
		// [><] == INNER JOIN
		op := str.Split(keys, " ")
		switch op[0] {
		case "[>]":
			result = result.LeftJoin(op[1] + " ON " + val.(string))
		case "[<]":
			result = result.RightJoin(op[1] + " ON " + val.(string))
		case "[><]":
			result = result.Join(op[1] + " ON " + val.(string))
		default:
			result = result.Where(ConditionSql(keys, val))
		}
	}
	sql, args, err := result.ToSql()
	return sql, args, err
}

func UpdateSql(table string, data map[string]interface{}, where map[string]interface{}) (string, []interface{}, error) {
	sqlResult := sq.Update(table)
	for keys, value := range data {
		sqlResult = sqlResult.Set(keys, value)
	}

	for keys, value := range where {
		sqlResult = sqlResult.Where(ConditionSql(keys, value))
	}

	sql, args, err := sqlResult.ToSql()
	return sql, args, err
}

func DeleteSql(table string, where map[string]interface{}) (string, []interface{}, error) {
	sqlResult := sq.Delete(table)
	for keys, value := range where {
		sqlResult = sqlResult.Where(ConditionSql(keys, value))
	}

	sql, args, err := sqlResult.ToSql()
	return sql, args, err
}

func ConditionSql(column string, value interface{}) (result sq.Sqlizer) {
	op := str.Split(column, " ")
	// operator sama dengan
	if len(op) < 2 {
		if column != "OR" {
			result = sq.Eq{column: value}
		} else {
			result = OperatorOrSql(value.(map[string]interface{}))
		}
	} else if len(op) == 2 {
		// op[0] = nama kolom
		// op[1] = operator
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
	}
	return result
}

func OperatorOrSql(columnList map[string]interface{}) (result sq.Or) {
	for keys, val := range columnList {
		result = append(result, ConditionSql(keys, val)) // semua isi kolom OR diproses rekursif oleh fungsi ini
	}
	return result
}
