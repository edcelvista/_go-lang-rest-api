package lib

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"path/filepath"
	model "pkg/model"
	"reflect"
	"strings"
	"time"
)

type CertsAndKeys struct {
	Cert string
	Key  string
}

var debug = false

func init() {
	if strings.ToLower(os.Getenv("IS_DEBUG")) == "true" {
		debug = true
	}

	log.Printf("💡 Debug enabled: %v", debug)
	// •	%v → Print the values
	// •	%+v → Print field names and values
	// •	%#v → Print Go syntax (main.Person{Name:"Alice", Age:30})
}

func Debug(v string) {
	if debug {
		log.Printf("🚨 [DEBUG] %v", v)
	}
}

func DefaultIfEmpty(s, def string) string {
	if s == "" {
		return def
	}
	return s
}

func (m *CertsAndKeys) CheckCerts() {
	// Load certificate
	certPath := m.Cert
	certData, err := os.ReadFile(certPath)
	if err != nil {
		log.Fatalf("‼️ Error reading certificate: %v %v", err, m.Cert)
	}

	// Decode the PEM data
	block, _ := pem.Decode(certData)
	if block == nil {
		log.Fatalf("Failed to decode PEM block")
	}

	// Parse certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatalf("‼️ Error parsing certificate: %v %v", err, m.Cert)
	}

	// Check if the certificate is expired
	currentTime := time.Now()
	if currentTime.After(cert.NotAfter) {
		log.Println("‼️ Certificate has expired.")
	} else {
		log.Println("💡 Certificate is valid.")
	}

	// Check the NotBefore field (if the certificate is not yet valid)
	if currentTime.Before(cert.NotBefore) {
		log.Println("‼️ Certificate is not yet valid.")
	} else {
		log.Println("💡 Certificate is within the valid period.")
	}

	// Optionally, you could also validate other parts, like the issuer and subject
	log.Println("🔑 Issuer:", cert.Issuer)
	log.Println("🔑 Subject:", cert.Subject)
	log.Println("🔑 Not After:", cert.NotAfter)
	log.Println("🔑 Not Before:", cert.NotBefore)
}

func CountInterfaceValues[T model.Number](v interface{}) (count T) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return T(rv.Len())
	default:
		return 1 // For single values
	}
}

func GetAppRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Abs(wd)
}
