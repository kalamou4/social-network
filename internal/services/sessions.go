package services

import (
	"social-network/internal/models"
	"social-network/internal/repository"
)

type SessionService struct {
	repo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) *SessionService {
	return &SessionService{
		repo: repo,
	}
}

func (s *SessionService) GetSessionByID(sessionID string) (*models.Session, error) {
	return s.repo.GetByID(sessionID)
}
