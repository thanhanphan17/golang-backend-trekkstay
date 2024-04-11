package migration

import (
	"gorm.io/gorm"
	destination "trekkstay/modules/destination/domain/entity"
	hotel "trekkstay/modules/hotel/domain/entity"
	hotelEmp "trekkstay/modules/hotel_emp/domain/entity"
	hotelRoom "trekkstay/modules/hotel_room/domain/entity"
	region "trekkstay/modules/region/domain/entity"
	reservation "trekkstay/modules/reservation/domain/entity"
	user "trekkstay/modules/user/domain/entity"
)

// Migration is a function that performs migrations on the given database.
//
// It takes a pointer to a gorm.DB object as a parameter.
// It returns an error indicating the success or failure of the migration operation.
func Migration(db *gorm.DB) error {
	err := db.AutoMigrate(
		user.UserEntity{},
		hotel.HotelEntity{},
		hotelEmp.HotelEmpEntity{},
		hotelRoom.HotelRoomEntity{},
		region.ProvinceEntity{},
		region.DistrictEntity{},
		region.WardEntity{},
		reservation.ReservationEntity{},
		destination.DestinationEntity{},
	)

	return err
}
