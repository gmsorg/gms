package discovery

type IDiscovery interface {
	GetServer() ([]string, error)
}
