package main

import (
	"net/http"
	"strconv"

	"github.com/Adilfarooque/todolist/db"
	"github.com/Adilfarooque/todolist/utils/models"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

func main() {
	db.InitDB()

	r := gin.Default()
	r.LoadHTMLFiles("tmpl/*.html")

	//Display th todo list
	r.GET("/", func(c *gin.Context) {
		var todos []models.Todo
		result := db.DB.Find(&todos)
		if result.Error != nil {
			c.String(http.StatusInternalServerError, "Error fetching todos")
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	//Add a new todo
	r.POST("/add", func(c *gin.Context) {
		title := c.PostForm("title")
		newtodo := models.Todo{Title: title}

		db.DB.Create(&newtodo)
		c.Redirect(http.StatusFound, "/")
	})

	//Mark a todo as done or undone
	r.POST("/toggle/:id", func(c *gin.Context) {
		todoID, _ := strconv.Atoi(c.Param("id"))
		var todo models.Todo

		if err := db.DB.First(&todo, todoID).Error; err == nil {
			todo.Done = !todo.Done
		}
		db.DB.Save(&todo) //Update the todo in the database

		c.Redirect(http.StatusFound, "/")
	})

	//edit todo list
	r.GET("/edit/:id", func(c *gin.Context) {
		todoID, _ := strconv.Atoi(c.Param("id"))
		var todo models.Todo

		if err := db.DB.First(&todo, todoID).Error; err == nil {
			todo.Title = c.PostForm("title") //Update the title from input
			db.DB.Save(&todo)
		}

		c.Redirect(http.StatusFound, "/")
	})

	//Delete todo list
	r.POST("/delete/:id", func(c *gin.Context) {
		todoID, _ := strconv.Atoi(c.Param("id"))
		db.DB.Delete(&models.Todo{}, todoID)

		c.Redirect(http.StatusFound, "/")
	})
	r.Run(port)
}
