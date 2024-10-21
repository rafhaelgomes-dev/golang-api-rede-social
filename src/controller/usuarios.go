package controller

import (
	"encoding/json"
	"golang-rede-social/src/banco"
	"golang-rede-social/src/model"
	"golang-rede-social/src/repository"
	"golang-rede-social/src/response"
	"io"
	"net/http"
	"strings"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, error := io.ReadAll(r.Body)
	if error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var usuario model.Usuario

	if error = json.Unmarshal(corpoRequest, &usuario); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = usuario.Preparar(); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := banco.Conectar()

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)

	ID, error := repositorio.CriarUsuario(usuario)

	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	usuario.ID = uint(ID)

	response.JSON(w, http.StatusCreated, usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, error := banco.Conectar()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuarios, error := repositorio.Buscar(nomeOuNick)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	response.JSON(w, http.StatusOK, usuarios)
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
