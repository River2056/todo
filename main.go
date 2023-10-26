package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"todo/assets"
	"todo/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var todoList []models.TodoItem
var editTodo models.TodoItem
var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("./data.db"), &gorm.Config{})
	if err != nil {
		// database not exists
		panic(err)
	}

	db.AutoMigrate(&models.TodoItem{})
	DB = db
}

func main() {
	todoList = make([]models.TodoItem, 0)

	allHtmls := assets.GetAllTemplates()
	router := gin.Default()
	router.LoadHTMLFiles(allHtmls...)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/fetch-todos", func(c *gin.Context) {
		DB.Find(&todoList)
		c.HTML(http.StatusOK, "todolist.html", gin.H{
			"todoList": todoList,
		})
	})

	router.GET("/fetch-add-todo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addtodo.html", gin.H{
			"editTodo": models.TodoItem{},
			"isEdit":   false,
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
			Content: jsonArr[1],
			IsDone:  false,
		}
		DB.Create(&t)
		DB.Find(&todoList)
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

		DB.Delete(&models.TodoItem{}, id)
		DB.Find(&todoList)

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
					"isEdit":   true,
				})
				return
			}
		}
	})

	router.Run(":9000")
}
