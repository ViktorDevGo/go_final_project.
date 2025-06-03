package main

import (
	"fmt"

	dbase "github.com/ViktorDevGo/go_final_project/pkg/db"
	"github.com/ViktorDevGo/go_final_project/pkg/server"
)

// go clean -testcache
func main() {
	var err error
	dbFile := "scheduler.db"

	err = dbase.Init(dbFile)

	fmt.Println("DB= ", dbase.DB)
	if err != nil {
		fmt.Println(err)
	}

	defer dbase.DB.Close()

	fmt.Println("Запускаем сервер")
	err = server.Run()

	if err != nil {
		panic(err)
	}

}
