package http

import (
	"context"
	"getCurs/internal/entity"
	"net/http"
)

type receiverService interface {
	GetAllRates(ctx context.Context) ([]entity.Rate, error)
	GetRateByCurrency(ctx context.Context, currency string) (entity.Rate, error)
	CalculateStatistics(ctx context.Context, currency string) (minPrice, maxPrice, hourChange float64, err error)
}

type updaterService interface {
	StartAutoUpdate(minutes int) error
	StopAutoUpdate() error
}

type Server struct {
	receiver receiverService
	updater  updaterService
}

func NewServer(receiver receiverService, updater updaterService) *Server {
	return &Server{receiver: receiver, updater: updater}
}

func (s *Server) GetAllRates(w http.ResponseWriter, r *http.Request) {
}
func (s *Server) GetRateByCurrency(w http.ResponseWriter, r *http.Request) {}
func (s *Server) CalculateStatistics(w http.ResponseWriter, r *http.Request) {
}
func (s *Server) StartAutoUpdate(w http.ResponseWriter, r *http.Request) {
}
func (s *Server) StopAutoUpdate(w http.ResponseWriter, r *http.Request) {
}
