package service

import (
	"context"
	"errors"
	"project/internal/auth"
	"project/internal/models"
	"project/internal/repository"
)

type Service struct {
	UserRepo repository.UserRepo
	auth     *auth.Auth
}

type UserService interface {
	UserSignup(ctx context.Context, userData models.NewUser) (models.User, error)
	UserSignIn(ctx context.Context, userData models.NewUser) (string, error)

	AddCompanyDetails(ctx context.Context, companyData models.Company) (models.Company, error)
	ViewAllCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyDetails(ctx context.Context, cid uint64) (models.Company, error)
	ViewJob(ctx context.Context, cid uint64) ([]models.Jobs, error)

	AddJobDetails(ctx context.Context, jobData models.Jobs, cid uint64) (models.Jobs, error)
	ViewAllJobs(ctx context.Context) ([]models.Jobs, error)
	ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error)
}

func NewService(userRepo repository.UserRepo, a *auth.Auth) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		UserRepo: userRepo,
		auth:     a,
	}, nil
}
