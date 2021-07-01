package business

import (
	"testing"

	"github.com/Mondongo-cl/http-rest-echo-go/dataaccess"
)

func TestCreateResponseItem(t *testing.T) {

	tests := []struct {
		name string
	}{
		{name: "TestCreateResponseItem"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mensaje := "Este es un mensaje"
			actual, err := CreateResponseItem(1, &mensaje)
			if err != nil {
				t.Fail()
			} else {
				if actual.Id != 1 {
					t.Error("Incorrect ID Value")
				}
				if actual.Message != mensaje {
					t.Error("Incorrect Message")
				}
			}
		})
	}
}

func Test_CreateResponseItemList(t *testing.T) {
	tests := []struct {
		name string
	}{

		{name: "Test_CreateResponseItemList"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response []dataaccess.MessageRow = make([]dataaccess.MessageRow, 10)
			actual := CreateResponseItemList(response)
			if len(*actual) != 10 {
				t.Error("worng number of items in result")
			}

		})
	}
}
