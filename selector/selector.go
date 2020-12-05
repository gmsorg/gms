package selector

import (
	"github.com/gmsorg/gms/connection"
)

type ISelector interface {
	Select() (string, error)                     // SelectFunc
	SelectConn() (connection.IConnection, error) // SelectFunc
	// UpdateServer(servers map[string]string)
}
