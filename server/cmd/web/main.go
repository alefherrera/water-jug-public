package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"water-jug/internal"
	"water-jug/internal/models"
)

func main() {
	router := mux.NewRouter()
	solver := internal.NewSolver()
	router.HandleFunc("/solve", func(writer http.ResponseWriter, request *http.Request) {
		var input models.Input

		if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		solution, err := solver.Do(input.X, input.Y, input.Z)

		if err != nil {
			if err == internal.NoSolution {
				http.Error(writer, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("content-type", "application/json")

		if err = json.NewEncoder(writer).Encode(solution); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
