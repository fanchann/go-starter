package code

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/fanchann/go-starter/helpers"
)

type service struct {
	Image         string            `yaml:"image"`
	ContainerName string            `yaml:"container_name"`
	Ports         []string          `yaml:"ports"`
	Environment   map[string]string `yaml:"environment"`
	Volumes       []string          `yaml:"volumes"`
}

type volume struct {
	Name string `yaml:"name"`
}

type volumeSelection interface {
	isVolumeSelection()
}
type postgresVolume struct {
	PostgresData volume `yaml:"postgres_volume"`
}

func (p *postgresVolume) isVolumeSelection() {}

type mysqlVolume struct {
	MysqlData volume `yaml:"mysql_volume"`
}

func (m *mysqlVolume) isVolumeSelection() {}

type dockerCompose struct {
	Version  string             `yaml:"version"`
	Services map[string]service `yaml:"services"`
	Volumes  volumeSelection    `yaml:"volumes"`
}

func GenerateDockerCompose(driver, dbUsername, dbPassword, dbName, dbPort string) string {
	var composeFormat dockerCompose

	switch driver {
	case "postgres":
		posEnv := map[string]string{
			"POSTGRES_DB":       dbName,
			"POSTGRES_USER":     dbUsername,
			"POSTGRES_PASSWORD": dbPassword,
		}
		composeFormat = createDockerCompose("3.8", "bitnami/postgresql:latest", "postgres_database", dbName, dbPort, dbPassword, posEnv, "/var/lib/postgresql/data", &postgresVolume{PostgresData: volume{Name: "postgres_volume"}})
		composeByte, err := yaml.Marshal(composeFormat)
		helpers.ErrorWithLog(err)
		return string(composeByte)
	case "mysql":
		mysqlEnv := map[string]string{
			"MYSQL_ROOT_PASSWORD": dbPassword,
		}
		composeFormat = createDockerCompose("3.8", "mysql:latest", "mysql_database", dbName, dbPort, dbPassword, mysqlEnv, "/var/lib/mysql", &mysqlVolume{MysqlData: volume{Name: "mysql_volume"}})
		composeFormat.Volumes.isVolumeSelection()
		composeByte, err := yaml.Marshal(composeFormat)
		helpers.ErrorWithLog(err)
		return string(composeByte)
	default:
		return ""
	}
}

func createDockerCompose(version, imageName, containerName, dbName, dbPort, dbPassword string, environment map[string]string, savedData string, volSelect volumeSelection) dockerCompose {
	return dockerCompose{
		Version: version,
		Services: map[string]service{
			containerName: {
				Image:         imageName,
				ContainerName: containerName,
				Ports:         []string{fmt.Sprintf("%v:%v", dbPort, dbPort)},
				Environment:   environment,
				Volumes:       []string{fmt.Sprintf("%s:%s", volSelect, savedData)},
			},
		},
		Volumes: volSelect,
	}
}
