package server

import (
	"fmt"
	"net/http"

	nextdate "github.com/ViktorDevGo/go_final_project/pkg/api"
)

func Run() error {
	//nextdate.Init()
	newr := nextdate.Init()
	port := 7540
	fmt.Println("Запуск сервера на порту %d\n", port)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), newr)
}
