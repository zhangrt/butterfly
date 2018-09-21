package user

import (
	"alien"
	_ "github.com/lib/pq"
)

func Add(id string)  {
	db := alien.GetApplication().GetDb()
	db.Create(&)
}