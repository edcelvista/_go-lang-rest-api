package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var e bool
var ee error

var DB_URI string

func init() {
	rootDir := fmt.Sprintf("%v/.env", os.Getenv("ROOT_DIR"))
	if err := godotenv.Load(rootDir); err != nil {
		log.Printf("Error ROOT_DIR defaulting to getenv")

		DB_URI = os.Getenv("DB_URI")
		if DB_URI == "" {
			log.Fatalf("Error router DB_URI")
		}
	} else {
		DB_URI, e = os.LookupEnv("DB_URI")
		if !e {
			log.Fatalf("Error router DB_URI")
		}
	}

	dBCreateClient()
}
