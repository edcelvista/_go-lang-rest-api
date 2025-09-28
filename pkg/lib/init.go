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
		log.Fatalf("Error ROOT_DIR")
	}

	DB_URI, e = os.LookupEnv("DB_URI")
	if !e {
		log.Fatalf("Error router DB_URI")
	}

	dBCreateClient()
}
