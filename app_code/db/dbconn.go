package db

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// get mongodb db
func GetMGODB() *mgo.Database {
	//session, err := mgo.Dial("127.0.0.1:27017")
	session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		fmt.Println(time.Now())
		fmt.Println("CommonDBBase > db error!!!")
		panic(err)
	}
	// defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("goagodb")
	return db
}

func GetMGODB_DT(dtname string) *mgo.Collection {

	return GetMGODB().C(dtname)
}
