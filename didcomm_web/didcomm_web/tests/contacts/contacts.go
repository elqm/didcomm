package contacts

import "fmt"

type Contacts struct {
	Cs []Contact `json:"cs"`
}

type Contact struct {
	Did   string `json:"did"`
	Label string `json:"label"`
}

func NewContacts() *Contacts {
	return &Contacts{
		Cs: []Contact{},
	}
}

func (cs *Contacts) GetContacts() []Contact {
	// 연락처 목록 확인 테스트 용도
	// if len(cs.Cs) == 0 {
	// 	contact := &Contact{
	// 		Did:   "did",
	// 		Label: "Peter",
	// 	}
	// 	fmt.Printf("등록된 연락처 없음.\n테스트용 연락처 추가.\n")
	// 	cs.Cs = append(cs.Cs, *contact)
	// }
	fmt.Printf("연락처 목록(contacts): %s\n", cs.Cs)

	return cs.Cs
}

func (cs *Contacts) GetContact(did string) *Contact {
	contact := &Contact{}
	if len(cs.Cs) != 0 {
		for i := 0; i < len(cs.Cs); i++ {
			if cs.Cs[i].Did == did {
				contact.Did = did
				return contact
			}
		}
	}
	fmt.Printf("연락처 반환: %s\n", contact)
	return nil
}

func (cs *Contacts) AddContact(contact Contact) {
	cs.Cs = append(cs.Cs, contact)
}
