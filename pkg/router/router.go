package router

import (
	"crypto/tls"
	"log"
	"net/http"

	lib "pkg/lib"
	routes "pkg/routes"

	"github.com/gorilla/mux"
)

func registerRouters(router *routes.Router) {
	router.PingRoutes()
	router.CrudRoutes()
}

func bindRouters(muxRouter *mux.Router) {
	// Bind the router to the muxRouter
	http.Handle("/", muxRouter)
}

func Run() {
	tlsCertsAndKey := lib.CertsAndKeys{
		Cert: SSL_CERT,
		Key:  SSL_KEY,
	}
	tlsCertsAndKey.CheckCerts()

	// Init router
	muxRouter := mux.NewRouter()
	router := routes.Router{
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
		Addr:      PORT,
		Handler:   muxRouter,
		TLSConfig: cfg,
	}

	log.Printf("💡 ⚡️ Mux API Running 📦 %s with 🔑 %v %v \n", PORT, SSL_CERT, SSL_KEY)
	// err = http.ListenAndServe(port, muxRouter)

	// TLS config
	if err := server.ListenAndServeTLS(SSL_CERT, SSL_KEY); err != nil {
		log.Fatalf("‼️ Failed to start router %s with %v %v", err, SSL_CERT, SSL_KEY)
	}
}
