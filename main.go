package main

import (
	"fmt"
	"golang-rede-social/src/router"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Rodando API!")
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":5000", r))
}
