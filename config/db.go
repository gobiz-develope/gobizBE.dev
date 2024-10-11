package config

import (
	"gobizdevelop/helper/atdb"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

var mongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "gobizdev",
}

var Mongoconn, ErrorMongoconn = atdb.MongoConnect(mongoinfo)
