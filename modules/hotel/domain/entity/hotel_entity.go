package entity

import "trekkstay/core"

type HotelEntity struct {
	core.Entity   `json:",inline"`
	Name          string              `json:"name" gorm:"not null;"`
	OwnerID       string              `json:"owner_id" gorm:"not null;"`
	Email         string              `json:"email" gorm:"uniqueIndex;not null;"`
	Phone         string              `json:"phone" gorm:"uniqueIndex;not null;"`
	CheckInTime   string              `json:"check_in_time" gorm:"not null;default:14:00"`
	CheckOutTime  string              `json:"check_out_time" gorm:"not null;default:12:00"`
	ProvinceCode  string              `json:"-" gorm:"not null;"`
	DistrictCode  string              `json:"-" gorm:"not null;"`
	WardCode      string              `json:"-" gorm:"not null;"`
	AddressDetail string              `json:"address_detail" gorm:"not null;"`
	Description   string              `json:"description" gorm:"not null;"`
	Status        string              `json:"status" gorm:"not null;default:active"`
	Rooms         []RoomEntity        `json:"rooms" gorm:"foreignKey:HotelID;references:ID"`
	Owner         HotelEmployeeEntity `json:"owner" gorm:"foreignKey:OwnerID;references:ID"`
	HotelFacility HotelFacilityEntity `json:"hotel_facility" gorm:"foreignKey:HotelID;references:ID"`
	Province      ProvinceEntity      `json:"province" gorm:"foreignKey:ProvinceCode;references:Code"`
	District      DistrictEntity      `json:"district" gorm:"foreignKey:DistrictCode;references:Code"`
	Ward          WardEntity          `json:"ward" gorm:"foreignKey:WardCode;references:Code"`
	Coordinates   CoordinatesEntity   `json:"coordinates" gorm:"type:jsonb;default:null"`
	Videos        MediaObject         `json:"videos" gorm:"type:jsonb;default:null"`
	Images        MediaObject         `json:"images" gorm:"type:jsonb;default:null"`
}

func (HotelEntity) TableName() string {
	return "hotels"
}