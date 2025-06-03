package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func nextDayHandler(res http.ResponseWriter, req *http.Request) {
	gnow := req.URL.Query().Get("now")
	if gnow == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("now missing"))
		return
	}
	parsTime, err := time.Parse(DateFormat, gnow)
	if err != nil {
		fmt.Println("Ошибка при разборе даты dstart", err)
		return
	}

	gdate := req.URL.Query().Get("date")
	if gdate == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("date missing"))
		return
	}

	grepeat := req.URL.Query().Get("repeat")
	if grepeat == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("repeat missing"))
		return
	}

	str, err := NextDate(parsTime, gdate, grepeat)
	if err != nil {
		fmt.Println("Ошибка NextDate:", err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("NextDate missing"))
		return
	}

	res.Write([]byte(str))

}

func Init() *chi.Mux {

	rout := chi.NewRouter()
	rout.Handle("/*", http.FileServer(http.Dir("web")))
	rout.Get("/api/nextdate", nextDayHandler)
	rout.Post("/api/task", addTaskHandler)
	rout.Get("/api/tasks", getTaskHandler)
	rout.Get("/api/task", getoneTaskHandler)
	rout.Put("/api/task", putTaskHandler)
	rout.Delete("/api/task", delTaskHandler)
	rout.Post("/api/task/done", doneTaskHandler)

	return rout
}
