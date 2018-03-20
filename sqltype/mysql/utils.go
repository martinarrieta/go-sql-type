package mysqltype

import (
	"database/sql"
	"log"

	"github.com/martinarrieta/db-sql-type/sqltype"
)

func ProcessColumn(rows sql.Rows) (*sqltype.Column, error) {
	column := new(sqltype.Column)

	temp := new(string)
	err := rows.Scan(&column.Name, &column.OrdinalPosition, &column.Default, &temp,
		&column.DataType, &column.CharacterMaxLength, &column.CharacterOctetLength,
		&column.NumericPrecision, &column.NumericScale, &column.DateTimePrecision,
		&column.CharacterSetName, &column.CollationName, &column.ColumnKey,
		&column.Extra)

	if err != nil {
		return nil, err
	}

	if *temp == "YES" {
		column.IsNullable = true
	} else {
		column.IsNullable = false
	}

	if column.ColumnKey == "PRI" {
		column.IsPrimaryKey = true
	} else {
		column.IsPrimaryKey = false
	}

	return column, nil
}

func NewTable(schema string, name string, db *sql.DB) *sqltype.Table {

	query := `SELECT
		COLUMN_NAME,
		ORDINAL_POSITION,
		COLUMN_DEFAULT,
		IS_NULLABLE,
		DATA_TYPE,
		CHARACTER_MAXIMUM_LENGTH,
		CHARACTER_OCTET_LENGTH,
		NUMERIC_PRECISION,
		NUMERIC_SCALE,
		DATETIME_PRECISION,
		CHARACTER_SET_NAME,
		COLLATION_NAME,
		COLUMN_KEY,
		EXTRA
	 FROM information_schema.columns WHERE table_schema=? AND table_name=? ORDER BY ORDINAL_POSITION`

	rows, err := db.Query(query, schema, name)

	if err != nil && err != sql.ErrNoRows {
		return nil
	}

	var columns []*sqltype.Column
	columnsOrdinals := make(map[string]int)
	var pk *sqltype.Column

	for rows.Next() {

		column, err := ProcessColumn(*rows)
		if err != nil {
			log.Fatal("error", err.Error())
		}
		columns = append(columns, column)
		if column.IsPrimaryKey {
			pk = column
		}
		columnsOrdinals[column.Name] = column.OrdinalPosition - 1
	}

	table := &sqltype.Table{
		Name:            name,
		Schema:          schema,
		Columns:         columns,
		ColumnsOrdinals: columnsOrdinals,
		PrimaryKey:      pk,
	}

	return table
}
