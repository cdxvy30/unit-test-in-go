package main

import (
	"testing"
)

type StubSearcher struct {
	phone string
}

func (ss StubSearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

// 測試目標物件: main.go 中的 Find function (Phonebook 的 method)
// 以 StubSearcher 模擬 Searcher 行為
func TestFindReturnsPerson(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}

	phone, _ := phonebook.Find(StubSearcher{fakePhone}, "David", "Doe")

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}
}
