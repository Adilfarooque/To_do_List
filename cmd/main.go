package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)



func main() {
	r := gin.Default()
	r.LoadHTMLFiles("tmpl/*.html")

	//Display th todo list

}
