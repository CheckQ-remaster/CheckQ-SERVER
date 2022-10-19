package models

import (
	"time"
)

type User struct {
	Id    string `gorm:"size:64; primaryKey; NOT NULL" json:"id"`
	Pw    string `gorm:"size:128" json:"pw"`
	Name  string `gorm:"size:64" json:"name"`
	Email string `gorm:"size:64" json:"email"`
	Phone string `gorm:"size:64s" json:"phone"`
	Role  string `gorm:"size:32" json:"role"` // 사용자 역할; 사용자와 기업으로 나뉨, 미들웨어를 만들어서 접근 가능한 부분과 불가능한 부분을 나눌 예정

	CreatedAt time.Time

	Hotels []Hotel `gorm:"foreignKey:UserId"`
}

type Hotel struct {
	UserId      string `gorm:"size:64" json:"userid"`      // 예약자; 았을경우 예약이 되어 있는걸로 판단
	Room        string `gorm:"size:64" json:"room"`        // 방 번호
	Personnel   int    `gorm:"size:64" json:"personnel"`   // 방 인원
	Price       int    `gorm:"size:1024" json:"price"`     // 방 가격
	Hotelname   string `gorm:"size:64" json:"hotelname"`   // 호텔 이름
	Location    string `gorm:"size:64" json:"location"`    // 호텔 위치
	PhoneNumber string `gorm:"size:64" json:"phoneNumber"` // 호텔 연락처
	CheckIn     string `gorm:"size:64" json:"checkin"`     // 체크 인
	CheckOut    string `gorm:"size:64" json:"checkout"`    // 체크 인
	Alpha       int    `gorm:"size:64" json:"alpha"`       // 알파값

	HotelImage string `gorm:"size:128"`
}
