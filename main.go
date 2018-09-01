package main

import (
	"alien"
	"butterfly/controller/system/user"
)

func main()  {
	alien.BuildApplication()
	alien.GetApplication().GetRouter().Get("/users", user.Add)
	alien.GetApplication().Listen()
}