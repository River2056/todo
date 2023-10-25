package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"todo/assets"
	"todo/models"

	"github.com/gin-gonic/gin"
)

var editTodo models.TodoItem

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

    router.GET("/fetch-add-todo", func(c *gin.Context) {
        c.HTML(http.StatusOK, "addtodo.html", gin.H{
            "editTodo": models.TodoItem{},
            "isEdit": false,
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
			Id:      rand.Int(),
			Content: jsonArr[1],
			IsDone:  false,
		}
		todoList = append([]models.TodoItem{t}, todoList...)
        editTodo = models.TodoItem{}

		c.HTML(http.StatusOK, "todolist.html", gin.H{
			"todoList": todoList,
		})
	})

	router.DELETE("/delete/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		for idx, todo := range todoList {
			if todo.Id == id {
				todoList = append(todoList[:idx], todoList[idx+1:]...)
			}
		}

		c.HTML(http.StatusOK, "todolist.html", gin.H{
			"todoList": todoList,
		})
	})

    router.PUT("/edit/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            panic(err)
        }

        for idx, todo := range todoList {
            if todo.Id == id {
                todoList = append(todoList[:idx], todoList[idx+1:]...)
                c.HTML(http.StatusOK, "addtodo.html", gin.H{
                    "editTodo": todo,
                    "isEdit": true,
                })
                return
            }
        }
    })

	router.Run(":9000")
}
