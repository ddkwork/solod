package main

type HttpStatus int

const (
	StatusOK       HttpStatus = 200
	StatusNotFound HttpStatus = 404
	StatusError    HttpStatus = 500
)

// Methods are only supported for struct types.
// func (s HttpStatus) String() string {
// 	switch s {
// 	case StatusOK:
// 		return "OK"
// 	case StatusNotFound:
// 		return "Not Found"
// 	case StatusError:
// 		return "Error"
// 	default:
// 		return "Unknown"
// 	}
// }

type ServerState string

const (
	StateIdle      ServerState = "idle"
	StateConnected ServerState = "connected"
	StateError     ServerState = "error"
)

func main() {
	status := StatusOK
	println(status)

	state := StateConnected
	if state != StateIdle {
		println(true)
	}
}
