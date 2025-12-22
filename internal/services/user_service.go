package services

import (
	"gin-user-api/internal/models"
	"gin-user-api/internal/repositories"
)

type UserService struct {
	Repo repositories.UserRepository
}

func (s *UserService) GetAll(users *[]models.User) error {
	return s.Repo.FindAll(users)
}

func (s *UserService) GetByID(user *models.User, id uint) error {
	return s.Repo.FindByID(user, id)
}

func (s *UserService) Create(user *models.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) Update(user *models.User) error {
	return s.Repo.Update(user)
}

func (s *UserService) Delete(user *models.User) error {
	return s.Repo.Delete(user)
}
