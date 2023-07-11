package getapis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	"fmt"
	"path/filepath"
)

// Creating Struct for Employees
type Employee struct {
	EmployeeID		int		`json:"empid"`
	FullName		string	`json:"fullname"`
	Gender			string	`json:"gender"`
	ExpStart		string	`json:"expstart"`
	ExpEnd			string	`json:"expend"`
	Mobile			string	`json:"mobile"`
	Team			string	`json:"team"`
	Resume			[]byte	`json:"resume"`
	Email			string	`json:"email"`
}


// Function to get all employees data
func GetAllEmployees(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT empid, fullname, gender, expstart, expend, mobile, team, resume, email FROM employee.details")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer rows.Close()

		var employees []Employee
	
		for rows.Next() {
			var employee Employee
	
			if err := rows.Scan(&employee.EmployeeID, &employee.FullName, &employee.Gender, &employee.ExpStart, &employee.ExpEnd, &employee.Mobile, &employee.Team, &employee.Resume, &employee.Email); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			employees = append(employees, employee)
		}
	
		c.JSON(http.StatusOK, employees)
	}
}


func GetResumeOfEmployee(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		empid := c.Param("empid")

		var resume []byte
		err := db.QueryRow("SELECT resume FROM employee.details WHERE empid = $1", empid).Scan(&resume)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		contentType := http.DetectContentType(resume)
		fileExtension := filepath.Ext(contentType)
	
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=file%s", fileExtension))
		c.Header("Content-Type", contentType)
		c.Data(http.StatusOK, contentType, resume)
	}
}
