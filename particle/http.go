package particle

import (
	"fmt"
	"net/http"
)

func ParticleHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	fmt.Println("Can I handle this???")
	fmt.Println("Checking... [" + r.URL.String() + "]")
	if r.URL.String() == "" && r.Method == "POST" {
		Trigger()
		fmt.Println("Triggered it")
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
