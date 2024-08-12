package tests

import "testing"

func TestController(t *testing.T) {
	err := Controller()
	if err != nil {
		t.Error(err)
	}
}
