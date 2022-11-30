package pingService

import "context"

type PingService interface {
	GetPing(ctx context.Context) (string, error)
}

type pingService struct {
}

func NewPingService() pingService {
	return pingService{}
}

func (ps pingService) GetPing(ctx context.Context) (string, error) {
	return "Hello from the server", nil
}
