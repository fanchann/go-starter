package code

var MigrateCode = `package helpers

func MigrateDB(models ...interface{}) error {
	return InitDatabase().DatabaseConnection().AutoMigrate(models...)
}
`
