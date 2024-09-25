package controller

import (
	"encoding/json"
	"fmt"
	"golang-rede-social/src/banco"
	"golang-rede-social/src/model"
	"golang-rede-social/src/repository"
	"io"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, error := io.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var usuario model.Usuario

	if error = json.Unmarshal(corpoRequest, &usuario); error != nil {
		log.Fatal(error)
	}

	db, error := banco.Conectar()

	if error != nil {
		log.Fatal(error)
	}

	repositorio := repository.NovoRepositorioDeUsuarios(db)

	ID, error := repositorio.CriarUsuario(usuario)

	if error != nil {
		log.Fatal(error)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", ID)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os Usuários!"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um Usuário!"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
