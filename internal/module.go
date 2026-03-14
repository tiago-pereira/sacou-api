package internal

import "net/http"

type Module struct {
	Router http.Handler
}
