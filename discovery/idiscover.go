package discovery

type IDiscover interface {
	GetServer() ([]string, error)
}
