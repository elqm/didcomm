package profile

import (
	"fmt"
	"math/rand"
)

type Profile struct {
	Label string `json:"label"`
	Did   string `json:"did"`
}

var DefaultProfile = &Profile{}

var DefaultActors = []string{
	"Alice",
	"Bob",
	"Carol",
	"Dave",
	"Eve",
	"Faythe",
	"Grace",
	"Heidi",
	"Ivan",
	"Judy",
	"Karl",
	"Lloyd",
	"Mallory",
	"Nia",
	"Oscar",
	"Peggy",
	"Quentin",
	"Rupert",
	"Sybil",
	"Trent",
	"Ursula",
	"Victor",
	"Walter",
	"Xavier",
	"Yvonne",
	"Zoe",
}

func getRandomActor() string {
	ranNum := rand.Float64()
	ranIdx := int(ranNum * float64(len(DefaultActors)))
	ranActor := DefaultActors[ranIdx]
	return ranActor
}

func GenProfile(content string) *Profile {
	fmt.Println("프로필 생성 시작.")
	fmt.Printf("입력한 label(content): %v\n", content)
	if content == "" {
		DefaultProfile.Label = getRandomActor()
	} else {
		DefaultProfile.Label = content
	}
	fmt.Printf("label: %v\n", DefaultProfile)
	return DefaultProfile
}
