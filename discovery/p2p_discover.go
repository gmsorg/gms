package discovery

import "errors"

type P2PDiscover struct {
	address []string
}

/*
NewP2PDiscover 初始化点对点服务发现
*/
func NewP2PDiscover(address []string) IDiscover {
	return &P2PDiscover{address: address}
}

/*
GetServer 获取所有服务
*/
func (p *P2PDiscover) GetServer() ([]string, error) {
	if len(p.address) < 1 {
		return nil, errors.New("no server")
	}
	return p.address, nil
}
