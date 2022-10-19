package route

import (
	"example.com/m/controller"
	"example.com/m/middleware"
	"github.com/gin-gonic/gin"
	// gin-swagger middleware
	// swagger embed files
)

func Route(e *gin.Engine) {

	e.Use(middleware.CORSMiddleware()) // cors 허용

	e.POST("/login", controller.Login)
	e.POST("/signup", controller.Signup)
	e.POST("/addhotel", controller.AddHotel)       // 호텔 추가
	e.POST("/deletehotel", controller.DeleteHotel) // 호텔 삭제
	e.POST("/reservation", controller.Reservation) // 파라미터 id, 호텛 정보(호텔이름, 방번호)

	e.POST("/uploadimg", controller.UploadIMG)
	e.GET("/uploadimg", controller.UploadIMG)
	//e.POST("/bookhotel", controller.BookHotel)

	e.GET("/logout", middleware.TokenAuthMiddleware(), controller.Logout)

	e.GET("/getmyreservation", controller.GetMyReservation) // id를 파라미터로 보내주면 유저의 정보와 호텔 예약정보를 객체배열로 보내줌
	e.GET("/gethotelinfo", controller.GetHotelInfo)         // hotel과 room을 보내주면 그 호텔의 room의 정보를 보내줌
	e.GET("/hotelsearch", controller.HotelSearch)           // 검색결과를 hotel파라미터로 보내주면 그 호텔의 방을 모두 보내줌
	e.GET("/getuser", controller.GetUser)                   // id를 파라미터로 보내주면 유저정보를 보내줌
	e.GET("/gethotel", controller.GetHotel)                 // hotel을 파라미터로 보내주면 그 호텔의 방을 모두 보내줌
	e.GET("/gethotels", controller.GetHotels)               // 호텔을 랜덤하게 20개 보내줌
	e.GET("/getalpha", controller.GetAlpha)                 // hotel과 room을 파라미터로 받아 alpha값 반환
	//e.GET("/barcord/*any", controller.Barcord)              // 바코드
	//e.GET("/barcord", controller.Barcord)
}
