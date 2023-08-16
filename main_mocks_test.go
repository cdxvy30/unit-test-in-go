package main

import (
	"testing"
)

// methodsToCall 存放所有期望要 call 的 methods
type MockSearcher struct {
	phone         string
	methodsToCall map[string]bool
}

// A mock object register calls it receives, allowing us to assert how to use it in the real code.
// A mock do not have working implementaions.
func (ms *MockSearcher) Search(people []*Person, firstName, lastName string) *Person {
	ms.methodsToCall["Search"] = true

	// Pre-defined behaviour
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearcher) ExpectToCall(methodName string) {
	if ms.methodsToCall == nil {
		ms.methodsToCall = make(map[string]bool)
	}
	ms.methodsToCall[methodName] = false
}

func (ms *MockSearcher) Verify(t *testing.T) {
	for methodName, called := range ms.methodsToCall {
		if !called {
			t.Errorf("Expect to call '%s', but it wasn't.", methodName)
		}
	}
}

func TestFindCallsSearchAndReturnsPersonUsingMock(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	mock := &MockSearcher{phone: fakePhone}
	mock.ExpectToCall("Search")

	phone, _ := phonebook.Find(mock, "Jane", "Doe")

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}

	mock.Verify(t)
}
