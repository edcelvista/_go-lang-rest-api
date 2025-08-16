# Requires Environment Variable
.env
```
PORT=":8443"
SSL_CERT="../tls.crt"
SSL_KEY="../tls.key"
IS_DEBUG="true"

DB_URI="mongodb+srv://<username>:<password>@<mongodb-url>.mongodb.net/?retryWrites=true&w=majority&appName=<app-name>"
```

## Start
```
./_go-run.sh
```

## PPOF
```
go tool pprof -http=localhost:6061 "http://localhost:6060/debug/pprof/profile?seconds=30"
```
Requires: `raphviz`
```
brew install graphviz
```