package server

import (
	"fmt"
	"net/http"

	api "github.com/ViktorDevGo/go_final_project/pkg/api"
)

func Run() error {
	//nextdate.Init()
	newr := api.Init()
	port := 7540
	fmt.Println("Запуск сервера на порту %d\n", port)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), newr)
}
