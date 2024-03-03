package routes

import (
	"go-30/todo/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func contentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func Routes() {
	router := mux.NewRouter()
	router.Use(contentTypeApplicationJsonMiddleware)

	router.HandleFunc("/todo", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todo", controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", controllers.DeleteTodo).Methods("DELETE")

	http.ListenAndServe(":80", router)
}
