package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

//Função intermediária responsável por coletar e printar os logs referentes as ações dos usuários do sistema (método/URI/Host)
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

//Authenticate verifica se o token é válido e passa para a próxima função
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
