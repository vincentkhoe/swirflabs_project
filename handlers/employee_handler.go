package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"swirflabstest/models"
	"swirflabstest/repository"

	"github.com/gorilla/mux"
)

type EmployeeHandler struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeHandler(repo *repository.EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{Repo: repo}
}

func (h *EmployeeHandler) CreateNewEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var emp models.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}

	if strings.TrimSpace(emp.Name) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Name is required"})
		return
	}

	if strings.TrimSpace(emp.IdentificationNumber) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Identification number is required"})
		return
	}

	if strings.TrimSpace(emp.Occupation) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Occupation is required"})
		return
	}

	if strings.TrimSpace(emp.DateOfBirth) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Date of birth is required"})
		return
	}

	emp.UniqueKey = strings.ReplaceAll(strings.TrimSpace(emp.Name), " ", "") + "_" + strings.TrimSpace(emp.IdentificationNumber)

	age, err := models.CountAge(emp.DateOfBirth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid date of birth format"})
		return
	}

	emp.Age = age

	exist, err := h.Repo.CheckDuplicate(emp.UniqueKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database check duplicate error"})
		return
	}

	if exist {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Employee with this name and identification number already exist"})
		return
	}

	err = h.Repo.CreateEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to create employee"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Message: "Employee created successfully",
		Data:    emp,
	})
}

func (h *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	employees, err := h.Repo.GetAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to load employees"})
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func (h *EmployeeHandler) GetEmployeeByUniqueKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	name := strings.TrimSpace(vars["name"])
	identificationNumber := strings.TrimSpace(vars["in"])
	uniqueKey := name + "_" + identificationNumber

	employee, err := h.Repo.GetByUniqueKey(uniqueKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invaled employee name or in"})
		return
	}

	if employee == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Employee name and in not found"})
		return
	}

	json.NewEncoder(w).Encode(employee)
}

func (h *EmployeeHandler) DeleteEmployeeByUniqueKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")

	vars := mux.Vars(r)
	name := strings.TrimSpace(vars["name"])
	identificationNumber := strings.TrimSpace(vars["in"])
	uniqueKey := name + "_" + identificationNumber

	err := h.Repo.DeleteEmployee(uniqueKey)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid employee name or identification number"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to delete employee"})
		return
	}

	json.NewEncoder(w).Encode(models.SuccessResponse{
		Message: "Employee deleted successfully",
	})
}
