package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/hotel_emp/domain/usecase"
)

type HotelEmpHandler interface {
	HandleCreateHotelEmp(c *gin.Context)
	HandleCreateHotelOwner(c *gin.Context)
	HandleLoginHotelEmp(c *gin.Context)
}

type hotelEmpHandler struct {
	requestValidator        *validator.Validate
	createHotelEmpUseCase   usecase.CreateHotelEmpUseCase
	createHotelOwnerUseCase usecase.CreateHotelOwnerUseCase
	loginHotelEmpUseCase    usecase.LoginHotelEmpUseCase
}

func NewHotelEmpHandler(
	requestValidator *validator.Validate,
	createHotelEmpUseCase usecase.CreateHotelEmpUseCase,
	createHotelOwnerUseCase usecase.CreateHotelOwnerUseCase,
	loginHotelEmpUseCase usecase.LoginHotelEmpUseCase,
) HotelEmpHandler {
	return &hotelEmpHandler{
		requestValidator:        requestValidator,
		createHotelEmpUseCase:   createHotelEmpUseCase,
		createHotelOwnerUseCase: createHotelOwnerUseCase,
		loginHotelEmpUseCase:    loginHotelEmpUseCase,
	}
}