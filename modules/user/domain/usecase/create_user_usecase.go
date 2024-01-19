package usecase

import (
	"context"
	"log/slog"
	"trekkstay/modules/user/constant"
	"trekkstay/modules/user/domain/entity"
	"trekkstay/pkgs/log"
)

type CreateUserUseCase interface {
	ExecCreateUser(ctx context.Context, userEntity entity.UserEntity) error
}

type createUserUseCaseImpl struct {
	hashAlgo   HashAlgo
	readerRepo userReaderRepository
	writerRepo userWriterRepository
}

var _ CreateUserUseCase = (*createUserUseCaseImpl)(nil)

func NewCreateUserUseCase(hashAlgo HashAlgo, readerRepo userReaderRepository,
	writerRepo userWriterRepository) CreateUserUseCase {
	return &createUserUseCaseImpl{
		hashAlgo:   hashAlgo,
		readerRepo: readerRepo,
		writerRepo: writerRepo,
	}
}

func (useCase createUserUseCaseImpl) ExecCreateUser(ctx context.Context, userEntity entity.UserEntity) error {
	// Check if user already exists
	user, err := useCase.readerRepo.FindUserByCondition(ctx, map[string]interface{}{
		"email": userEntity.Email,
	})
	if user != nil {
		log.JsonLogger.Error("ExecCreateEmp.email_already_exists",
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorEmailAlreadyExists(err)
	}

	// Check if phone already exists
	user, err = useCase.readerRepo.FindUserByCondition(ctx, map[string]interface{}{
		"phone": userEntity.Phone,
	})
	if user != nil {
		log.JsonLogger.Error("ExecCreateEmp.phone_already_exists",
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorPhoneAlreadyExists(err)
	}

	// Hash password
	hashedPassword, err := useCase.hashAlgo.HashAndSalt([]byte(userEntity.Password))
	if err != nil {
		log.JsonLogger.Error("ExecCreateEmp.hash_password",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorInternalServerError(err)
	}
	userEntity.Password = hashedPassword

	// Insert user
	err = useCase.writerRepo.InsertUser(ctx, userEntity)
	if err != nil {
		log.JsonLogger.Error("ExecCreateEmp.insert_user",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorInternalServerError(err)
	}

	return nil
}
