package discovery

import "errors"

type P2PDiscovery struct {
	address string
}

/*
NewP2PDiscovery 初始化点对点服务发现
*/
func NewP2PDiscovery(address string) IDiscovery {
	return &P2PDiscovery{address: address}
}

/*
GetServer 获取所有服务
*/
func (p *P2PDiscovery) GetServer() ([]string, error) {
	if len(p.address) < 1 {
		return nil, errors.New("no server")
	}
	return []string{p.address}, nil
}
