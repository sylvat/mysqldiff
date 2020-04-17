package internal

import (
	"fmt"
	"strings"
)

var addTables []string
var removeTables []string
var addColumns map[string][]string
var removeColumns map[string][]string
var updateColumns map[string][]map[string]*Column

func HandleDiff() error {
	addColumns = make(map[string][]string)
	removeColumns = make(map[string][]string)
	updateColumns = make(map[string][]map[string]*Column)

	tableMap := make(map[string][]string)
	for k, v := range DBMap {
		tables, err := GetTables(v)
		if err != nil {
			return err
		}
		tableMap[k] = tables
	}
	if err := diffTable(tableMap[Src], tableMap[Dest]); err != nil {
		return err
	}
	printResult()
	return nil
}

func diffTable(srcTable, destTable []string) error {
	for _, s := range srcTable {
		exist := false
		for _, d := range destTable {
			if s == d {
				exist = true
			}
		}
		if !exist {
			addTables = append(addTables, s)
		} else {
			if err := diffColumn(s); err != nil {
				return err
			}
		}
	}
	for _, d := range destTable {
		exist := false
		for _, s := range srcTable {
			if d == s {
				exist = true
			}
		}
		if !exist {
			removeTables = append(removeTables, d)
		}
	}
	return nil
}

func diffColumn(tableName string) error {
	columnMap := make(map[string][]*Column)
	for k, v := range DBMap {
		columns, err := GetColumns(v, tableName)
		if err != nil {
			return err
		}
		columnMap[k] = columns
	}
	for _, s := range columnMap[Src] {
		exist := false
		for _, d := range columnMap[Dest] {
			if s.ColumnName == d.ColumnName {
				exist = true
				if s.ColumnType != d.ColumnType || string(s.ColumnDefault) != string(d.ColumnDefault) {
					uc := make(map[string]*Column)
					uc[Src] = s
					uc[Dest] = d
					updateColumns[tableName] = append(updateColumns[tableName], uc)
				}
			}
		}
		if !exist {
			addColumns[tableName] = append(addColumns[tableName], s.ColumnName)
		}
	}
	for _, d := range columnMap[Dest] {
		exist := false
		for _, s := range columnMap[Src] {
			if d.ColumnName == s.ColumnName {
				exist = true
			}
		}
		if !exist {
			removeColumns[tableName] = append(removeColumns[tableName], d.ColumnName)
		}
	}
	return nil
}

func printResult() {

	fmt.Println("=========================数据库差异对比如下==================================")
	if len(addTables) > 0 {
		fmt.Println(Dest + "数据库缺少数据表：" + strings.Join(addTables, ","))
	}

	if len(removeTables) > 0 {
		fmt.Println(Dest + "数据库增加数据表：" + strings.Join(removeTables, ","))
	}

	fmt.Println("====================================================================")
	fmt.Println(Dest + "数据库缺少字段：")
	for k, v := range addColumns {
		fmt.Println(RedText(k) + "表缺少字段：" + strings.Join(v, ","))
	}
	fmt.Println("====================================================================")
	fmt.Println(Dest + "数据库增加字段：")
	for k, v := range removeColumns {
		fmt.Println(RedText(k) + "表增加字段：" + strings.Join(v, ","))
	}
	fmt.Println("====================================================================")
	fmt.Println(fmt.Sprintf("%s数据库和%s数据库表结构差异：", Src, Dest))
	for km, vm := range updateColumns {
		fmt.Println(RedText(km) + "表：")
		for _, v := range vm {
			fmt.Println(fmt.Sprintf("字段%s:%s-类型为%s,默认值为%s;%s-类型为%s,默认值为%s",
				RedText(v[Src].ColumnName), Src, v[Src].ColumnType, v[Src].ColumnDefault, Dest, v[Dest].ColumnType, v[Dest].ColumnDefault))
		}
	}
}
