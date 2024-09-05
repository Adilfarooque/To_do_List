package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)



func main() {

	r := gin.Default()
	r.LoadHTMLFiles("tmpl/*.html")

	//Display th todo list
	r.GET("/",func (c *gin.Context)  {
		
	})
}
