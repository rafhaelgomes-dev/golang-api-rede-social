package main

import (
	"fmt"
	"golang-rede-social/src/config"
	"golang-rede-social/src/router"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
