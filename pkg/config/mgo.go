package config

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

// Mgo hold our Mongodb session
var Mgo *mgo.Session

func init() {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost/chat"
	}

	session, err := mgo.Dial(mongoURI)
	if err != nil {
		panic(err)
	}
	Mgo = session

	// Set safe to receive errors back
	Mgo.SetSafe(&mgo.Safe{WMode: "majority"})

	// Ensure some Index
	err = Mgo.DB("").C("messages").EnsureIndexKey("channel", "timestamp")
	if err != nil {
		log.Print("Error when ensure index: ", err)
	}
}
