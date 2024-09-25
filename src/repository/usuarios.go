package repository

import (
	"database/sql"
	"golang-rede-social/src/model"
)

type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositórios de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// CriarUsuario insere um usuário no banco de dados
func (repositorio usuarios) CriarUsuario(usuario model.Usuario) (uint64, error) {
	statement, error := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")

	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if error != nil {
		return 0, error
	}

	ID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(ID), nil
}
