package agent

import (
	"didcomm_web/tests/profile"
	"encoding/json"
	"fmt"
	"log"
)

type Agent struct {
	Profile *profile.Profile `json:"profile"`
}

var DefaultAgent = &Agent{}

func SetupProfile(profile *profile.Profile) {
	DefaultAgent.Profile = profile
	// 구조체 확인용 json 변환
	jsonAgent, err := json.Marshal(DefaultAgent)
	if err != nil {
		log.Printf("err: %v\n", err)
	}
	fmt.Printf("agent: %s\n", jsonAgent)
}

func OnDidGenerated(did string) {
	DefaultAgent.Profile.Did = did
	// 구조체 확인용 json 변환
	jsonAgent, err := json.Marshal(DefaultAgent)
	if err != nil {
		log.Printf("err: %v\n", err)
	}
	fmt.Printf("Agent DID: %s\n", jsonAgent)
}

func OnConnect() {
	// ProfilePage 클래스의 connected 변수를 true로 저장해야 함
}

func OnDisconnect() {
	// ProfilePage 클래스의 connected 변수를 false로 저장해야 함
}
