/*
* Package allowing you to setup a jet DB connection based on the app.info config file.
*  The following config settings are expected:
*  - db.adapter (mysql or postgres)
*  - db.username
*  - db.password
*
* Optional ones:
*  - db.hostname (defaults to localhost)
*  - db.port (defaults to 3306 for MySQL)
*
* Usage:
* In app/controllers/init.go
*
* Import the DB driver you need, Postgres or MySQL:
*
* _ "github.com/bmizerany/pq"
* _ "github.com/go-sql-driver/mysql"
*
* import this module:
*
*  "github.com/mattetti/revel_addons/jetDb"
*
* And in the init function, add the following:
*
*   revel.OnAppStart(func() { jetDb.Setup() })
*
* Then from your code, you can access the db via jetDb.DB
*
 */

package jetDb

import (
	"fmt"
	"github.com/eaigner/jet"
	"github.com/robfig/revel"
)

var (
	DB           jet.Db
	Adapter      string
	databaseName string
	username     string
	password     string
	port         string
	hostname     string
	dataSource   string
)

func Setup() (err error) {

	configRequired := func(key string) string {
		value, found := revel.Config.String(key)
		if !found {
			revel.ERROR.Fatal(fmt.Sprintf("Configuration for %s missing in app.conf.", key))
		}
		return value
	}

	// Read configuration.
	Adapter = configRequired("db.adapter")
	databaseName = configRequired("db.database_name")
	username = configRequired("db.username")
	password = configRequired("db.password")
	hostname = revel.Config.StringDefault("db.hostname", "localhost")

	if Adapter == "postgres" {
		dataSource = "user=" + username + " dbname=" + databaseName + " sslmode=disable"
	} else if Adapter == "mysql" {
		port = revel.Config.StringDefault("db.port", "3306")
		dataSource = username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + databaseName + "?charset=utf8"
	}

	DB, err = jet.Open(Adapter, dataSource)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	return err
}
