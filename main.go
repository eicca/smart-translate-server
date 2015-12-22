package main

import (
	"github.com/eicca/translate-server/api"
)

func main() {
	api.ListenAndServeRest(":8080")
}
