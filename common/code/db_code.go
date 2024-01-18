package code

type AppConfiguration struct {
	Dependencies      string
	DatabaseDriver    string
	ConfigurationFile map[string]interface{}
	DockerCompose     map[string]interface{}
	MainApp           string
}

var MysqlDBConfig = `package config

import (
	"fmt"
	"time"

	"{{.Package}}/internals/helpers"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConnection(v *viper.Viper) *gorm.DB {
	username := v.GetString("database.username")
	password := v.GetString("database.password")
	host := v.GetString("database.host")
	port := v.GetInt("database.port")
	database := v.GetString("database.name")
	idleConnection := v.GetInt("database.pool.idle")
	maxConnection := v.GetInt("database.pool.max")
	maxLifeTimeConnection := v.GetInt("database.pool.lifetime")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helpers.ErrorLogger(err)

	connection, err := db.DB()
	helpers.ErrorLogger(err)

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}
`

var PostgresDBConfig = `package config

import (
	"{{.Package}}/internals/helpers"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(v *viper.Viper) *gorm.DB {
	username := v.GetString("database.username")
	password := v.GetString("database.password")
	host := v.GetString("database.host")
	port := v.GetInt("database.port")
	database := v.GetString("database.name")
	sslMode := v.GetString("database.sslmode")
	timeZone := v.GetString("database.time_zone")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, username, password, database, port, sslMode, timeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.ErrorLogger(err)

	return db
}
`

var MongoDBConfig = `package config

import (
	"{{.Package}}/internals/helpers"
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(v *viper.Viper) *mongo.Database {
	ctx := context.Background()

	url := v.GetString("database.url")
	clientOptions := options.Client()
	clientOptions.ApplyURI(url)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		helpers.ErrorLogger(err)

	}

	err = client.Connect(ctx)
	if err != nil {
		helpers.ErrorLogger(err)
	}

	return client.Database(v.GetString("database.name"))
}
`

func NewDatabaseOptions(driver string) *AppConfiguration {
	switch driver {
	case "mongodb":
		return &AppConfiguration{
			Dependencies:      MongoDBDependencies,
			DatabaseDriver:    MongoDBConfig,
			DockerCompose:     ComposeCodeGenerate(driver),
			MainApp:           MongoMain,
			ConfigurationFile: ConfigurationFileGenerate(driver),
		}
	case "mysql":
		return &AppConfiguration{
			Dependencies:      MysqlDependencies,
			DatabaseDriver:    MysqlDBConfig,
			DockerCompose:     ComposeCodeGenerate(driver),
			MainApp:           MysqlMain,
			ConfigurationFile: ConfigurationFileGenerate(driver),
		}
	case "postgres":
		return &AppConfiguration{
			Dependencies:      PostgresDependencies,
			DatabaseDriver:    PostgresDBConfig,
			DockerCompose:     ComposeCodeGenerate(driver),
			MainApp:           PostgresMain,
			ConfigurationFile: ConfigurationFileGenerate(driver),
		}
	default:
		return &AppConfiguration{
			Dependencies:      MysqlDependencies,
			DatabaseDriver:    MysqlDBConfig,
			DockerCompose:     ComposeCodeGenerate(driver),
			MainApp:           MysqlMain,
			ConfigurationFile: ConfigurationFileGenerate(driver),
		}
	}
}
