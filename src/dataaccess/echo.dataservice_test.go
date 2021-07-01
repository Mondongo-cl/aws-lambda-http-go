package dataaccess

import (
	"fmt"
	"testing"
)

func init_test() {

	data = map[int32]*MessageRow{
		1: {1, "message 1"},
		2: {2, "message 2"},
		3: {3, "message 3"},
		4: {4, "message 4"},
		5: {5, "message 5"},
	}
}

func Test_Find(t *testing.T) {
	init_test()
	var msg string = "message 1"
	tests := []struct {
		name string
	}{
		{name: "Find"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, _ := Find(1)

			if *v != msg {
				t.Error("Found value is not expected")
			}
		})
	}
}

func Test_Add(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Add"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row, err := Add("Test x")
			if err != nil {
				t.Error(err.Error())
				return
			}
			if row.Id == 0 {
				t.Error("Assigned ID is not Valid")
				return
			}
			if row.Message != "Test x" {
				t.Error("Value is not expected")
			}
		})
	}
}

func Test_Remove(t *testing.T) {
	init_test()
	tests := []struct {
		name string
	}{
		{name: "Remove"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Remove(5); err != nil {
				t.Error("Delete Operation Failed ", err.Error())
			}
			if len(data) != 4 {
				t.Error("Delete Operation failed, length is not expected the list contains ", len(data), " elements")
				return
			}
			for _, v := range data {
				if v.Message != fmt.Sprintf("message %d", v.Id) {
					t.Error(fmt.Sprintf("Item with id %d is not valid", v.Id))
				}
			}

		})
	}
}

func Test_GetAll(t *testing.T) {
	init_test()
	tests := []struct {
		name string
	}{
		{name: "GetAll"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetAll()
			if err != nil {
				t.Error(err.Error())
				return
			}
			if len(actual) != 5 {
				t.Error(fmt.Sprintf("Collection no get correct lenght, the size returned is %d", len(actual)))
				return
			}
			for _, v := range actual {
				if v.Message != fmt.Sprintf("message %d", v.Id) {
					t.Error(fmt.Sprintf("Item Message is not correct, the value get is %s but the expected is %s", v.Message, fmt.Sprintf("message %d", v.Id)))
				}
			}
		})
	}
}
