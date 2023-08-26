package code

var LoadConfigCode = `package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func InitDatabase() *DB {
	return &DB{
		Driver:   GetString("database.driver"),
		Host:     GetString("database.host"),
		Username: GetString("database.username"),
		Password: GetString("database.password"),
		Port:     GetInt("database.port"),
		DBName:   GetString("database.name"),
		SSlmode:  GetString("database.sslmode"),
	}
}

type Configuration struct {
	file string
}

func NewConfiguration(filename string) *Configuration {
	return &Configuration{file: filename}
}

// Initialize initializes the Configuration object by reading configuration
// values from a file and mapping them to the corresponding fields in the
// Configuration struct. It returns an error if the file is not found or if
// there is an error reading or unmarshaling the config values.
func (c *Configuration) Initialize() error {

	configName := filepath.Base(c.file)
	configExtension := filepath.Ext(c.file)
	configExtension = strings.TrimPrefix(configExtension, ".")

	viper.SetConfigName(configName)
	viper.SetConfigType(configExtension)
	viper.AddConfigPath(filepath.Dir(c.file))

	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return err
		}
		return err
	}

	var dbStructure DB
	err = viper.Unmarshal(&dbStructure, func(c *mapstructure.DecoderConfig) {
		c.TagName = "mapstructure"
	})
	if err != nil {
		return err
	}

	return nil
}

// checkKey checks if the given configuration key exists in the viper configuration.
func checkKey(key string) {
	if !viper.IsSet(key) {
		log.Fatalf("Configuration key %s not found; aborting \n", key)
		os.Exit(1)
	}
}

// GetString returns the string value associated with the given key.
func GetString(key string) string {
	checkKey(key)
	return viper.GetString(key)
}

// GetInt retrieves the integer value associated with the given key from viper configuration.
func GetInt(key string) int {
	checkKey(key)
	return viper.GetInt(key)
}

// GetBool returns the boolean value of the key in the viper configuration file.
func GetBool(key string) bool {
	checkKey(key)
	return viper.GetBool(key)
}
`
var ConfigCodeEnv = `package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	if err != nil {
		panic(err)
	}
	return &configImpl{}
}
`
