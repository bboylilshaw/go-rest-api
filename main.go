package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func TodoIndex(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	todos := Todos{
		Todo{Name: "Write some code"},
		Todo{Name: "Clean up house"},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos", TodoIndex)

	log.Fatal(http.ListenAndServe(":8080", router))
}
