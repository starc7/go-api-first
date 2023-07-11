package postapis

import (
	"io"
	"github.com/gin-gonic/gin"
	"database/sql"
	"net/http"
	_ "github.com/lib/pq"
)

func PostEmployeeData(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		empid := c.PostForm("empid")
		fullname := c.PostForm("fullname")
		gender := c.PostForm("gender")
		expstart := c.PostForm("expstart")
		expend := c.PostForm("expend")
		mobile := c.PostForm("mobile")
		team := c.PostForm("team")
		email := c.PostForm("email")
		resume, err := c.FormFile("resume")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return		
		}

		uploadedFile, err := resume.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer uploadedFile.Close()

		fileData, err := io.ReadAll(uploadedFile)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    	    return
		}

		_, err = db.Exec("INSERT INTO employee.details (empid, fullname, gender, expstart, expend, mobile, team, resume, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", empid, fullname, gender, expstart, expend, mobile, team, fileData, email)
    	if err != nil {
    	    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    	    return
    	}
		c.JSON(http.StatusOK, gin.H{"message": "File submitted successfully"})
	}
}