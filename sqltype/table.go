package sqltype

import (
	"fmt"
)

type ColumnsMap map[string]int

type Table struct {
	Name            string
	Schema          string
	PrimaryKey      *Column
	Columns         []*Column
	ColumnsOrdinals ColumnsMap
}

func (this *Table) AddColumn(column *Column) {
	this.Columns = append(this.Columns, column)
}

func (this *Table) GetSchameAndTable() string {
	return fmt.Sprintf("%s.%s", this.Schema, this.Name)
}
