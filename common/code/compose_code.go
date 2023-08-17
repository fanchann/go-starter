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
		composeFormat = dockerCompose{
			Version: "3.8",
			Services: map[string]service{
				"postgres_database": {
					Image:         "bitnami/postgresql:latest",
					ContainerName: "postgres_database",
					Ports:         []string{fmt.Sprintf("%v:%v", dbPort, dbPort)},
					Environment: map[string]string{
						"POSTGRES_DB":       dbName,
						"POSTGRES_USER":     dbUsername,
						"POSTGRES_PASSWORD": dbPassword,
					},
					Volumes: []string{"postgres_volume:/var/lib/postgresql/data"},
				},
			},
			Volumes: &postgresVolume{
				PostgresData: volume{Name: "postgres_volume"},
			},
		}
		composeByte, err := yaml.Marshal(composeFormat)
		helpers.ErrorWithLog(err)
		return string(composeByte)
	case "mysql":
		composeFormat = dockerCompose{
			Version: "3.8",
			Services: map[string]service{
				"mysql_database": {
					Image:         "mysql:latest",
					ContainerName: "mysql_database",
					Ports:         []string{fmt.Sprintf("%v:%v", dbPort, dbPort)},
					Environment: map[string]string{
						"MYSQL_ROOT_PASSWORD": dbPassword,
					},
					Volumes: []string{"mysql_volume:/var/lib/mysql"},
				},
			},
			Volumes: &mysqlVolume{
				MysqlData: volume{Name: "mysql_volume"},
			},
		}
		composeByte, err := yaml.Marshal(composeFormat)
		helpers.ErrorWithLog(err)
		return string(composeByte)
	default:
		return ""
	}
}
