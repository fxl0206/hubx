package modelx

type ServiceInfo struct {
	Name string
	ServerIp string
	ClusterIp string
	Ports []Port
}

type Port struct {
	Protocol string
	Name string
	Url string
	Target string
}
