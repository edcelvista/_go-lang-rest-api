package lib

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

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

func DBInit() (client any) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("‼️ Error loading .env file")
	} else {
		connString, _ := os.LookupEnv("DB_URI")
		dbConfig := DBOpts{
			ConnectionString: connString,
			ContextDeadline:  20,
		}

		ctx, err := dbConfig.Init()
		log.Println("💡 Database Initialized 🗄️ ...")
		if err != nil {
			log.Fatalf("‼️ Error initializing database context: %v", err)
		}

		client, err = dbConfig.Connect(ctx)
		if err != nil {
			log.Fatalf("‼️ Error Connecting database context: %v", err)
		}
	}

	return client
}
