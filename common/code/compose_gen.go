package code

var (
	mysqlCompose = map[string]interface{}{
		"version": "3.8",
		"services": map[string]interface{}{
			"mysql_db": map[string]interface{}{
				"image":          "mysql:latest",
				"container_name": "mysql_db",
				"restart":        "on-failure",
				"ports": []interface{}{
					"3306:3306",
				},
				"environment": []string{
					"MYSQL_ROOT_PASSWORD=root",
				},
				"command": []string{
					"mysqld",
					"--character-set-server=utf8mb4",
					"--collation-server=utf8mb4_unicode_ci",
					"--default-time-zone=+09:00",
				},
				"networks": []string{
					"mysql_network",
				},
			},
		},
		"networks": map[string]interface{}{
			"mysql_network": map[string]interface{}{
				"driver": "bridge",
				"name":   "mysql_network",
			},
		},
	}
	postgresCompose = map[string]interface{}{
		"version": "3.8",
		"services": map[string]interface{}{
			"postgres": map[string]interface{}{
				"image":   "postgres:latest",
				"restart": "always",
				"environment": []string{
					"POSTGRES_USER=postgres",
					"POSTGRES_PASSWORD=postgres",
					"POSTGRES_DB=test",
				},
				"logging": map[string]interface{}{
					"options": map[string]interface{}{
						"max-size": "10m",
						"max-file": "3",
					},
				},
				"ports": []string{"5432:5432"},
			},
		},
	}
	mongoCompose = map[string]interface{}{
		"version": "3.8",
		"services": map[string]interface{}{
			"mongodb": map[string]interface{}{
				"image":          "mongo:4.2.2-bionic",
				"container_name": "mongodb",
				"ports":          []string{"27017:27017"},
				"environment": []string{
					"MONGO_INITDB_DATABASE=test",
					"MONGO_INITDB_ROOT_USERNAME=admin",
					"MONGO_INITDB_ROOT_PASSWORD=admin",
				},
				"networks": []string{"mongo_net"},
			},
		},
		"networks": map[string]interface{}{
			"mongo_net": map[string]interface{}{
				"driver": "bridge",
			},
		},
	}
)

func ComposeCodeGenerate(driver string) map[string]interface{} {
	switch driver {
	case "mysql":
		return mysqlCompose
	case "postgres":
		return postgresCompose
	case "mongodb":
		return mongoCompose
	default:
		return mysqlCompose
	}
}
