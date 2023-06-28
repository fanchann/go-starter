## Go-Starter
![App Screenshot](/assets/GO-STARTER.png)

## What is Go-Starter?
This tool helps you to do the tedious work of setting configuration and creating layers for the REST API

## Running Go-Starter
```sh
go-starter -pkg=awesomeProject  -config={configurationFile}
```
or 
```sh
go-starter -help
```
configurationFile: 
- json
- yaml
- toml [default]

![go-starter](/assets/go-starter.gif)

## About Go-Starter
This tool adopt [gorm](https://gorm.io/)\
Database support:
- mysql
- postgres

set database driver in configuration file\
`config.example.json` :
```json
{
  "database": {
    "driver": "db_driver", // database driver [mysql / postgres]
    "host": "url", // <- database url
    "port": 3306, // <- port 
    "username": "root", // <- your username
    "password": "root", // <- your password
    "name": "db_name",// <- your database to connect
    "sslmode": "disable",
    /**
    database pooling
    **/
    "max_idle_conn": 30,
    "max_open_conn": 20,
    "life_time_conn": 40
  }
}
```
## Database Driver in Go-Starter
| `Driver`      | `Support` |
| ----------- | ----------- |
| Mysql       | ✅      |
| Postgres    | ✅      |

## Installation
```sh
go install github.com/fanchann/go-starter@latest
```
## Structure
Structure reference
- [Golang Project Structure](https://github.com/Mindinventory/Golang-Project-Structure)
```sh
├── cmd
│   └── main.go
├── common
│   ├── config
│   │   └── load.go
│   ├── database
│   │   └── database.go
│   ├── helpers
│   │   ├── db.go
│   │   └── migrate.go
│   ├── middleware
│   └── types
│       └── dsn.go
├── config.json
├── docs
├── go.mod
├── models
└── repository
```
## Running Your App
After generating template, install dependencies using the following command
```sh
go mod tidy
```
This command will download and install the required dependencies for your project\
After success install dependencies, run app with command :
```sh
go run cmd/main.go -c config.toml
```
## Authors

- [@fanchann](https://github.com/fanchann)

