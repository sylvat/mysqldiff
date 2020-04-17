package internal

type Column struct {
	ColumnName    string `db:"COLUMN_NAME"`
	ColumnType    string `db:"COLUMN_TYPE"`
	ColumnDefault []byte `db:"COLUMN_DEFAULT"`
	ColumnComment []byte `db:"COLUMN_COMMENT"`
}

func GetTables(db *DBData) ([]string, error) {
	var tables []string
	query := "select table_name from information_schema.tables where table_schema=?"
	err := db.DB.Select(&tables, query, db.DBName)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func GetColumns(db *DBData, tableName string) ([]*Column, error) {
	var columns []*Column
	query := "select column_name,column_type,column_default,column_comment from information_schema.columns where table_schema=? and table_name=?"
	err := db.DB.Select(&columns, query, db.DBName, tableName)
	if err != nil {
		return nil, err
	}
	return columns, nil
}
