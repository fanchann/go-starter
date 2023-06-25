## Go-Starter
![App Screenshot](/assets/GO-STARTER.png)

## What is Go-Starter?
This tool helps you to do the tedious work of setting configuration and creating layers for the REST API

## Running Go-Starter
```sh
go-starter -pkg=awesomeProject  -config={configurationFile}
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
json example:
```json
{
  "database": {
    "driver": "your_db_driver", // your driver
    "host": "url", // <- database url
    "port": 3306, // <- port 
    "username": "root", // <- your username
    "password": "root", // <- your password
    "name": "your_app",// <- your database to connect
    "sslmode": "disabled", // <- if you use postgres, recomended to enabled
    /**
    database pooling
    **/
    "max_idle_conn": 30,
    "max_open_conn": 20,
    "life_time_conn": 40
  }
}
```
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
│   │   └── db.go
│   ├── middleware
│   └── types
│       └── dsn.go
├── config.toml
├── docs
├── go.mod
├── go.sum
├── models
├── repository
```
## Running Your App
```sh
go run cmd/main.go -c config.toml
```
## Authors

- [@fanchann](https://github.com/fanchann)

