package routes

import (
	"go-30/todo/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	router.HandleFunc("/todo", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todo", controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", controllers.DeleteTodo).Methods("DELETE")

	// only load the .env file when running locally
	// check for a RAILWAY_ENVIRONMENT, if not found, code is running locally
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, router)
}
