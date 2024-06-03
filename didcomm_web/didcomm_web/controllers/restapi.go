package controllers

import (
	"didcomm_web/tests/agent"
	"didcomm_web/tests/contacts"
	"didcomm_web/tests/profile"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var bsUrl = "http://" + config.BackServerURL

var cs = contacts.NewContacts()

/* 웹페이지 띄우기 */
func OpenPage(c *gin.Context) {
	c.File(config.PathHtml) // HTML 파일 경로 입력. 현재 상대 경로 사용 중
	fmt.Println("openPage 종료.")
}

/* 프로필 생성 */
func GenProfile(c *gin.Context) {
	var label string
	c.BindJSON(&label)
	// 나중에 go back-server로 http 요청 코드로 변경
	// profile, _ := utils.HttpRequest(bsUrl + "/genProfile", label)
	profile := profile.GenProfile(label)
	c.JSON(http.StatusOK, profile)
}

/* 사용자(agent) 프로필 설정 */
func SetupProfile(c *gin.Context) {
	var profile *profile.Profile
	c.BindJSON(&profile)
	fmt.Printf("profile: %s\n", profile)
	// 나중에 go back-server로 http 요청 코드로 변경
	// utils.HttpRequest(bsUrl + "/genProfile", profile)
	agent.SetupProfile(profile)
	c.JSON(http.StatusOK, "")
}

/* ScopedEventBus 생성 요청 */
func NewScopedEventBus(c *gin.Context) {
	// var content string
	// c.BindJSON(&content)
	// go back-server로 http 요청
	// result, _ := utils.HttpRequest(bsUrl + "newScopedEventBus", "")
	c.JSON(http.StatusOK, "")
}

/* 연락처 목록 반환 */
func GetContacts(c *gin.Context) {
	// var content string
	// c.BindJSON(&content)
	// 나중에 go back-server로 http 요청 코드로 변경
	contacts := cs.GetContacts()
	c.JSON(http.StatusOK, contacts)
}

/* 연락처 반환 */
func GetContact(c *gin.Context) {
	var did string
	c.BindJSON(&did)
	contact := cs.GetContact(did)
	c.JSON(http.StatusOK, contact)
}

/* 연락처 추가 */
func AddContact(c *gin.Context) {
	var contact contacts.Contact
	c.BindJSON(&contact)
	cs.AddContact(contact)
	c.JSON(http.StatusOK, "")
}

// /* DID 생성 이벤트 버스 등록 */
// func OnDid(c *gin.Context) {
// 	var content string
// 	c.BindJSON(&content)
// 	eventbus.On("didGenerated", agent.OnDidGenerated)
// 	eventbus.Emit("didGenerated", "sgsgs") // 실행 후 웹페이지의 UI를 업데이트하여 화면을 다시 그려야 함. (프론트엔드 동작)
// 	c.JSON(http.StatusOK, "")
// }

// /* 통신 연결 이벤트 버스 등록 */
// func OnConnect(c *gin.Context) {
// 	var content string
// 	c.BindJSON(&content)
// 	eventbus.On("connected", agent.OnDidGenerated)
// 	c.JSON(http.StatusOK, "")
// }

// /* 통신 연결 해제 이벤트 버스 등록 */
// func OnDisconnect(c *gin.Context) {
// 	var content string
// 	c.BindJSON(&content)
// 	eventbus.On("disconnected", agent.OnDidGenerated)
// 	c.JSON(http.StatusOK, "")
// }

// // 테스트 용도 (최초 실행 관련 X)
// func Calculator() {
// 	a := 10
// 	b := 5
// 	fmt.Printf("a + b = %d\n", a+b)
// }

// func On(c *gin.Context) {
// 	eventbus.On("Cal", Calculator)
// }

// func Emit(c *gin.Context) {
// 	eventbus.Emit("Cal")
// }

// func Off(c *gin.Context) {
// 	eventbus.Off("Cal", Calculator)
// }
