package helper

import (
	sq "github.com/Masterminds/squirrel"
	str "strings"
)

func SelectSql(table string, column []string, whereList map[string]interface{}) (string, []interface{}, error) {
	result := sq.Select(column...).From(table)
	for keys, val := range whereList {
		op := str.Split(keys, " ")
		if len(op) == 2 {
			switch op[0] {
			case "[>]": // [>] == LEFT JOIN
				result = result.LeftJoin(op[1] + " ON " + val.(string))
			case "[<]": // [<] == RIGH JOIN
				result = result.RightJoin(op[1] + " ON " + val.(string))
			case "[><]": // [><] == INNER JOIN
				result = result.Join(op[1] + " ON " + val.(string))
			default: // TANPA JOIN, LANGSUNG WHERE
				result = result.Where(WhereSql(keys, val))
			}
		} else if len(op) == 1 {
			switch keys {
			case "ORDER":
				orderList := val.([]string)
				result = result.OrderBy(orderList...)
			case "GROUP":
				groupList := val.([]string)
				result = result.GroupBy(groupList...)
			case "HAVING":
				dataHaving := val.(map[string]interface{})
				for keys, value := range dataHaving {
					result = result.Having(WhereSql(keys, value))
				}
			case "LIMIT":
				result = result.Limit(uint64(val.(int)))
			case "OFFSET":
				result = result.Offset(uint64(val.(int)))
			}
		} else {
			panic("Format Kolom SQL Salah! Nama Kolom maksimal 1 Spasi")
		}

	}
	sql, args, err := result.ToSql()
	return sql, args, err
}

func InsertSql(table string, data map[string]interface{}) (string, []interface{}, error) {
	sqlResult := sq.Insert(table)

	for keys, value := range data {
		sqlResult = sqlResult.Values(keys, value)
	}
	sql, args, err := sqlResult.ToSql()
	return sql, args, err
}

func UpdateSql(table string, data map[string]interface{}, where map[string]interface{}) (string, []interface{}, error) {
	sqlResult := sq.Update(table)
	for keys, value := range data {
		sqlResult = sqlResult.Set(keys, value)
	}

	for keys, value := range where {
		sqlResult = sqlResult.Where(WhereSql(keys, value))
	}

	sql, args, err := sqlResult.ToSql()
	return sql, args, err
}

func DeleteSql(table string, where map[string]interface{}) (string, []interface{}, error) {
	sqlResult := sq.Delete(table)
	for keys, value := range where {
		sqlResult = sqlResult.Where(WhereSql(keys, value))
	}

	sql, args, err := sqlResult.ToSql()
	return sql, args, err
}

func WhereSql(column string, value interface{}) (result sq.Sqlizer) {
	op := str.Split(column, " ")
	// operator sama dengan
	if len(op) == 1 {
		switch column {
		case "OR":
			dataOr := value.(map[string]interface{})
			sqlOr := sq.Or{}
			for keys, val := range dataOr {
				sqlOr = append(sqlOr, WhereSql(keys, val))
			}
			result = sqlOr
		default:
			result = sq.Eq{column: value}
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
	} else {
		panic("Format SQL Where Salah!")
	}
	return result
}
