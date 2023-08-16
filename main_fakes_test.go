package main

import (
	"testing"
)

type FakeSearcher struct{}

// A fake has an actual implementation
// 模擬真實搜尋資料的情況，但沒有真的跟資料庫連線
func (fs FakeSearcher) Search(people []*Person, firstName string, lastName string) *Person {
	if len(people) == 0 {
		return nil
	}

	return people[0]
}

func TestFindCallsSearchAndReturnsEmptyStringForNoPerson(t *testing.T) {
	phonebook := &Phonebook{}
	fake := &FakeSearcher{}

	phone, _ := phonebook.Find(fake, "Jane", "Doe")

	if phone != "" {
		t.Errorf("Wanted '', got '%s'", phone)
	}
}
