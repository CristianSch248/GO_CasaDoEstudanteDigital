package handlers

import (
	"fmt"
	"net/http"
)

type User struct{}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Criado usuário")
}
func (u *User) Read(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando usuário")
}
func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Editando usuário")
}
func (u *User) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deletando usuário")
}

/*
	w http.ResponseWriter:
	Este parâmetro é usado para construir a resposta HTTP que será enviada de volta ao cliente.
	Através dele, você pode escrever o corpo da resposta, definir cabeçalhos, entre outras coisas.
*/
/*
	r *http.Request: Este parâmetro representa a requisição HTTP que o servidor recebeu.
	Ele contém todas as informações sobre a requisição, como método, URL, cabeçalhos, corpo, etc.
*/
