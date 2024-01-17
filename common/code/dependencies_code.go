package code

var (
	MongoDBDependencies = `module {{.Package}}

	go {{.GoVersion}}
	
	require (
		github.com/spf13/viper v1.18.2
		go.mongodb.org/mongo-driver v1.13.1
	)
	
	require (
		github.com/fsnotify/fsnotify v1.7.0 // indirect
		github.com/golang/snappy v0.0.1 // indirect
		github.com/hashicorp/hcl v1.0.0 // indirect
		github.com/klauspost/compress v1.17.0 // indirect
		github.com/magiconair/properties v1.8.7 // indirect
		github.com/mitchellh/mapstructure v1.5.0 // indirect
		github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
		github.com/pelletier/go-toml/v2 v2.1.0 // indirect
		github.com/sagikazarmark/locafero v0.4.0 // indirect
		github.com/sagikazarmark/slog-shim v0.1.0 // indirect
		github.com/sourcegraph/conc v0.3.0 // indirect
		github.com/spf13/afero v1.11.0 // indirect
		github.com/spf13/cast v1.6.0 // indirect
		github.com/spf13/pflag v1.0.5 // indirect
		github.com/subosito/gotenv v1.6.0 // indirect
		github.com/xdg-go/pbkdf2 v1.0.0 // indirect
		github.com/xdg-go/scram v1.1.2 // indirect
		github.com/xdg-go/stringprep v1.0.4 // indirect
		github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
		go.uber.org/atomic v1.9.0 // indirect
		go.uber.org/multierr v1.9.0 // indirect
		golang.org/x/crypto v0.16.0 // indirect
		golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
		golang.org/x/sync v0.5.0 // indirect
		golang.org/x/sys v0.15.0 // indirect
		golang.org/x/text v0.14.0 // indirect
		gopkg.in/ini.v1 v1.67.0 // indirect
		gopkg.in/yaml.v3 v3.0.1 // indirect
	)
	`
	PostgresDependencies = `module {{.Package}}

go {{.GoVersion}}

require (
	github.com/spf13/viper v1.18.2
	gorm.io/gorm v1.25.5
)

require (
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/crypto v0.16.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/postgres v1.5.4 // indirect
)
`

	MysqlDependencies = `module {{.Package}}

go {{.GoVersion}}

require (
	github.com/spf13/viper v1.18.2
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.5
)

require (
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
`
)
