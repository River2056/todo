package models

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Id      int
	Content string
	IsDone  bool
}
