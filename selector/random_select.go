package selector

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gmsorg/gms/connection"
	"github.com/gmsorg/gms/discovery"
)

type RandomSelect struct {
	rw         sync.RWMutex
	discovery  discovery.IDiscover
	connection map[string]connection.IConnection
}

func NewRandomSelect(discovery discovery.IDiscover) ISelector {
	return &RandomSelect{
		discovery:  discovery,
		connection: make(map[string]connection.IConnection),
	}
}

func (r *RandomSelect) Select() (string, error) {
	servers, err := r.discovery.GetServer()
	if err != nil {
		return "", err
	}
	size := len(servers)
	if size == 0 {
		return "", errors.New("no server")
	}

	rand.Seed(time.Now().UnixNano())
	return servers[rand.Intn(size)], nil
}

func (r *RandomSelect) SelectConn() (connection.IConnection, error) {
	servers, err := r.discovery.GetServer()
	if err != nil {
		return nil, err
	}
	size := len(servers)
	if size == 0 {
		return nil, errors.New("no server")
	}

	r.rw.Lock()
	defer r.rw.Unlock()

	rand.Seed(time.Now().UnixNano())
	address := servers[rand.Intn(size)]

	key := fmt.Sprintf("%v-%v", rand.Intn(10), address)
	// key := address

	if gmsConn, ok := r.connection[key]; ok {
		// fmt.Println("get ok", ok)
		return gmsConn, nil
	}

	gmsConn := connection.NewConnection(address)
	r.connection[key] = gmsConn

	return gmsConn, err
}
