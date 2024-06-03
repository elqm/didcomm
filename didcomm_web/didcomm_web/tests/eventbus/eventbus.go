package eventbus

import (
	"fmt"

	"github.com/asaskevich/EventBus"
)

var DefaultEB *EB

type EB struct {
	Eb EventBus.Bus
}

func Event() {
	DefaultEB = &EB{}
	DefaultEB.Eb = NewEvent()
	fmt.Println("이벤트 버스 생성 완료.")
}

func NewEvent() EventBus.Bus {
	bus := EventBus.New()
	return bus
}

func On(pattern string, callback interface{}) {
	if DefaultEB.Eb.HasCallback(pattern) {
		fmt.Println("이벤트가 이미 등록되어 있음.")
	} else {
		DefaultEB.Eb.Subscribe(pattern, callback)
		fmt.Printf("%s 이벤트 등록 완료.\n", pattern)
	}
}

// ...interface{} - 함수 정의할 때 사용. 임의의 개수(매개변수가 몇 개든 상관 없음) 만큼 매개변수를 받을 수 있음
// args... - 함수 호출할 때 사용. 가변 인수를 슬라이스로 확장하여 함수에 전달이 가능함. 입력한 만큼 모두 매개변수로 사용할 수 있음
func Emit(pattern string, args ...interface{}) {
	if DefaultEB.Eb.HasCallback(pattern) {
		DefaultEB.Eb.Publish(pattern, args...)
		fmt.Printf("%s 이벤트 실행 완료.\n", pattern)

	} else {
		fmt.Println("등록된 이벤트 없음.")
	}
}

func Off(pattern string, callback interface{}) {
	if DefaultEB.Eb.HasCallback(pattern) {
		DefaultEB.Eb.Unsubscribe(pattern, callback)
		fmt.Printf("%s 이벤트 제거 완료.\n", pattern)
	} else {
		fmt.Println("제거할 이벤트 없음.")
	}
}
