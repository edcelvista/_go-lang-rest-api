package model

type Pong struct {
	Pong    string
	Headers interface{}
	Message map[string]string
}

type Ping struct {
	Ping    string
	Message map[string]string
}

type Echo struct {
	EchoHeaders interface{}
	EchoData    interface{}
}
