package routes

import (
	"database/sql"
	"net/http"
	"swirflabstest/handlers"
	"swirflabstest/repository"

	"github.com/gorilla/mux"
)

func Routing(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	employeeRepo := repository.NewEmployeeRepository(db)
	employeeHandler := handlers.NewEmployeeHandler(employeeRepo)

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/employee", employeeHandler.CreateNewEmployee).Methods("POST")
	api.HandleFunc("/employee", employeeHandler.GetAllEmployees).Methods("GET")
	api.HandleFunc("/employee/{name}/{in}", employeeHandler.GetEmployeeByUniqueKey).Methods("GET")
	api.HandleFunc("/employee/{name}/{in}", employeeHandler.DeleteEmployeeByUniqueKey).Methods("DELETE")

	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/").Handler(fs)

	return router
}
