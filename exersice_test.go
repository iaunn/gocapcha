package golang

import (
	"testing"
)

func TestCaptcha(t *testing.T) {
	c := []int{1, 1, 1, 1}

	rs := getCaptcha(c)

	if rs != "1+one" {
		t.Error("it should be 1 + one")
	}

}
