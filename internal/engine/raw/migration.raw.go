package raw

type ExtendedMigrationField struct {
	Name     string
	SQLType  string
	IsUnique bool
}

type ExtendedMigrationModel struct {
	NameUC          string
	NameLC          string
	Fields          []ExtendedMigrationField
	PrimaryKeyField string
}

type MigrationRawData struct {
	Models []ExtendedMigrationModel
}
