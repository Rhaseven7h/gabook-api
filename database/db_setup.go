package gabookApiDatabase

import (
	"log"

	"gopkg.in/mgo.v2"
)

const (
	// DatabaseName is the application's database name, and that's all!
	DatabaseName = "gabook"
)

// GaBookDB holds references
type GaBookDB struct {
	Session  *mgo.Session
	Database *mgo.Database
}

func newGaBookDB() *GaBookDB {
	gbdb := new(GaBookDB)
	var err error
	gbdb.Session, err = mgo.Dial("localhost")
	if err != nil {
		log.Println("Error occurred while opening database")
		panic(err)
	}
	gbdb.Session.SetMode(mgo.Monotonic, true)
	gbdb.Database = gbdb.Session.DB(DatabaseName)
	return gbdb
}

// Close method closes database session
func (gbdb *GaBookDB) Close() {
	gbdb.Session.Close()
}

var gabookdb *GaBookDB

// GetGaBookDB is the singleton
func GetGaBookDB() *GaBookDB {
	if gabookdb == nil {
		gabookdb = newGaBookDB()
	}
	return gabookdb
}
