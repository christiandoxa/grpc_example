package repository

import (
	"time"
)

const TableNameTodo = "todo"

type TodoModel struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Title       string    `gorm:"column:title;not null" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Reminder    time.Time `gorm:"column:reminder" json:"reminder"`
}

func (*TodoModel) TableName() string {
	return TableNameTodo
}
