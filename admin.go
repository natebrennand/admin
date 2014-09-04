package admin

import (
	"log"
	"net/http"
	"os"
)

// AdminPortString is the string respresentation of the port that the /ping endpoint will listen on.
var AdminPortString string

func init() {
	AdminPortString = os.Getenv("ADMIN_PORT")
	if AdminPortString == "" {
		log.Print("ADMIN_PORT not set so listening on 8001")
		AdminPortString = "8000"
	}

	go func() {
		http.ListenAndServe(":"+AdminPortString, http.HandlerFunc(healthcheck))
	}()
}

func healthcheck(resp http.ResponseWriter, req *http.Request) {
	n, err := resp.Write([]byte(`pong`))
	if n != 4 || err != nil {
		log.Print("Admin: Failed to properly respond to healthcheck")
	}
}
