package main

import (
	"github.com/akka/gms"
	"github.com/akka/gms/example/V1/server/handler/user"
)

func main() {
	gms := gms.NewGms()

	gms.AddRouter("user.Add", user.UserAdd)
	
	gms.Run()
}
