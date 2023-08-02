## Go-Starter
![App Screenshot](/assets/GO-STARTER.png)

## What is Go-Starter?
This tool helps you to do the tedious work of setting configuration and creating layers for the REST API

## Running Go-Starter
```sh
go-starter
```

![go-starter](/assets/go-starter.gif)

## About Go-Starter
This tool adopt [gorm](https://gorm.io/)\
Database support:
- mysql
- postgres
## Installation
```sh
go install github.com/fanchann/go-starter@latest
```
## Structure
Structure reference
- [Golang standards project layout](https://github.com/golang-standards/project-layout/)
```sh
├── app
│   ├── controllers
│   ├── domain
│   │   ├── models
│   │   └── types
│   ├── middlewares
│   ├── repositories
│   ├── routers
│   └── services
├── cmd
│   └── main.go
├── config
│   └── toml_reader.go // -> filename follows your configuration
├── config.toml
├── go.mod
├── go.sum
├── lib
│   └── mysql.go // -> filename follows your configuration 
└── utils
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

