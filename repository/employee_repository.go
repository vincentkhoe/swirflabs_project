package repository

import (
	"database/sql"
	"fmt"
	"swirflabstest/models"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) CreateEmployee(emp *models.Employee) error {
	query := `insert into employees (unique_key, name, identification_number, age, address, occupation, place_of_birth, date_of_birth)
	values (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.DB.Exec(query,
		emp.UniqueKey,
		emp.Name,
		emp.IdentificationNumber,
		emp.Age,
		emp.Address,
		emp.Occupation,
		emp.PlaceOfBirth,
		emp.DateOfBirth)
	if err != nil {
		return fmt.Errorf("error inserting employee: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert id: %v", err)
	}

	emp.ID = int(id)
	return nil
}

func (r *EmployeeRepository) GetAllEmployees() ([]models.EmployeeResponse, error) {
	query := `select id, unique_key, name, identification_number, age, address, occupation, place_of_birth, date_of_birth
		from employees
		order by id desc`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error query all employee: %v", err)
	}
	defer rows.Close()

	var allEmployees []models.EmployeeResponse

	for rows.Next() {
		var emp models.EmployeeResponse
		var dob sql.NullString
		err := rows.Scan(&emp.ID, &emp.UniqueKey, &emp.Name, &emp.IdentificationNumber, &emp.Age, &emp.Address, &emp.Occupation, &emp.PlaceOfBirth, &dob)

		if err != nil {
			return nil, fmt.Errorf("error scanning employee: %v", err)
		}

		if dob.Valid {
			emp.DateOfBirth = dob.String
		}

		allEmployees = append(allEmployees, emp)
	}

	return allEmployees, nil
}

func (r *EmployeeRepository) GetByUniqueKey(uniqueKey string) (*models.EmployeeResponse, error) {
	query := `select id, unique_key, name, identification_number, age, address, occupation, place_of_birth, date_of_birth
		from employees
		where unique_key = ?`

	var emp models.EmployeeResponse
	var dob sql.NullString
	err := r.DB.QueryRow(query, uniqueKey).Scan(&emp.ID, &emp.UniqueKey, &emp.Name, &emp.IdentificationNumber, &emp.Age, &emp.Address, &emp.Occupation, &emp.PlaceOfBirth, &dob)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting employee: %v", err)
	}

	if dob.Valid {
		emp.DateOfBirth = dob.String
	}

	return &emp, nil
}

func (r *EmployeeRepository) DeleteEmployee(unique_key string) error {
	query := `delete from employees where unique_key = ?`

	result, err := r.DB.Exec(query, unique_key)
	if err != nil {
		return fmt.Errorf("error deleting employee: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("employee with identification number %v not found", unique_key)
	}

	return nil
}

func (r *EmployeeRepository) CheckDuplicate(uniqueKey string) (bool, error) {
	query := `select exists(select 1 from employees where unique_key = ?)`
	var exist bool
	err := r.DB.QueryRow(query, uniqueKey).Scan(&exist)
	if err != nil {
		return false, fmt.Errorf("error checking duplicate: %v", err)
	}

	return exist, nil
}
