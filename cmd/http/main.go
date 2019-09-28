package main

import (
	"github.com/mattdsteele/particle-function-cloud-function/particle"

	"net/http"
)

func main() {
	http.HandleFunc("/", particle.ParticleHandler)
	http.ListenAndServe(":8081", nil)
}
