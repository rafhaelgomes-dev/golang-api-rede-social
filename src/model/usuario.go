package model

import (
	"errors"
	"golang-rede-social/src/seguranca"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint      `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar usuários recebidos
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if error := usuario.formatar(etapa); error != nil {
		return error
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("nome é obrigatório")
	}

	if usuario.Nick == "" {
		return errors.New("nick é obrigatório")
	}

	if usuario.Email == "" {
		return errors.New("email é obrigatório")
	}

	if error := checkmail.ValidateFormat(usuario.Email); error != nil {
		return errors.New("email inserido inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("senha é obrigatório")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, error := seguranca.Hash(usuario.Senha)

		if error != nil {
			return error
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
