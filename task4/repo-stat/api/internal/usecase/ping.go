package usecase

import (
	"context"
	"repo-stat/api/internal/domain"
)

type Pinger interface {
	Ping(ctx context.Context) domain.PingStatus
}

type Ping struct {
	services map[string]Pinger
}

func NewPing(processor, subscriber Pinger) *Ping {
	return &Ping{
		services: map[string]Pinger{
			"processor":  processor,
			"subscriber": subscriber,
		},
	}
}

func (u *Ping) Execute(ctx context.Context) (string, []domain.ServiceInfo) {
	var result []domain.ServiceInfo
	allUp := true
	for name, pinger := range u.services {
		status := pinger.Ping(ctx)
		if status == domain.PingStatusDown {
			allUp = false
		}
		result = append(result, domain.ServiceInfo{
			Name:   name,
			Status: status,
		})
	}
	if !allUp {
		return "degraded", result
	}
	return "ok", result
}
