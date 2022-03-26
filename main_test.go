package  main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_should_return_error_when_first_name_is_empty(t *testing.T) {

	got := addContact("", "last", "09876543210")
	want := errors.New("first name cannot be blank")

	if !reflect.DeepEqual(got, want) {
		t.Error("should fail when firstname is empty")
	}

}

func Test_addContact(t *testing.T) {

	got := addContact("", "last", "09876543210")
	want := errors.New("first name cannot be blank")

	if !reflect.DeepEqual(got, want) {
		t.Error("should fail when firstname is empty")
	}

}