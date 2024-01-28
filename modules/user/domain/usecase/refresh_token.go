package usecase

import (
	"context"
	"log/slog"
	"trekkstay/api/middlewares/constant"
	"trekkstay/core"
	"trekkstay/pkgs/log"
)

type RefreshTokenUseCase interface {
	ExecRefreshToken(ctx context.Context) (*string, *string, error)
}

type refreshTokenUseCaseImpl struct {
	tokenProvider    TokenProvider
	accessTokenTime  int
	refreshTokenTime int
}

var _ RefreshTokenUseCase = (*refreshTokenUseCaseImpl)(nil)

func NewRefreshTokenUseCase(tokenProvider TokenProvider,
	accessTokenTime int, refreshTokenTime int) RefreshTokenUseCase {
	return &refreshTokenUseCaseImpl{
		tokenProvider:    tokenProvider,
		accessTokenTime:  accessTokenTime,
		refreshTokenTime: refreshTokenTime,
	}
}

func (useCase refreshTokenUseCaseImpl) ExecRefreshToken(ctx context.Context) (*string, *string, error) {
	userID := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID()

	// Generate access token
	accessToken, err := useCase.tokenProvider.Generate(
		map[string]interface{}{
			"user_id": userID,
		},
		useCase.accessTokenTime,
	)
	if err != nil {
		log.JsonLogger.Error("ExecLoginUser.generate_access_token",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, nil, constant.ErrInternal(err)
	}

	// Generate refresh token
	refreshToken, err := useCase.tokenProvider.Generate(
		map[string]interface{}{
			"user_id": userID,
		},
		useCase.refreshTokenTime,
	)
	if err != nil {
		log.JsonLogger.Error("ExecLoginUser.generate_refresh_token",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, nil, constant.ErrInternal(err)
	}

	accessTokenString := accessToken["token"].(string)
	refreshTokenString := refreshToken["token"].(string)

	return &accessTokenString, &refreshTokenString, nil
}
