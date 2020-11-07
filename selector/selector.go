package selector

type ISelector interface {
	Select() (string, error) // SelectFunc
	// UpdateServer(servers map[string]string)
}
