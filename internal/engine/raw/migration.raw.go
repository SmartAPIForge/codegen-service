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
}

type MigrationRawData struct {
	Models []ExtendedMigrationModel
}
