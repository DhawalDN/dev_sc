package controllers

import (
	"dev_sc/forms"
	"dev_sc/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

var studentModel = new(models.StudentModel)

// studentModel:= models.StudentModel{}

type StudentController struct{}

func (student *StudentController) Create(c *gin.Context) {
	var data forms.CreateStudentCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := studentModel.Create(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Student could not be registered", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Student Registered", "form": data})
}

func (movie *StudentController) Find(c *gin.Context) {
	list, err := studentModel.Find()
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

func (movie *StudentController) Get(c *gin.Context) {
	id := c.Param("id")
	profile, err := studentModel.Get(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Student not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": profile})
	}
}

func (movie *StudentController) Update(c *gin.Context) {
	id := c.Param("id")
	data := forms.CreateStudentCommand{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid Parameters"})
		c.Abort()
		return
	}
	err := studentModel.Update(id, data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Student Data Could Not Be Updated", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Student Data Updated"})
}

func (movie *StudentController) Delete(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	err := studentModel.Delete(id)
	if err != nil {
		c.JSON(406, gin.H{"message": "Student Could Not Be Deleted", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Student Deleted"})
}
