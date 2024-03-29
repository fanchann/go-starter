package code

var (
	MysqlMain = `package main

	import (
		"flag"
		"fmt"

		"{{.Package}}/internals/config"
		"{{.Package}}/internals/helpers"
	)
	
	var stage *string
	
	func init() {
		stage = flag.String("s", "dev", "❌ please input stage !")
		flag.Parse()
	}
	
	func main() {
		v := config.NewViper(*stage)
		db := config.NewMysqlConnection(v)
	
		d, err := db.DB()
		helpers.ErrorLogger(err)
	
		err2 := d.Ping()
		helpers.ErrorLogger(err2)
	
		fmt.Println("✅ connected")
	}
	`

	PostgresMain = `package main

	import (
		"{{.Package}}/internals/config"
		"{{.Package}}/internals/helpers"
		"flag"
		"fmt"
	)
	
	var stage *string
	
	func init() {
		stage = flag.String("s", "dev", "❌ please input stage !")
		flag.Parse()
	}
	
	func main() {
		v := config.NewViper(*stage)
		db := config.NewPostgresConnection(v)
	
		d, err := db.DB()
		helpers.ErrorLogger(err)
	
		err2 := d.Ping()
		helpers.ErrorLogger(err2)
	
		fmt.Println("✅ connected")
	}
	`

	MongoMain = `package main

	import (
		"{{.Package}}/internals/config"
		"{{.Package}}/internals/helpers"
		"context"
		"flag"
		"fmt"
		"time"
	
		"go.mongodb.org/mongo-driver/mongo/readpref"
	)
	
	var stage *string
	
	func init() {
		stage = flag.String("s", "dev", "❌ please input stage !")
		flag.Parse()
	}
	
	func main() {
		v := config.NewViper(*stage)
		db := config.NewMongoConnection(v)


		ctx, errTimeout := context.WithTimeout(context.Background(), time.Second*5)
		defer errTimeout()
	
		err := db.Client().Ping(ctx, &readpref.ReadPref{})
		helpers.ErrorLogger(err)
	
		fmt.Println("✅ connected")
	}
	`
)
