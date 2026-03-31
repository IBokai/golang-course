package dto

import "repo-stat/api/internal/domain"

type ServiceDTO struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type PingResponse struct {
	Status   string       `json:"status"`
	Services []ServiceDTO `json:"services"`
}

func NewPingResponse(status string, services []domain.ServiceInfo) PingResponse {
	dtos := make([]ServiceDTO, len(services))
	for i, s := range services {
		dtos[i] = ServiceDTO{
			Name:   s.Name,
			Status: string(s.Status),
		}
	}
	return PingResponse{
		Status:   status,
		Services: dtos,
	}
}
