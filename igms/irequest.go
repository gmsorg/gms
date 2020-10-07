package igms

import "github.com/akka/gms/context"

type Controller func(c *context.Context) error

type IRequest interface {
}
