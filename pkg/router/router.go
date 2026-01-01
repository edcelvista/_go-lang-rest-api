package router

import (
	"compress/gzip"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	lib "pkg/lib"
	routes "pkg/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

// Middleware  Rate limit
type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
)

func shouldSkip(path string) bool {
	return strings.HasPrefix(path, "/healthz") ||
		strings.HasPrefix(path, "/auth/sso-") ||
		strings.HasPrefix(path, "/public")
}

func getLimiter(token string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	c, exists := clients[token]
	if !exists {
		c = &client{
			limiter:  rate.NewLimiter(5, 10), // 5 req/sec, burst 10
			lastSeen: time.Now(),
		}
		clients[token] = c
	}

	c.lastSeen = time.Now()
	return c.limiter
}

func rateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if shouldSkip(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		token := extractToken(r)
		if token == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		if !getLimiter(token).Allow() {
			w.Header().Set("Retry-After", "1")
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	if strings.HasPrefix(auth, "ey") {
		return auth
	}
	return ""
}

func cleanupClients() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for token, c := range clients {
			if time.Since(c.lastSeen) > 10*time.Minute {
				delete(clients, token)
			}
		}
		mu.Unlock()
	}
}

// Middleware Compression
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func gzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		next.ServeHTTP(gzw, r)
	})
}

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
	muxRouter.Use(gzipMiddleware) // attach middleware
	muxRouter.Use(rateLimit)      // rate limiting
	router := routes.Router{
		R: muxRouter,
	}

	// register routers
	registerRouters(&router)

	// bind routers
	bindRouters(muxRouter)

	// 🔥 Wrap the router with CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"https://localhost:5173",
		}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(muxRouter)

	// TLS config
	cfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// TLS config
	server := &http.Server{
		Addr:      PORT,
		Handler:   corsHandler,
		TLSConfig: cfg,
	}

	log.Printf("💡 ⚡️ Mux API Running 📦 %s with 🔑 %v %v \n", PORT, SSL_CERT, SSL_KEY)
	// err = http.ListenAndServe(port, muxRouter)

	go cleanupClients()

	// TLS config
	if err := server.ListenAndServeTLS(SSL_CERT, SSL_KEY); err != nil {
		log.Fatalf("‼️ Failed to start router %s with %v %v", err, SSL_CERT, SSL_KEY)
	}
}
