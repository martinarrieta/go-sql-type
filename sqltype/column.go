package sqltype

type Column struct {
	Name                 string
	OrdinalPosition      int
	Default              string
	IsNullable           bool
	DataType             string
	CharacterMaxLength   int
	CharacterOctetLength int
	NumericPrecision     int
	NumericScale         int
	DateTimePrecision    int
	CharacterSetName     string
	CollationName        string
	ColumnKey            string
	Extra                string
	IsPrimaryKey         bool
}
