package user

import (
	"fmt"

	"github.com/akka/gms/context"
)

func UserAdd(c *context.Context) error {
	fmt.Println("call userAdd...")
	return nil
}
