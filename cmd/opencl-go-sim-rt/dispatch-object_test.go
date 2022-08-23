package main

import (
	"testing"
)

func TestDispatchObjectReturnsReference(t *testing.T) {
	someString := "test"
	obj := NewDispatchObject(someString)
	defer obj.Delete()

	gotString, ok := obj.Ref().(string)
	if !ok {
		t.Errorf("could not retrieve string")
	}
	if gotString != someString {
		t.Errorf("Retrieved wrong string: got: '%s', want: '%s'", gotString, someString)
	}
}

func TestDispatchObjectReturnsHandleAndID(t *testing.T) {
	someString := "test"
	obj1 := NewDispatchObject(someString)
	defer obj1.Delete()
	obj2 := NewDispatchObject(someString)
	defer obj2.Delete()

	gotHandle1 := obj1.Handle()
	gotHandle2 := obj2.Handle()
	if gotHandle1 == gotHandle2 {
		t.Errorf("handles are equal: got: %v", gotHandle1)
	}

	gotID1 := obj1.ID()
	gotID2 := obj2.ID()
	if gotID1 == gotID2 {
		t.Errorf("IDs are equal: got: %v", gotID1)
	}
}

func TestDispatchObjectFreeResetsProperties(t *testing.T) {
	obj := NewDispatchObject("abcd")
	obj.Delete()

	gotHandle := obj.Handle()
	if gotHandle != 0 {
		t.Errorf("handle is not reset: got: %v", gotHandle)
	}
	gotID := obj.ID()
	if gotID != 0 {
		t.Errorf("ID is not reset: got: %v", gotID)
	}
}
