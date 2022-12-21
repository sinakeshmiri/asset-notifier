package ports
type APIPort interface {
	CheckDNS() (error)
}