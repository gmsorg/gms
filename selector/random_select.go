package selector

import (
	"errors"
	"math/rand"
	"time"

	"github.com/akkagao/gms/discovery"
)

type RandomSelect struct {
	discovery discovery.IDiscover
}

func NewRandomSelect(discovery discovery.IDiscover) ISelector {
	return &RandomSelect{discovery: discovery}
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
