package routes

import (
	"didcomm_web/controllers"
	"didcomm_web/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(config *utils.Config) *gin.Engine {
	r := gin.Default()

	r.Static("/static2", config.PathBundle) // 번들 파일 경로 입력. 현재 상대 경로 사용 중

	/* 웹소켓 */
	ws := controllers.NewWebSocket()
	r.GET("/ws", ws.HandleWebSocket2)
	// r.GET("/ws", controllers.HandleWebSocket)

	/* HTTP */
	r.GET("/", controllers.OpenPage)
	r.POST("/genProfile", controllers.GenProfile)
	r.POST("/setupProfile", controllers.SetupProfile)
	r.POST("/getContacts", controllers.GetContacts)

	r.POST("/newScopedEventBus", controllers.NewScopedEventBus)
	r.POST("/getContact", controllers.GetContact)
	r.POST("/addContact", controllers.AddContact)

	return r
}
