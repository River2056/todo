package main

import (
    "math/rand"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"todo/assets"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func main() {
    todoList := make([]models.TodoItem, 0)

    allHtmls := assets.GetAllTemplates()
	router := gin.Default()
	router.LoadHTMLFiles(allHtmls...)

	router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{})
	})

    router.GET("/fetch-todos", func(c *gin.Context) {
        c.HTML(http.StatusOK, "todolist.html", gin.H{
            "todoList": todoList,
        })
    })

    router.POST("/add", func(c *gin.Context) {
        jsonBytes, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            log.Fatalf("error while reading post body, %v\n", err)
        }
        jsonData := string(jsonBytes)
        jsonData = strings.Replace(jsonData, "%20", " ", -1)
        jsonArr := strings.Split(jsonData, "=")
        t := models.TodoItem{
            Id: rand.Int(),
            Content: jsonArr[1],
            IsDone: false,
        }
        todoList = append([]models.TodoItem{t}, todoList...)

        c.HTML(http.StatusOK, "todolist.html", gin.H{
            "todoList": todoList,
        })
    })

	router.Run(":8080")
}
