package usecase

import "context"

type DefaultUsecaseInterface interface {
	WelcomeUsecase(ctx context.Context) string
}
