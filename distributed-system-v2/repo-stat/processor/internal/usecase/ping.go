package usecase

import "context"

type PingUseCase struct{}

func NewPing() *PingUseCase {
	return &PingUseCase{}
}

func (uc *PingUseCase) Ping(ctx context.Context) string {
	return "ok"
}
