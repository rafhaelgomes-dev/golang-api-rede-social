package rotas

import (
	"golang-rede-social/src/controller"
	"net/http"
)

var RotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controller.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controller.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
