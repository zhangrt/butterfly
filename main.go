package main

import (
	"alien"
	"butterfly/controller/system/user"
	"fmt"
)

func main()  {
	alien.BuildApplication()
	_, err := alien.CreateDbConnPool()
	if err != nil {
		fmt.Println("数据库连接失败，服务关闭中")
		return
	}
	defer alien.CloseDbConn()
	alien.GetApplication().GetRouter().Get("/users", user.Add)

	alien.GetApplication().Listen()
}