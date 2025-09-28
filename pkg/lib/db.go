package lib

import (
	"log"
	"time"
)

var dbClient any

func dBCreateClient() {
	dbConfig := DBOpts{
		ConnectionString: DB_URI,
		ContextDeadline:  20,
	}

	ctx, err := dbConfig.Init()
	log.Println("🗄️ Database Initialized")
	if err != nil {
		log.Fatalf("‼️ Error initializing database context: %v", err)
	}

	dbClient, err = dbConfig.Connect(ctx)
	if err != nil {
		log.Fatalf("‼️ Error Connecting database context: %v", err)
	}
	log.Println("🛜 Database Connected")
}

func GetDBClient() any {
	return dbClient
}

type DBOpts struct {
	ConnectionString string
	ContextDeadline  time.Duration
}
type DB interface {
	Execute() (any, error)
}

func DBProcess(db DB) (res any, err error) {
	data, err := db.Execute()
	return data, err
}
