package user

import (
	"alien"
	"fmt"
)

func Add(ctx *alien.Context)  {
	fmt.Fprintf(ctx.ResponseWriter, "Users")
	return
}