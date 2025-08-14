package router

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	Lib "pkg/lib"
	Routes "pkg/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func environ() (cert string, key string, port string) {
	// Load .env file
	err := godotenv.Load()
	port = ":8443"

	// TLS config
	cert = "./tls.crt"
	key = "./tls.key"

	if err != nil {
		log.Println("⚠️ Error loading .env file")
		// Get environment variables
		if os.Getenv("PORT") != "" {
			port = os.Getenv("PORT")
		}
		if os.Getenv("SSL_CERT") != "" {
			cert = os.Getenv("SSL_CERT")
		}
		if os.Getenv("SSL_KEY") != "" {
			key = os.Getenv("SSL_KEY")
		}
		log.Printf("⚠️ Defaulting to Port %v", port)
	} else {
		port, _ = os.LookupEnv("PORT")
		cert, _ = os.LookupEnv("SSL_CERT")
		key, _ = os.LookupEnv("SSL_KEY")
		log.Println("💡 Found .env")
	}

	tlsCertsAndKey := Lib.CertsAndKeys{
		Cert: cert,
		Key:  key,
	}

	tlsCertsAndKey.CheckCerts()
	Lib.DebuggerInit()

	return cert, key, port
}

func registerRouters(router *Routes.Router) {
	router.PingRoutes()
	router.CrudRoutes()
}

func bindRouters(muxRouter *mux.Router) {
	// Bind the router to the muxRouter
	http.Handle("/", muxRouter)

}

func Run() {
	cert, key, port := environ()

	// Init router
	muxRouter := mux.NewRouter()
	router := Routes.Router{
		R: muxRouter,
	}

	// register routers
	registerRouters(&router)

	// bind routers
	bindRouters(muxRouter)

	// TLS config
	cfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// TLS config
	server := &http.Server{
		Addr:      port,
		Handler:   muxRouter,
		TLSConfig: cfg,
	}

	log.Printf("💡 ⚡️ Mux API Running 📦 %s with 🔑 %v %v \n", port, cert, key)
	// err = http.ListenAndServe(port, muxRouter)

	// TLS config
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalf("‼️ Failed to start router %s with %v %v", err, cert, key)
	}
}
