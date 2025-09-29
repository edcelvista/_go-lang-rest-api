package router

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var e bool
var ee error

var PORT = ":8443"
var SSL_CERT = "./tls.crt"
var SSL_KEY = "./tls.key"
var ROOT_DIR string

func init() {
	rootDir := fmt.Sprintf("%v/.env", os.Getenv("ROOT_DIR"))
	if err := godotenv.Load(rootDir); err != nil {
		log.Printf("Error ROOT_DIR defaulting to getenv")

		PORT = os.Getenv("PORT")
		if PORT == "" {
			log.Printf("Error router PORT")
		}

		SSL_CERT = os.Getenv("SSL_CERT")
		if SSL_CERT == "" {
			log.Fatalf("Error router SSL_CERT")
		}

		SSL_KEY = os.Getenv("SSL_KEY")
		if SSL_KEY == "" {
			log.Fatalf("Error router SSL_KEY")
		}
	} else {
		PORT, e = os.LookupEnv("PORT")
		if !e {
			log.Fatalf("Error router PORT")
		}

		SSL_CERT, e = os.LookupEnv("SSL_CERT")
		if !e {
			log.Fatalf("Error router SSL_CERT")
		}

		SSL_KEY, e = os.LookupEnv("SSL_KEY")
		if !e {
			log.Fatalf("Error router SSL_KEY")
		}
	}
}
