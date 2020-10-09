package main

import (
	"github.com/akkagao/gms"
	"github.com/akkagao/gms/example/V1/server/handler/user"
)

func main() {
	gms := gms.NewGms()

	gms.AddRouter("user.Add", user.UserAdd)

	gms.Run()
}
