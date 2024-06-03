package didcomm_web

import (
	"didcomm_web/routes"
	"didcomm_web/utils"
	"runtime"
)

func Init() {
	/* CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정 */
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 설정 파일 정보 불러오기
	config := utils.InitConfig()

	r := routes.SetupRouter(config)
	r.Run(":" + config.Port)
}
