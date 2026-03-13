Project Task by Swirflabs

**Main Tasks**
1. Create a form page for filling out the Employee Data, which includes Name, Identification Number, Age, Address, Occupation (Unemployed, Programmer, Designer, Architect,
Artist), Place, and Date of Birth, also a table of employees according to the following rules:
  * The entry of Identification Number and Name serves as a unique key that differentiates each data entry.
  * Age is determined by the system based on the date of birth.
  * The fields Name, Identification Number, Age, and Occupation are mandatory and must be validated.
  * The form is submitted to the system using the submit button.
  * When the form fails to submit, display the reason for the failure.
  * When the form is successfully submitted, all fields are cleared and the participant list is updated.
    
2. The table in point #1 must be created in the following manners:
  * The column header row has a Navy Blue background color with White text, and the text size is twice as large as the text size of the table contents.
  * The Name and Occupation columns are right-aligned.
  * The Age column is center-aligned.
  * The Age column has a text size that is 1.2 times larger.
  * The width proportions of the Name, Age, Address, and Occupation columns are 30%, 10%, 40%, and 20% respectively.

**Bonus**
  * The form submitted on the HTML page should be processed asynchronously to prevent the user from being redirected to another page.
  * Integrating a back-end process using PHP and MySQL for data transactions further elevates this approach by enabling dynamic content management and robust data handling.
  * Provide robust protection for user data while maintaining quick and efficient processing speeds.

**Routes**
  * Main page | http://localhost:8080
  * Create new employee | http://localhost:8080/api/employee | METHOD "POST"
  * Get all employees | http://localhost:8080/api/employee | METHOD "GET"
  * Search employee by uniqueKey | http://localhost:8080/api/employee/{name}/{in} | METHOD "GET"
  * Delete employee by uniqueKey | http://localhost:8080/api/employee/{name}/{in} | METHOD "DELETE"
  
**Notes**
  * The fornt-end is written in Javascript, back-end is written in Go-Lang, database is using MySQL.
  * Run cmd/main.go
  * Back-end Server connection:
      - Host name = Localhost
      - Port = 8080
  * Database Env:
      - DB_USER=
      - DB_PASS=
      - DB_HOST=localhost
      - DB_PORT=3306
      - DB_NAME=employee_db
