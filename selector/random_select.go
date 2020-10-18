package selector

import (
	"math/rand"
)

type RandomSelect struct {
	servers []string
}

func NewRandomSelect(servers []string) ISelector {
	return &RandomSelect{servers: servers}
}

func (r *RandomSelect) Select() string {
	size := len(r.servers)
	if size == 0 {
		return ""
	}
	return r.servers[rand.Intn(size)]
}
