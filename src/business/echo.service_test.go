package business

import (
	"testing"
)

func TestCreateResponseItem(t *testing.T) {
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
}
