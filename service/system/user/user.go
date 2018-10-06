package user

import (
	"alien"
	"butterfly/entity/system/user"
	_ "github.com/lib/pq"
)

func Add(id string)  {
	db := alien.GetApplication().GetDb()
	db.Create(&user.User{})
}