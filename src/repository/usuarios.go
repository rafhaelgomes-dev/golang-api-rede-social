package repository

import (
	"database/sql"
	"fmt"
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

// Buscar tras todos os usuáiros que atende o filtro de nome ou nick
func (repositorio usuarios) Buscar(usuario string) ([]model.Usuario, error) {
	usuario = fmt.Sprintf("%%%s%%", usuario) //%usuario%

	linhas, error := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?", usuario, usuario)
	if error != nil {
		return nil, error
	}
	defer linhas.Close()

	var usuarios []model.Usuario

	for linhas.Next() {
		var usuario model.Usuario
		if error = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); error != nil {
			return nil, error
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarById traz um usuário do banco de dados
func (repositorio usuarios) BuscarById(ID uint64) (model.Usuario, error) {
	linhas, error := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", ID)

	if error != nil {
		return model.Usuario{}, error
	}

	defer linhas.Close()

	var usuario model.Usuario

	if linhas.Next() {
		if error = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); error != nil {
			return model.Usuario{}, error
		}
	}

	return usuario, nil
}

func (repositorio usuarios) AtualizarById(ID uint64, usuario model.Usuario) error {
	statement, error := repositorio.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); error != nil {
		return error
	}

	return nil
}

func (repositorio usuarios) DeletarById(ID uint64) error {
	statement, error := repositorio.db.Prepare("delete from usuarios where id = ?")

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(ID); error != nil {
		return error
	}

	return nil
}
