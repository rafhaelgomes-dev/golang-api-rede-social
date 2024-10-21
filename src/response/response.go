package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

func Error(w http.ResponseWriter, statusCOde int, erro error) {
	JSON(w, statusCOde, struct {
		Erro string `json:"error"`
	}{
		Erro: erro.Error(),
	})
}
