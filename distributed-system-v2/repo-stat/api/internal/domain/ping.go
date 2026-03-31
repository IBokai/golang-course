package domain

type PingStatus string

const (
	PingStatusUp   PingStatus = "up"
	PingStatusDown PingStatus = "down"
)

type ServiceInfo struct {
	Name   string
	Status PingStatus
}
