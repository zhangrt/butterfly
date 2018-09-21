package user

import (
	"alien"
	"butterfly/service/system/user"
	"fmt"
)

func Add(ctx *alien.Context)  {
	user.Add("11")
	fmt.Fprintf(ctx.ResponseWriter, "Users")
	return
}