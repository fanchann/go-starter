## Go-Starter
![App Screenshot](/assets/GO-STARTER.png)

## What is Go-Starter?
This tool helps you to do the tedious work of setting configuration and creating layers for the REST API

## go-starter command
```sh
 go-starter help

	┌─┐┌─┐   ┌─┐┌┬┐┌─┐┬─┐┌┬┐┌─┐┬─┐
	│ ┬│ │───└─┐ │ ├─┤├┬┘ │ ├┤ ├┬┘
	└─┘└─┘   └─┘ ┴ ┴ ┴┴└─ ┴ └─┘┴└─
	 
Usage: go-starter [OPTIONS]

Options:
  -f string
    	starter configuration file (default "starter.yaml")
  help
    	show help message

Examples:
  go-starter new      Generate configuration file
  go-starter          Run the application
  go-starter -f=configuration.yaml[default: starter.yaml] Specify a custom configuration file
```
## generate starter configuration
```sh
go-starter new

	┌─┐┌─┐   ┌─┐┌┬┐┌─┐┬─┐┌┬┐┌─┐┬─┐
	│ ┬│ │───└─┐ │ ├─┤├┬┘ │ ├┤ ├┬┘
	└─┘└─┘   └─┘ ┴ ┴ ┴┴└─ ┴ └─┘┴└─
	 
2024/01/18 16:46:54 Success generate starter.yaml
```

`starter.yaml` is a go-starter configuration:
```yaml
version: "1"
package: your-name-app
database: database
```
database support :
- mysql
- mongodb
- postgres

## generate structure project
```sh
 go-starter

	┌─┐┌─┐   ┌─┐┌┬┐┌─┐┬─┐┌┬┐┌─┐┬─┐
	│ ┬│ │───└─┐ │ ├─┤├┬┘ │ ├┤ ├┬┘
	└─┘└─┘   └─┘ ┴ ┴ ┴┴└─ ┴ └─┘┴└─
	 
layer 'cmd' created successfully.
layer 'api' created successfully.
layer 'db/migrations' created successfully.
layer 'internals/config' created successfully.
layer 'internals/delivery/http' created successfully.
layer 'internals/delivery/messaging' created successfully.
layer 'internals/gateway' created successfully.
layer 'internals/models' created successfully.
layer 'internals/repository' created successfully.
layer 'internals/usecase' created successfully.
layer 'internals/helpers' created successfully.
layer 'tests' created successfully.
```

## Structure
Structure reference
- [Golang standards project layout](https://github.com/golang-standards/project-layout/)
```sh
├── api
├── cmd
│   └── main.go
├── config.dev.yaml
├── db
│   └── migrations
├── docker-compose.yaml
├── go.mod
├── internals
│   ├── config
│   │   ├── mysql.go ## filename following the database
│   │   └── viper.go
│   ├── delivery
│   │   ├── http
│   │   └── messaging
│   ├── gateway
│   ├── helpers
│   │   └── error.go
│   ├── models
│   ├── repository
│   └── usecase
├── starter.yaml
└── tests

```

## About go-starter
This tool adopt \
[gorm](https://gorm.io/)\
[mongo client](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo)\
[viper](https://pkg.go.dev/github.com/dvln/viper)

Database support:
| Database              | Support |
| :---------------- | :------: |
| MySQL        |   ✅   |
| PostgreSQL           |   ✅   |
| MongoDB    |  ✅   |

## Installation
```sh
go install github.com/fanchann/go-starter@latest
```

see older version here [version](https://github.com/fanchann/go-starter/tags)

## Running Your App
After generating the template, enter the folder that has been generated by `go-starter`, and then add dependencies using the following command.
```sh
go mod tidy
```
This command will download and install the required dependencies for your project\
After success install dependencies, run app with command :
```sh
go run cmd/main.go
```
## Authors

- [@fanchann](https://github.com/fanchann)

