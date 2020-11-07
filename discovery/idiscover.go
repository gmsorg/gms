package discovery

type IDiscover interface {
	GetServer() ([]string, error)
	DeleteServer(key string)
}
