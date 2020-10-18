package selector

type ISelector interface {
	Select() string // SelectFunc
	// UpdateServer(servers map[string]string)
}
